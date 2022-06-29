package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	gcasbin "github.com/maxwellhertz/gin-casbin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRouter struct {
	Router     *gin.Engine
	MySqlConn  *gorm.DB
	MsSqlConn  *gorm.DB
	AppConfig  *ini.File
	DBConfig   *ini.File
	AuthCasbin *gcasbin.CasbinMiddleware
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group("/api/v1/account")
	{
		r.GET(
			"/:id",
			cls.AuthCasbin.RequiresPermissions([]string{"account:show"}, gcasbin.WithLogic(gcasbin.AND)),
			//middlewares.JwtCheck(cls.MySqlConn),
			func(ctx *gin.Context) {
				id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "id必须填写整数")

				account := (&models.AccountModel{
					BaseModel: models.BaseModel{
						DB:       cls.MySqlConn,
						Preloads: []string{clause.Associations},
					},
				}).
					FindOneById(id)
				tools.ThrowErrorWhenIsEmpty(account, models.AccountModel{}, "用户")

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account": account}))
			})
	}
}
