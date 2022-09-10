package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationWorkAreaTypeRouter 工区类型路由
type OrganizationWorkAreaTypeRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationWorkAreaTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkAreaType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaTypeController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaTypeController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaTypeController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaTypeController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaTypeController).I(ctx) })
	}
}
