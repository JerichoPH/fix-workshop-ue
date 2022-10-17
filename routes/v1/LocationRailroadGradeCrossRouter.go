package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationRailroadRouter 道口路由
type LocationRailroadRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (LocationRailroadRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationRailroadGradeCross",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationRailroadController).N(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadController).R(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadController).E(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadController).D(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationRailroadController).L(ctx) })
	}
}
