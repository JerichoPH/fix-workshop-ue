package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationWorkAreaRouter 工区路由
type OrganizationWorkAreaRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (OrganizationWorkAreaRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkArea",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaController).I(ctx) })
	}
}
