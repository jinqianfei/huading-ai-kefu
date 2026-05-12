/**
 * SKU映射扩展模块
 * 实现4层匹配策略：精确匹配→历史匹配→模糊匹配→语义匹配
 */

package tools

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"ai-cs-support/tools/db"
)

// SystemSkuCandidate 系统SKU候选结果
type SystemSkuCandidate struct {
	SkuCode    string  `json:"sku_code"`
	SkuName    string  `json:"sku_name"`
	Confidence float64 `json:"confidence"`
	MatchType  string  `json:"match_type"`
}

// MatchLevel 匹配级别定义
const (
	MatchLevelExact    = "EXACT"     // 精确匹配
	MatchLevelHistory  = "HISTORY"   // 历史匹配
	MatchLevelFuzzy    = "FUZZY"     // 模糊匹配
	MatchLevelSemantic = "SEMANTIC"  // 语义匹配
	MatchLevelNone     = "NONE"      // 无匹配
)

// GetSystemSkuByProductName 根据商品名称获取系统SKU候选列表
// 实现4层匹配策略：精确匹配→历史匹配→模糊匹配→语义匹配
func GetSystemSkuByProductName(productName, customerID string) ([]SystemSkuCandidate, error) {
	if productName == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}

	ctx := context.Background()
	candidates := make([]SystemSkuCandidate, 0)
	seen := make(map[string]bool)

	// 第1层：精确匹配（product_name_lower字段完全匹配）
	exactMatches, err := matchExact(ctx, productName, customerID, seen)
	if err != nil {
		return nil, fmt.Errorf("exact match failed: %w", err)
	}
	candidates = append(candidates, exactMatches...)

	// 第2层：历史匹配（查询该客户的历史匹配记录）
	if len(candidates) == 0 || candidates[0].Confidence < 0.95 {
		historyMatches, err := matchHistory(ctx, productName, customerID, seen)
		if err != nil {
			// 历史匹配失败不影响整体流程
			fmt.Printf("history match warning: %v\n", err)
		}
		candidates = append(candidates, historyMatches...)
	}

	// 第3层：模糊匹配（基于关键词和相似度）
	if len(candidates) == 0 || candidates[0].Confidence < 0.85 {
		fuzzyMatches, err := matchFuzzy(ctx, productName, customerID, seen)
		if err != nil {
			fmt.Printf("fuzzy match warning: %v\n", err)
		}
		candidates = append(candidates, fuzzyMatches...)
	}

	// 第4层：语义匹配（基于向量相似度）
	if len(candidates) == 0 || candidates[0].Confidence < 0.70 {
		semanticMatches, err := matchSemantic(ctx, productName, customerID, seen)
		if err != nil {
			fmt.Printf("semantic match warning: %v\n", err)
		}
		candidates = append(candidates, semanticMatches...)
	}

	// 按置信度排序并去重
	sortedCandidates := sortAndDeduplicateCandidates(candidates)

	// 限制返回数量
	if len(sortedCandidates) > 10 {
		sortedCandidates = sortedCandidates[:10]
	}

	return sortedCandidates, nil
}

// matchExact 第1层：精确匹配
func matchExact(ctx context.Context, productName, customerID string, seen map[string]bool) ([]SystemSkuCandidate, error) {
	candidates := make([]SystemSkuCandidate, 0)

	// 1.1 查询customer_sku_mapping中的精确匹配
	mapping, err := db.GetCustomerSkuMappingByName(ctx, customerID, productName)
	if err == nil && mapping != nil {
		if !seen[mapping.SystemSkuCode] {
			seen[mapping.SystemSkuCode] = true
			candidates = append(candidates, SystemSkuCandidate{
				SkuCode:    mapping.SystemSkuCode,
				SkuName:    mapping.SystemSkuName,
				Confidence: mapping.Confidence,
				MatchType:  MatchLevelExact,
			})
		}
	}

	// 1.2 查询system_sku表中的product_name_lower精确匹配
	productNameLower := strings.ToLower(strings.TrimSpace(productName))
	skus, err := db.GetSystemSkuByProductNameLower(ctx, productNameLower)
	if err == nil {
		for _, sku := range skus {
			if !seen[sku.SkuCode] {
				seen[sku.SkuCode] = true
				candidates = append(candidates, SystemSkuCandidate{
					SkuCode:    sku.SkuCode,
					SkuName:    sku.SkuName,
					Confidence: 1.0,
					MatchType:  MatchLevelExact,
				})
			}
		}
	}

	return candidates, nil
}

// matchHistory 第2层：历史匹配
func matchHistory(ctx context.Context, productName, customerID string, seen map[string]bool) ([]SystemSkuCandidate, error) {
	candidates := make([]SystemSkuCandidate, 0)

	// 查询该客户的历史匹配记录
	historyList, err := db.GetMatchHistoryByProductName(ctx, customerID, productName)
	if err != nil {
		return candidates, err
	}

	for _, history := range historyList {
		if seen[history.SystemSkuCode] {
			continue
		}
		seen[history.SystemSkuCode] = true

		// 计算置信度（历史匹配降权）
		confidence := history.Confidence * 0.95
		if history.Corrected {
			// 人工修正过的记录置信度更高
			confidence = 0.98
		}

		candidates = append(candidates, SystemSkuCandidate{
			SkuCode:    history.SystemSkuCode,
			SkuName:    history.SystemSkuName,
			Confidence: confidence,
			MatchType:  MatchLevelHistory,
		})
	}

	return candidates, nil
}

// matchFuzzy 第3层：模糊匹配
func matchFuzzy(ctx context.Context, productName, customerID string, seen map[string]bool) ([]SystemSkuCandidate, error) {
	candidates := make([]SystemSkuCandidate, 0)

	// 提取关键词
	keywords := extractKeywords(productName)
	if len(keywords) == 0 {
		keywords = []string{productName}
	}

	// 3.1 基于关键词匹配
	for _, keyword := range keywords {
		if len(keyword) < 2 {
			continue
		}

		skus, err := db.GetSystemSkuByKeywords(ctx, keyword)
		if err != nil {
			continue
		}

		for _, sku := range skus {
			if seen[sku.SkuCode] {
				continue
			}

			// 计算相似度
			similarity := calculateSimilarity(productName, sku.SkuName)
			if similarity >= 0.6 {
				seen[sku.SkuCode] = true
				candidates = append(candidates, SystemSkuCandidate{
					SkuCode:    sku.SkuCode,
					SkuName:    sku.SkuName,
					Confidence: similarity,
					MatchType:  MatchLevelFuzzy,
				})
			}
		}
	}

	// 3.2 基于字符编辑距离的模糊匹配
	allSkus, err := db.GetSystemSkuCandidates()
	if err == nil {
		for _, sku := range allSkus {
			if seen[sku.SkuCode] {
				continue
			}

			similarity := calculateLevenshteinSimilarity(productName, sku.SkuName)
			if similarity >= 0.7 {
				seen[sku.SkuCode] = true
				candidates = append(candidates, SystemSkuCandidate{
					SkuCode:    sku.SkuCode,
					SkuName:    sku.SkuName,
					Confidence: similarity,
					MatchType:  MatchLevelFuzzy,
				})
			}
		}
	}

	return candidates, nil
}

// matchSemantic 第4层：语义匹配
func matchSemantic(ctx context.Context, productName, customerID string, seen map[string]bool) ([]SystemSkuCandidate, error) {
	candidates := make([]SystemSkuCandidate, 0)

	// 获取向量嵌入（如果有的话）
	embedding, err := db.GetProductEmbedding(ctx, productName)
	if err != nil || embedding == nil {
		// 如果没有嵌入向量，使用基于词向量的简单语义匹配
		return matchSemanticByWords(ctx, productName, customerID, seen)
	}

	// 使用向量相似度搜索
	semanticMatches, err := db.SearchByEmbedding(ctx, embedding, 10)
	if err != nil {
		return candidates, err
	}

	for _, match := range semanticMatches {
		if seen[match.SkuCode] {
			continue
		}
		seen[match.SkuCode] = true

		// 向量相似度转换为置信度
		confidence := match.Score
		if confidence < 0.5 {
			continue
		}

		candidates = append(candidates, SystemSkuCandidate{
			SkuCode:    match.SkuCode,
			SkuName:    match.SkuName,
			Confidence: confidence,
			MatchType:  MatchLevelSemantic,
		})
	}

	return candidates, nil
}

// matchSemanticByWords 基于关键词的语义匹配（降级方案）
func matchSemanticByWords(ctx context.Context, productName, customerID string, seen map[string]bool) ([]SystemSkuCandidate, error) {
	candidates := make([]SystemSkuCandidate, 0)

	// 提取核心词
	coreWords := extractCoreWords(productName)
	if len(coreWords) == 0 {
		return candidates, nil
	}

	// 查询所有SKU
	allSkus, err := db.GetSystemSkuCandidates()
	if err != nil {
		return candidates, err
	}

	for _, sku := range allSkus {
		if seen[sku.SkuCode] {
			continue
		}

		// 计算核心词匹配度
		skuCoreWords := extractCoreWords(sku.SkuName)
		matchScore := calculateWordMatchScore(coreWords, skuCoreWords)

		if matchScore >= 0.5 {
			seen[sku.SkuCode] = true
			candidates = append(candidates, SystemSkuCandidate{
				SkuCode:    sku.SkuCode,
				SkuName:    sku.SkuName,
				Confidence: matchScore,
				MatchType:  MatchLevelSemantic,
			})
		}
	}

	return candidates, nil
}

// ==================== 辅助函数 ====================

// extractKeywords 提取关键词
func extractKeywords(productName string) []string {
	// 去除空格和特殊字符
	cleaned := cleanProductName(productName)

	// 按常见分隔符分割
	separators := []string{" ", "-", "_", "|", "，", ",", "、", "（", "）", "(", ")"}
	words := []string{cleaned}

	for _, sep := range separators {
		var newWords []string
		for _, w := range words {
			parts := strings.Split(w, sep)
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p != "" && len(p) >= 2 {
					newWords = append(newWords, p)
				}
			}
		}
		words = newWords
	}

	return words
}

// extractCoreWords 提取核心词（去除量词、单位等）
func extractCoreWords(productName string) []string {
	// 常见量词和单位
	stopWords := map[string]bool{
		"个": true, "件": true, "箱": true, "盒": true, "瓶": true, "袋": true,
		"kg": true, "g": true, "ml": true, "l": true, "升": true, "克": true,
		"大": true, "中": true, "小": true, "新": true, "老": true,
	}

	words := extractKeywords(productName)
	var coreWords []string

	for _, w := range words {
		w = strings.ToLower(w)
		if !stopWords[w] && len(w) >= 2 {
			coreWords = append(coreWords, w)
		}
	}

	return coreWords
}

// cleanProductName 清理商品名称
func cleanProductName(name string) string {
	// 去除前后空格
	name = strings.TrimSpace(name)

	// 去除多余空格
	var result strings.Builder
	prevSpace := false
	for _, r := range name {
		if unicode.IsSpace(r) {
			if !prevSpace {
				result.WriteRune(' ')
				prevSpace = true
			}
		} else {
			result.WriteRune(r)
			prevSpace = false
		}
	}

	return result.String()
}

// calculateSimilarity 计算字符串相似度（基于最长公共子串）
func calculateSimilarity(s1, s2 string) float64 {
	if s1 == "" || s2 == "" {
		return 0
	}

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if s1 == s2 {
		return 1.0
	}

	// 计算最长公共子串长度
	maxLen := longestCommonSubstring(s1, s2)

	// 参考长度
	refLen := len(s1)
	if len(s2) > refLen {
		refLen = len(s2)
	}

	if refLen == 0 {
		return 0
	}

	return float64(maxLen) / float64(refLen)
}

// longestCommonSubstring 计算最长公共子串长度
func longestCommonSubstring(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return 0
	}

	// 动态规划
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}

// calculateLevenshteinSimilarity 计算编辑距离相似度
func calculateLevenshteinSimilarity(s1, s2 string) float64 {
	if s1 == "" && s2 == "" {
		return 1.0
	}
	if s1 == "" || s2 == "" {
		return 0
	}

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	distance := levenshteinDistance(s1, s2)
	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	return 1.0 - float64(distance)/float64(maxLen)
}

// levenshteinDistance 计算编辑距离
func levenshteinDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 使用滚动数组优化空间
	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for j := 0; j <= n; j++ {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr[0] = i
		for j := 1; j <= n; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			curr[j] = min(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}

	return prev[n]
}

// min 返回三个整数中的最小值
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// calculateWordMatchScore 计算词匹配分数
func calculateWordMatchScore(words1, words2 []string) float64 {
	if len(words1) == 0 || len(words2) == 0 {
		return 0
	}

	matchCount := 0
	for _, w1 := range words1 {
		for _, w2 := range words2 {
			if w1 == w2 || strings.Contains(w1, w2) || strings.Contains(w2, w1) {
				matchCount++
				break
			}
		}
	}

	// 基于较小集合计算匹配率
	minLen := len(words1)
	if len(words2) < minLen {
		minLen = len(words2)
	}

	return float64(matchCount) / float64(minLen)
}

// sortAndDeduplicateCandidates 排序并去重候选结果
func sortAndDeduplicateCandidates(candidates []SystemSkuCandidate) []SystemSkuCandidate {
	// 按置信度降序排序（冒泡排序）
	n := len(candidates)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if candidates[j].Confidence > candidates[i].Confidence {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	// 去重（保留置信度最高的）
	seen := make(map[string]bool)
	var result []SystemSkuCandidate
	for _, c := range candidates {
		if !seen[c.SkuCode] {
			seen[c.SkuCode] = true
			result = append(result, c)
		}
	}

	return result
}

// SaveSkuMatchResult 保存SKU匹配结果到历史记录
func SaveSkuMatchResult(customerID, customerSkuName, systemSkuCode, systemSkuName, matchType string, confidence float64) error {
	history := &db.MatchHistory{
		CustomerID:      customerID,
		CustomerSkuCode: "", // 商品名称匹配时可能没有编码
		CustomerSkuName: customerSkuName,
		SystemSkuCode:   systemSkuCode,
		SystemSkuName:   systemSkuName,
		MatchType:       matchType,
		Confidence:      confidence,
		Corrected:       false,
	}

	return db.SaveMatchHistory(history)
}

// UpdateSkuMappingFromFeedback 根据用户反馈更新SKU映射
func UpdateSkuMappingFromFeedback(customerID, customerSkuName, correctSkuCode string) error {
	// 获取正确的SKU信息
	ctx := context.Background()
	sku, err := db.GetSystemSkuByCode(ctx, correctSkuCode)
	if err != nil || sku == nil {
		return fmt.Errorf("invalid SKU code: %s", correctSkuCode)
	}

	// 保存或更新映射关系
	mapping := &db.CustomerSkuMapping{
		CustomerID:         customerID,
		CustomerSkuCode:    "", // 商品名称匹配时可能没有编码
		CustomerSkuName:    customerSkuName,
		SystemSkuCode:      correctSkuCode,
		SystemSkuName:      sku.SkuName,
		Confidence:         1.0, // 人工确认，置信度为1
		Source:             "manual_feedback",
		UnitConversionRule: "",
		Status:             "active",
	}

	if err := db.SaveCustomerSkuMapping(mapping); err != nil {
		return fmt.Errorf("failed to save mapping: %w", err)
	}

	// 更新历史记录为已修正
	return db.UpdateMatchHistoryCorrected(customerID, customerSkuName, correctSkuCode)
}
