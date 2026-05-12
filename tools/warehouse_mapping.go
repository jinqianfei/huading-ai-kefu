// warehouse_mapping.go - 仓库映射工具
// AI客服赋能系统 - Task 1.3: get_warehouse_code 工具实现
//
// 功能说明：
// 根据客户ID和华鼎门店编码查询对应的仓库字典编码
// 用于订单处理时确定货物应从哪个仓库发出

package tools

import (
	"database/sql"
	"fmt"
	"strings"
)

// WarehouseMappingResult 仓库映射结果结构体
type WarehouseMappingResult struct {
	WarehouseDictCode string  `json:"warehouse_dict_code"`  // 仓库字典编码
	WarehouseDictName string  `json:"warehouse_dict_name,omitempty"` // 仓库字典名称
	WarehouseCode     string  `json:"warehouse_code,omitempty"`      // 原始仓库编码
	WarehouseName     string  `json:"warehouse_name,omitempty"`      // 仓库名称
	CustomerID        string  `json:"customer_id"`          // 客户ID
	HuadingStoreCode  string  `json:"huading_store_code"`   // 华鼎门店编码
	MatchType         string  `json:"match_type"`           // 匹配类型: DIRECT/STORE_MAPPED/DEFAULT/NOT_FOUND
	Confidence        float64 `json:"confidence"`           // 匹配置信度 0.0-1.0
	Priority          int     `json:"priority,omitempty"`   // 仓库优先级
	Message           string  `json:"message,omitempty"`    // 附加信息
}

// GetWarehouseCodeByStore 通过华鼎门店编码获取仓库字典编码（工具入口函数）
//
// 参数:
//   - customerID: 客户唯一标识
//   - huadingStoreCode: 华鼎系统门店编号
//
// 返回:
//   - string: 仓库字典编码
//   - error: 错误信息
//
// 查询逻辑：
// 1. 通过 customer_warehouse_mapping 表查询
// 2. 按 customer_id + huading_store_code 关联
// 3. 优先返回 warehouse_dict_code，如为空则返回 warehouse_code
// 4. 按 priority 排序，取优先级最高的
func GetWarehouseCodeByStore(customerID, huadingStoreCode string) (string, error) {
	if customerID == "" {
		return "", fmt.Errorf("customer_id is required")
	}

	if huadingStoreCode == "" {
		return "", fmt.Errorf("huading_store_code is required")
	}

	// 清理输入
	customerID = strings.TrimSpace(customerID)
	huadingStoreCode = strings.TrimSpace(huadingStoreCode)

	// 调用数据库查询
	warehouseDictCode, err := getWarehouseCodeFromDB(customerID, huadingStoreCode)
	if err != nil {
		return "", err
	}

	if warehouseDictCode == "" {
		return "", fmt.Errorf("no warehouse mapping found for customer_id=%s, huading_store_code=%s", customerID, huadingStoreCode)
	}

	return warehouseDictCode, nil
}

// GetWarehouseCodeByStoreWithDetail 通过华鼎门店编码获取仓库字典编码（带详细结果）
//
// 参数:
//   - customerID: 客户唯一标识
//   - huadingStoreCode: 华鼎系统门店编号
//
// 返回:
//   - *WarehouseMappingResult: 详细的仓库映射结果
//   - error: 错误信息
func GetWarehouseCodeByStoreWithDetail(customerID, huadingStoreCode string) (*WarehouseMappingResult, error) {
	if customerID == "" {
		return nil, fmt.Errorf("customer_id is required")
	}

	if huadingStoreCode == "" {
		return nil, fmt.Errorf("huading_store_code is required")
	}

	// 清理输入
	customerID = strings.TrimSpace(customerID)
	huadingStoreCode = strings.TrimSpace(huadingStoreCode)

	// 调用数据库查询
	warehouseDictCode, err := getWarehouseCodeFromDB(customerID, huadingStoreCode)
	if err != nil {
		// 如果没有找到门店特定的映射，尝试获取客户默认仓库
		defaultCode, defaultErr := getDefaultWarehouseCode(customerID)
		if defaultErr == nil && defaultCode != "" {
			return &WarehouseMappingResult{
				WarehouseDictCode: defaultCode,
				CustomerID:        customerID,
				HuadingStoreCode:  huadingStoreCode,
				MatchType:         "DEFAULT",
				Confidence:        0.6,
				Message:           fmt.Sprintf("using default warehouse for customer %s", customerID),
			}, nil
		}

		return &WarehouseMappingResult{
			CustomerID:       customerID,
			HuadingStoreCode: huadingStoreCode,
			MatchType:        "NOT_FOUND",
			Confidence:       0.0,
			Message:          fmt.Sprintf("no warehouse mapping found for customer_id=%s, huading_store_code=%s", customerID, huadingStoreCode),
		}, fmt.Errorf("no warehouse mapping found")
	}

	return &WarehouseMappingResult{
		WarehouseDictCode: warehouseDictCode,
		CustomerID:        customerID,
		HuadingStoreCode:  huadingStoreCode,
		MatchType:         "STORE_MAPPED",
		Confidence:        0.95,
		Message:           "warehouse found via store mapping",
	}, nil
}

// GetWarehouseCodeByCustomer 通过客户ID获取默认仓库编码
//
// 参数:
//   - customerID: 客户唯一标识
//
// 返回:
//   - string: 仓库字典编码
//   - error: 错误信息
func GetWarehouseCodeByCustomer(customerID string) (string, error) {
	if customerID == "" {
		return "", fmt.Errorf("customer_id is required")
	}

	customerID = strings.TrimSpace(customerID)

	warehouseDictCode, err := getDefaultWarehouseCode(customerID)
	if err != nil {
		return "", err
	}

	return warehouseDictCode, nil
}

// getWarehouseCodeFromDB 从数据库查询仓库编码（通过门店关联）
func getWarehouseCodeFromDB(customerID, huadingStoreCode string) (string, error) {
	// 使用 db.go 中定义的函数
	return GetWarehouseCodeByStore(customerID, huadingStoreCode)
}

// getDefaultWarehouseCode 获取客户默认仓库编码
func getDefaultWarehouseCode(customerID string) (string, error) {
	// 使用 db.go 中定义的简化查询函数
	return GetWarehouseCodeByStoreDirect(customerID, "")
}

// BatchGetWarehouseCodes 批量获取仓库编码
//
// 参数:
//   - customerID: 客户唯一标识
//   - huadingStoreCodes: 华鼎门店编码列表
//
// 返回:
//   - map[string]*WarehouseMappingResult: 门店编码到结果的映射
//   - []error: 错误列表
func BatchGetWarehouseCodes(customerID string, huadingStoreCodes []string) (map[string]*WarehouseMappingResult, []error) {
	results := make(map[string]*WarehouseMappingResult)
	var errors []error

	for _, storeCode := range huadingStoreCodes {
		storeCode = strings.TrimSpace(storeCode)
		if storeCode == "" {
			continue
		}

		result, err := GetWarehouseCodeByStoreWithDetail(customerID, storeCode)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to get warehouse for store %s: %w", storeCode, err))
		}

		results[storeCode] = result
	}

	return results, errors
}

// ValidateWarehouseCode 验证仓库编码是否有效
//
// 参数:
//   - warehouseCode: 仓库字典编码
//
// 返回:
//   - bool: 是否有效
//   - error: 错误信息
func ValidateWarehouseCode(warehouseCode string) (bool, error) {
	if warehouseCode == "" {
		return false, fmt.Errorf("warehouse_code is required")
	}

	// 查询仓库字典表验证
	query := `
		SELECT COUNT(*) 
		FROM warehouse_dict 
		WHERE warehouse_code = $1 
		  AND enabled = true
	`

	var count int
	err := DB.QueryRow(query, warehouseCode).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to validate warehouse code: %w", err)
	}

	return count > 0, nil
}

// GetWarehouseInfo 获取仓库详细信息
//
// 参数:
//   - warehouseCode: 仓库字典编码
//
// 返回:
//   - map[string]interface{}: 仓库信息
//   - error: 错误信息
func GetWarehouseInfo(warehouseCode string) (map[string]interface{}, error) {
	if warehouseCode == "" {
		return nil, fmt.Errorf("warehouse_code is required")
	}

	query := `
		SELECT warehouse_code, warehouse_name, region, enabled
		FROM warehouse_dict
		WHERE warehouse_code = $1
	`

	var code, name, region string
	var enabled bool

	err := DB.QueryRow(query, warehouseCode).Scan(&code, &name, &region, &enabled)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("warehouse not found: %s", warehouseCode)
		}
		return nil, fmt.Errorf("failed to get warehouse info: %w", err)
	}

	return map[string]interface{}{
		"warehouse_code": code,
		"warehouse_name": name,
		"region":         region,
		"enabled":        enabled,
	}, nil
}

// GetAllWarehouses 获取所有可用仓库列表
//
// 返回:
//   - []map[string]interface{}: 仓库列表
//   - error: 错误信息
func GetAllWarehouses() ([]map[string]interface{}, error) {
	query := `
		SELECT warehouse_code, warehouse_name, region
		FROM warehouse_dict
		WHERE enabled = true
		ORDER BY warehouse_code
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query warehouses: %w", err)
	}
	defer rows.Close()

	var warehouses []map[string]interface{}
	for rows.Next() {
		var code, name, region string
		if err := rows.Scan(&code, &name, &region); err != nil {
			return nil, fmt.Errorf("failed to scan warehouse row: %w", err)
		}

		warehouses = append(warehouses, map[string]interface{}{
			"warehouse_code": code,
			"warehouse_name": name,
			"region":         region,
		})
	}

	return warehouses, nil
}

// GetCustomerWarehouseMappings 获取客户的所有仓库映射
//
// 参数:
//   - customerID: 客户唯一标识
//
// 返回:
//   - []map[string]interface{}: 仓库映射列表
//   - error: 错误信息
func GetCustomerWarehouseMappings(customerID string) ([]map[string]interface{}, error) {
	if customerID == "" {
		return nil, fmt.Errorf("customer_id is required")
	}

	query := `
		SELECT cwm.warehouse_code, cwm.warehouse_name, cwm.warehouse_dict_code,
		       cwm.delivery_type, cwm.default_region, cwm.priority, cwm.enabled,
		       wd.warehouse_name as dict_name, wd.region
		FROM customer_warehouse_mapping cwm
		LEFT JOIN warehouse_dict wd ON wd.warehouse_code = COALESCE(cwm.warehouse_dict_code, cwm.warehouse_code)
		WHERE cwm.customer_id = $1
		ORDER BY cwm.priority ASC, cwm.warehouse_code ASC
	`

	rows, err := DB.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query customer warehouse mappings: %w", err)
	}
	defer rows.Close()

	var mappings []map[string]interface{}
	for rows.Next() {
		var warehouseCode, warehouseName, warehouseDictCode string
		var deliveryType, defaultRegion string
		var priority int
		var enabled bool
		var dictName, region sql.NullString

		err := rows.Scan(
			&warehouseCode, &warehouseName, &warehouseDictCode,
			&deliveryType, &defaultRegion, &priority, &enabled,
			&dictName, &region,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan mapping row: %w", err)
		}

		mappings = append(mappings, map[string]interface{}{
			"warehouse_code":      warehouseCode,
			"warehouse_name":      warehouseName,
			"warehouse_dict_code": warehouseDictCode,
			"delivery_type":       deliveryType,
			"default_region":      defaultRegion,
			"priority":            priority,
			"enabled":             enabled,
			"dict_name":           dictName.String,
			"region":              region.String,
		})
	}

	return mappings, nil
}


