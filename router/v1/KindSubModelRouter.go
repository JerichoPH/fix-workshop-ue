package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type KindSubModelRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *KindSubModelRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/kind",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		// 型号详情
		r.GET("subType/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindSubType := (&model.KindSubTypeModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"KindCategory",
						"KindEntireType",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(kindSubType, model.KindSubTypeModel{}, "型号")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"kind_sub_type": kindSubType}))
		})
	}
}
