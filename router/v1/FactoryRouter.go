package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type FactoryRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *FactoryRouter) Load() {
	r := cls.Router.Group(
		"/api/v1",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		// 供应商详情
		r.GET("factory/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			factory := (&model.FactoryModel{}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(factory, model.FactoryModel{}, "供应商")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"factory": factory}))
		})
	}
}
