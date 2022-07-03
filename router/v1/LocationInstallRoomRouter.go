package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type LocationInstallRoomRouter struct {
	Router    *gin.Engine
}

// Load 加载路由
func (cls *LocationInstallRoomRouter) Load() {
	r := cls.Router.Group("/api/v1/location")
	{
		r.GET("installRoom/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			locationInstallRoom := (&model.LocationInstallRoomModel{
				BaseModel: model.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			ctx.JSON(tool.CorrectIns("").OK(gin.H{"location_install_room": locationInstallRoom}))
		})
	}
}
