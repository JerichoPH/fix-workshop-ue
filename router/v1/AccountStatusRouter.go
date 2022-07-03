package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type AccountStatusRouter struct {
	Router *gin.Engine
}

func (cls *AccountStatusRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/accountStatus",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		// GET 用户状态详情
		r.GET(
			"/:unique_code",

			func(ctx *gin.Context) {
				uniqueCode := ctx.Param("unique_code")
				accountStatus := (&model.AccountStatusModel{
					BaseModel: model.BaseModel{
						Preloads: []string{},
					},
				}).FindOneByUniqueCode(uniqueCode)

				ctx.JSON(tool.CorrectIns("").OK(gin.H{"account_status": accountStatus}))
			})
	}
}
