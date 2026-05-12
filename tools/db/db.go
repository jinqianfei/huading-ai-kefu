/**
 * OpenClaw Agent Core 数据库模块
 * 负责所有与数据库相关的操作
 */

package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"unicode"
)

// 全局数据库连接（由 main.go 初始化）
var db *sql.DB

// InitDB 初始化数据库连接
func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// ==================== 客户相关 ====================

// Customer 客户结构
type Customer struct {
	ID        string    `json:"customer_id"`
	Name      string    `json:"customer_name"`
	Code      string    `json:"customer_code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllCustomers 获取所有客户
func GetAllCustomers() ([]Customer, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT customer_id, customer_name, customer_code, status, created_at
		FROM customer WHERE status = 'active' ORDER BY customer_name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.Status, &c.CreatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}

// GetCustomer 获取单个客户
func GetCustomer(customerID string) (*Customer, error) {
	var c Customer
	err := db.QueryRowContext(context.Background(), `
		SELECT customer_id, customer_name, customer_code, status, created_at
		FROM customer WHERE customer_id = $1
	`, customerID).Scan(&c.ID, &c.Name, &c.Code, &c.Status, &c.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// ==================== SKU映射相关 ====================

// CustomerSkuMapping 客户SKU映射结构
type CustomerSkuMapping struct {
	ID                  string    `json:"id"`
	CustomerID           string    `json:"customer_id"`
	CustomerSkuCode     string    `json:"customer_sku_code"`
	CustomerSkuName     string    `json:"customer_sku_name"`
	SystemSkuCode       string    `json:"system_sku_code"`
	SystemSkuName       string    `json:"system_sku_name"`
	Confidence          float64   `json:"confidence"`
	Source              string    `json:"source"`
	UnitConversionRule  string    `json:"unit_conversion_rule"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// GetCustomerSkuMapping 获取客户商品编码→系统SKU的精确映射
func GetCustomerSkuMapping(customerID, customerSkuCode string) (*CustomerSkuMapping, error) {
	var m CustomerSkuMapping
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
			   system_sku_code, system_sku_name, confidence, source,
			   unit_conversion_rule, status, created_at, updated_at
		FROM customer_sku_mapping
		WHERE customer_id = $1 AND customer_sku_code = $2 AND status = 'active'
	`, customerID, customerSkuCode).Scan(
		&m.ID, &m.CustomerID, &m.CustomerSkuCode, &m.CustomerSkuName,
		&m.SystemSkuCode, &m.SystemSkuName, &m.Confidence, &m.Source,
		&m.UnitConversionRule, &m.Status, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// SaveCustomerSkuMapping 保存客户SKU映射
func SaveCustomerSkuMapping(m *CustomerSkuMapping) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO customer_sku_mapping (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, confidence, source,
			unit_conversion_rule, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (customer_id, customer_sku_code) DO UPDATE SET
			system_sku_code = EXCLUDED.system_sku_code,
			system_sku_name = EXCLUDED.system_sku_name,
			confidence = EXCLUDED.confidence,
			updated_at = NOW()
	`, m.CustomerID, m.CustomerSkuCode, m.CustomerSkuName,
		m.SystemSkuCode, m.SystemSkuName, m.Confidence, m.Source,
		m.UnitConversionRule, m.Status)
	return err
}

// GetCustomerSkuMappingByName 根据客户商品名称获取SKU映射
func GetCustomerSkuMappingByName(ctx context.Context, customerID, customerSkuName string) (*CustomerSkuMapping, error) {
	var m CustomerSkuMapping
	err := db.QueryRowContext(ctx, `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
		       system_sku_code, system_sku_name, confidence, source,
		       unit_conversion_rule, status, created_at, updated_at
		FROM customer_sku_mapping
		WHERE customer_id = $1 AND customer_sku_name = $2 AND status = 'active'
		ORDER BY confidence DESC
		LIMIT 1
	`, customerID, customerSkuName).Scan(
		&m.ID, &m.CustomerID, &m.CustomerSkuCode, &m.CustomerSkuName,
		&m.SystemSkuCode, &m.SystemSkuName, &m.Confidence, &m.Source,
		&m.UnitConversionRule, &m.Status, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ==================== 门店别名相关 ====================

// StoreAliasMapping 门店别名映射结构
type StoreAliasMapping struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	Alias       string  `json:"alias"`
	FullName    string  `json:"full_name"`
	MatchType   string  `json:"match_type"`
	Confidence  float64 `json:"confidence"`
	Address     string  `json:"address"`
	Contact     string  `json:"contact"`
	Phone       string  `json:"phone"`
}

// GetStoreAliasMapping 获取门店简称映射
func GetStoreAliasMapping(alias string) (*StoreAliasMapping, error) {
	var m StoreAliasMapping
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, alias, full_name, match_type,
			   confidence, address, contact, phone
		FROM customer_alias_mapping
		WHERE alias = $1 AND status = 'active'
		LIMIT 1
	`, alias).Scan(&m.ID, &m.CustomerID, &m.Alias, &m.FullName,
		&m.MatchType, &m.Confidence, &m.Address, &m.Contact, &m.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ==================== 匹配历史 ====================

// MatchHistory 匹配历史结构
type MatchHistory struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	CustomerSkuCode string    `json:"customer_sku_code"`
	CustomerSkuName string    `json:"customer_sku_name"`
	SystemSkuCode   string    `json:"system_sku_code"`
	SystemSkuName   string    `json:"system_sku_name"`
	MatchType       string    `json:"match_type"`
	Confidence      float64   `json:"confidence"`
	Corrected       bool      `json:"corrected"`
	CorrectedBy     string    `json:"corrected_by"`
	CreatedAt       time.Time `json:"created_at"`
}

// GetMatchHistory 获取匹配历史
func GetMatchHistory(customerID, customerSkuCode string) (*MatchHistory, error) {
	var h MatchHistory
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
			   system_sku_code, system_sku_name, match_type, confidence,
			   corrected, corrected_by, created_at
		FROM match_history
		WHERE customer_id = $1 AND customer_sku_code = $2
		ORDER BY created_at DESC LIMIT 1
	`, customerID, customerSkuCode).Scan(
		&h.ID, &h.CustomerID, &h.CustomerSkuCode, &h.CustomerSkuName,
		&h.SystemSkuCode, &h.SystemSkuName, &h.MatchType, &h.Confidence,
		&h.Corrected, &h.CorrectedBy, &h.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// SaveMatchHistory 保存匹配历史
func SaveMatchHistory(h *MatchHistory) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO match_history (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, match_type, confidence
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, h.CustomerID, h.CustomerSkuCode, h.CustomerSkuName,
		h.SystemSkuCode, h.SystemSkuName, h.MatchType, h.Confidence)
	return err
}

// GetMatchHistoryByProductName 根据商品名称获取匹配历史
func GetMatchHistoryByProductName(ctx context.Context, customerID, productName string) ([]MatchHistory, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
		       system_sku_code, system_sku_name, match_type, confidence,
		       corrected, corrected_by, created_at
		FROM match_history
		WHERE customer_id = $1 AND customer_sku_name ILIKE '%' || $2 || '%'
		ORDER BY created_at DESC
		LIMIT 20
	`, customerID, productName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []MatchHistory
	for rows.Next() {
		var h MatchHistory
		if err := rows.Scan(&h.ID, &h.CustomerID, &h.CustomerSkuCode, &h.CustomerSkuName,
			&h.SystemSkuCode, &h.SystemSkuName, &h.MatchType, &h.Confidence,
			&h.Corrected, &h.CorrectedBy, &h.CreatedAt); err != nil {
			continue
		}
		histories = append(histories, h)
	}
	return histories, rows.Err()
}

// UpdateMatchHistoryCorrected 更新匹配历史为已修正状态
func UpdateMatchHistoryCorrected(customerID, customerSkuName, correctSkuCode string) error {
	_, err := db.ExecContext(context.Background(), `
		UPDATE match_history
		SET corrected = true,
		    corrected_by = 'manual_feedback',
		    system_sku_code = $3,
		    updated_at = NOW()
		WHERE customer_id = $1 AND customer_sku_name = $2
	`, customerID, customerSkuName, correctSkuCode)
	return err
}

// ==================== 系统SKU ====================

// SystemSku 系统SKU结构（扩展版）
type SystemSku struct {
	ID                string    `json:"id"`
	SkuCode           string    `json:"sku_code"`
	SkuName           string    `json:"sku_name"`
	ProductNameLower  string    `json:"product_name_lower"`
	ProductKeywords   string    `json:"product_keywords"`
	Category          string    `json:"category"`
	Unit              string    `json:"unit"`
	WeightKg          float64   `json:"weight_kg"`
	FeeCoefficient    float64   `json:"fee_coefficient"`
	WarehouseCode     string    `json:"warehouse_code"`
	EmbeddingVectorID string    `json:"embedding_vector_id"`
	UnitConfig        string    `json:"unit_config"`
	Status            string    `json:"status"`
}

// SystemSkuCandidate SKU候选结构
type SystemSkuCandidate struct {
	SkuCode    string  `json:"sku_code"`
	SkuName    string  `json:"sku_name"`
	Confidence float64 `json:"confidence"`
	Score      float64 `json:"score"`
}

// GetSystemSkuCandidates 获取所有可用的系统SKU候选
func GetSystemSkuCandidates() ([]SystemSku, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, sku_code, sku_name, category, unit, status
		FROM system_sku WHERE status = 'active'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skus []SystemSku
	for rows.Next() {
		var s SystemSku
		if err := rows.Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.Category, &s.Unit, &s.Status); err != nil {
			return nil, err
		}
		skus = append(skus, s)
	}
	return skus, rows.Err()
}

// GetSystemSkuByProductName 根据商品名称获取系统SKU候选列表
// 实现4层匹配策略：精确匹配→历史匹配→模糊匹配→语义匹配
func GetSystemSkuByProductName(productName string, customerID string) ([]SystemSkuCandidate, error) {
	if productName == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}

	ctx := context.Background()
	candidates := make([]SystemSkuCandidate, 0)
	seen := make(map[string]bool)

	// 第1层：精确匹配（product_name_lower字段完全匹配）
	productNameLower := strings.ToLower(strings.TrimSpace(productName))
	exactMatches, err := GetSystemSkuByProductNameLower(ctx, productNameLower)
	if err == nil {
		for _, sku := range exactMatches {
			if !seen[sku.SkuCode] {
				seen[sku.SkuCode] = true
				candidates = append(candidates, SystemSkuCandidate{
					SkuCode:    sku.SkuCode,
					SkuName:    sku.SkuName,
					Confidence: 1.0,
					Score:      1.0,
				})
			}
		}
	}

	// 第2层：查询客户历史匹配记录
	if len(candidates) == 0 {
		historyList, err := GetMatchHistoryByProductName(ctx, customerID, productName)
		if err == nil {
			for _, history := range historyList {
				if !seen[history.SystemSkuCode] {
					seen[history.SystemSkuCode] = true
					confidence := history.Confidence * 0.95
					if history.Corrected {
						confidence = 0.98
					}
					candidates = append(candidates, SystemSkuCandidate{
						SkuCode:    history.SystemSkuCode,
						SkuName:    history.SystemSkuName,
						Confidence: confidence,
						Score:      confidence,
					})
				}
			}
		}
	}

	// 第3层：基于关键词模糊匹配
	if len(candidates) == 0 || candidates[0].Confidence < 0.85 {
		keywords := extractKeywords(productName)
		for _, keyword := range keywords {
			if len(keyword) < 2 {
				continue
			}
			skus, err := GetSystemSkuByKeywords(ctx, keyword)
			if err != nil {
				continue
			}
			for _, sku := range skus {
				if !seen[sku.SkuCode] {
					seen[sku.SkuCode] = true
					similarity := calculateNameSimilarity(productName, sku.SkuName)
					if similarity >= 0.6 {
						candidates = append(candidates, SystemSkuCandidate{
							SkuCode:    sku.SkuCode,
							SkuName:    sku.SkuName,
							Confidence: similarity,
							Score:      similarity,
						})
					}
				}
			}
		}
	}

	// 按置信度排序
	sortCandidates(candidates)

	// 限制返回数量
	if len(candidates) > 10 {
		candidates = candidates[:10]
	}

	return candidates, nil
}

// GetSystemSkuByProductNameLower 根据product_name_lower精确查询
func GetSystemSkuByProductNameLower(ctx context.Context, productNameLower string) ([]SystemSku, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, sku_code, sku_name, product_name_lower, product_keywords,
		       category, unit, status
		FROM system_sku
		WHERE product_name_lower = $1 AND status = 'active'
	`, productNameLower)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skus []SystemSku
	for rows.Next() {
		var s SystemSku
		if err := rows.Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.ProductNameLower,
			&s.ProductKeywords, &s.Category, &s.Unit, &s.Status); err != nil {
			continue
		}
		skus = append(skus, s)
	}
	return skus, rows.Err()
}

// GetSystemSkuByKeywords 根据关键词查询SKU
func GetSystemSkuByKeywords(ctx context.Context, keyword string) ([]SystemSku, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, sku_code, sku_name, product_name_lower, product_keywords,
		       category, unit, status
		FROM system_sku
		WHERE status = 'active'
		  AND (product_keywords ILIKE '%' || $1 || '%'
		       OR sku_name ILIKE '%' || $1 || '%'
		       OR product_name_lower ILIKE '%' || $1 || '%')
		LIMIT 50
	`, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skus []SystemSku
	for rows.Next() {
		var s SystemSku
		if err := rows.Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.ProductNameLower,
			&s.ProductKeywords, &s.Category, &s.Unit, &s.Status); err != nil {
			continue
		}
		skus = append(skus, s)
	}
	return skus, rows.Err()
}

// GetSystemSkuByCode 根据SKU编码获取系统SKU详情
func GetSystemSkuByCode(ctx context.Context, skuCode string) (*SystemSku, error) {
	var s SystemSku
	err := db.QueryRowContext(ctx, `
		SELECT id, sku_code, sku_name, product_name_lower, product_keywords,
		       category, unit, weight_kg, fee_coefficient, warehouse_code,
		       embedding_vector_id, status, unit_config
		FROM system_sku WHERE sku_code = $1 AND status = 'active'
	`, skuCode).Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.ProductNameLower, &s.ProductKeywords,
		&s.Category, &s.Unit, &s.WeightKg, &s.FeeCoefficient, &s.WarehouseCode,
		&s.EmbeddingVectorID, &s.Status, &s.UnitConfig)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// UpdateSystemSkuSearchFields 更新system_sku表的product_name_lower和product_keywords字段
func UpdateSystemSkuSearchFields(skuCode, skuName string) error {
	productNameLower := strings.ToLower(strings.TrimSpace(skuName))
	productKeywords := generateProductKeywords(skuName)

	_, err := db.ExecContext(context.Background(), `
		UPDATE system_sku
		SET product_name_lower = $1,
		    product_keywords = $2,
		    updated_at = NOW()
		WHERE sku_code = $3
	`, productNameLower, productKeywords, skuCode)

	return err
}

// BatchUpdateSystemSkuSearchFields 批量更新SKU搜索字段
func BatchUpdateSystemSkuSearchFields() error {
	// 获取所有需要更新的SKU
	rows, err := db.QueryContext(context.Background(), `
		SELECT sku_code, sku_name
		FROM system_sku
		WHERE status = 'active'
		  AND (product_name_lower IS NULL OR product_name_lower = ''
		       OR product_keywords IS NULL OR product_keywords = '')
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	type skuInfo struct {
		code string
		name string
	}
	var skus []skuInfo

	for rows.Next() {
		var s skuInfo
		if err := rows.Scan(&s.code, &s.name); err != nil {
			continue
		}
		skus = append(skus, s)
	}

	// 批量更新
	for _, sku := range skus {
		if err := UpdateSystemSkuSearchFields(sku.code, sku.name); err != nil {
			fmt.Printf("Failed to update SKU %s: %v\n", sku.code, err)
		}
	}

	return nil
}

// generateProductKeywords 生成商品关键词
func generateProductKeywords(skuName string) string {
	keywords := extractKeywords(skuName)
	return strings.Join(keywords, " ")
}

// extractKeywords 提取关键词
func extractKeywords(productName string) []string {
	// 清理商品名称
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

// cleanProductName 清理商品名称
func cleanProductName(name string) string {
	name = strings.TrimSpace(name)

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

// calculateNameSimilarity 计算名称相似度
func calculateNameSimilarity(s1, s2 string) float64 {
	if s1 == "" || s2 == "" {
		return 0
	}

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if s1 == s2 {
		return 1.0
	}

	// 最长公共子串
	maxLen := longestCommonSubstring(s1, s2)

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

// sortCandidates 按置信度排序
func sortCandidates(candidates []SystemSkuCandidate) {
	for i := 0; i < len(candidates)-1; i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[j].Confidence > candidates[i].Confidence {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}
}

// ==================== 向量嵌入相关 ====================

// EmbeddingVector 向量嵌入结构
type EmbeddingVector struct {
	ID        string    `json:"id"`
	SkuCode   string    `json:"sku_code"`
	Vector    []float64 `json:"vector"`
	CreatedAt time.Time `json:"created_at"`
}

// SemanticMatch 语义匹配结果
type SemanticMatch struct {
	SkuCode   string  `json:"sku_code"`
	SkuName   string  `json:"sku_name"`
	Score     float64 `json:"score"`
	VectorID  string  `json:"vector_id"`
}

// GetProductEmbedding 获取商品名称的向量嵌入
// 如果数据库中不存在，则返回nil
func GetProductEmbedding(ctx context.Context, productName string) (*EmbeddingVector, error) {
	// 首先查找是否有缓存的向量
	var ev EmbeddingVector
	var vectorData []byte

	err := db.QueryRowContext(ctx, `
		SELECT id, sku_code, vector_data, created_at
		FROM product_embeddings
		WHERE product_name_hash = MD5($1)
		LIMIT 1
	`, productName).Scan(&ev.ID, &ev.SkuCode, &vectorData, &ev.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// 解析向量数据（简化实现，实际可能需要更复杂的解析）
	ev.Vector = parseVectorData(vectorData)

	return &ev, nil
}

// SaveProductEmbedding 保存商品名称的向量嵌入
func SaveProductEmbedding(ctx context.Context, productName, skuCode string, vector []float64) error {
	vectorData := encodeVectorData(vector)

	_, err := db.ExecContext(ctx, `
		INSERT INTO product_embeddings (product_name_hash, sku_code, vector_data)
		VALUES (MD5($1), $2, $3)
		ON CONFLICT (product_name_hash) DO UPDATE SET
			sku_code = EXCLUDED.sku_code,
			vector_data = EXCLUDED.vector_data,
			created_at = NOW()
	`, productName, skuCode, vectorData)

	return err
}

// SearchByEmbedding 基于向量相似度搜索SKU
func SearchByEmbedding(ctx context.Context, embedding *EmbeddingVector, limit int) ([]SemanticMatch, error) {
	// 简化实现：使用向量ID关联查询
	// 实际实现应该使用向量数据库（如pgvector）进行相似度计算

	rows, err := db.QueryContext(ctx, `
		SELECT s.sku_code, s.sku_name, pe.id as vector_id
		FROM system_sku s
		JOIN product_embeddings pe ON s.sku_code = pe.sku_code
		WHERE s.status = 'active' AND pe.id != $1
		LIMIT $2
	`, embedding.ID, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []SemanticMatch
	for rows.Next() {
		var m SemanticMatch
		if err := rows.Scan(&m.SkuCode, &m.SkuName, &m.VectorID); err != nil {
			continue
		}
		// 简化：使用固定分数，实际应该计算向量相似度
		m.Score = 0.75
		matches = append(matches, m)
	}

	return matches, rows.Err()
}

// parseVectorData 解析向量数据（简化实现）
func parseVectorData(data []byte) []float64 {
	// 实际实现应该根据存储格式解析
	// 这里返回空切片作为占位
	return []float64{}
}

// encodeVectorData 编码向量数据（简化实现）
func encodeVectorData(vector []float64) []byte {
	// 实际实现应该根据存储格式编码
	// 这里返回空字节作为占位
	return []byte{}
}

// ==================== 合同规则相关 ====================

// ContractRule 合同规则结构
type ContractRule struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	ContractNo   string    `json:"contract_no"`
	RuleCategory string    `json:"rule_category"`
	FieldName    string    `json:"field_name"`
	FieldValue   string    `json:"field_value"`
	FieldUnit    string    `json:"field_unit"`
	Source       string    `json:"source"`
	Confidence   float64   `json:"confidence"`
	Status       string    `json:"status"`
	EffDate      time.Time `json:"effective_date"`
	ExpDate      time.Time `json:"expiry_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GetContractRules 获取客户的合同规则
func GetContractRules(customerID string) ([]ContractRule, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, customer_id, contract_no, rule_category, field_name,
			   field_value, field_unit, source, confidence, status,
			   effective_date, expiry_date, created_at, updated_at
		FROM customer_contract_rules
		WHERE customer_id = $1
			AND status = 'active'
			AND effective_date <= NOW()
			AND expiry_date >= NOW()
		ORDER BY rule_category, field_name
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []ContractRule
	for rows.Next() {
		var r ContractRule
		if err := rows.Scan(&r.ID, &r.CustomerID, &r.ContractNo, &r.RuleCategory,
			&r.FieldName, &r.FieldValue, &r.FieldUnit, &r.Source, &r.Confidence,
			&r.Status, &r.EffDate, &r.ExpDate, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		rules = append(rules, r)
	}
	return rules, rows.Err()
}

// SaveContractRule 保存合同规则
func SaveContractRule(r *ContractRule) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO customer_contract_rules (
			customer_id, contract_no, rule_category, field_name,
			field_value, field_unit, source, confidence, status,
			effective_date, expiry_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (customer_id, contract_no, field_name) DO UPDATE SET
			field_value = EXCLUDED.field_value,
			confidence = EXCLUDED.confidence,
			updated_at = NOW()
	`, r.CustomerID, r.ContractNo, r.RuleCategory, r.FieldName,
		r.FieldValue, r.FieldUnit, r.Source, r.Confidence, r.Status,
		r.EffDate, r.ExpDate)
	return err
}

// ==================== 出库记录相关 ====================

// OutboundRecord 出库记录结构
type OutboundRecord struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	OrderNo      string    `json:"order_no"`
	OutboundDate time.Time `json:"outbound_date"`
	StoreName    string    `json:"store_name"`
	SkuCode      string    `json:"sku_code"`
	SkuName      string    `json:"sku_name"`
	Quantity     float64   `json:"quantity"`
	Unit         string    `json:"unit"`
	Weight       float64   `json:"weight"`
	Region       string    `json:"region"`
}

// GetOutboundRecords 获取出库记录
func GetOutboundRecords(customerID, startDate, endDate string) ([]OutboundRecord, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, customer_id, order_no, outbound_date, store_name,
			   sku_code, sku_name, quantity, unit, weight, region
		FROM outbound_records
		WHERE customer_id = $1
			AND outbound_date >= $2
			AND outbound_date <= $3
		ORDER BY outbound_date DESC
	`, customerID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []OutboundRecord
	for rows.Next() {
		var r OutboundRecord
		if err := rows.Scan(&r.ID, &r.CustomerID, &r.OrderNo, &r.OutboundDate,
			&r.StoreName, &r.SkuCode, &r.SkuName, &r.Quantity, &r.Unit,
			&r.Weight, &r.Region); err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, rows.Err()
}
