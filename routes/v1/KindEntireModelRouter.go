package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
)

type KindEntireModelRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *KindEntireModelRouter) Load() {
	r := cls.Router.Group("/api/v1/kind")
	{
		// 类型详情
		r.GET("entireType/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindEntireType := (&models.KindEntireTypeModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"KindCategoryModel",
						"KindSubTypes",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindEntireType, models.KindEntireTypeModel{}, "类型")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_entire_type": kindEntireType}))
		})
	}
}
