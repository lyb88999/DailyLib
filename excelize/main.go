package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	NewWorkbook()
	ReadWorkbook()
	ModifyWorkbook()
	ReadWorkbook()
	SetCellStyle()
	InsertChart()
}

// NewWorkbook 创建一个工作簿
func NewWorkbook() {
	f := excelize.NewFile()
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close the file")
		}
	}(f)

	// 设置工作表"sheet1"中"A1"单元格的值
	err := f.SetCellValue("sheet1", "A1", "Hello Excel!")
	if err != nil {
		log.Fatalf("failed to set cell value: %v", err)
	}
	log.Println("Successfully set cell value")

	// 保存文件
	err = f.SaveAs("test.xlsx")
	if err != nil {
		log.Fatalf("failed to save file: %v", err)
	}
	log.Println("Successfully saved file")
}

// ReadWorkbook 读取一个现有的工作簿
func ReadWorkbook() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close the file")
		}
	}(f)
	// 读取"sheet1"工作表的"A1"单元格的值
	cell, err := f.GetCellValue("sheet1", "A1")
	if err != nil {
		log.Fatalf("failed to get cell value: %v", err)
	}
	log.Printf("Cell A1 value: %v", cell)
}

// ModifyWorkbook 修改现有的工作簿的单元格内容
func ModifyWorkbook() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close the file")
		}
	}(f)
	// 修改"sheet1"工作表的"A1"单元格的值
	err = f.SetCellValue("sheet1", "A1", "Hello World!")
	if err != nil {
		log.Fatalf("failed to set cell value: %v", err)
	}
	log.Println("Successfully update cell value")
	err = f.Save()
	if err != nil {
		log.Fatalf("failed to save file: %v", err)
	}
	log.Println("Successfully saved file")
}

// SetCellStyle 设置单元格样式
func SetCellStyle() {
	f := excelize.NewFile()
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close the file")
		}
	}(f)

	// 设置单元格A1的字体样式
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  12,
			Color: "FF0000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		}})
	if err != nil {
		log.Fatalf("failed to create style: %v", err)
	}

	// 应用样式到单元格A1
	err = f.SetCellValue("Sheet1", "A1", "Styled Cell")
	if err != nil {
		log.Fatalf("failed to set cell value: %v", err)
	}
	log.Println("Successfully set cell value")

	err = f.SetCellStyle("Sheet1", "A1", "A1", style)
	if err != nil {
		log.Fatalf("failed to set cell style: %v", err)
	}
	log.Println("Successfully set cell style")
	err = f.SaveAs("styled.xlsx")
	if err != nil {
		log.Fatalf("failed to save file: %v", err)
	}
	log.Println("Successfully saved file")
}

// InsertChart 插入一个图表
func InsertChart() {
	f := excelize.NewFile()
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close the file")
		}
	}(f)

	// 向sheet1中插入数据
	for i := 1; i <= 10; i++ {
		err := f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i), i)
		if err != nil {
			log.Fatalf("failed to set cell value: %v", err)
		}
		err = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i), i*i)
		if err != nil {
			log.Fatalf("failed to set cell value: %v", err)
		}
	}
	log.Println("Successfully set cell value")
	err := f.AddChart("sheet1", "D1", &excelize.Chart{
		Type: excelize.Scatter,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$B$1",
				Categories: "Sheet1!$A$1:$A$10",
				Values:     "Sheet1!$B$1:$B$10"},
		},
	})
	if err != nil {
		log.Fatalf("failed to add chart: %v", err)
	}
	log.Println("Successfully added chart")

	err = f.SaveAs("chart.xlsx")
	if err != nil {
		log.Fatalf("failed to save file: %v", err)
	}
	log.Println("Successfully saved file")
}
