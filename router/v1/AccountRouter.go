package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group(
		"/api/v1",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		r.GET(
			"account/:id",
			func(ctx *gin.Context) {
				id := tool.ThrowErrorWhenIsNotInt(ctx.Param("id"), "id必须填写整数")

				account := (&model.AccountModel{
					BaseModel: model.BaseModel{
						Preloads: []string{
							"AccountStatus",
						},
					},
				}).
					FindOneById(id)
				tool.ThrowErrorWhenIsEmpty(account, model.AccountModel{}, "用户")

				ctx.JSON(tool.CorrectIns("").OK(gin.H{"account": account}))
			})
	}
}
