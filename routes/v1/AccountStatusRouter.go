package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type AccountStatusRouter struct {
	Router *gin.Engine
}

func (cls *AccountStatusRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/accountStatus",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// GET 用户状态详情
		r.GET(
			"/:unique_code",

			func(ctx *gin.Context) {
				uniqueCode := ctx.Param("unique_code")
				accountStatus := (&models.AccountStatusModel{
					BaseModel: models.BaseModel{
						Preloads: []string{},
					},
				}).FindOneByUniqueCode(uniqueCode)

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account_status": accountStatus}))
			})
	}
}
