package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// KindCategoryRouter 种类路由
type KindCategoryRouter struct{}

func (KindCategoryRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kindCategory",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.KindCategoryController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.KindCategoryController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.KindCategoryController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.KindCategoryController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.KindCategoryController).I(ctx) })
	}
}
