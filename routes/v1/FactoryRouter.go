package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type FactoryRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *FactoryRouter) Load() {
	r := cls.Router.Group(
		"/api/v1",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// 供应商详情
		r.GET("factory/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			factory := (&models.FactoryModel{}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(factory, models.FactoryModel{}, "供应商")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"factory": factory}))
		})
	}
}
