package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
)

type KindCategoryRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *KindCategoryRouter) Load() {
	r := cls.Router.Group("/api/v1/kindCategory")
	{
		// 种类详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindCategory := (&models.KindCategoryModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"KindEntireModels",
						"KindEntireModels.KindSubModels",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindCategory, models.KindCategoryModel{}, "种类")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_category": kindCategory}))
		})
	}
}
