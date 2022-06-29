package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LocationInstallRoomRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *LocationInstallRoomRouter) Load() {
	r := cls.Router.Group("/api/v1/locationInstallRoom")
	{
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			locationInstallRoom := (&models.LocationInstallRoom{
				BaseModel: models.BaseModel{
					DB:       cls.MySqlConn,
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			ctx.JSON(tools.CorrectIns("").OK(gin.H{"location_install_room": locationInstallRoom}))
		})
	}
}
