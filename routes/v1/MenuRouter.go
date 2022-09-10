package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

// MenuRouter 菜单路由
type MenuRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (MenuRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/menu",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建菜单
		r.POST("", func(ctx *gin.Context) { new(controllers.MenuController).C(ctx) })

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) { new(controllers.MenuController).D(ctx) })

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) { new(controllers.MenuController).U(ctx) })

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) { new(controllers.MenuController).S(ctx) })

		// 列表
		r.GET("", func(ctx *gin.Context) { new(controllers.MenuController).I(ctx) })
	}
}
