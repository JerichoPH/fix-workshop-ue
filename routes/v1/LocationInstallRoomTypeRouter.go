package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type LocationInstallRoomTypeRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *LocationInstallRoomTypeRouter) Load() {
	r := cls.Router.Group("/api/v1/locationInstallRoomType")
	{
		// 机房类型详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			locationInstallRoomType := (&models.LocationInstallRoomTypeModel{
				BaseModel: models.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			ctx.JSON(tools.CorrectIns("").OK(gin.H{"location_install_room_type": locationInstallRoomType}))
		})
	}
}
