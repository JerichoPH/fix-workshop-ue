package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// KindEntireTypeRouter 类型路由
type KindEntireTypeRouter struct{}

// Load 加载路由
//  @receiver KindEntireTypeRouter
//  @param engine
func (KindEntireTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kindEntireType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.KindEntireTypeController).N(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.KindEntireTypeController).R(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.KindEntireTypeController).E(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.KindEntireTypeController).D(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.KindEntireTypeController).L(ctx) })
	}
}
