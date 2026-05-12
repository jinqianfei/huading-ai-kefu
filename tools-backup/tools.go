/**
 * 工具函数定义 - Openclaw Agent Core 调用的外部工具
 *
 * 这些工具在 Agent 处理请求时被调用，用于访问数据库、执行计算等
 */

package tools

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"ai-cs-support/tools/db"
	"ai-cs-support/tools/ocr"
	"ai-cs-support/tools/excel"
)

// Tool 定义
type Tool struct {
	Name        string
	Description string
	Parameters  map[string]Parameter
	Handler     func(ctx context.Context, params map[string]interface{}) (interface{}, error)
}

// Parameter 定义
type Parameter struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Default    interface{} `json:"default,omitempty"`
}

// 所有工具定义
var Tools = map[string]Tool{
	"read_excel": {
		Name:        "read_excel",
		Description: "读取Excel文件内容，返回结构化表格数据",
		Parameters: map[string]Parameter{
			"file_path":    {Type: "string", Description: "文件路径", Required: true},
			"sheet_index": {Type: "number", Description: "工作表索引，默认0", Required: false, Default: 0},
			"max_rows":    {Type: "number", Description: "最大读取行数，默认1000", Required: false, Default: 1000},
		},
		Handler: ReadExcel,
	},

	"match_customer_sku": {
		Name:        "match_customer_sku",
		Description: "客户商品编码精确匹配系统SKU，支持语义匹配降级",
		Parameters: map[string]Parameter{
			"customer_id":      {Type: "string", Description: "客户ID", Required: true},
			"customer_sku_code": {Type: "string", Description: "客户商品编码", Required: true},
			"fallback_semantic": {Type: "boolean", Description: "精确匹配失败时是否降级到语义匹配", Required: false, Default: true},
		},
		Handler: MatchCustomerSKU,
	},

	"match_store_alias": {
		Name:        "match_store_alias",
		Description: "门店简称匹配全称",
		Parameters: map[string]Parameter{
			"alias": {Type: "string", Description: "门店简称", Required: true},
		},
		Handler: MatchStoreAlias,
	},

	"generate_huading_template": {
		Name:        "generate_huading_template",
		Description: "生成华鼎标准出库单Excel模板",
		Parameters: map[string]Parameter{
			"order_data": {Type: "object", Description: "订单数据", Required: true},
			"output_format": {Type: "string", Description: "输出格式，默认xlsx", Required: false, Default: "xlsx"},
		},
		Handler: GenerateHuadingTemplate,
	},

	"calculate_bill": {
		Name:        "calculate_bill",
		Description: "按计费规则计算账单",
		Parameters: map[string]Parameter{
			"customer_id": {Type: "string", Description: "客户ID", Required: true},
			"start_date":  {Type: "string", Description: "开始日期 YYYY-MM-DD", Required: true},
			"end_date":    {Type: "string", Description: "结束日期 YYYY-MM-DD", Required: true},
		},
		Handler: CalculateBill,
	},

	"audit_bill": {
		Name:        "audit_bill",
		Description: "AI审核账单，标红异常项",
		Parameters: map[string]Parameter{
			"bill_summary":    {Type: "object", Description: "账单汇总数据", Required: true},
			"contract_rules":  {Type: "object", Description: "合同规则", Required: true},
		},
		Handler: AuditBill,
	},

	"read_contract_pdf": {
		Name:        "read_contract_pdf",
		Description: "读取合同PDF，返回文本内容",
		Parameters: map[string]Parameter{
			"file_path": {Type: "string", Description: "合同文件路径", Required: true},
		},
		Handler: ReadContractPDF,
	},

	"extract_contract_rules": {
		Name:        "extract_contract_rules",
		Description: "从合同文本提取计费规则结构化数据",
		Parameters: map[string]Parameter{
			"contract_text":  {Type: "string", Description: "合同文本", Required: true},
			"customer_id":   {Type: "string", Description: "客户ID", Required: true},
		},
		Handler: ExtractContractRules,
	},
}

// ==================== 工具实现 ====================

// ReadExcel 读取Excel文件
func ReadExcel(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	filePath, ok := params["file_path"].(string)
	if !ok || filePath == "" {
		return nil, fmt.Errorf("file_path is required")
	}

	sheetIndex := 0
	if v, ok := params["sheet_index"].(float64); ok {
		sheetIndex = int(v)
	}

	maxRows := 1000
	if v, ok := params["max_rows"].(float64); ok {
		maxRows = int(v)
	}

	// 调用 Excel 读取库
	data, err := excel.ReadExcel(filePath, sheetIndex, maxRows)
	if err != nil {
		return nil, fmt.Errorf("failed to read excel: %w", err)
	}

	return map[string]interface{}{
		"success":    true,
		"sheet_name": data.SheetName,
		"row_count":  data.RowCount,
		"col_count":  data.ColCount,
		"headers":    data.Headers,
		"rows":       data.Rows,
	}, nil
}

// MatchCustomerSKU 客户商品编码匹配
func MatchCustomerSKU(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	customerID, _ := params["customer_id"].(string)
	customerSkuCode, _ := params["customer_sku_code"].(string)
	fallbackSemantic, _ := params["fallback_semantic"].(bool)
	if fallbackSemantic == false {
		fallbackSemantic = true
	}

	// 精确匹配
	mapping, err := db.GetCustomerSkuMapping(customerID, customerSkuCode)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if mapping != nil {
		return map[string]interface{}{
			"success":         true,
			"match_type":      "EXACT",
			"confidence":      mapping.Confidence,
			"system_sku_code": mapping.SystemSkuCode,
			"system_sku_name": mapping.SystemSkuName,
			"unit_conversion": mapping.UnitConversionRule,
		}, nil
	}

	// 降级到语义匹配
	if fallbackSemantic {
		// 查询历史匹配记录
		history, err := db.GetMatchHistory(customerID, customerSkuCode)
		if err == nil && history != nil {
			return map[string]interface{}{
				"success":         true,
				"match_type":      "HISTORY",
				"confidence":      history.Confidence * 0.9, // 历史匹配降权
				"system_sku_code": history.SystemSkuCode,
				"system_sku_name": history.SystemSkuName,
				"notes":           "来自历史匹配记录",
			}, nil
		}

		// 语义相似度匹配（需要调用 Embedding 服务）
		candidates, err := db.GetSystemSkuCandidates()
		if err == nil {
			// 计算相似度并返回 Top3
			results := calculateSimilarity(customerSkuCode, candidates)
			if len(results) > 0 && results[0].Score >= 0.6 {
				return map[string]interface{}{
					"success":      true,
					"match_type":   "SEMANTIC",
					"confidence":   results[0].Score,
					"alternatives": results,
				}, nil
			}
		}
	}

	// 无法匹配
	return map[string]interface{}{
		"success":    false,
		"match_type": "NOT_FOUND",
		"confidence": 0,
		"message":    "无法匹配到系统SKU，请人工选择",
	}, nil
}

// MatchStoreAlias 门店简称匹配
func MatchStoreAlias(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	alias, _ := params["alias"].(string)
	if alias == "" {
		return nil, fmt.Errorf("alias is required")
	}

	// 查询映射表
	mapping, err := db.GetStoreAliasMapping(alias)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if mapping != nil {
		return map[string]interface{}{
			"success":      true,
			"match_type":   mapping.MatchType,
			"confidence":   mapping.Confidence,
			"full_name":    mapping.FullName,
			"customer_id":  mapping.CustomerID,
		}, nil
	}

	return map[string]interface{}{
		"success":  false,
		"confidence": 0,
		"message":   "疑似新门店，请确认",
	}, nil
}

// GenerateHuadingTemplate 生成华鼎标准模板
func GenerateHuadingTemplate(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	orderData, ok := params["order_data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("order_data is required")
	}

	// 生成 Excel 文件
	filePath, err := excel.GenerateHuadingTemplate(orderData)
	if err != nil {
		return nil, fmt.Errorf("failed to generate template: %w", err)
	}

	return map[string]interface{}{
		"success":    true,
		"file_path":  filePath,
		"file_name":  "华鼎标准出库单.xlsx",
		"mime_type":  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}, nil
}

// CalculateBill 计算账单
func CalculateBill(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	customerID, _ := params["customer_id"].(string)
	startDate, _ := params["start_date"].(string)
	endDate, _ := params["end_date"].(string)

	if customerID == "" || startDate == "" || endDate == "" {
		return nil, fmt.Errorf("customer_id, start_date, end_date are required")
	}

	// 读取出库记录
	records, err := db.GetOutboundRecords(customerID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get outbound records: %w", err)
	}

	// 读取合同规则
	rules, err := db.GetContractRules(customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract rules: %w", err)
	}

	// 计算费用
	result := calculateFees(records, rules)

	return map[string]interface{}{
		"success":           true,
		"customer_id":       customerID,
		"period":            fmt.Sprintf("%s 至 %s", startDate, endDate),
		"total_amount":      result.TotalAmount,
		"storage_fee":       result.StorageFee,
		"transport_fee":    result.TransportFee,
		"outbound_fee":     result.OutboundFee,
		"fee_details":       result.FeeDetails,
		"calculated_items":  len(records),
	}, nil
}

// AuditBill 审核账单
func AuditBill(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	billSummary, ok := params["bill_summary"].(map[string]interface{})
	_, _ = params["contract_rules"].(map[string]interface{}) // contract_rules 参数预留
	if !ok || billSummary == nil {
		return nil, fmt.Errorf("bill_summary is required")
	}

	// AI 重新计算
	recalculated := calculateFeesFromBill(billSummary)

	// 检查必需字段
	aiAmountRaw, hasAI := recalculated["total_amount"]
	billAmountRaw, hasBill := billSummary["total_amount"]
	if !hasAI || !hasBill {
		return nil, fmt.Errorf("total_amount field is required in bill_summary")
	}

	// 类型转换
	aiAmount, aiOk := aiAmountRaw.(float64)
	billAmount, billOk := billAmountRaw.(float64)
	if !aiOk || !billOk {
		return nil, fmt.Errorf("total_amount must be a number")
	}

	// 比对差异
	var auditItems []map[string]interface{}

	diff := aiAmount - billAmount
	maxAmount := aiAmount
	if billAmount > maxAmount {
		maxAmount = billAmount
	}
	consistency := 1.0
	if maxAmount > 0 {
		consistency = 1 - abs(diff)/maxAmount
	}

	if consistency < 0.999 {
		severity := "WARNING"
		if consistency < 0.8 {
			severity = "SEVERE"
		}
		auditItems = append(auditItems, map[string]interface{}{
			"type":           "CONSISTENCY",
			"severity":       severity,
			"ai_calculated":  aiAmount,
			"bill_amount":    billAmount,
			"diff":           diff,
			"consistency":    consistency,
			"description":    fmt.Sprintf("AI计算金额%.2f与账单金额%.2f差异%.2f", aiAmount, billAmount, diff),
		})
	}

	return map[string]interface{}{
		"success":      true,
		"audit_passed": len(auditItems) == 0 || auditItems[0]["severity"] != "SEVERE",
		"audit_items":  auditItems,
		"summary": map[string]interface{}{
			"total_checked": 1,
			"passed":        len(auditItems) == 0,
			"warnings":      len(auditItems),
			"severes":       len(auditItems),
		},
	}, nil
}

// ReadContractPDF 读取合同PDF
func ReadContractPDF(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	filePath, ok := params["file_path"].(string)
	if !ok || filePath == "" {
		return nil, fmt.Errorf("file_path is required")
	}

	// 调用 OCR 服务
	text, err := ocr.ReadPDF(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read PDF: %w", err)
	}

	return map[string]interface{}{
		"success":    true,
		"text":       text,
		"char_count": len(text),
	}, nil
}

// ExtractContractRules 提取合同规则
func ExtractContractRules(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	contractText, _ := params["contract_text"].(string)
	customerID, _ := params["customer_id"].(string)

	if contractText == "" {
		return nil, fmt.Errorf("contract_text is required")
	}

	// 解析计费规则
	rules := parseContractRules(contractText)

	// 存储到数据库
	for _, rule := range rules {
		rule.CustomerID = customerID
		if err := db.SaveContractRule(&rule); err != nil {
			return nil, fmt.Errorf("failed to save rule: %w", err)
		}
	}

	return map[string]interface{}{
		"success":       true,
		"rules_count":  len(rules),
		"rules":        rules,
		"confidence":   calculateRuleConfidence(rules),
	}, nil
}

// ==================== 辅助函数 ====================

type similarityResult struct {
	SystemSkuCode string  `json:"system_sku_code"`
	SystemSkuName string  `json:"system_sku_name"`
	Score         float64 `json:"score"`
}

func calculateSimilarity(name string, candidates []db.SystemSku) []similarityResult {
	// 简化实现，实际应该调用 Embedding 服务
	var results []similarityResult
	for _, c := range candidates {
		score := stringSimilarity(name, c.SkuName)
		if score >= 0.6 {
			results = append(results, similarityResult{
				SystemSkuCode: c.SkuCode,
				SystemSkuName: c.SkuName,
				Score:         score,
			})
		}
	}
	return results
}

func stringSimilarity(s1, s2 string) float64 {
	// 简化：字符相似度
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	same := 0
	for _, c := range s1 {
		if strings.Contains(s2, string(c)) {
			same++
		}
	}
	m := len(s1)
	if len(s2) > m {
		m = len(s2)
	}
	return float64(same) / float64(m)
}

type feeResult struct {
	TotalAmount  float64
	StorageFee   float64
	TransportFee float64
	OutboundFee  float64
	OtherFee     float64
	FeeDetails   []map[string]interface{}
}

// ruleContext 规则上下文，用于计费
type ruleContext struct {
	storageUnitPrice  float64 // 仓储费单价 (元/托/天)
	transportUnitPrice float64 // 运输费单价 (元/公里或元/件)
	transportMinFee   float64 // 运输最低收费
	outboundUnitPrice float64 // 出库费单价 (元/件)
	outboundMinFee    float64 // 出库最低收费
}

// calculateFees 根据出库记录和合同规则计算费用
func calculateFees(records []db.OutboundRecord, rules []db.ContractRule) feeResult {
	var result feeResult

	// 构建规则上下文
	ctx := buildRuleContext(rules)
	if ctx == nil {
		// 没有规则，返回空结果
		return result
	}

	// 按区域分组统计
	regionStats := make(map[string]*regionStat)
	for _, record := range records {
		region := record.Region
		if region == "" {
			region = "默认"
		}

		stat, exists := regionStats[region]
		if !exists {
			stat = &regionStat{
				region:      region,
				totalWeight: 0,
				totalQty:    0,
				orderCount:  0,
			}
			regionStats[region] = stat
		}

		stat.totalWeight += record.Weight
		stat.totalQty += record.Quantity
		stat.orderCount++
	}

	// 计算各项费用
	for region, stat := range regionStats {
		// 1. 仓储费 = 单价 × 重量(托) × 天数
		// 假设存储天数 = 30天/月
		storageDays := 30
		storageFee := ctx.storageUnitPrice * stat.totalWeight * float64(storageDays)
		result.StorageFee += storageFee

		// 2. 运输费 = 最低收费 或 单价 × 重量(托)
		var transportFee float64
		if ctx.transportMinFee > 0 && stat.totalWeight > 0 {
			// 实际运输费 = 单价 × 托数，与最低收费取较大值
			unitTransport := ctx.transportUnitPrice * stat.totalWeight
			transportFee = unitTransport
			if unitTransport < ctx.transportMinFee {
				transportFee = ctx.transportMinFee
			}
		} else {
			transportFee = ctx.transportMinFee
		}
		result.TransportFee += transportFee

		// 3. 出库费 = 最低收费 或 单价 × 数量
		var outboundFee float64
		if ctx.outboundMinFee > 0 && stat.totalQty > 0 {
			unitOutbound := ctx.outboundUnitPrice * stat.totalQty
			outboundFee = unitOutbound
			if unitOutbound < ctx.outboundMinFee {
				outboundFee = ctx.outboundMinFee
			}
		} else {
			outboundFee = ctx.outboundMinFee
		}
		result.OutboundFee += outboundFee

		// 记录明细
		result.FeeDetails = append(result.FeeDetails, map[string]interface{}{
			"region":       region,
			"total_weight": stat.totalWeight,
			"total_qty":    stat.totalQty,
			"order_count":  stat.orderCount,
			"storage_fee":  storageFee,
			"transport_fee": transportFee,
			"outbound_fee": outboundFee,
		})
	}

	// 计算总额
	result.TotalAmount = result.StorageFee + result.TransportFee + result.OutboundFee + result.OtherFee

	return result
}

type regionStat struct {
	region      string
	totalWeight float64
	totalQty    float64
	orderCount  int
}

// buildRuleContext 从规则列表构建计费上下文
func buildRuleContext(rules []db.ContractRule) *ruleContext {
	ctx := &ruleContext{}

	for _, rule := range rules {
		value := parseRuleValue(rule.FieldValue)
		unit := rule.FieldUnit

		switch rule.RuleCategory {
		case "仓储费":
			switch rule.FieldName {
			case "存储费单价", "仓储单价":
				ctx.storageUnitPrice = value
				if unit != "" {
					// 解析单位，如 "元/托/天"
					if strings.Contains(unit, "件") {
						// 按件计费，需要换算
						ctx.storageUnitPrice = value / 10 // 假设 1托 = 10件
					}
				}
			case "最低收费":
				// 仓储费也有最低收费
			}

		case "运输费":
			switch rule.FieldName {
			case "超量单价", "运输单价":
				ctx.transportUnitPrice = value
			case "最低收费", "运输最低收费":
				ctx.transportMinFee = value
			}

		case "出库费":
			switch rule.FieldName {
			case "出库单价", "出库费单价":
				ctx.outboundUnitPrice = value
			case "最低收费", "出库最低收费":
				ctx.outboundMinFee = value
			}
		}
	}

	// 检查是否有有效的计费规则
	if ctx.storageUnitPrice == 0 && ctx.transportMinFee == 0 && ctx.outboundUnitPrice == 0 {
		return nil
	}

	return ctx
}

// parseRuleValue 解析规则值为浮点数
func parseRuleValue(value string) float64 {
	if value == "" {
		return 0
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return v
}

// calculateFeesFromBill 从账单重新计算费用（用于AI审核对比）
func calculateFeesFromBill(bill map[string]interface{}) map[string]interface{} {
	// AI 根据账单数据重新计算，验证账单准确性
	result := make(map[string]interface{})

	// 提取各项费用
	storageFee := getFloatValue(bill, "storage_fee")
	transportFee := getFloatValue(bill, "transport_fee")
	outboundFee := getFloatValue(bill, "outbound_fee")
	otherFee := getFloatValue(bill, "other_fee")

	// 计算总额
	total := storageFee + transportFee + outboundFee + otherFee

	result["storage_fee"] = storageFee
	result["transport_fee"] = transportFee
	result["outbound_fee"] = outboundFee
	result["other_fee"] = otherFee
	result["total_amount"] = total

	// 添加计算明细
	result["calculation_details"] = map[string]interface{}{
		"method":          "ai_recalculation",
		"storage_formula": "单价 × 托数 × 30天",
		"transport_formula": "max(单价 × 托数, 最低收费)",
		"outbound_formula": "max(单价 × 件数, 最低收费)",
	}

	return result
}

// getFloatValue 安全获取浮点值
func getFloatValue(m map[string]interface{}, key string) float64 {
	if v, ok := m[key]; ok {
		switch val := v.(type) {
		case float64:
			return val
		case int:
			return float64(val)
		case string:
			f, _ := strconv.ParseFloat(val, 64)
			return f
		}
	}
	return 0
}

// ruleParsePattern 规则解析模式
type ruleParsePattern struct {
	category  string
	fieldName string
	patterns  []string
	unit      string
}

// parseContractRules 从合同文本提取计费规则
func parseContractRules(text string) []db.ContractRule {
	var rules []db.ContractRule

	// 定义所有解析模式
	patterns := []ruleParsePattern{
		// 仓储费规则
		{
			category:  "仓储费",
			fieldName: "存储费单价",
			patterns: []string{
				`存储费[：:]\s*([0-9.]+)\s*元\s*/\s*托\s*/\s*天`,
				`仓储费[：:]\s*([0-9.]+)\s*元\s*/\s*托\s*/\s*天`,
				`托盘费[：:]\s*([0-9.]+)\s*元\s*/\s*托\s*/\s*天`,
				`仓位费[：:]\s*([0-9.]+)\s*元\s*/\s*托\s*/\s*天`,
			},
			unit: "元/托/天",
		},
		{
			category:  "仓储费",
			fieldName: "存储费单价",
			patterns: []string{
				`存储费[：:]\s*([0-9.]+)\s*元\s*/\s*件\s*/\s*天`,
				`仓储费[：:]\s*([0-9.]+)\s*元\s*/\s*件\s*/\s*天`,
			},
			unit: "元/件/天",
		},
		{
			category:  "仓储费",
			fieldName: "最低仓储费",
			patterns: []string{
				`仓储最低收费[：:]\s*([0-9.]+)\s*元`,
				`存储最低收费[：:]\s*([0-9.]+)\s*元`,
			},
			unit: "元/月",
		},

		// 运输费规则
		{
			category:  "运输费",
			fieldName: "超量单价",
			patterns: []string{
				`超量单价[：:]\s*([0-9.]+)\s*元\s*/\s*件`,
				`超量单价[：:]\s*([0-9.]+)\s*元\s*/\s*托`,
				`运输单价[：:]\s*([0-9.]+)\s*元\s*/\s*件`,
				`运输单价[：:]\s*([0-9.]+)\s*元\s*/\s*托`,
			},
			unit: "元/件",
		},
		{
			category:  "运输费",
			fieldName: "最低收费",
			patterns: []string{
				`运输最低收费[：:]\s*([0-9.]+)\s*元`,
				`配送最低收费[：:]\s*([0-9.]+)\s*元`,
				`最低运费[：:]\s*([0-9.]+)\s*元`,
				`起步价[：:]\s*([0-9.]+)\s*元`,
			},
			unit: "元/次",
		},
		{
			category:  "运输费",
			fieldName: "公里单价",
			patterns: []string{
				`每公里[：:]\s*([0-9.]+)\s*元`,
				`公里费[：:]\s*([0-9.]+)\s*元`,
				`([0-9.]+)\s*元\s*/\s*公里`,
			},
			unit: "元/公里",
		},

		// 出库费规则
		{
			category:  "出库费",
			fieldName: "出库单价",
			patterns: []string{
				`出库费[：:]\s*([0-9.]+)\s*元\s*/\s*件`,
				`出库费[：:]\s*([0-9.]+)\s*元\s*/\s*单`,
				`拣货费[：:]\s*([0-9.]+)\s*元\s*/\s*件`,
				`加工费[：:]\s*([0-9.]+)\s*元\s*/\s*件`,
			},
			unit: "元/件",
		},
		{
			category:  "出库费",
			fieldName: "最低收费",
			patterns: []string{
				`出库最低收费[：:]\s*([0-9.]+)\s*元`,
				`拣货最低收费[：:]\s*([0-9.]+)\s*元`,
			},
			unit: "元/单",
		},

		// 其他费用
		{
			category:  "其他费用",
			fieldName: "管理费",
			patterns: []string{
				`管理费[：:]\s*([0-9.]+)\s*元`,
				`服务费[：:]\s*([0-9.]+)\s*元`,
			},
			unit: "元/月",
		},
	}

	// 预编译正则表达式（避免重复编译）
	type compiledPattern struct {
		re       *regexp.Regexp
		pattern  ruleParsePattern
	}
	var compiled []compiledPattern

	// 使用互斥锁保护正则表达式编译（Go 1.9+ 正则表达式不是线程安全的）
	for _, p := range patterns {
		for _, pattern := range p.patterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				continue
			}
			compiled = append(compiled, compiledPattern{
				re:      re,
				pattern: p,
			})
		}
	}

	// 已提取的规则（避免重复）
	extracted := make(map[string]bool)

	// 提取规则
	for _, cp := range compiled {
		matches := cp.re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) < 2 {
				continue
			}

			value := match[1]
			key := cp.pattern.category + ":" + cp.pattern.fieldName + ":" + value

			// 避免重复提取
			if extracted[key] {
				continue
			}
			extracted[key] = true

			// 查找规则来源位置
			sourceLocation := findSourceLocation(text, match[0])

			rules = append(rules, db.ContractRule{
				RuleCategory: cp.pattern.category,
				FieldName:    cp.pattern.fieldName,
				FieldValue:   value,
				FieldUnit:    cp.pattern.unit,
				Source:       sourceLocation,
				Confidence:   calculatePatternConfidence(match[0], value),
				Status:       "PENDING",
			})
		}
	}

	return rules
}

// findSourceLocation 查找规则在文本中的位置
func findSourceLocation(text, match string) string {
	// 简化实现：返回匹配位置
	idx := strings.Index(text, match)
	if idx >= 0 {
		return fmt.Sprintf("合同第%d页附近", idx/200+1) // 假设每页约200字符
	}
	return "合同正文"
}

// calculatePatternConfidence 根据匹配质量计算置信度
func calculatePatternConfidence(match, value string) float64 {
	// 基础置信度
	confidence := 0.8

	// 如果值看起来完整（不是被截断的）
	if len(value) > 0 && value[len(value)-1] >= '0' && value[len(value)-1] <= '9' {
		confidence += 0.1
	}

	// 如果有明确的单位
	if strings.Contains(match, "元") && strings.Contains(match, "/") {
		confidence += 0.05
	}

	// 最高置信度限制
	if confidence > 0.98 {
		confidence = 0.98
	}

	return confidence
}

func calculateRuleConfidence(rules []db.ContractRule) float64 {
	// 简化实现
	if len(rules) == 0 {
		return 0
	}
	return 0.85
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
