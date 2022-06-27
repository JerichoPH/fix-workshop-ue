package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group("/api/v1/account")
	{
		r.GET("/:id", func(ctx *gin.Context) {
			id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "id必须填写整数")

			account := (&models.Account{
				Preloads: []string{clause.Associations},
			}).FindOneById(cls.MySqlConn, id)
			tools.ThrowErrorWhenIsEmpty(account, models.Account{}, "用户")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"account": account}))
		})
	}
}
