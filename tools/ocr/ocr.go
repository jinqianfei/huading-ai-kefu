/**
 * OCR 工具模块
 * 负责 PDF 和图片的文字识别
 */

package ocr

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ReadPDF 读取PDF文件文本内容
func ReadPDF(filePath string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file not found: %s", filePath)
	}

	// 尝试使用 pdftotext（需要安装 poppler-utils）
	cmd := exec.Command("pdftotext", "-layout", filePath, "-")
	output, err := cmd.Output()
	if err != nil {
		// 如果 pdftotext 不可用，尝试使用其他方式
		return fallbackPDFRead(filePath)
	}

	return string(output), nil
}

// fallbackPDFRead PDF读取降级方案
func fallbackPDFRead(filePath string) (string, error) {
	// 尝试使用 strings.Replace 模拟（实际应该使用 PDF 解析库）
	// 这里只是一个占位实现
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// 尝试提取可读文本
	text := extractReadableText(string(content))
	return text, nil
}

// extractReadableText 从二进制内容中提取可读文本
func extractReadableText(content string) string {
	var result strings.Builder
	inText := false

	for _, c := range content {
		// 简单判断是否为可打印字符
		if (c >= 32 && c <= 126) || c == '\n' || c == '\r' || c == '\t' {
			if !inText {
				inText = true
			}
			result.WriteRune(c)
		} else {
			if inText && result.Len() > 0 {
				last := result.Len() - 1
				if rune(result.String()[last]) != '\n' {
					result.WriteRune('\n')
				}
			}
			inText = false
		}
	}

	return result.String()
}

// ReadImageOCR 图片OCR识别
func ReadImageOCR(imagePath string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file not found: %s", imagePath)
	}

	// TODO: 集成阿里云表格OCR服务
	// 目前返回占位文本
	return "OCR识别功能开发中...\n图片路径: " + imagePath, nil
}

// ReadTableFromImage 从图片中提取表格
func ReadTableFromImage(imagePath string) ([][]string, error) {
	// TODO: 集成阿里云表格OCR
	// 目前返回空表格
	return [][]string{}, nil
}
