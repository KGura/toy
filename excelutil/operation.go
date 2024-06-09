package excelutil

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)
// 创建excel文件
func CreateExcel(filePath string) {
	f := excelize.NewFile()
	sheet, _ := f.NewSheet("Video")
	f.SetActiveSheet(sheet)
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("创建成功")
	}
	f.Close()
}
