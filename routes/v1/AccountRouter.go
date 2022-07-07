package v1

import (
	"fix-workshop-ue/errors"
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

				var account models.AccountModel
				if ret := (&models.BaseModel{
					Preloads: []string{"AccountStatus"},
					Wheres:   map[string]interface{}{"id": id},
				}).
					Prepare().
					First(&account); ret.Error != nil {
					panic(errors.ThrowEmpty("用户不存在"))
				}

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account": account}))
			})
	}
}
