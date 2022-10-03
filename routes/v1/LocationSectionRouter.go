package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// LocationSectionRouter 区间路由
type LocationSectionRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (LocationSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationSection",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.LocationSectionController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.LocationSectionController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.LocationSectionController).U(ctx) })

		// 区间绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) { new(controllers.LocationSectionController).PutBindLines(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.LocationSectionController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.LocationSectionController).I(ctx) })
	}
}
