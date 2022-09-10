package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorTierRouter 室内上道位置柜架层路由
type PositionIndoorTierRouter struct{}

// Load 加载路由
//  @receiver PositionIndoorTierRouter
//  @param engine
func (PositionIndoorTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/position",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionIndoorTierController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorTierController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorTierController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorTierController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionIndoorTierController).I(ctx) })
	}
}
