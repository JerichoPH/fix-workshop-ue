package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group(
		"/api/v1",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		r.GET(
			"account/:id",
			func(ctx *gin.Context) {
				id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "id必须填写整数")

				account := (&models.AccountModel{
					BaseModel: models.BaseModel{
						Preloads: []string{
							"AccountStatus",
						},
					},
				}).
					FindOneById(id)
				tools.ThrowErrorWhenIsEmpty(account, models.AccountModel{}, "用户")

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account": account}))
			})
	}
}
