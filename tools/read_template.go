package main

import (
	"fmt"
	"log"

	"ai-cs-support/tools/excel"
)

func main() {
	filePath := "/Users/jinqianfei/Library/Mobile Documents/com~apple~CloudDocs/ProductManagement/AI客服赋能/资料/出库单导入模板 (2).xls"
	data, err := excel.ReadExcel(filePath, 0, 50)
	if err != nil {
		log.Fatalf("Failed to read Excel: %v", err)
	}

	fmt.Printf("Sheet Name: %s\n", data.SheetName)
	fmt.Printf("Row Count: %d\n", data.RowCount)
	fmt.Printf("Col Count: %d\n", data.ColCount)
	fmt.Println("Headers:")
	for i, header := range data.Headers {
		fmt.Printf("  %d: %s\n", i, header)
	}
	fmt.Println("First 10 rows:")
	for i, row := range data.Rows {
		if i >= 10 {
			break
		}
		fmt.Printf("  Row %d: %v\n", i, row)
	}
}