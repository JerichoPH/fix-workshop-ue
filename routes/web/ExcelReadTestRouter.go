package web

import (
	"fix-workshop-ue/controllers"
	"github.com/gin-gonic/gin"
)

type ExcelReadTestRouter struct{}

func (ExcelReadTestRouter) Load(engine *gin.Engine) {
	r := engine.Group("excelReadTest")
	{
		r.GET("/", func(ctx *gin.Context) { (&controllers.ExcelReadTestController{}).L(ctx) })
	}
}
