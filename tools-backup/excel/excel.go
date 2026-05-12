/**
 * Excel 工具模块
 * 负责 Excel 文件的读取和生成
 */

package excel

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExcelData Excel数据结构
type ExcelData struct {
	SheetName string        `json:"sheet_name"`
	RowCount  int           `json:"row_count"`
	ColCount  int           `json:"col_count"`
	Headers  []string       `json:"headers"`
	Rows     [][]string     `json:"rows"`
	Metadata map[string]any `json:"metadata,omitempty"`
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

// TemplateFilePath 出库单导入模板文件路径
const TemplateFilePath = "/Users/jinqianfei/openclaw-workspaces/ai-kefu/data/出库单导入模板 (2).xls"

// GenerateHuadingTemplate 生成华鼎标准出库单（基于模板填充）
func GenerateHuadingTemplate(orderData map[string]interface{}) (string, error) {
	// 读取模板文件
	f, err := excelize.OpenFile(TemplateFilePath)
	if err != nil {
		// 如果模板不存在，从零创建
		return generateHuadingTemplateFromScratch(orderData)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	// 提取订单数据
	var orderNo, customerName, contactPerson, phone, address, notes, deliveryDate string
	var items []map[string]interface{}

	if v, ok := orderData["order_no"].(string); ok {
		orderNo = v
	}
	if v, ok := orderData["customer_name"].(string); ok {
		customerName = v
	}
	if v, ok := orderData["contact"].(string); ok {
		contactPerson = v
	}
	if v, ok := orderData["phone"].(string); ok {
		phone = v
	}
	if v, ok := orderData["address"].(string); ok {
		address = v
	}
	if v, ok := orderData["notes"].(string); ok {
		notes = v
	}
	if v, ok := orderData["delivery_date"].(string); ok {
		deliveryDate = v
	}
	if v, ok := orderData["pickup_date"].(string); ok && deliveryDate == "" {
		deliveryDate = v
	}
	if v, ok := orderData["rows"].([]interface{}); ok {
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				items = append(items, m)
			}
		}
	}

	// 根据模板结构填充数据
	// 注意：具体填充位置需要根据实际《出库单导入模板》的表头结构调整
	// 以下是示例填充位置，请根据实际模板修改：
	// 假设模板结构为：
	//   A列: 商品编码, B列: 商品名称, C列: 规格, D列: 数量, E列: 单位, F列: 备注
	// 基本信息区域假设在第1-8行

	// 填充商品明细（从第9行开始，根据模板调整）
	startRow := 9
	for idx, item := range items {
		row := startRow + idx
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), getStringValue(item, "sku_code"))
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), getStringValue(item, "sku_name"))
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), getStringValue(item, "spec"))
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), getFloatValue(item, "quantity"))
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), getStringValue(item, "unit"))
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), getStringValue(item, "notes"))
	}

	// 保存填充后的文件
	outputDir := "/tmp"
	fileName := fmt.Sprintf("%s/华鼎出库单_%s.xlsx", outputDir, time.Now().Format("20060102150405"))
	if err := f.SaveAs(fileName); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	_ = orderNo      // 避免未使用警告
	_ = customerName // 避免未使用警告
	_ = contactPerson
	_ = phone
	_ = address
	_ = notes
	_ = deliveryDate

	return fileName, nil
}

// generateHuadingTemplateFromScratch 从零创建出库单（模板不存在时的降级方案）
func generateHuadingTemplateFromScratch(orderData map[string]interface{}) (string, error) {
	f := excelize.NewFile()
	defer f.Close()

	sheetName := "华鼎标准出库单"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return "", fmt.Errorf("failed to create sheet: %w", err)
	}
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	// 设置列宽
	f.SetColWidth(sheetName, "A", "A", 15)
	f.SetColWidth(sheetName, "B", "B", 12)
	f.SetColWidth(sheetName, "C", "C", 25)
	f.SetColWidth(sheetName, "D", "D", 15)
	f.SetColWidth(sheetName, "E", "E", 30)
	f.SetColWidth(sheetName, "F", "F", 10)
	f.SetColWidth(sheetName, "G", "G", 10)

	// 表头样式
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 12},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#E0E0E0"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})

	// 数据样式
	_, _ = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})

	row := 1

	// 标题
	f.MergeCell(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "华鼎标准出库单")
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:    &excelize.Font{Bold: true, Size: 16},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row), titleStyle)
	row++

	// 提取订单数据
	orderNo := getStringValue(orderData, "order_no")
	customerName := getStringValue(orderData, "customer_name")
	contactPerson := getStringValue(orderData, "contact")
	phone := getStringValue(orderData, "phone")
	address := getStringValue(orderData, "address")
	notes := getStringValue(orderData, "notes")
	deliveryDate := getStringValue(orderData, "pickup_date")
	if deliveryDate == "" {
		deliveryDate = getStringValue(orderData, "delivery_date")
	}

	// 基本信息
	infoRows := [][]string{
		{"订单编号", orderNo, "客户名称", customerName},
		{"门店名称", customerName, "联系人", contactPerson},
		{"联系电话", phone, "收货地址", address},
		{"发货日期", deliveryDate, "备注", notes},
	}
	for _, info := range infoRows {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), info[0])
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), info[1])
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), info[2])
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), info[3])
		f.MergeCell(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row))
		f.MergeCell(sheetName, fmt.Sprintf("D%d", row), fmt.Sprintf("G%d", row))
		row++
	}

	row++ // 空行

	// 商品明细表头
	headers := []string{"商品编码", "商品名称", "规格", "数量", "单位", "备注", "系统SKU"}
	for col, header := range headers {
		colName := string(rune('A'+col)) + fmt.Sprintf("%d", row)
		f.SetCellValue(sheetName, colName, header)
	}
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row), headerStyle)
	row++

	// 商品明细数据
	var items []map[string]interface{}
	if v, ok := orderData["rows"].([]interface{}); ok {
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				items = append(items, m)
			}
		}
	}

	itemStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})

totalQuantity := 0
	for _, item := range items {
		skuCode := getStringValue(item, "sku_code")
		skuName

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), skuCode)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), skuName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), spec)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), unit)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "")
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), systemSku)

		for col := 0; col < 7; col++ {
			colName := string(rune('A'+col)) + fmt.Sprintf("%d", row)
			f.SetCellStyle(sheetName, colName, colName, itemStyle)
		}
		totalQuantity += quantity
		row++
	}

	// 如果没有数据，预留10行空行
	for i := len(items); i < 10; i++ {
		for col := 0; col < 7; col++ {
			colName := string(rune('A'+col)) + fmt.Sprintf("%d", row)
			f.SetCellStyle(sheetName, colName, colName, itemStyle)
		}
		row++
	}

	row++ // 空行

	// 汇总
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("商品种数：%d", len(items)))
	f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), fmt.Sprintf("总件数：%d", totalQuantity))

	row += 2

	// 签名区
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "制单人：")
	f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), "审核人：")
	f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), "日期：")

	// 保存文件
	fileName := fmt.Sprintf("/tmp/华鼎标准出库单_%s.xlsx", time.Now().Format("20060102150405"))
	if err := f.SaveAs(fileName); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return fileName, nil
}

// getStringValue 安全获取字符串值
func getStringValue(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

// getFloatValue 安全获取浮点数值
func getFloatValue(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	if v, ok := m[key].(int); ok {
		return float64(v)
	}
	return 0
}

// ExcelToJSON 将Excel数据转换为JSON
func ExcelToJSON(data *ExcelData) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
