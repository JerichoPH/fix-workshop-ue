package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationParagraphRouter 站段路由
type OrganizationParagraphRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (OrganizationParagraphRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/organizationParagraph",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationParagraphController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationParagraphController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationParagraphController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationParagraphController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationParagraphController).I(ctx) })
	}
}
