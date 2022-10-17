package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationCenterRouter 中心路由
type LocationCenterRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (LocationCenterRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationCenter",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationCenterController).N(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationCenterController).R(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationCenterController).E(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationCenterController).D(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationCenterController).L(ctx) })
	}
}
