/**
 * OpenClaw Agent Core 单元测试框架
 * 测试工具函数的核心功能
 */

package tools

import (
	"context"
	"fmt"
	"math"
	"testing"

	"ai-cs-support/tools/db"
)

// ==================== 工具函数测试 ====================

/**
 * 测试字符串相似度计算
 * 期望：字符相似度在 0-1 之间
 */
func TestStringSimilarity(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		minScore float64 // 期望的最低分数
		maxScore float64 // 期望的最高分数
	}{
		{
			name:     "完全相同",
			s1:       "1000塑杯",
			s2:       "1000塑杯",
			minScore: 0.5,
			maxScore: 1.0,
		},
		{
			name:     "高度相似",
			s1:       "1000塑杯",
			s2:       "1000塑料杯",
			minScore: 0.4,
			maxScore: 1.0,
		},
		{
			name:     "完全不同",
			s1:       "1000塑杯",
			s2:       "速冻桑葚",
			minScore: 0.0,
			maxScore: 0.4,
		},
		{
			name:     "空字符串",
			s1:       "",
			s2:       "1000塑杯",
			minScore: 0.0,
			maxScore: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := stringSimilarity(tt.s1, tt.s2)
			if score < tt.minScore || score > tt.maxScore {
				t.Errorf("stringSimilarity(%q, %q) = %v, want between %v and %v",
					tt.s1, tt.s2, score, tt.minScore, tt.maxScore)
			}
		})
	}
}

/**
 * 测试置信度计算
 * 期望：置信度在 0-1 之间
 */
func TestCalculateRuleConfidence(t *testing.T) {
	tests := []struct {
		name     string
		rules    int // 规则数量
		minScore float64
		maxScore float64
	}{
		{
			name:     "无规则",
			rules:    0,
			minScore: 0.0,
			maxScore: 0.0,
		},
		{
			name:     "有规则",
			rules:    5,
			minScore: 0.8,
			maxScore: 0.9,
		},
		{
			name:     "多规则",
			rules:    10,
			minScore: 0.8,
			maxScore: 0.9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rules []db.ContractRule
			for i := 0; i < tt.rules; i++ {
				rules = append(rules, db.ContractRule{
					FieldName:  fmt.Sprintf("field%d", i),
					FieldValue: "100",
					Confidence: 0.9,
				})
			}
			// 测试置信度计算
			score := calculateRuleConfidence(rules)
			if tt.rules == 0 {
				if score != 0 {
					t.Errorf("calculateRuleConfidence(nil) = %v, want 0", score)
				}
			} else {
				if score < tt.minScore || score > tt.maxScore {
					t.Errorf("calculateRuleConfidence() = %v, want between %v and %v",
						score, tt.minScore, tt.maxScore)
				}
			}
		})
	}
}

/**
 * 测试绝对值计算
 */
func TestAbs(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.0, 5.0},
		{-5.0, 5.0},
		{0.0, 0.0},
		{-0.001, 0.001},
		{999.99, 999.99},
	}

	for _, tt := range tests {
		result := abs(tt.input)
		if math.Abs(result-tt.expected) > 0.0001 {
			t.Errorf("abs(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

/**
 * 测试最大值计算
 */
func TestMax(t *testing.T) {
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{1.0, 2.0, 2.0},
		{2.0, 1.0, 2.0},
		{1.0, 1.0, 1.0},
		{-1.0, 1.0, 1.0},
		{0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		result := max(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("max(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

// ==================== 工具函数集成测试 ====================

/**
 * 测试 ReadExcel 工具 - 正常路径
 * 期望：能正确读取Excel文件
 */
func TestReadExcel_Success(t *testing.T) {
	// TODO: 需要准备测试用的Excel文件
	t.Skip("需要准备测试数据文件")

	ctx := context.Background()
	params := map[string]interface{}{
		"file_path":    "./test_data/test.xlsx",
		"sheet_index":  float64(0),
		"max_rows":     float64(1000),
	}

	result, err := ReadExcel(ctx, params)
	if err != nil {
		t.Fatalf("ReadExcel() error = %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("ReadExcel() result is not a map")
	}

	if resultMap["success"] != true {
		t.Errorf("ReadExcel() success = false, want true")
	}

	if resultMap["row_count"] == nil {
		t.Errorf("ReadExcel() row_count is nil")
	}
}

/**
 * 测试 ReadExcel 工具 - 文件路径为空
 * 期望：返回错误
 */
func TestReadExcel_EmptyPath(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"file_path": "",
	}

	_, err := ReadExcel(ctx, params)
	if err == nil {
		t.Error("ReadExcel() with empty file_path should return error")
	}
}

/**
 * 测试 ReadExcel 工具 - 文件不存在
 * 期望：返回错误
 */
func TestReadExcel_FileNotFound(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"file_path":    "/nonexistent/path/test.xlsx",
		"sheet_index":  float64(0),
		"max_rows":     float64(1000),
	}

	_, err := ReadExcel(ctx, params)
	if err == nil {
		t.Error("ReadExcel() with nonexistent file should return error")
	}
}

/**
 * 测试 MatchStoreAlias 工具 - 正常路径
 * 期望：能正确匹配门店别名
 */
func TestMatchStoreAlias_Success(t *testing.T) {
	// TODO: 需要准备测试数据
	t.Skip("需要准备测试数据")

	ctx := context.Background()
	params := map[string]interface{}{
		"alias": "测试门店",
	}

	result, err := MatchStoreAlias(ctx, params)
	if err != nil {
		t.Fatalf("MatchStoreAlias() error = %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("MatchStoreAlias() result is not a map")
	}

	if resultMap["success"] != true {
		t.Errorf("MatchStoreAlias() success = false, want true")
	}
}

/**
 * 测试 MatchStoreAlias 工具 - 别名为空
 * 期望：返回错误
 */
func TestMatchStoreAlias_EmptyAlias(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"alias": "",
	}

	_, err := MatchStoreAlias(ctx, params)
	if err == nil {
		t.Error("MatchStoreAlias() with empty alias should return error")
	}
}

/**
 * 测试 GenerateHuadingTemplate 工具 - 正常路径
 * 期望：能正确生成华鼎标准模板
 */
func TestGenerateHuadingTemplate_Success(t *testing.T) {
	// TODO: 需要准备测试数据
	t.Skip("需要准备测试数据")

	ctx := context.Background()
	params := map[string]interface{}{
		"order_data": map[string]interface{}{
			"order_no":      "TEST-001",
			"customer_name": "测试客户",
			"items":         []interface{}{},
		},
	}

	result, err := GenerateHuadingTemplate(ctx, params)
	if err != nil {
		t.Fatalf("GenerateHuadingTemplate() error = %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("GenerateHuadingTemplate() result is not a map")
	}

	if resultMap["success"] != true {
		t.Errorf("GenerateHuadingTemplate() success = false, want true")
	}

	if resultMap["file_path"] == nil {
		t.Errorf("GenerateHuadingTemplate() file_path is nil")
	}
}

/**
 * 测试 CalculateBill 工具 - 参数缺失
 * 期望：返回错误
 */
func TestCalculateBill_MissingParams(t *testing.T) {
	ctx := context.Background()

	tests := []map[string]interface{}{
		{"customer_id": "", "start_date": "2026-01-01", "end_date": "2026-01-31"},
		{"customer_id": "C001", "start_date": "", "end_date": "2026-01-31"},
		{"customer_id": "C001", "start_date": "2026-01-01", "end_date": ""},
		{"customer_id": "", "start_date": "", "end_date": ""},
	}

	for i, params := range tests {
		_, err := CalculateBill(ctx, params)
		if err == nil {
			t.Errorf("CalculateBill() test case %d should return error", i)
		}
	}
}

/**
 * 测试 AuditBill 工具 - 参数缺失
 * 期望：返回错误
 */
func TestAuditBill_MissingParams(t *testing.T) {
	ctx := context.Background()

	tests := []map[string]interface{}{
		{"bill_summary": nil, "contract_rules": map[string]interface{}{}},
		{"bill_summary": map[string]interface{}{}, "contract_rules": nil},
		{"bill_summary": nil, "contract_rules": nil},
	}

	for i, params := range tests {
		_, err := AuditBill(ctx, params)
		if err == nil {
			t.Errorf("AuditBill() test case %d should return error", i)
		}
	}
}

/**
 * 测试 ReadContractPDF 工具 - 文件路径为空
 * 期望：返回错误
 */
func TestReadContractPDF_EmptyPath(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"file_path": "",
	}

	_, err := ReadContractPDF(ctx, params)
	if err == nil {
		t.Error("ReadContractPDF() with empty file_path should return error")
	}
}

/**
 * 测试 ExtractContractRules 工具 - 合同文本为空
 * 期望：返回错误
 */
func TestExtractContractRules_EmptyText(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"contract_text": "",
		"customer_id":   "C001",
	}

	_, err := ExtractContractRules(ctx, params)
	if err == nil {
		t.Error("ExtractContractRules() with empty contract_text should return error")
	}
}

// ==================== 性能测试 ====================

/**
 * 性能测试：字符串相似度计算性能
 * 期望：在1秒内完成1000次计算
 */
func BenchmarkStringSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSimilarity("1000塑杯（600个装）", "1000塑料杯")
	}
}

/**
 * 性能测试：ReadExcel 工具性能
 * 期望：能在合理时间内完成
 */
func BenchmarkReadExcel(b *testing.B) {
	// TODO: 需要准备测试数据
	b.Skip("需要准备测试数据")
}

// ==================== 费用计算测试 ====================

/**
 * 测试费用计算 - 基础场景
 */
func TestCalculateFees(t *testing.T) {
	// 准备测试数据
	records := []db.OutboundRecord{
		{ID: "1", CustomerID: "TEST001", OrderNo: "ORD-001", StoreName: "王口镇店",
			SkuCode: "SKU001", Quantity: 100, Weight: 50, Region: "天津"},
		{ID: "2", CustomerID: "TEST001", OrderNo: "ORD-002", StoreName: "任丘店",
			SkuCode: "SKU002", Quantity: 200, Weight: 80, Region: "沧州"},
	}

	rules := []db.ContractRule{
		{RuleCategory: "仓储费", FieldName: "存储费单价", FieldValue: "3.2", FieldUnit: "元/托/天"},
		{RuleCategory: "运输费", FieldName: "最低收费", FieldValue: "100", FieldUnit: "元/次"},
		{RuleCategory: "出库费", FieldName: "出库单价", FieldValue: "1", FieldUnit: "元/件"},
	}

	result := calculateFees(records, rules)

	// 验证总额大于0
	if result.TotalAmount <= 0 {
		t.Errorf("TotalAmount should be positive, got %v", result.TotalAmount)
	}

	// 验证明细不为空
	if len(result.FeeDetails) == 0 {
		t.Error("FeeDetails should not be empty")
	}

	t.Logf("费用计算结果: 总计=%.2f, 仓储费=%.2f, 运输费=%.2f, 出库费=%.2f",
		result.TotalAmount, result.StorageFee, result.TransportFee, result.OutboundFee)
}

/**
 * 测试费用计算 - 无规则场景
 */
func TestCalculateFees_NoRules(t *testing.T) {
	records := []db.OutboundRecord{
		{ID: "1", CustomerID: "TEST001", StoreName: "测试店", Quantity: 100, Weight: 50},
	}

	result := calculateFees(records, []db.ContractRule{})

	// 没有规则时，费用应该为0
	if result.TotalAmount != 0 {
		t.Errorf("TotalAmount should be 0 when no rules, got %v", result.TotalAmount)
	}
}

/**
 * 测试费用计算 - 多区域场景
 */
func TestCalculateFees_MultiRegion(t *testing.T) {
	records := []db.OutboundRecord{
		{ID: "1", StoreName: "天津店1", Quantity: 100, Weight: 50, Region: "天津"},
		{ID: "2", StoreName: "天津店2", Quantity: 150, Weight: 60, Region: "天津"},
		{ID: "3", StoreName: "成都店", Quantity: 200, Weight: 80, Region: "成都"},
	}

	rules := []db.ContractRule{
		{RuleCategory: "仓储费", FieldName: "存储费单价", FieldValue: "3.0", FieldUnit: "元/托/天"},
		{RuleCategory: "运输费", FieldName: "最低收费", FieldValue: "50", FieldUnit: "元/次"},
		{RuleCategory: "出库费", FieldName: "出库单价", FieldValue: "0.5", FieldUnit: "元/件"},
	}

	result := calculateFees(records, rules)

	// 应该有两个区域的明细
	if len(result.FeeDetails) != 2 {
		t.Errorf("Should have 2 region details, got %d", len(result.FeeDetails))
	}

	// 天津区域应该累计重量 (50+60=110托)
	t.Logf("费用计算完成，区域数: %d", len(result.FeeDetails))
}

// ==================== 合同规则解析测试 ====================

/**
 * 测试合同规则解析 - 仓储费
 */
func TestParseContractRules_Storage(t *testing.T) {
	text := `
	仓储费：3.2元/托/天
	最低仓储费：50元/月
	`

	rules := parseContractRules(text)

	// 应该提取到仓储费规则
	var storageRules []db.ContractRule
	for _, r := range rules {
		if r.RuleCategory == "仓储费" {
			storageRules = append(storageRules, r)
		}
	}

	if len(storageRules) == 0 {
		t.Error("Should extract at least one storage rule")
	}

	// 验证解析的值
	for _, r := range storageRules {
		if r.FieldValue == "" {
			t.Errorf("Storage rule value should not be empty")
		}
		t.Logf("提取规则: %s - %s = %s %s (置信度: %.2f)",
			r.RuleCategory, r.FieldName, r.FieldValue, r.FieldUnit, r.Confidence)
	}
}

/**
 * 测试合同规则解析 - 运输费
 */
func TestParseContractRules_Transport(t *testing.T) {
	text := `
	超量单价：0.5元/件
	最低运费：100元
	`

	rules := parseContractRules(text)

	var transportRules []db.ContractRule
	for _, r := range rules {
		if r.RuleCategory == "运输费" {
			transportRules = append(transportRules, r)
		}
	}

	if len(transportRules) == 0 {
		t.Error("Should extract at least one transport rule")
	}
}

/**
 * 测试合同规则解析 - 出库费
 */
func TestParseContractRules_Outbound(t *testing.T) {
	text := `
	出库费：1元/件
	拣货费：0.5元/件
	`

	rules := parseContractRules(text)

	var outboundRules []db.ContractRule
	for _, r := range rules {
		if r.RuleCategory == "出库费" {
			outboundRules = append(outboundRules, r)
		}
	}

	if len(outboundRules) == 0 {
		t.Error("Should extract at least one outbound rule")
	}
}

/**
 * 测试合同规则解析 - 完整合同文本
 */
func TestParseContractRules_FullText(t *testing.T) {
	text := `
	合同编号：HT-2026-001

	一、仓储费
	存储费：3.2元/托/天
	最低收费：50元/月

	二、运输费
	超量单价：0.5元/件
	最低运费：100元/次
	每公里：2元

	三、出库费
	出库费：1元/件
	最低出库收费：30元/单

	四、其他费用
	管理费：200元/月
	服务费：100元/月
	`

	rules := parseContractRules(text)

	if len(rules) == 0 {
		t.Error("Should extract rules from full contract text")
	}

	// 分类统计
	categories := make(map[string]int)
	for _, r := range rules {
		categories[r.RuleCategory]++
	}

	t.Logf("提取规则统计:")
	for cat, count := range categories {
		t.Logf("  %s: %d条", cat, count)
	}
	t.Logf("总计提取 %d 条规则", len(rules))
}

/**
 * 测试合同规则解析 - 空文本
 */
func TestParseContractRules_Empty(t *testing.T) {
	rules := parseContractRules("")
	if len(rules) != 0 {
		t.Errorf("Empty text should return 0 rules, got %d", len(rules))
	}
}

/**
 * 测试合同规则解析 - 无效格式
 */
func TestParseContractRules_InvalidFormat(t *testing.T) {
	text := `
	这是一段普通文本，不包含计费规则。
	价格可能是3.2但是没有单位。
	或者只是一些数字12345。
	`

	rules := parseContractRules(text)
	// 可能提取不到任何规则，这是正常的
	t.Logf("无效格式文本提取到 %d 条规则", len(rules))
}

// ==================== 辅助函数测试 ====================

/**
 * 测试从账单重新计算
 */
func TestCalculateFeesFromBill(t *testing.T) {
	bill := map[string]interface{}{
		"storage_fee":   4800.0,
		"transport_fee": 1200.0,
		"outbound_fee":  800.0,
		"other_fee":     0.0,
	}

	result := calculateFeesFromBill(bill)

	total := result["total_amount"].(float64)
	if total != 6800.0 {
		t.Errorf("Expected total 6800.0, got %v", total)
	}

	// 验证明细包含计算方法
	details := result["calculation_details"].(map[string]interface{})
	if details["method"] != "ai_recalculation" {
		t.Error("Should contain calculation method")
	}

	t.Logf("AI重新计算结果: 总计=%.2f", total)
}
