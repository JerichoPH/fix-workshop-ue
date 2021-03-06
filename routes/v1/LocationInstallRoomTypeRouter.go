package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LocationInstallRoomTypeRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *LocationInstallRoomTypeRouter) Load() {
	r := cls.Router.Group("/api/v1/locationInstallRoomType")
	{
		// 机房类型详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			locationInstallRoomType := (&models.LocationInstallRoomTypeService{
				CTX:       ctx,
				MySqlConn: cls.MySqlConn,
				Preloads: []string{
					clause.Associations,
				},
			}).FindOneByUniqueCode(uniqueCode)
			ctx.JSON(tools.CorrectIns("").OK(gin.H{"location_install_room_type": locationInstallRoomType}))
		})
	}
}
