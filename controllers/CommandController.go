package controllers

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type CommandController struct{}

// ExcelHelperDemo 列表
func (cls *CommandController) ExcelHelperDemo(ctx *gin.Context) {
	dir := os.Getenv("PWD")
	operation := ctx.Query("operation")
	excelName := ctx.Query("excel_name")

	if operation == "read" {
		excelReader := (&tools.ExcelReader{}).
			OpenFile(fmt.Sprintf("%s/static/%s.xlsx", dir, excelName)).
			SetSheetName("Sheet1").
			ReadTitle().
			Read()

		fmt.Println(excelReader.GetTitle(), excelReader.ToList(), excelReader.ToMap())
		fmt.Println("----------")
		fmt.Println(excelReader.GetByDataFrameUseDefaultType().Records())

		ctx.JSON(
			tools.CorrectBootByDefault().
				Ok(
					tools.Map{
						"title":     excelReader.GetTitle(),
						"list":      excelReader.ToList(),
						"map":       excelReader.ToMap(),
						"dataframe": excelReader.GetByDataFrameUseDefaultType().Maps(),
					},
				),
		)
	} else if operation == "write" {
		// 写入Excel
		// 设置表头
		titleRow := new(tools.ExcelRow).
			SetRowNumber(1).
			SetCells([]*tools.ExcelCell{
				new(tools.ExcelCell).SetContent("姓名").SetFontColor("#FF0000", true),
				new(tools.ExcelCell).SetContent("年龄"),
				new(tools.ExcelCell).SetContent("性别"),
			})
		err := (&tools.ExcelWriter{}).
			Init(fmt.Sprintf("%s/static/%s.xlsx", dir, excelName)).
			ActiveSheetByIndex(0).
			SetRows([]*tools.ExcelRow{titleRow}).
			Save()
		if err != nil {
			wrongs.PanicForbidden(fmt.Sprintf("保存文件失败：%s", err.Error()))
		}
	}
}
