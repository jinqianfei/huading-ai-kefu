// db.go - 数据库查询函数
// AI客服赋能系统 - 工具函数数据库层

package tools

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// DB 数据库连接池
var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// HuadingStoreMapping 华鼎门店映射结构体
type HuadingStoreMapping struct {
	ID                 int64   `json:"id"`
	CustomerID         string  `json:"customer_id"`
	CustomerStoreCode  string  `json:"customer_store_code"`
	CustomerStoreName  string  `json:"customer_store_name"`
	CustomerAddress    string  `json:"customer_address"`
	HuadingStoreCode   string  `json:"huading_store_code"`
	HuadingStoreName   string  `json:"huading_store_name"`
	HuadingAddress     string  `json:"huading_address"`
	MatchConfidence    float64 `json:"match_confidence"`
	MatchType          string  `json:"match_type"`
	Status             string  `json:"status"`
}

// GetHuadingStoreCode 查询华鼎门店编码
// 查询逻辑：
// 1. 先精确匹配 customer_store_code
// 2. 再模糊匹配 customer_store_name
// 3. 最后地址匹配
func GetHuadingStoreCode(customerID, storeCodeOrName string) ([]HuadingStoreMapping, error) {
	if customerID == "" || storeCodeOrName == "" {
		return nil, fmt.Errorf("customer_id and store_code_or_name are required")
	}

	// 第一步：精确匹配 customer_store_code
	query := `
		SELECT id, customer_id, customer_store_code, customer_store_name, customer_address,
		       huading_store_code, huading_store_name, huading_address,
		       match_confidence, match_type, status
		FROM huading_store_mapping
		WHERE customer_id = $1
		  AND customer_store_code = $2
		  AND status = 'ACTIVE'
		ORDER BY match_confidence DESC
		LIMIT 1
	`

	rows, err := DB.Query(query, customerID, storeCodeOrName)
	if err != nil {
		return nil, fmt.Errorf("failed to query by store code: %w", err)
	}
	defer rows.Close()

	var results []HuadingStoreMapping
	for rows.Next() {
		var m HuadingStoreMapping
		err := rows.Scan(
			&m.ID, &m.CustomerID, &m.CustomerStoreCode, &m.CustomerStoreName, &m.CustomerAddress,
			&m.HuadingStoreCode, &m.HuadingStoreName, &m.HuadingAddress,
			&m.MatchConfidence, &m.MatchType, &m.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, m)
	}

	if len(results) > 0 {
		return results, nil
	}

	// 第二步：模糊匹配 customer_store_name
	fuzzyQuery := `
		SELECT id, customer_id, customer_store_code, customer_store_name, customer_address,
		       huading_store_code, huading_store_name, huading_address,
		       match_confidence, match_type, status
		FROM huading_store_mapping
		WHERE customer_id = $1
		  AND status = 'ACTIVE'
		  AND (
		      customer_store_name ILIKE $2
		      OR customer_store_name ILIKE $3
		      OR $4 ILIKE '%' || customer_store_name || '%'
		  )
		ORDER BY match_confidence DESC, 
		         CASE 
					 WHEN customer_store_name ILIKE $5 THEN 1
					 WHEN customer_store_name ILIKE $6 THEN 2
					 ELSE 3
				 END
		LIMIT 5
	`

	// 构建模糊匹配参数
	exactPattern := storeCodeOrName
	prefixPattern := storeCodeOrName + "%"
	suffixPattern := "%" + storeCodeOrName
	containsPattern := "%" + storeCodeOrName + "%"

	rows, err = DB.Query(fuzzyQuery, customerID, prefixPattern, suffixPattern, exactPattern, exactPattern, prefixPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to query by store name: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m HuadingStoreMapping
		err := rows.Scan(
			&m.ID, &m.CustomerID, &m.CustomerStoreCode, &m.CustomerStoreName, &m.CustomerAddress,
			&m.HuadingStoreCode, &m.HuadingStoreName, &m.HuadingAddress,
			&m.MatchConfidence, &m.MatchType, &m.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, m)
	}

	if len(results) > 0 {
		return results, nil
	}

	return nil, fmt.Errorf("no matching store found for customer_id=%s, store=%s", customerID, storeCodeOrName)
}

// GetHuadingStoreCodeByAddress 通过地址匹配查询华鼎门店编码
func GetHuadingStoreCodeByAddress(customerID, address string) ([]HuadingStoreMapping, error) {
	if customerID == "" || address == "" {
		return nil, fmt.Errorf("customer_id and address are required")
	}

	// 提取地址关键词进行匹配
	addressKeywords := extractAddressKeywords(address)
	if len(addressKeywords) == 0 {
		return nil, fmt.Errorf("no valid address keywords extracted")
	}

	// 构建动态查询条件
	var conditions []string
	var args []interface{}
	args = append(args, customerID)
	argIndex := 2

	for _, keyword := range addressKeywords {
		if len(keyword) >= 2 {
			conditions = append(conditions, fmt.Sprintf("customer_address ILIKE $%d", argIndex))
			args = append(args, "%"+keyword+"%")
			argIndex++
		}
	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no valid address conditions")
	}

	query := fmt.Sprintf(`
		SELECT id, customer_id, customer_store_code, customer_store_name, customer_address,
		       huading_store_code, huading_store_name, huading_address,
		       match_confidence, match_type, status
		FROM huading_store_mapping
		WHERE customer_id = $1
		  AND status = 'ACTIVE'
		  AND (%s)
		ORDER BY match_confidence DESC
		LIMIT 3
	`, strings.Join(conditions, " OR "))

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query by address: %w", err)
	}
	defer rows.Close()

	var results []HuadingStoreMapping
	for rows.Next() {
		var m HuadingStoreMapping
		err := rows.Scan(
			&m.ID, &m.CustomerID, &m.CustomerStoreCode, &m.CustomerStoreName, &m.CustomerAddress,
			&m.HuadingStoreCode, &m.HuadingStoreName, &m.HuadingAddress,
			&m.MatchConfidence, &m.MatchType, &m.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, m)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no matching store found by address for customer_id=%s", customerID)
	}

	return results, nil
}

// extractAddressKeywords 提取地址关键词
func extractAddressKeywords(address string) []string {
	// 去除常见地址连接词，提取关键地点信息
	stopWords := map[string]bool{
		"省": true, "市": true, "区": true, "县": true, "镇": true,
		"街道": true, "路": true, "号": true, "室": true, "栋": true,
	}

	var keywords []string
	// 按常见分隔符分割
	parts := strings.FieldsFunc(address, func(r rune) bool {
		return r == ' ' || r == ',' || r == '，' || r == '、' || r == '/' || r == '\\'
	})

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		// 去除停用词后缀
		for stopWord := range stopWords {
			if strings.HasSuffix(part, stopWord) && len(part) > len(stopWord) {
				keywords = append(keywords, part[:len(part)-len(stopWord)])
			}
		}
		keywords = append(keywords, part)
	}

	return keywords
}

// GetWarehouseCodeByStore 通过华鼎门店编码查询仓库字典编码
// 查询逻辑：通过 customer_warehouse_mapping 表，按 customer_id + huading_store_code 关联查询
func GetWarehouseCodeByStore(customerID, huadingStoreCode string) (string, error) {
	if customerID == "" || huadingStoreCode == "" {
		return "", fmt.Errorf("customer_id and huading_store_code are required")
	}

	query := `
		SELECT COALESCE(cwm.warehouse_dict_code, cwm.warehouse_code) as warehouse_dict_code
		FROM customer_warehouse_mapping cwm
		INNER JOIN huading_store_mapping hsm ON hsm.customer_id = cwm.customer_id
		WHERE cwm.customer_id = $1
		  AND hsm.huading_store_code = $2
		  AND cwm.enabled = true
		ORDER BY cwm.priority ASC
		LIMIT 1
	`

	var warehouseDictCode string
	err := DB.QueryRow(query, customerID, huadingStoreCode).Scan(&warehouseDictCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no warehouse mapping found for customer_id=%s, huading_store_code=%s", customerID, huadingStoreCode)
		}
		return "", fmt.Errorf("failed to query warehouse code: %w", err)
	}

	return warehouseDictCode, nil
}

// GetWarehouseCodeByStoreDirect 直接通过客户ID和门店编码查询仓库（简化版本）
func GetWarehouseCodeByStoreDirect(customerID, huadingStoreCode string) (string, error) {
	if customerID == "" || huadingStoreCode == "" {
		return "", fmt.Errorf("customer_id and huading_store_code are required")
	}

	// 简化查询：直接通过客户ID查询默认仓库
	query := `
		SELECT COALESCE(warehouse_dict_code, warehouse_code) as warehouse_dict_code
		FROM customer_warehouse_mapping
		WHERE customer_id = $1
		  AND enabled = true
		ORDER BY priority ASC
		LIMIT 1
	`

	var warehouseDictCode string
	err := DB.QueryRow(query, customerID).Scan(&warehouseDictCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no warehouse mapping found for customer_id=%s", customerID)
		}
		return "", fmt.Errorf("failed to query warehouse code: %w", err)
	}

	return warehouseDictCode, nil
}
