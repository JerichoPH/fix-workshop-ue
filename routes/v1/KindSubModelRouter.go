package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type KindSubModelRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *KindSubModelRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/kind",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// 型号详情
		r.GET("subType/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindSubType := (&models.KindSubTypeModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"KindCategory",
						"KindEntireType",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindSubType, models.KindSubTypeModel{}, "型号")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_sub_type": kindSubType}))
		})
	}
}
