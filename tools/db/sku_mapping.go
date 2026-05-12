package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// SkuCandidate SKU候选
type SkuCandidate struct {
	SkuCode    string  `json:"sku_code"`
	SkuName    string  `json:"sku_name"`
	Confidence float64 `json:"confidence"`
	Source     string  `json:"source"`
}

// GetSystemSkuByProductName 根据商品名称模糊搜索系统SKU
func GetSystemSkuByProductName(ctx context.Context, productName, customerID string) ([]SkuCandidate, error) {
	// 1. 先查customer_sku_mapping中的customer_sku_name
	query := `
        SELECT system_sku_code, system_sku_name, confidence
        FROM customer_sku_mapping
        WHERE customer_id = $1 AND status = 'ACTIVE'
          AND customer_sku_name ILIKE '%' || $2 || '%'
        ORDER BY confidence DESC LIMIT 10
    `
	rows, err := db.QueryContext(ctx, query, customerID, productName)
	if err != nil {
		return nil, fmt.Errorf("按商品名查询SKU映射失败: %w", err)
	}
	defer rows.Close()

	var candidates []SkuCandidate
	seen := make(map[string]bool)
	for rows.Next() {
		var code, name string
		var conf float64
		if err := rows.Scan(&code, &name, &conf); err != nil {
			continue
		}
		if seen[code] {
			continue
		}
		seen[code] = true
		candidates = append(candidates, SkuCandidate{
			SkuCode: code, SkuName: name, Confidence: conf, Source: "alias_name",
		})
	}

	// 2. 再查system_sku表
	skuQuery := `
        SELECT sku_code, sku_name FROM system_sku
        WHERE status = 'ACTIVE'
          AND (sku_name ILIKE '%' || $1 || '%' OR product_keywords ILIKE '%' || $1 || '%')
        ORDER BY sku_code LIMIT 10
    `
	skuRows, err := db.QueryContext(ctx, skuQuery, productName)
	if err == nil {
		defer skuRows.Close()
		for skuRows.Next() {
			var code, name string
			if err := skuRows.Scan(&code, &name); err != nil {
				continue
			}
			if seen[code] {
				continue
			}
			seen[code] = true
			// 简单相似度计算
			conf := calculateNameSimilarity(productName, name)
			if conf > 0.5 {
				candidates = append(candidates, SkuCandidate{
					SkuCode: code, SkuName: name, Confidence: conf, Source: "system_sku",
				})
			}
		}
	}

	// 按置信度排序
	sortCandidates(candidates)
	if len(candidates) > 10 {
		candidates = candidates[:10]
	}
	return candidates, nil
}

// GetAllSystemSKUs 获取所有系统SKU（用于无匹配时返回完整列表）
func GetAllSystemSKUs(ctx context.Context) ([]SkuCandidate, error) {
	rows, err := db.QueryContext(ctx, `
        SELECT sku_code, sku_name FROM system_sku WHERE status = 'ACTIVE' ORDER BY sku_code
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var skus []SkuCandidate
	for rows.Next() {
		var code, name string
		if err := rows.Scan(&code, &name); err != nil {
			continue
		}
		skus = append(skus, SkuCandidate{SkuCode: code, SkuName: name, Confidence: 0.3, Source: "all"})
	}
	return skus, rows.Err()
}

// GetSystemSkuByCode 根据SKU编码获取系统SKU详情
func GetSystemSkuByCode(ctx context.Context, skuCode string) (*SystemSku, error) {
	var s SystemSku
	err := db.QueryRowContext(ctx, `
        SELECT id, sku_code, sku_name, category, unit, weight_kg, fee_coefficient,
               warehouse_code, embedding_vector_id, status, unit_config
        FROM system_sku WHERE sku_code = $1
    `, skuCode).Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.Category, &s.Unit,
		&s.WeightKg, &s.FeeCoefficient, &s.WarehouseCode, &s.EmbeddingVectorID,
		&s.Status, &s.UnitConfig)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// calculateNameSimilarity 计算两个字符串的相似度（0-1）
func calculateNameSimilarity(s1, s2 string) float64 {
	if s1 == "" || s2 == "" {
		return 0
	}
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	if s1 == s2 {
		return 1.0
	}
	// 使用最长公共子串比例作为简单相似度
	maxLen := 0
	for i := 0; i < len(s1); i++ {
		for j := i + 1; j <= len(s1); j++ {
			substr := s1[i:j]
			if strings.Contains(s2, substr) && len(substr) > maxLen {
				maxLen = len(substr)
			}
		}
	}
	refLen := len(s1)
	if len(s2) > refLen {
		refLen = len(s2)
	}
	if refLen == 0 {
		return 0
	}
	return float64(maxLen) / float64(refLen)
}

// sortCandidates 按置信度降序排序
func sortCandidates(candidates []SkuCandidate) {
	for i := 0; i < len(candidates)-1; i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[j].Confidence > candidates[i].Confidence {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}
}
