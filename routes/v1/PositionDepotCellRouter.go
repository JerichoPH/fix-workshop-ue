package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotCellRouter 仓储柜架格位路由
type PositionDepotCellRouter struct{}

// Load 加载路由
//  @receiver PositionDepotCellRouter
//  @param router
func (PositionDepotCellRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotCell",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotCellController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCellController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCellController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotCellController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotCellController).I(ctx) })
	}
}
