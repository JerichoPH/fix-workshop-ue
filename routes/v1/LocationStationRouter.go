package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationStationRouter 站场路由
type LocationStationRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (LocationStationRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationStation",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationStationController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationStationController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationStationController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationStationController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationStationController).I(ctx) })
	}
}
