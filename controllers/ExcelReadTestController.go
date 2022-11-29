package controllers

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type ExcelReadTestController struct{}

// L 列表
func (cls *ExcelReadTestController) L(ctx *gin.Context) {
	dir := os.Getenv("PWD")

	fmt.Println(dir + "/static/Book1.xlsx")

	excelReader := (&tools.ExcelReader{}).
		OpenFile(dir + "/static/Book1.xlsx").
		SetSheetName("Sheet1").
		ReadTitle().
		Read()

	fmt.Println(excelReader.GetTitle(), excelReader.ToList(), excelReader.ToMap())

	// 设置表头
	titleRow := new(tools.ExcelRow).
		SetRowNumber(1).
		SetCells([]*tools.ExcelCell{
			new(tools.ExcelCell).SetContent("姓名").SetFontColor("#FF0000", true),
			new(tools.ExcelCell).SetContent("年龄"),
			new(tools.ExcelCell).SetContent("性别"),
		})

	// 写入Excel
	err := (&tools.ExcelWriter{}).
		Init(dir + "/static/Book2.xlsx").
		ActiveSheetByIndex(0).
		SetRows([]*tools.ExcelRow{titleRow}).
		Save()
	if err != nil {
		wrongs.PanicForbidden(fmt.Sprintf("保存文件失败：%s", err.Error()))
	}

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"dir": dir}))
}
