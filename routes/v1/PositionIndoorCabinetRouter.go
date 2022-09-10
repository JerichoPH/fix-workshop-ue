package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorCabinetRouter 室内上道位置机柜路由
type PositionIndoorCabinetRouter struct{}

// Load 加载路由
//  @receiver PositionIndoorCabinetRouter
//  @param router
func (PositionIndoorCabinetRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorCabinet",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionIndoorCabinetController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCabinetController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCabinetController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCabinetController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionIndoorCabinetController).I(ctx) })
	}
}
