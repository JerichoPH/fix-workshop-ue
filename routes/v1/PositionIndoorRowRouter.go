package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorRowRouter 上道位置排路由
type PositionIndoorRowRouter struct{}

// Load 加载路由
//  @receiver PositionIndoorRowRouter
//  @param router
func (PositionIndoorRowRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorRow",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { (&controllers.PositionIndoorRowController{}).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { (&controllers.PositionIndoorRowController{}).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { (&controllers.PositionIndoorRowController{}).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { (&controllers.PositionIndoorRowController{}).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { (&controllers.PositionIndoorRowController{}).I(ctx) })
	}
}
