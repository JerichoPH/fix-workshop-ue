package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// KindSubTypeRouter 型号路由
type KindSubTypeRouter struct{}

// Load 加载路由
//  @receiver KindSubTypeRouter
//  @param engine
func (KindSubTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kindSubType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.KindSubTypeController).N(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.KindSubTypeController).R(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.KindSubTypeController).E(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.KindSubTypeController).D(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.KindSubTypeController).L(ctx) })
	}
}
