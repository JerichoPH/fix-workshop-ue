package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// OrganizationWorkshopTypeRouter 车间类型路由
type OrganizationWorkshopTypeRouter struct{}

// Load 加载路由
//  @receiver OrganizationWorkshopTypeRouter
//  @param router
func (OrganizationWorkshopTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkshopType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建车间类型
		r.POST("", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopTypeController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopTypeController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopTypeController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopTypeController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.OrganizationWorkshopTypeController).I(ctx) })
	}
}
