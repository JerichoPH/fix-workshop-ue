package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionDepotRowRouter 仓储仓库排路由
type PositionDepotRowRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param engine
func (PositionDepotRowRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotRow",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionDepotRowController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotRowController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotRowController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionDepotRowController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionDepotRowController).I(ctx) })
	}
}
