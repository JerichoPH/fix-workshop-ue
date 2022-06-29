package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type AccountRouter struct {
	Router     *gin.Engine
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group("/api/v1/account")
	{
		r.GET(
			"/:id",
			//middlewares.JwtCheck(cls.MySqlConn),
			func(ctx *gin.Context) {
				id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "id必须填写整数")

				account := (&models.AccountModel{
					BaseModel: models.BaseModel{
						Preloads: []string{clause.Associations},
					},
				}).
					FindOneById(id)
				tools.ThrowErrorWhenIsEmpty(account, models.AccountModel{}, "用户")

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account": account}))
			})
	}
}
