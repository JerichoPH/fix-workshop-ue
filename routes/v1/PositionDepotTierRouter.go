package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotTierRouter 仓库柜架层路由
type PositionDepotTierRouter struct{}

// Load 加载路由
//  @receiver PositionDepotTierRouter
//  @param router
func (PositionDepotTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotTier",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotTierController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotTierController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotTierController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotTierController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotTierController).I(ctx) })
	}
}
