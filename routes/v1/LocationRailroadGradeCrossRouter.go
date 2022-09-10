package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationRailroadGradeCrossRouter 道口路由
type LocationRailroadGradeCrossRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationRailroadGradeCrossRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationRailroadGradeCross",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).U(ctx) })

		// 道口绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).PutBindLines(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationRailroadGradeCrossController).I(ctx) })
	}
}
