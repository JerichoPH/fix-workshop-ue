package web

import (
	"fix-workshop-ue/controllers"
	"github.com/gin-gonic/gin"
)

type CommandRouter struct{}

func (CommandRouter) Load(engine *gin.Engine) {
	r := engine.Group("command")
	{
		// ExcelHelper类演示
		r.GET("excelHelperDemo", func(ctx *gin.Context) { (&controllers.CommandController{}).ExcelHelperDemo(ctx) })

		// Command类演示
		r.GET("commandHelperDemo", func(ctx *gin.Context) { (&controllers.CommandController{}).CommandHelperDemo(ctx) })
	}
}
