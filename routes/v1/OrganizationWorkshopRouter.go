package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationWorkshopRouter 车间路由
type OrganizationWorkshopRouter struct{}

// Load 加载路由
//  @receiver OrganizationWorkshopRouter
//  @param router
func (OrganizationWorkshopRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkshop",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopController).I(ctx) })
	}
}
