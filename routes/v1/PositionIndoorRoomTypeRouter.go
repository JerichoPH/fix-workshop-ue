package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorRoomTypeRouter 机房路由
type PositionIndoorRoomTypeRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param engine
func (PositionIndoorRoomTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorRoomType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomTypeController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomTypeController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomTypeController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomTypeController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionIndoorRoomTypeController).I(ctx) })
	}
}
