package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotSectionRouter 仓储仓库区域路由
type PositionDepotSectionRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (ins PositionDepotSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotSection",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotSectionController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotSectionController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotSectionController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotSectionController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotSectionController).I(ctx) })
	}
}
