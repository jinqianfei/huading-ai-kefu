/**
 * 单位转换模块
 * 实现客户单位到华鼎标准单位类型的识别和映射
 */

package tools

import (
	"fmt"
	"strings"
	"unicode"
)

// UnitType 华鼎标准单位类型
type UnitType string

const (
	UnitTypePiece   UnitType = "PIECE"   // 件/个
	UnitTypeBox     UnitType = "BOX"     // 箱/盒
	UnitTypeBag     UnitType = "BAG"     // 袋/包
	UnitTypeBottle  UnitType = "BOTTLE"  // 瓶/罐
	UnitTypeWeight  UnitType = "WEIGHT"  // 重量(kg/g)
	UnitTypeVolume  UnitType = "VOLUME"  // 体积(L/ml)
	UnitTypePallet  UnitType = "PALLET"  // 托/托盘
	UnitTypeCase    UnitType = "CASE"    // 件(大包装)
	UnitTypeSet     UnitType = "SET"     // 套/组
	UnitTypeUnknown UnitType = "UNKNOWN" // 未知
)

// UnitMapping 单位映射规则
type UnitMapping struct {
	CustomerUnits []string // 客户可能使用的单位变体
	StandardType  UnitType // 标准单位类型
	Category      string   // 单位分类
	Priority      int      // 优先级（数字越小优先级越高）
}

// 单位映射表（支持常见变体）
var unitMappings = []UnitMapping{
	// 件/个类
	{
		CustomerUnits: []string{"件", "个", "只", "支", "条", "根", "台", "辆", "张", "把", "块", "片", "粒", "颗", "本", "份"},
		StandardType:  UnitTypePiece,
		Category:      "单品",
		Priority:      1,
	},
	// 箱/盒类
	{
		CustomerUnits: []string{"箱", "盒", " carton", "box", "case", "ctn"},
		StandardType:  UnitTypeBox,
		Category:      "包装",
		Priority:      1,
	},
	// 袋/包类
	{
		CustomerUnits: []string{"袋", "包", " sack", "bag", "pkg", "package"},
		StandardType:  UnitTypeBag,
		Category:      "包装",
		Priority:      1,
	},
	// 瓶/罐类
	{
		CustomerUnits: []string{"瓶", "罐", "听", "杯", "桶", "壶", "bottle", "can", "jar", "drum"},
		StandardType:  UnitTypeBottle,
		Category:      "容器",
		Priority:      1,
	},
	// 重量类 - 千克
	{
		CustomerUnits: []string{"千克", "公斤", "kg", "kgs", "kilogram", "kilograms", "kilo"},
		StandardType:  UnitTypeWeight,
		Category:      "重量",
		Priority:      1,
	},
	// 重量类 - 克
	{
		CustomerUnits: []string{"克", "g", "gram", "grams", "gm", "gms"},
		StandardType:  UnitTypeWeight,
		Category:      "重量",
		Priority:      2,
	},
	// 重量类 - 吨
	{
		CustomerUnits: []string{"吨", "t", "ton", "tons", "tonne", "tonnes"},
		StandardType:  UnitTypeWeight,
		Category:      "重量",
		Priority:      3,
	},
	// 重量类 - 斤/两
	{
		CustomerUnits: []string{"斤", "市斤", "两", "市两"},
		StandardType:  UnitTypeWeight,
		Category:      "重量(市制)",
		Priority:      4,
	},
	// 体积类 - 升
	{
		CustomerUnits: []string{"升", "l", "L", "liter", "liters", "litre", "litres"},
		StandardType:  UnitTypeVolume,
		Category:      "体积",
		Priority:      1,
	},
	// 体积类 - 毫升
	{
		CustomerUnits: []string{"毫升", "ml", "mL", "milliliter", "milliliters"},
		StandardType:  UnitTypeVolume,
		Category:      "体积",
		Priority:      2,
	},
	// 托盘类
	{
		CustomerUnits: []string{"托", "托盘", "板", "pallet", "plt", "plts"},
		StandardType:  UnitTypePallet,
		Category:      "运输",
		Priority:      1,
	},
	// 大件类
	{
		CustomerUnits: []string{"大件", "大包装", "bulk"},
		StandardType:  UnitTypeCase,
		Category:      "包装",
		Priority:      2,
	},
	// 套/组类
	{
		CustomerUnits: []string{"套", "组", "副", "对", "双", "set", "sets", "kit", "kits", "pair", "pairs"},
		StandardType:  UnitTypeSet,
		Category:      "组合",
		Priority:      1,
	},
}

// UnitConversionRule 单位转换规则
type UnitConversionRule struct {
	FromUnit     string  // 源单位
	ToUnit       string  // 目标单位
	FromType     UnitType
	ToType       UnitType
	Ratio        float64 // 转换比例
	IsStandard   bool    // 是否为标准转换
	Description  string  // 说明
}

// 常见单位转换规则
var conversionRules = []UnitConversionRule{
	// 重量转换
	{FromUnit: "g", ToUnit: "kg", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 0.001, IsStandard: true, Description: "克转千克"},
	{FromUnit: "kg", ToUnit: "g", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 1000, IsStandard: true, Description: "千克转克"},
	{FromUnit: "t", ToUnit: "kg", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 1000, IsStandard: true, Description: "吨转千克"},
	{FromUnit: "kg", ToUnit: "t", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 0.001, IsStandard: true, Description: "千克转吨"},
	{FromUnit: "斤", ToUnit: "kg", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 0.5, IsStandard: true, Description: "市斤转千克"},
	{FromUnit: "两", ToUnit: "kg", FromType: UnitTypeWeight, ToType: UnitTypeWeight, Ratio: 0.05, IsStandard: true, Description: "市两转千克"},

	// 体积转换
	{FromUnit: "ml", ToUnit: "l", FromType: UnitTypeVolume, ToType: UnitTypeVolume, Ratio: 0.001, IsStandard: true, Description: "毫升转升"},
	{FromUnit: "l", ToUnit: "ml", FromType: UnitTypeVolume, ToType: UnitTypeVolume, Ratio: 1000, IsStandard: true, Description: "升转毫升"},
}

// GetUnitType 解析客户单位，返回华鼎标准单位类型
// 函数签名: GetUnitType(customerUnit string) (string, error)
func GetUnitType(customerUnit string) (string, error) {
	if customerUnit == "" {
		return "", fmt.Errorf("customer unit cannot be empty")
	}

	// 标准化输入
	unit := normalizeUnit(customerUnit)

	// 精确匹配
	unitType, found := matchUnitExact(unit)
	if found {
		return string(unitType), nil
	}

	// 模糊匹配
	unitType, found = matchUnitFuzzy(unit)
	if found {
		return string(unitType), nil
	}

	// 尝试提取单位（去除数字和特殊字符）
	extractedUnit := extractUnitFromString(unit)
	if extractedUnit != "" && extractedUnit != unit {
		unitType, found = matchUnitExact(extractedUnit)
		if found {
			return string(unitType), nil
		}
	}

	return string(UnitTypeUnknown), fmt.Errorf("unable to recognize unit: %s", customerUnit)
}

// GetUnitTypeWithConfidence 解析客户单位，返回单位类型和置信度
func GetUnitTypeWithConfidence(customerUnit string) (UnitType, float64, error) {
	if customerUnit == "" {
		return UnitTypeUnknown, 0, fmt.Errorf("customer unit cannot be empty")
	}

	// 标准化输入
	unit := normalizeUnit(customerUnit)

	// 精确匹配
	unitType, found := matchUnitExact(unit)
	if found {
		return unitType, 1.0, nil
	}

	// 模糊匹配
	unitType, confidence := matchUnitFuzzyWithScore(unit)
	if confidence >= 0.7 {
		return unitType, confidence, nil
	}

	return UnitTypeUnknown, confidence, fmt.Errorf("unable to recognize unit: %s", customerUnit)
}

// GetStandardUnitName 获取标准单位名称
func GetStandardUnitName(unitType UnitType) string {
	switch unitType {
	case UnitTypePiece:
		return "件"
	case UnitTypeBox:
		return "箱"
	case UnitTypeBag:
		return "袋"
	case UnitTypeBottle:
		return "瓶"
	case UnitTypeWeight:
		return "千克"
	case UnitTypeVolume:
		return "升"
	case UnitTypePallet:
		return "托"
	case UnitTypeCase:
		return "大件"
	case UnitTypeSet:
		return "套"
	default:
		return "未知"
	}
}

// GetUnitCategory 获取单位分类
func GetUnitCategory(unitType UnitType) string {
	switch unitType {
	case UnitTypePiece:
		return "单品"
	case UnitTypeBox, UnitTypeBag, UnitTypeCase:
		return "包装"
	case UnitTypeBottle:
		return "容器"
	case UnitTypeWeight:
		return "重量"
	case UnitTypeVolume:
		return "体积"
	case UnitTypePallet:
		return "运输"
	case UnitTypeSet:
		return "组合"
	default:
		return "未知"
	}
}

// ConvertUnit 单位转换
func ConvertUnit(value float64, fromUnit, toUnit string) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value cannot be negative")
	}

	fromType, err := GetUnitType(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("invalid from unit: %w", err)
	}

	toType, err := GetUnitType(toUnit)
	if err != nil {
		return 0, fmt.Errorf("invalid to unit: %w", err)
	}

	// 同类型直接查找转换规则
	if fromType == toType {
		for _, rule := range conversionRules {
			if strings.EqualFold(normalizeUnit(rule.FromUnit), normalizeUnit(fromUnit)) &&
				strings.EqualFold(normalizeUnit(rule.ToUnit), normalizeUnit(toUnit)) {
				return value * rule.Ratio, nil
			}
		}
	}

	return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
}

// GetAllSupportedUnits 获取所有支持的单位列表
func GetAllSupportedUnits() []string {
	unitSet := make(map[string]bool)
	for _, mapping := range unitMappings {
		for _, unit := range mapping.CustomerUnits {
			unitSet[unit] = true
		}
	}

	var units []string
	for unit := range unitSet {
		units = append(units, unit)
	}
	return units
}

// GetUnitsByType 获取指定类型的所有单位
func GetUnitsByType(unitType UnitType) []string {
	var units []string
	for _, mapping := range unitMappings {
		if mapping.StandardType == unitType {
			units = append(units, mapping.CustomerUnits...)
		}
	}
	return units
}

// AddCustomUnitMapping 添加自定义单位映射（用于扩展）
func AddCustomUnitMapping(customerUnit string, standardType UnitType, category string) error {
	if customerUnit == "" {
		return fmt.Errorf("customer unit cannot be empty")
	}

	normalized := normalizeUnit(customerUnit)

	// 检查是否已存在
	for _, mapping := range unitMappings {
		for _, unit := range mapping.CustomerUnits {
			if strings.EqualFold(unit, normalized) {
				return fmt.Errorf("unit %s already exists", customerUnit)
			}
		}
	}

	// 添加到映射表
	unitMappings = append(unitMappings, UnitMapping{
		CustomerUnits: []string{normalized},
		StandardType:  standardType,
		Category:      category,
		Priority:      5, // 自定义映射优先级较低
	})

	return nil
}

// IsCompatibleUnit 检查两个单位是否兼容（同类型）
func IsCompatibleUnit(unit1, unit2 string) bool {
	type1, err1 := GetUnitType(unit1)
	type2, err2 := GetUnitType(unit2)

	if err1 != nil || err2 != nil {
		return false
	}

	return type1 == type2
}

// ==================== 内部辅助函数 ====================

// normalizeUnit 标准化单位字符串
func normalizeUnit(unit string) string {
	// 去除前后空格
	unit = strings.TrimSpace(unit)

	// 转换为小写（英文部分）
	var result strings.Builder
	for _, r := range unit {
		if r >= 'A' && r <= 'Z' {
			result.WriteRune(r + 32) // 转小写
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// matchUnitExact 精确匹配单位
func matchUnitExact(unit string) (UnitType, bool) {
	for _, mapping := range unitMappings {
		for _, candidate := range mapping.CustomerUnits {
			if strings.EqualFold(unit, candidate) {
				return mapping.StandardType, true
			}
		}
	}
	return UnitTypeUnknown, false
}

// matchUnitFuzzy 模糊匹配单位
func matchUnitFuzzy(unit string) (UnitType, bool) {
	unitType, confidence := matchUnitFuzzyWithScore(unit)
	if confidence >= 0.7 {
		return unitType, true
	}
	return UnitTypeUnknown, false
}

// matchUnitFuzzyWithScore 模糊匹配单位并返回置信度
func matchUnitFuzzyWithScore(unit string) (UnitType, float64) {
	bestType := UnitTypeUnknown
	bestScore := 0.0

	for _, mapping := range unitMappings {
		for _, candidate := range mapping.CustomerUnits {
			score := calculateUnitSimilarity(unit, candidate)
			if score > bestScore {
				bestScore = score
				bestType = mapping.StandardType
			}
		}
	}

	return bestType, bestScore
}

// calculateUnitSimilarity 计算单位相似度
func calculateUnitSimilarity(s1, s2 string) float64 {
	if s1 == "" || s2 == "" {
		return 0
	}

	// 完全匹配
	if s1 == s2 {
		return 1.0
	}

	// 包含关系
	if strings.Contains(s1, s2) || strings.Contains(s2, s1) {
		longer := float64(len(s1))
		shorter := float64(len(s2))
		if len(s2) > len(s1) {
			longer = float64(len(s2))
			shorter = float64(len(s1))
		}
		return 0.8 + 0.2*(shorter/longer)
	}

	// 编辑距离相似度
	distance := levenshteinDistanceForUnit(s1, s2)
	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	if maxLen == 0 {
		return 0
	}

	similarity := 1.0 - float64(distance)/float64(maxLen)
	return similarity
}

// levenshteinDistanceForUnit 计算编辑距离（优化版）
func levenshteinDistanceForUnit(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 使用滚动数组
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
			curr[j] = minForUnit(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}

	return prev[n]
}

// minForUnit 返回三个整数中的最小值
func minForUnit(a, b, c int) int {
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

// extractUnitFromString 从字符串中提取单位部分
func extractUnitFromString(s string) string {
	// 去除数字和小数点
	var result strings.Builder
	for _, r := range s {
		if !unicode.IsDigit(r) && r != '.' && r != ',' && r != ' ' {
			result.WriteRune(r)
		}
	}
	return strings.TrimSpace(result.String())
}

// ParseQuantityWithUnit 解析带单位的数量字符串
func ParseQuantityWithUnit(input string) (float64, string, error) {
	if input == "" {
		return 0, "", fmt.Errorf("input cannot be empty")
	}

	// 提取数字部分
	var numStr strings.Builder
	var unitStr strings.Builder
	foundDigit := false

	for i, r := range input {
		if unicode.IsDigit(r) || r == '.' || r == ',' {
			// 处理逗号作为千位分隔符的情况
			if r == ',' {
				continue
			}
			numStr.WriteRune(r)
			foundDigit = true
		} else if foundDigit {
			// 数字之后的内容都是单位
			unitStr.WriteString(input[i:])
			break
		}
	}

	if numStr.Len() == 0 {
		return 0, "", fmt.Errorf("no number found in input: %s", input)
	}

	quantity, err := parseFloat(numStr.String())
	if err != nil {
		return 0, "", fmt.Errorf("failed to parse quantity: %w", err)
	}

	unit := strings.TrimSpace(unitStr.String())
	if unit == "" {
		// 尝试从整个字符串提取单位
		unit = extractUnitFromString(input)
	}

	return quantity, unit, nil
}

// parseFloat 解析浮点数
func parseFloat(s string) (float64, error) {
	var result float64
	var divisor float64 = 1
	decimalFound := false

	for _, r := range s {
		if r == '.' {
			decimalFound = true
			continue
		}
		if r >= '0' && r <= '9' {
			digit := float64(r - '0')
			if decimalFound {
				divisor *= 10
				result += digit / divisor
			} else {
				result = result*10 + digit
			}
		}
	}

	return result, nil
}

// ValidateUnitConversion 验证单位转换是否可行
func ValidateUnitConversion(fromUnit, toUnit string) error {
	fromType, err := GetUnitType(fromUnit)
	if err != nil {
		return fmt.Errorf("invalid source unit: %w", err)
	}

	toType, err := GetUnitType(toUnit)
	if err != nil {
		return fmt.Errorf("invalid target unit: %w", err)
	}

	if fromType != toType {
		return fmt.Errorf("cannot convert between different unit types: %s -> %s", fromType, toType)
	}

	// 检查是否有转换规则
	found := false
	for _, rule := range conversionRules {
		if strings.EqualFold(normalizeUnit(rule.FromUnit), normalizeUnit(fromUnit)) &&
			strings.EqualFold(normalizeUnit(rule.ToUnit), normalizeUnit(toUnit)) {
			found = true
			break
		}
	}

	if !found && fromUnit != toUnit {
		return fmt.Errorf("no conversion rule found for %s -> %s", fromUnit, toUnit)
	}

	return nil
}

// GetConversionRatio 获取单位转换比例
func GetConversionRatio(fromUnit, toUnit string) (float64, error) {
	for _, rule := range conversionRules {
		if strings.EqualFold(normalizeUnit(rule.FromUnit), normalizeUnit(fromUnit)) &&
			strings.EqualFold(normalizeUnit(rule.ToUnit), normalizeUnit(toUnit)) {
			return rule.Ratio, nil
		}
	}
	return 0, fmt.Errorf("no conversion ratio found for %s -> %s", fromUnit, toUnit)
}

// BatchConvertUnits 批量转换单位
func BatchConvertUnits(values []float64, fromUnit, toUnit string) ([]float64, error) {
	results := make([]float64, len(values))
	for i, v := range values {
		converted, err := ConvertUnit(v, fromUnit, toUnit)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value at index %d: %w", i, err)
		}
		results[i] = converted
	}
	return results, nil
}
