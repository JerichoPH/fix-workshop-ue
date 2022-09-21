package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// AccountRouter 用户路由
type AccountRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (AccountRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/account",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) { new(controllers.AccountController).N(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.AccountController).E(ctx) })

		// 修改密码
		r.PUT(":uuid/updatePassword", func(ctx *gin.Context) { new(controllers.AccountController).PutPassword(ctx) })

		// 删除用户
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.AccountController).R(ctx) })

		// 用户详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.AccountController).D(ctx) })

		// 用户列表
		r.GET("", func(ctx *gin.Context) { new(controllers.AccountController).L(ctx) })
	}
}
