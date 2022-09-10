package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationRailwayRouter 路局路由
type OrganizationRailwayRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationRailwayRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/organizationRailway",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建路局
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationRailwayController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationRailwayController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationRailwayController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationRailwayController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationRailwayController).I(ctx) })
	}
}
