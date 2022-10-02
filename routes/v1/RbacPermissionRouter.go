package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// RbacPermissionRouter 权限路由
type RbacPermissionRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (RbacPermissionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/rbacPermission",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.RbacPermissionController).N(ctx) })

		// 批量添加资源权限
		r.POST("resource", func(ctx *gin.Context) { new(controllers.RbacPermissionController).PostResource(ctx) })

		// 删除权限
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionController).R(ctx) })

		// 编辑权限
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionController).E(ctx) })

		// 权限详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionController).D(ctx) })

		// 权限列表
		r.GET("", func(ctx *gin.Context) { new(controllers.RbacPermissionController).L(ctx) })
	}

}
