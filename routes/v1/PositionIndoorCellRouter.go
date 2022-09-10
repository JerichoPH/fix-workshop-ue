package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// PositionIndoorCellRouter 室内上道位置-柜架格位
type PositionIndoorCellRouter struct{}

// Load 加载路由
//  @receiver PositionIndoorCellRouter
//  @param engine
func (PositionIndoorCellRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorCell",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.PositionIndoorCellController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCellController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCellController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.PositionIndoorCellController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.PositionIndoorCellController).I(ctx) })
	}
}
