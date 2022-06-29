package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
)

type KindSubModelRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *KindSubModelRouter) Load() {
	r := cls.Router.Group("/api/v1/kindSubModel")
	{
		// 型号详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindSubModel := (&models.KindSubTypeModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"KindCategoryModel",
						"KindEntireTypeModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindSubModel, models.KindSubTypeModel{}, "型号")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_sub_model": kindSubModel}))
		})
	}
}
