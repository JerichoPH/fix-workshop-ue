package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationWorkAreaProfessionRouter 工区专业路由
type OrganizationWorkAreaProfessionRouter struct{}

// Load 加载路由
//  @receiver OrganizationWorkAreaProfessionRouter
//  @param engine
func (OrganizationWorkAreaProfessionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkAreaProfession",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaProfessionController).C(ctx) })

		// 删除
		r.DELETE("/:uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaProfessionController).D(ctx) })

		// 编辑
		r.PUT("/:uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaProfessionController).U(ctx) })

		// 详情
		r.GET("/:uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaProfessionController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationWorkAreaProfessionController).I(ctx) })
	}
}
