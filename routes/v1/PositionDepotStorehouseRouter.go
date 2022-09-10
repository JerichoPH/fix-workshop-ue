package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotStorehouseRouter 仓储仓库路由
type PositionDepotStorehouseRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls PositionDepotStorehouseRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotStorehouse",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotStorehouseController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotStorehouseController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotStorehouseController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotStorehouseController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotStorehouseController).I(ctx) })
	}
}
