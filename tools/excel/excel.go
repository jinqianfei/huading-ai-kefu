/**
 * Excel 工具模块
 * 负责 Excel 文件的读取和生成
 */

package excel

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExcelData Excel数据结构
type ExcelData struct {
	SheetName string         `json:"sheet_name"`
	RowCount  int            `json:"row_count"`
	ColCount  int            `json:"col_count"`
	Headers   []string       `json:"headers"`
	Rows      [][]string     `json:"rows"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// ReadExcel 读取Excel文件
func ReadExcel(filePath string, sheetIndex int, maxRows int) (*ExcelData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// 获取工作表名称
	sheets := f.GetSheetList()
	if sheetIndex >= len(sheets) {
		return nil, fmt.Errorf("sheet index out of range")
	}
	sheetName := sheets[sheetIndex]

	// 读取行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	// 限制行数
	if len(rows) > maxRows {
		rows = rows[:maxRows]
	}

	if len(rows) == 0 {
		return &ExcelData{
			SheetName: sheetName,
			RowCount:  0,
			ColCount:  0,
			Headers:   []string{},
			Rows:      [][]string{},
		}, nil
	}

	// 提取表头和数据行
	headers := rows[0]
	dataRows := rows[1:]

	return &ExcelData{
		SheetName: sheetName,
		RowCount:  len(dataRows),
		ColCount:  len(headers),
		Headers:   headers,
		Rows:      dataRows,
		Metadata: map[string]any{
			"file_path":  filePath,
			"total_rows": len(rows),
			"read_at":    time.Now().Format(time.RFC3339),
		},
	}, nil
}

// ==================== 华鼎模板生成相关 ====================

// HuadingTemplateConfig 华鼎模板生成配置
type HuadingTemplateConfig struct {
	// 输出目录，默认为/tmp
	OutputDir string
	// 模板文件路径（可选，如果提供则基于模板填充）
	TemplateFilePath string
	// 客户ID，用于生成三方单号
	CustomerID string
	// 订单号，用于生成三方单号
	OrderNo string
}

// HuadingField 华鼎模板字段定义
type HuadingField struct {
	Name         string      // 字段名（中文）
	ColumnIndex  int         // 列索引（从0开始）
	ColumnLetter string      // 列字母（A, B, C...）
	Required     bool        // 是否必填
	DefaultValue interface{} // 默认值
	DataType     string      // 数据类型：string, number, decimal, date
}

// HuadingRowData 华鼎模板单行数据
type HuadingRowData struct {
	// 序号 (1)
	SeqNo int
	// 门店编号 (2)
	StoreCode string
	// 门店三方编码 (3)
	StoreThirdPartyCode string
	// 仓库编码 (4)
	WarehouseCode string
	// 加急程度 (5) - 默认0
	UrgencyLevel string
	// 商品编码 (6)
	ProductCode string
	// 商品名称 (7)
	ProductName string
	// 规格 (8)
	Spec string
	// 数量 (9)
	Quantity int
	// 单位 (10)
	Unit string
	// 商品SKU编号 (11)
	SkuCode string
	// 商品三方SPEC编号 (12)
	SpecCode string
	// 单位类型 (13)
	UnitType string
	// 指定库存状态 (14) - 默认"正常"
	InventoryStatus string
	// 出库类型 (15) - 默认201
	OutboundType string
	// 配送方式 (16)
	DeliveryMethod string
	// 指定车型 (17)
	VehicleType string
	// 是否垫付 (18) - 默认"否"
	IsAdvancePayment string
	// 付款方式 (19) - 默认"预付"
	PaymentMethod string
	// 快递公司 (20)
	LogisticsCompany string
	// 单价 (21)
	UnitPrice float64
	// 总金额 (22)
	TotalAmount float64
	// 是否制定批次 (23) - 默认"否"
	IsBatchSpecified string
	// 批次号 (24)
	BatchNo string
	// 生产日期 (25)
	ProductionDate string
	// 生产厂家编号 (26)
	ManufacturerCode string
	// 门店收货地址编码 (27)
	StoreAddressCode string
	// 三方单号 (28) - 自动生成
	ThirdPartyOrderNo string
	// 业务模式 (29) - 默认"B2C"
	BusinessModel string
	// 业务类型 (30)
	BusinessType string
	// 收货人 (31)
	ReceiverName string
	// 联系电话 (32)
	ReceiverPhone string
	// 收货地址 (33)
	ReceiverAddress string
	// C端快递公司 (34)
	CLogisticsCompany string
	// 备注 (35)
	Notes string
}

// PendingField 待确认字段信息
type PendingField struct {
	FieldName string
	Column    string
	Row       int
	Reason    string
}

// ValidationError 验证错误
type ValidationError struct {
	Row       int
	Column    string
	FieldName string
	Message   string
}

// GenerateResult 生成结果
type GenerateResult struct {
	FilePath       string
	RowCount       int
	PendingFields  []PendingField
	ValidationErrors []ValidationError
	Warnings       []string
}

// 华鼎模板31字段定义（实际为35列，包含商品编码、商品名称、规格、单位等扩展字段）
func getHuadingFieldDefinitions() []HuadingField {
	return []HuadingField{
		{Name: "序号", ColumnIndex: 0, ColumnLetter: "A", Required: true, DataType: "number"},
		{Name: "门店编号", ColumnIndex: 1, ColumnLetter: "B", Required: true, DataType: "string"},
		{Name: "门店三方编码", ColumnIndex: 2, ColumnLetter: "C", Required: false, DataType: "string"},
		{Name: "仓库编码", ColumnIndex: 3, ColumnLetter: "D", Required: true, DataType: "string"},
		{Name: "加急程度", ColumnIndex: 4, ColumnLetter: "E", Required: false, DefaultValue: "0", DataType: "string"},
		{Name: "商品编码", ColumnIndex: 5, ColumnLetter: "F", Required: false, DataType: "string"},
		{Name: "商品名称", ColumnIndex: 6, ColumnLetter: "G", Required: false, DataType: "string"},
		{Name: "规格", ColumnIndex: 7, ColumnLetter: "H", Required: false, DataType: "string"},
		{Name: "数量", ColumnIndex: 8, ColumnLetter: "I", Required: true, DataType: "number"},
		{Name: "单位", ColumnIndex: 9, ColumnLetter: "J", Required: false, DataType: "string"},
		{Name: "商品SKU编号", ColumnIndex: 10, ColumnLetter: "K", Required: false, DataType: "string"},
		{Name: "商品三方SPEC编号", ColumnIndex: 11, ColumnLetter: "L", Required: true, DataType: "string"},
		{Name: "单位类型", ColumnIndex: 12, ColumnLetter: "M", Required: true, DataType: "string"},
		{Name: "指定库存状态", ColumnIndex: 13, ColumnLetter: "N", Required: false, DefaultValue: "正常", DataType: "string"},
		{Name: "出库类型", ColumnIndex: 14, ColumnLetter: "O", Required: false, DefaultValue: "201", DataType: "string"},
		{Name: "配送方式", ColumnIndex: 15, ColumnLetter: "P", Required: true, DataType: "string"},
		{Name: "指定车型", ColumnIndex: 16, ColumnLetter: "Q", Required: false, DataType: "string"},
		{Name: "是否垫付", ColumnIndex: 17, ColumnLetter: "R", Required: false, DefaultValue: "否", DataType: "string"},
		{Name: "付款方式", ColumnIndex: 18, ColumnLetter: "S", Required: false, DefaultValue: "预付", DataType: "string"},
		{Name: "快递公司", ColumnIndex: 19, ColumnLetter: "T", Required: false, DataType: "string"},
		{Name: "单价", ColumnIndex: 20, ColumnLetter: "U", Required: false, DataType: "decimal"},
		{Name: "总金额", ColumnIndex: 21, ColumnLetter: "V", Required: false, DataType: "decimal"},
		{Name: "是否制定批次", ColumnIndex: 22, ColumnLetter: "W", Required: false, DefaultValue: "否", DataType: "string"},
		{Name: "批次号", ColumnIndex: 23, ColumnLetter: "X", Required: false, DataType: "string"},
		{Name: "生产日期", ColumnIndex: 24, ColumnLetter: "Y", Required: false, DataType: "date"},
		{Name: "生产厂家编号", ColumnIndex: 25, ColumnLetter: "Z", Required: false, DataType: "string"},
		{Name: "门店收货地址编码", ColumnIndex: 26, ColumnLetter: "AA", Required: false, DataType: "string"},
		{Name: "三方单号", ColumnIndex: 27, ColumnLetter: "AB", Required: false, DataType: "string"},
		{Name: "业务模式", ColumnIndex: 28, ColumnLetter: "AC", Required: false, DefaultValue: "B2C", DataType: "string"},
		{Name: "业务类型", ColumnIndex: 29, ColumnLetter: "AD", Required: true, DataType: "string"},
		{Name: "收货人", ColumnIndex: 30, ColumnLetter: "AE", Required: true, DataType: "string"},
		{Name: "联系电话", ColumnIndex: 31, ColumnLetter: "AF", Required: true, DataType: "string"},
		{Name: "收货地址", ColumnIndex: 32, ColumnLetter: "AG", Required: true, DataType: "string"},
		{Name: "C端快递公司", ColumnIndex: 33, ColumnLetter: "AH", Required: false, DataType: "string"},
		{Name: "备注", ColumnIndex: 34, ColumnLetter: "AI", Required: false, DataType: "string"},
	}
}

// 获取有默认值的字段列表（共8个）
func getDefaultValueFields() map[string]interface{} {
	return map[string]interface{}{
		"加急程度":       "0",
		"指定库存状态":   "正常",
		"出库类型":       "201",
		"是否垫付":       "否",
		"付款方式":       "预付",
		"是否制定批次":   "否",
		"业务模式":       "B2C",
		"仓库编码":       "5", // 默认杭州仓
	}
}

// GenerateThirdPartyOrderNo 生成三方单号
// 格式：CUS_{customer_id}_{order_no}_{timestamp}
func GenerateThirdPartyOrderNo(customerID, orderNo string) string {
	ts := time.Now().UnixMilli()
	// 清理特殊字符
	cleanCID := strings.ReplaceAll(customerID, " ", "_")
	cleanOrder := strings.ReplaceAll(orderNo, " ", "_")
	return fmt.Sprintf("CUS_%s_%s_%d", cleanCID, cleanOrder, ts)
}

// GenerateHuadingTemplate 生成华鼎标准出库单模板
// 支持完整的31列华鼎标准模板
// 支持pending字段标记（黄色高亮+批注）
// 支持默认值填充（8个字段有默认值）
// 支持三方单号生成格式：CUS_{cid}_{order}_{ts}
// 移除硬编码路径，使用配置或参数传入
// 添加字段验证和错误处理
func GenerateHuadingTemplate(orderData map[string]interface{}, config *HuadingTemplateConfig) (*GenerateResult, error) {
	if config == nil {
		config = &HuadingTemplateConfig{
			OutputDir: "/tmp",
		}
	}

	// 验证必要参数
	if config.CustomerID == "" {
		return nil, fmt.Errorf("customer_id is required for generating third-party order number")
	}

	// 如果提供了模板文件路径，尝试基于模板填充
	if config.TemplateFilePath != "" {
		return generateFromTemplate(orderData, config)
	}

	// 否则从零创建
	return generateFromScratch(orderData, config)
}

// generateFromScratch 从零创建华鼎标准出库单
func generateFromScratch(orderData map[string]interface{}, config *HuadingTemplateConfig) (*GenerateResult, error) {
	result := &GenerateResult{
		PendingFields:    []PendingField{},
		ValidationErrors: []ValidationError{},
		Warnings:         []string{},
	}

	f := excelize.NewFile()
	defer f.Close()

	sheetName := "华鼎标准出库单"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to create sheet: %w", err)
	}
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	// 设置列宽
	setColumnWidths(f, sheetName)

	// 创建样式
	headerStyle, err := createHeaderStyle(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create header style: %w", err)
	}

	dataStyle, err := createDataStyle(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create data style: %w", err)
	}

	pendingStyle, err := createPendingStyle(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create pending style: %w", err)
	}

	// 写入表头
	fields := getHuadingFieldDefinitions()
	row := 1
	for _, field := range fields {
		cell := fmt.Sprintf("%s%d", field.ColumnLetter, row)
		f.SetCellValue(sheetName, cell, field.Name)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
	}
	row++

	// 提取订单数据
	rows := extractOrderRows(orderData)
	if len(rows) == 0 {
		result.Warnings = append(result.Warnings, "no order items found, generating empty template with default values")
		// 至少生成一行带默认值的示例
		rows = []map[string]interface{}{{"_is_example": true}}
	}

	// 获取订单号用于生成三方单号
	orderNo := getStringValue(orderData, "order_no")
	if orderNo == "" {
		orderNo = getStringValue(orderData, "order_number")
	}
	if orderNo == "" {
		orderNo = "UNKNOWN"
	}

	// 生成三方单号（所有行共用）
	thirdPartyOrderNo := GenerateThirdPartyOrderNo(config.CustomerID, orderNo)

	// 处理每一行数据
	for i, item := range rows {
		seqNo := i + 1
		dataRow := buildHuadingRowData(item, orderData, seqNo, thirdPartyOrderNo, config.CustomerID)

		// 验证数据
		validationErrors := validateRowData(dataRow, row)
		result.ValidationErrors = append(result.ValidationErrors, validationErrors...)

		// 写入数据并标记pending字段
		pendingFields := writeRowData(f, sheetName, row, dataRow, fields, dataStyle, pendingStyle)
		result.PendingFields = append(result.PendingFields, pendingFields...)

		row++
	}

	result.RowCount = len(rows)

	// 保存文件
	outputDir := config.OutputDir
	if outputDir == "" {
		outputDir = "/tmp"
	}
	fileName := fmt.Sprintf("%s/华鼎出库单_%s_%s.xlsx", outputDir, config.CustomerID, time.Now().Format("20060102150405"))
	if err := f.SaveAs(fileName); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	result.FilePath = fileName
	return result, nil
}

// generateFromTemplate 基于现有模板填充
func generateFromTemplate(orderData map[string]interface{}, config *HuadingTemplateConfig) (*GenerateResult, error) {
	result := &GenerateResult{
		PendingFields:    []PendingField{},
		ValidationErrors: []ValidationError{},
		Warnings:         []string{},
	}

	f, err := excelize.OpenFile(config.TemplateFilePath)
	if err != nil {
		// 模板不存在时降级为从零创建
		result.Warnings = append(result.Warnings, fmt.Sprintf("template not found at %s, falling back to scratch generation", config.TemplateFilePath))
		return generateFromScratch(orderData, config)
	}
	defer f.Close()

	// 获取第一个工作表
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in template")
	}
	sheetName := sheets[0]

	// 查找表头行
	headerRow := findHeaderRow(f, sheetName)
	if headerRow == 0 {
		result.Warnings = append(result.Warnings, "header row not found, assuming row 1")
		headerRow = 1
	}

	// 创建样式
	pendingStyle, err := createPendingStyle(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create pending style: %w", err)
	}

	// 提取订单数据
	rows := extractOrderRows(orderData)
	if len(rows) == 0 {
		result.Warnings = append(result.Warnings, "no order items found")
		return result, nil
	}

	// 获取订单号
	orderNo := getStringValue(orderData, "order_no")
	if orderNo == "" {
		orderNo = "UNKNOWN"
	}
	thirdPartyOrderNo := GenerateThirdPartyOrderNo(config.CustomerID, orderNo)

	// 从表头行下一行开始填充数据
	dataStartRow := headerRow + 1
	for i, item := range rows {
		currentRow := dataStartRow + i
		seqNo := i + 1
		dataRow := buildHuadingRowData(item, orderData, seqNo, thirdPartyOrderNo, config.CustomerID)

		// 验证数据
		validationErrors := validateRowData(dataRow, currentRow)
		result.ValidationErrors = append(result.ValidationErrors, validationErrors...)

		// 写入数据并标记pending字段
		pendingFields := writeRowToTemplate(f, sheetName, currentRow, dataRow, pendingStyle)
		result.PendingFields = append(result.PendingFields, pendingFields...)
	}

	result.RowCount = len(rows)

	// 保存文件
	outputDir := config.OutputDir
	if outputDir == "" {
		outputDir = "/tmp"
	}
	fileName := fmt.Sprintf("%s/华鼎出库单_%s_%s.xlsx", outputDir, config.CustomerID, time.Now().Format("20060102150405"))
	if err := f.SaveAs(fileName); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	result.FilePath = fileName
	return result, nil
}

// setColumnWidths 设置列宽
func setColumnWidths(f *excelize.File, sheetName string) {
	// 根据字段定义设置列宽
	widths := map[string]float64{
		"A": 8,   // 序号
		"B": 15,  // 门店编号
		"C": 15,  // 门店三方编码
		"D": 12,  // 仓库编码
		"E": 10,  // 加急程度
		"F": 15,  // 商品编码
		"G": 25,  // 商品名称
		"H": 15,  // 规格
		"I": 10,  // 数量
		"J": 10,  // 单位
		"K": 18,  // 商品SKU编号
		"L": 20,  // 商品三方SPEC编号
		"M": 12,  // 单位类型
		"N": 14,  // 指定库存状态
		"O": 12,  // 出库类型
		"P": 12,  // 配送方式
		"Q": 15,  // 指定车型
		"R": 12,  // 是否垫付
		"S": 12,  // 付款方式
		"T": 15,  // 快递公司
		"U": 12,  // 单价
		"V": 12,  // 总金额
		"W": 14,  // 是否制定批次
		"X": 15,  // 批次号
		"Y": 14,  // 生产日期
		"Z": 15,  // 生产厂家编号
		"AA": 18, // 门店收货地址编码
		"AB": 30, // 三方单号
		"AC": 12, // 业务模式
		"AD": 12, // 业务类型
		"AE": 12, // 收货人
		"AF": 15, // 联系电话
		"AG": 35, // 收货地址
		"AH": 15, // C端快递公司
		"AI": 30, // 备注
	}

	for col, width := range widths {
		f.SetColWidth(sheetName, col, col, width)
	}
}

// createHeaderStyle 创建表头样式
func createHeaderStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 11, Color: "#000000"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#4472C4"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
}

// createDataStyle 创建数据样式
func createDataStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
}

// createPendingStyle 创建待确认字段样式（黄色高亮）
func createPendingStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#FFFF00"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
}

// extractOrderRows 提取订单行数据
func extractOrderRows(orderData map[string]interface{}) []map[string]interface{} {
	var rows []map[string]interface{}

	if v, ok := orderData["rows"].([]interface{}); ok {
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}

	// 也支持 "items" 键
	if len(rows) == 0 {
		if v, ok := orderData["items"].([]interface{}); ok {
			for _, item := range v {
				if m, ok := item.(map[string]interface{}); ok {
					rows = append(rows, m)
				}
			}
		}
	}

	return rows
}

// buildHuadingRowData 构建华鼎行数据
func buildHuadingRowData(item, orderData map[string]interface{}, seqNo int, thirdPartyOrderNo, customerID string) *HuadingRowData {
	data := &HuadingRowData{
		SeqNo:             seqNo,
		ThirdPartyOrderNo: thirdPartyOrderNo,
	}

	// 从item提取字段（行级别）
	data.ProductCode = getStringValue(item, "sku_code")
	if data.ProductCode == "" {
		data.ProductCode = getStringValue(item, "product_code")
	}
	data.ProductName = getStringValue(item, "sku_name")
	if data.ProductName == "" {
		data.ProductName = getStringValue(item, "product_name")
	}
	data.Spec = getStringValue(item, "spec")
	data.Quantity = int(getFloatValue(item, "quantity"))
	data.Unit = getStringValue(item, "unit")
	data.SkuCode = getStringValue(item, "system_sku_code")
	if data.SkuCode == "" {
		data.SkuCode = getStringValue(item, "sku_code")
	}
	data.SpecCode = getStringValue(item, "spec_code")
	if data.SpecCode == "" {
		data.SpecCode = getStringValue(item, "system_sku_code")
	}
	data.UnitType = getStringValue(item, "unit_type")
	data.UnitPrice = getFloatValue(item, "unit_price")
	data.TotalAmount = getFloatValue(item, "total_amount")
	data.BatchNo = getStringValue(item, "batch_no")
	data.ProductionDate = getStringValue(item, "production_date")
	data.ManufacturerCode = getStringValue(item, "manufacturer_code")
	data.Notes = getStringValue(item, "notes")
	if data.Notes == "" {
		data.Notes = getStringValue(item, "remark")
	}

	// 从orderData提取字段（订单级别）
	data.StoreCode = getStringValue(orderData, "store_code")
	if data.StoreCode == "" {
		data.StoreCode = getStringValue(item, "store_code")
	}
	data.StoreThirdPartyCode = getStringValue(orderData, "store_third_party_code")
	data.WarehouseCode = getStringValue(orderData, "warehouse_code")
	data.UrgencyLevel = getStringValue(orderData, "urgency_level")
	data.InventoryStatus = getStringValue(orderData, "inventory_status")
	data.OutboundType = getStringValue(orderData, "outbound_type")
	data.DeliveryMethod = getStringValue(orderData, "delivery_method")
	data.VehicleType = getStringValue(orderData, "vehicle_type")
	data.IsAdvancePayment = getStringValue(orderData, "is_advance_payment")
	data.PaymentMethod = getStringValue(orderData, "payment_method")
	data.LogisticsCompany = getStringValue(orderData, "logistics_company")
	data.IsBatchSpecified = getStringValue(orderData, "is_batch_specified")
	data.StoreAddressCode = getStringValue(orderData, "store_address_code")
	data.BusinessModel = getStringValue(orderData, "business_model")
	data.BusinessType = getStringValue(orderData, "business_type")
	data.ReceiverName = getStringValue(orderData, "receiver_name")
	if data.ReceiverName == "" {
		data.ReceiverName = getStringValue(orderData, "contact")
	}
	if data.ReceiverName == "" {
		data.ReceiverName = getStringValue(orderData, "receiver")
	}
	data.ReceiverPhone = getStringValue(orderData, "receiver_phone")
	if data.ReceiverPhone == "" {
		data.ReceiverPhone = getStringValue(orderData, "phone")
	}
	if data.ReceiverPhone == "" {
		data.ReceiverPhone = getStringValue(orderData, "contact_phone")
	}
	data.ReceiverAddress = getStringValue(orderData, "receiver_address")
	if data.ReceiverAddress == "" {
		data.ReceiverAddress = getStringValue(orderData, "address")
	}
	data.CLogisticsCompany = getStringValue(orderData, "c_logistics_company")

	// 应用默认值
	applyDefaultValues(data)

	// 计算总金额（如果未提供）
	if data.TotalAmount == 0 && data.UnitPrice > 0 && data.Quantity > 0 {
		data.TotalAmount = data.UnitPrice * float64(data.Quantity)
	}

	return data
}

// applyDefaultValues 应用默认值
func applyDefaultValues(data *HuadingRowData) {
	defaults := getDefaultValueFields()

	if data.UrgencyLevel == "" {
		data.UrgencyLevel = defaults["加急程度"].(string)
	}
	if data.InventoryStatus == "" {
		data.InventoryStatus = defaults["指定库存状态"].(string)
	}
	if data.OutboundType == "" {
		data.OutboundType = defaults["出库类型"].(string)
	}
	if data.IsAdvancePayment == "" {
		data.IsAdvancePayment = defaults["是否垫付"].(string)
	}
	if data.PaymentMethod == "" {
		data.PaymentMethod = defaults["付款方式"].(string)
	}
	if data.IsBatchSpecified == "" {
		data.IsBatchSpecified = defaults["是否制定批次"].(string)
	}
	if data.BusinessModel == "" {
		data.BusinessModel = defaults["业务模式"].(string)
	}
	if data.WarehouseCode == "" {
		data.WarehouseCode = defaults["仓库编码"].(string)
	}
}

// writeRowData 写入行数据并标记pending字段
func writeRowData(f *excelize.File, sheetName string, row int, data *HuadingRowData, fields []HuadingField, dataStyle, pendingStyle int) []PendingField {
	var pendingFields []PendingField

	// 构建值映射
	values := map[string]interface{}{
		"A":  data.SeqNo,
		"B":  data.StoreCode,
		"C":  data.StoreThirdPartyCode,
		"D":  data.WarehouseCode,
		"E":  data.UrgencyLevel,
		"F":  data.ProductCode,
		"G":  data.ProductName,
		"H":  data.Spec,
		"I":  data.Quantity,
		"J":  data.Unit,
		"K":  data.SkuCode,
		"L":  data.SpecCode,
		"M":  data.UnitType,
		"N":  data.InventoryStatus,
		"O":  data.OutboundType,
		"P":  data.DeliveryMethod,
		"Q":  data.VehicleType,
		"R":  data.IsAdvancePayment,
		"S":  data.PaymentMethod,
		"T":  data.LogisticsCompany,
		"U":  data.UnitPrice,
		"V":  data.TotalAmount,
		"W":  data.IsBatchSpecified,
		"X":  data.BatchNo,
		"Y":  data.ProductionDate,
		"Z":  data.ManufacturerCode,
		"AA": data.StoreAddressCode,
		"AB": data.ThirdPartyOrderNo,
		"AC": data.BusinessModel,
		"AD": data.BusinessType,
		"AE": data.ReceiverName,
		"AF": data.ReceiverPhone,
		"AG": data.ReceiverAddress,
		"AH": data.CLogisticsCompany,
		"AI": data.Notes,
	}

	// 写入每个单元格
	for _, field := range fields {
		cell := fmt.Sprintf("%s%d", field.ColumnLetter, row)
		value := values[field.ColumnLetter]

		// 设置值
		if value != nil && value != "" {
			f.SetCellValue(sheetName, cell, value)
		}

		// 检查是否为pending字段（必填但为空）
		isEmpty := isEmptyValue(value)
		if field.Required && isEmpty {
			// 标记为pending（黄色高亮）
			f.SetCellStyle(sheetName, cell, cell, pendingStyle)
			// 添加批注
			comment := fmt.Sprintf("待确认：%s为必填项，但当前值为空", field.Name)
			f.AddComment(sheetName, excelize.Comment{
				Cell:   cell,
				Author: "系统",
				Text:   comment,
			})

			pendingFields = append(pendingFields, PendingField{
				FieldName: field.Name,
				Column:    field.ColumnLetter,
				Row:       row,
				Reason:    "必填字段为空",
			})
		} else {
			f.SetCellStyle(sheetName, cell, cell, dataStyle)
		}
	}

	return pendingFields
}

// writeRowToTemplate 写入数据到现有模板
func writeRowToTemplate(f *excelize.File, sheetName string, row int, data *HuadingRowData, pendingStyle int) []PendingField {
	var pendingFields []PendingField

	// 列映射（基于华鼎标准模板）
	columnMapping := map[string]interface{}{
		"A":  data.SeqNo,
		"B":  data.StoreCode,
		"C":  data.StoreThirdPartyCode,
		"D":  data.WarehouseCode,
		"E":  data.UrgencyLevel,
		"F":  data.ProductCode,
		"G":  data.ProductName,
		"H":  data.Spec,
		"I":  data.Quantity,
		"J":  data.Unit,
		"K":  data.SkuCode,
		"L":  data.SpecCode,
		"M":  data.UnitType,
		"N":  data.InventoryStatus,
		"O":  data.OutboundType,
		"P":  data.DeliveryMethod,
		"Q":  data.VehicleType,
		"R":  data.IsAdvancePayment,
		"S":  data.PaymentMethod,
		"T":  data.LogisticsCompany,
		"U":  data.UnitPrice,
		"V":  data.TotalAmount,
		"W":  data.IsBatchSpecified,
		"X":  data.BatchNo,
		"Y":  data.ProductionDate,
		"Z":  data.ManufacturerCode,
		"AA": data.StoreAddressCode,
		"AB": data.ThirdPartyOrderNo,
		"AC": data.BusinessModel,
		"AD": data.BusinessType,
		"AE": data.ReceiverName,
		"AF": data.ReceiverPhone,
		"AG": data.ReceiverAddress,
		"AH": data.CLogisticsCompany,
		"AI": data.Notes,
	}

	// 必填字段列表
	requiredFields := map[string]string{
		"B": "门店编号",
		"D": "仓库编码",
		"I": "数量",
		"L": "商品三方SPEC编号",
		"M": "单位类型",
		"P": "配送方式",
		"AD": "业务类型",
		"AE": "收货人",
		"AF": "联系电话",
		"AG": "收货地址",
	}

	for col, value := range columnMapping {
		cell := fmt.Sprintf("%s%d", col, row)

		// 设置值
		if value != nil && value != "" {
			f.SetCellValue(sheetName, cell, value)
		}

		// 检查必填字段
		if fieldName, isRequired := requiredFields[col]; isRequired && isEmptyValue(value) {
			// 应用pending样式
			f.SetCellStyle(sheetName, cell, cell, pendingStyle)
			// 添加批注
			comment := fmt.Sprintf("待确认：%s为必填项，但当前值为空", fieldName)
			f.AddComment(sheetName, excelize.Comment{
				Cell:   cell,
				Author: "系统",
				Text:   comment,
			})

			pendingFields = append(pendingFields, PendingField{
				FieldName: fieldName,
				Column:    col,
				Row:       row,
				Reason:    "必填字段为空",
			})
		}
	}

	return pendingFields
}

// validateRowData 验证行数据
func validateRowData(data *HuadingRowData, row int) []ValidationError {
	var errors []ValidationError

	// 验证数量必须大于0
	if data.Quantity <= 0 {
		errors = append(errors, ValidationError{
			Row:       row,
			Column:    "I",
			FieldName: "数量",
			Message:   "数量必须大于0",
		})
	}

	// 验证单价不能为负数
	if data.UnitPrice < 0 {
		errors = append(errors, ValidationError{
			Row:       row,
			Column:    "U",
			FieldName: "单价",
			Message:   "单价不能为负数",
		})
	}

	// 验证总金额不能为负数
	if data.TotalAmount < 0 {
		errors = append(errors, ValidationError{
			Row:       row,
			Column:    "V",
			FieldName: "总金额",
			Message:   "总金额不能为负数",
		})
	}

	// 验证联系电话格式（如果有值）
	if data.ReceiverPhone != "" {
		if !isValidPhone(data.ReceiverPhone) {
			errors = append(errors, ValidationError{
				Row:       row,
				Column:    "AF",
				FieldName: "联系电话",
				Message:   "联系电话格式不正确",
			})
		}
	}

	// 验证生产日期格式（如果有值）
	if data.ProductionDate != "" {
		if !isValidDate(data.ProductionDate) {
			errors = append(errors, ValidationError{
				Row:       row,
				Column:    "Y",
				FieldName: "生产日期",
				Message:   "生产日期格式不正确，应为YYYY-MM-DD",
			})
		}
	}

	return errors
}

// isEmptyValue 检查值是否为空
func isEmptyValue(v interface{}) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case string:
		return strings.TrimSpace(val) == ""
	case int:
		return val == 0
	case float64:
		return val == 0
	case float32:
		return val == 0
	default:
		return false
	}
}

// isValidPhone 验证电话格式
func isValidPhone(phone string) bool {
	// 简单验证：手机号11位或固话7-8位
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, " ", "")

	if len(phone) == 11 && strings.HasPrefix(phone, "1") {
		return true
	}
	if len(phone) >= 7 && len(phone) <= 8 {
		return true
	}
	// 支持区号+号码格式
	if len(phone) >= 10 && len(phone) <= 12 {
		return true
	}
	return false
}

// isValidDate 验证日期格式
func isValidDate(date string) bool {
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		"20060102",
	}

	for _, format := range formats {
		if _, err := time.Parse(format, date); err == nil {
			return true
		}
	}
	return false
}

// findHeaderRow 查找表头行
func findHeaderRow(f *excelize.File, sheetName string) int {
	// 简单策略：查找包含"序号"或"门店编号"的行
	for row := 1; row <= 10; row++ {
		cellA, _ := f.GetCellValue(sheetName, fmt.Sprintf("A%d", row))
		cellB, _ := f.GetCellValue(sheetName, fmt.Sprintf("B%d", row))

		if cellA == "序号" || cellB == "门店编号" {
			return row
		}
	}
	return 0
}

// ==================== 辅助函数 ====================

// getStringValue 安全获取字符串值
func getStringValue(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	// 尝试其他类型转换
	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// getFloatValue 安全获取浮点数值
func getFloatValue(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	if v, ok := m[key].(float32); ok {
		return float64(v)
	}
	if v, ok := m[key].(int); ok {
		return float64(v)
	}
	if v, ok := m[key].(int64); ok {
		return float64(v)
	}
	if v, ok := m[key].(string); ok {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}

// ExcelToJSON 将Excel数据转换为JSON
func ExcelToJSON(data *ExcelData) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

// ==================== 旧函数兼容层（已废弃，保留用于向后兼容） ====================

// TemplateFilePath 出库单导入模板文件路径（已废弃，使用HuadingTemplateConfig.TemplateFilePath替代）
// Deprecated: 使用 HuadingTemplateConfig 替代
const TemplateFilePath = ""

// generateHuadingTemplateFromScratch 旧函数（已废弃，保留向后兼容）
// Deprecated: 使用 generateFromScratch 替代
func generateHuadingTemplateFromScratch(orderData map[string]interface{}) (string, error) {
	config := &HuadingTemplateConfig{
		OutputDir:  "/tmp",
		CustomerID: "UNKNOWN",
	}
	result, err := generateFromScratch(orderData, config)
	if err != nil {
		return "", err
	}
	return result.FilePath, nil
}
