package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotCabinetRouter 仓储仓库柜架路由
type PositionDepotCabinetRouter struct{}

// Load 加载路由
//  @receiver PositionDepotCabinetRouter
//  @param router
func (PositionDepotCabinetRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotCabinet",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotCabinetController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCabinetController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCabinetController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCabinetController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotCabinetController).I(ctx) })
	}
}
