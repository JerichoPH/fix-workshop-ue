package v1

import (
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// TagRouter 赋码路由
type TagRouter struct{}

func (TagRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/tag",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 下载赋码模板Excel
		r.GET("template", func(ctx *gin.Context) {
		})
	}
}
