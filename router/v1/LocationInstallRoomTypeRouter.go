package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type LocationInstallRoomTypeRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *LocationInstallRoomTypeRouter) Load() {
	r := cls.Router.Group("/api/v1/location")
	{
		// 机房类型详情
		r.GET("installRoomType/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			locationInstallRoomType := (&model.LocationInstallRoomTypeModel{
				BaseModel: model.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			ctx.JSON(tool.CorrectIns("").OK(gin.H{"location_install_room_type": locationInstallRoomType}))
		})
	}
}
