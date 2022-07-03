package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type KindEntireModelRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *KindEntireModelRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/kind",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		// 类型详情
		r.GET("entireType/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindEntireType := (&model.KindEntireTypeModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"KindCategory",
						"KindSubTypes",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(kindEntireType, model.KindEntireTypeModel{}, "类型")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"kind_entire_type": kindEntireType}))
		})
	}
}
