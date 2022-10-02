package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

type RbacRoleRouter struct{}

// Load 加载路由
//  @receiver RbacRoleRouter
//  @param engine
func (RbacRoleRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/rbacRole",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建角色
		r.POST("", func(ctx *gin.Context) { new(controllers.RbacRoleController).N(ctx) })

		// 删除角色
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.RbacRoleController).R(ctx) })

		// 编辑角色
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.RbacRoleController).E(ctx) })

		// 绑定用户
		r.PUT(":uuid/bindAccounts", func(ctx *gin.Context) { new(controllers.RbacRoleController).PutBindAccounts(ctx) })

		// 绑定权限
		r.PUT(":uuid/bindRbacPermissions", func(ctx *gin.Context) { new(controllers.RbacRoleController).PutBindRbacPermissions(ctx) })

		// 绑定权限（根据权限分组）
		r.PUT(":uuid/bindRbacPermissionsByRbacPermissionGroup", func(ctx *gin.Context) {
			new(controllers.RbacRoleController).PutBindRbacPermissionsByRbacPermissionGroup(ctx)
		})

		// 角色详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.RbacRoleController).D(ctx) })

		// 角色列表
		r.GET("", func(ctx *gin.Context) { new(controllers.RbacRoleController).L(ctx) })
	}
}
