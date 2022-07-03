package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type KindCategoryRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *KindCategoryRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/kind",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		// 种类详情
		r.GET("category/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindCategory := (&model.KindCategoryModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"KindEntireTypes",
						"KindEntireTypes.KindSubTypes",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(kindCategory, model.KindCategoryModel{}, "种类")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"kind_category": kindCategory}))
		})
	}
}
