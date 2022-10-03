package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// RbacPermissionGroupRouter 权限分组路由
type RbacPermissionGroupRouter struct{}

// Load 加载路由
//  @receiver ins
//  @param router
func (RbacPermissionGroupRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/rbacPermissionGroup",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 创建权限分组
		r.POST("", func(ctx *gin.Context) { new(controllers.RbacPermissionGroupController).N(ctx) })

		// 删除用户分组
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionGroupController).R(ctx) })

		// 编辑权限分组
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionGroupController).E(ctx) })

		// 权限分组详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.RbacPermissionGroupController).D(ctx) })

		// 权限分组列表
		r.GET("", func(ctx *gin.Context) { new(controllers.RbacPermissionGroupController).L(ctx) })
	}
}
