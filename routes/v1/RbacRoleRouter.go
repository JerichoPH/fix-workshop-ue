package v1

import (
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
		r.POST("", func(ctx *gin.Context) {

		})

		// 删除角色
		r.DELETE(":uuid", func(ctx *gin.Context) {

		})

		// 编辑角色
		r.PUT(":uuid", func(ctx *gin.Context) {

		})

		// 绑定用户
		r.PUT("role/:uuid/bindAccounts", func(ctx *gin.Context) {

		})

		// 绑定权限
		r.PUT("role/:uuid/bindPermissions", func(ctx *gin.Context) {

		})

		// 角色详情
		r.GET(":uuid", func(ctx *gin.Context) {

		})

		// 角色列表
		r.GET("", func(ctx *gin.Context) {

		})
	}
}
