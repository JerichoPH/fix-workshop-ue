package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationLineRouter 线别路由
type LocationLineRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationLineRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/locationLine",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationLineController).N(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationLineController).R(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationLineController).E(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationLineController).D(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationLineController).L(ctx) })
	}
}
