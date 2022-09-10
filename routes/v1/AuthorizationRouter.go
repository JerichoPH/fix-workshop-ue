package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

type AuthorizationRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (AuthorizationRouter) Load(engine *gin.Engine) {
	r := engine.Group("/api/v1/authorization")
	{
		// 注册
		r.POST("register", func(ctx *gin.Context) { new(controllers.AuthorizationController).PostRegister(ctx) })

		// 登录
		r.POST("login", func(ctx *gin.Context) { new(controllers.AuthorizationController).PostLogin(ctx) })

		// 获取当前账号相关菜单
		r.GET(
			"menus",
			middlewares.CheckJwt(),
			func(ctx *gin.Context) { new(controllers.AuthorizationController).GetMenus(ctx) },
		)
	}
}
