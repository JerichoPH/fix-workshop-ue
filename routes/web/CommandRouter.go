package web

import (
	"fix-workshop-ue/controllers"
	"github.com/gin-gonic/gin"
)

type CommandRouter struct{}

func (CommandRouter) Load(engine *gin.Engine) {
	r := engine.Group("")
	{
		// ExcelHelper类演示
		r.GET("excelHelper", func(ctx *gin.Context) { (&controllers.CommandController{}).ExcelHelperDemo(ctx) })
	}
}
