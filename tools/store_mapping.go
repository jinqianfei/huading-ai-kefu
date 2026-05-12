// store_mapping.go - 门店映射工具
// AI客服赋能系统 - Task 1.2: get_huading_store_code 工具实现
//
// 功能说明：
// 将客户侧的门店代码或名称映射到华鼎系统的门店编号
// 支持三级匹配策略：精确匹配 → 模糊匹配 → 地址匹配

package tools

import (
	"fmt"
	"strings"
)

// StoreMappingResult 门店映射结果结构体
type StoreMappingResult struct {
	HuadingStoreCode  string  `json:"huading_store_code"`   // 华鼎门店编号
	HuadingStoreName  string  `json:"huading_store_name"`   // 华鼎门店名称
	MatchType         string  `json:"match_type"`           // 匹配类型: EXACT/FUZZY/ADDRESS/NOT_FOUND
	Confidence        float64 `json:"confidence"`           // 匹配置信度 0.0-1.0
	MatchedSource     string  `json:"matched_source"`       // 匹配来源字段
	CustomerStoreCode string  `json:"customer_store_code"`  // 客户侧门店代码
	CustomerStoreName string  `json:"customer_store_name"`  // 客户侧门店名称
	Message           string  `json:"message,omitempty"`    // 附加信息
}

// GetHuadingStoreCode 获取华鼎门店编码（工具入口函数）
//
// 参数:
//   - customerID: 客户唯一标识
//   - storeCodeOrName: 客户侧门店代码或名称
//   - address: 门店地址（可选，用于地址匹配兜底）
//
// 返回:
//   - string: 华鼎门店编号
//   - error: 错误信息
//
// 查询逻辑（按优先级）：
// 1. 精确匹配 customer_store_code
// 2. 模糊匹配 customer_store_name
// 3. 地址匹配（当提供address参数时）
func GetHuadingStoreCode(customerID, storeCodeOrName, address string) (string, error) {
	if customerID == "" {
		return "", fmt.Errorf("customer_id is required")
	}

	if storeCodeOrName == "" {
		return "", fmt.Errorf("store_code_or_name is required")
	}

	// 清理输入
	storeCodeOrName = strings.TrimSpace(storeCodeOrName)
	address = strings.TrimSpace(address)

	// 第一步：尝试精确匹配门店代码
	result, err := matchByStoreCode(customerID, storeCodeOrName)
	if err == nil && result.MatchType == "EXACT" {
		return result.HuadingStoreCode, nil
	}

	// 第二步：尝试模糊匹配门店名称
	result, err = matchByStoreName(customerID, storeCodeOrName)
	if err == nil {
		// 高置信度模糊匹配直接返回
		if result.Confidence >= 0.85 {
			return result.HuadingStoreCode, nil
		}
		// 中置信度返回但记录日志
		if result.Confidence >= 0.6 {
			// TODO: 记录低置信度匹配日志，供后续优化
			return result.HuadingStoreCode, nil
		}
	}

	// 第三步：如果提供了地址，尝试地址匹配
	if address != "" {
		result, err = matchByAddress(customerID, address)
		if err == nil {
			return result.HuadingStoreCode, nil
		}
	}

	// 所有匹配策略都失败
	return "", fmt.Errorf("unable to find matching huading store for customer_id=%s, input=%s", customerID, storeCodeOrName)
}

// GetHuadingStoreCodeWithDetail 获取华鼎门店编码（带详细结果）
//
// 参数:
//   - customerID: 客户唯一标识
//   - storeCodeOrName: 客户侧门店代码或名称
//   - address: 门店地址（可选）
//
// 返回:
//   - *StoreMappingResult: 详细的映射结果
//   - error: 错误信息
func GetHuadingStoreCodeWithDetail(customerID, storeCodeOrName, address string) (*StoreMappingResult, error) {
	if customerID == "" {
		return nil, fmt.Errorf("customer_id is required")
	}

	if storeCodeOrName == "" {
		return nil, fmt.Errorf("store_code_or_name is required")
	}

	// 清理输入
	storeCodeOrName = strings.TrimSpace(storeCodeOrName)
	address = strings.TrimSpace(address)

	// 第一步：尝试精确匹配门店代码
	result, err := matchByStoreCode(customerID, storeCodeOrName)
	if err == nil && result.MatchType == "EXACT" {
		return result, nil
	}

	// 第二步：尝试模糊匹配门店名称
	result, err = matchByStoreName(customerID, storeCodeOrName)
	if err == nil {
		return result, nil
	}

	// 第三步：如果提供了地址，尝试地址匹配
	if address != "" {
		result, err = matchByAddress(customerID, address)
		if err == nil {
			return result, nil
		}
	}

	// 所有匹配策略都失败，返回 NOT_FOUND 结果
	return &StoreMappingResult{
		MatchType:         "NOT_FOUND",
		Confidence:        0.0,
		CustomerStoreName: storeCodeOrName,
		Message:           fmt.Sprintf("no matching store found for customer_id=%s, input=%s", customerID, storeCodeOrName),
	}, fmt.Errorf("unable to find matching huading store")
}

// matchByStoreCode 通过门店代码精确匹配
func matchByStoreCode(customerID, storeCode string) (*StoreMappingResult, error) {
	mappings, err := GetHuadingStoreCode(customerID, storeCode)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return nil, fmt.Errorf("no store found by code")
	}

	// 取第一个（置信度最高的）
	m := mappings[0]

	// 判断是否为精确匹配
	matchType := "FUZZY"
	confidence := m.MatchConfidence
	if m.CustomerStoreCode == storeCode {
		matchType = "EXACT"
		confidence = 1.0
	}

	return &StoreMappingResult{
		HuadingStoreCode:  m.HuadingStoreCode,
		HuadingStoreName:  m.HuadingStoreName,
		MatchType:         matchType,
		Confidence:        confidence,
		MatchedSource:     "customer_store_code",
		CustomerStoreCode: m.CustomerStoreCode,
		CustomerStoreName: m.CustomerStoreName,
	}, nil
}

// matchByStoreName 通过门店名称模糊匹配
func matchByStoreName(customerID, storeName string) (*StoreMappingResult, error) {
	mappings, err := GetHuadingStoreCode(customerID, storeName)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return nil, fmt.Errorf("no store found by name")
	}

	// 取第一个（置信度最高的）
	m := mappings[0]

	// 计算模糊匹配置信度
	confidence := calculateNameMatchConfidence(storeName, m.CustomerStoreName)

	return &StoreMappingResult{
		HuadingStoreCode:  m.HuadingStoreCode,
		HuadingStoreName:  m.HuadingStoreName,
		MatchType:         "FUZZY",
		Confidence:        confidence,
		MatchedSource:     "customer_store_name",
		CustomerStoreCode: m.CustomerStoreCode,
		CustomerStoreName: m.CustomerStoreName,
	}, nil
}

// matchByAddress 通过地址匹配
func matchByAddress(customerID, address string) (*StoreMappingResult, error) {
	mappings, err := GetHuadingStoreCodeByAddress(customerID, address)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return nil, fmt.Errorf("no store found by address")
	}

	// 取第一个（置信度最高的）
	m := mappings[0]

	return &StoreMappingResult{
		HuadingStoreCode:  m.HuadingStoreCode,
		HuadingStoreName:  m.HuadingStoreName,
		MatchType:         "ADDRESS",
		Confidence:        0.7, // 地址匹配基础置信度
		MatchedSource:     "customer_address",
		CustomerStoreCode: m.CustomerStoreCode,
		CustomerStoreName: m.CustomerStoreName,
	}, nil
}

// calculateNameMatchConfidence 计算名称匹配置信度
func calculateNameMatchConfidence(input, matched string) float64 {
	input = strings.ToLower(strings.TrimSpace(input))
	matched = strings.ToLower(strings.TrimSpace(matched))

	if input == matched {
		return 1.0
	}

	// 完全包含
	if strings.Contains(matched, input) || strings.Contains(input, matched) {
		return 0.9
	}

	// 计算编辑距离相似度（简化版）
	distance := levenshteinDistance(input, matched)
	maxLen := len(input)
	if len(matched) > maxLen {
		maxLen = len(matched)
	}

	if maxLen == 0 {
		return 1.0
	}

	similarity := 1.0 - float64(distance)/float64(maxLen)
	return similarity
}

// levenshteinDistance 计算编辑距离
func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// 创建动态规划矩阵
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	// 初始化边界
	for i := 0; i <= len(s1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		dp[0][j] = j
	}

	// 填充矩阵
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1,      // 删除
				min(dp[i][j-1]+1,              // 插入
					dp[i-1][j-1]+cost)) // 替换
		}
	}

	return dp[len(s1)][len(s2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// BatchGetHuadingStoreCodes 批量获取华鼎门店编码
//
// 参数:
//   - customerID: 客户唯一标识
//   - storeInputs: 门店输入列表，每个元素包含 code/name/address
//
// 返回:
//   - map[string]*StoreMappingResult: 输入到结果的映射
//   - []error: 错误列表
func BatchGetHuadingStoreCodes(customerID string, storeInputs []map[string]string) (map[string]*StoreMappingResult, []error) {
	results := make(map[string]*StoreMappingResult)
	var errors []error

	for _, input := range storeInputs {
		key := input["code"]
		if key == "" {
			key = input["name"]
		}

		result, err := GetHuadingStoreCodeWithDetail(
			customerID,
			input["code"],
			input["address"],
		)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to map store %s: %w", key, err))
		}

		results[key] = result
	}

	return results, errors
}
