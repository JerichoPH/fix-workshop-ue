package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorRoomRouter 室内上道位置机房路由
type PositionIndoorRoomRouter struct{}

// Load 加载路由
//  @receiver PositionIndoorRoomRouter
//  @param router
func (PositionIndoorRoomRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorRoom",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomController).I(ctx) })
	}
}
