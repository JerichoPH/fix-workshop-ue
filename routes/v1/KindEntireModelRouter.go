package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

type KindEntireModelRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *KindEntireModelRouter) Load() {
	r := cls.Router.Group("/api/v1/kindEntireModel")
	{
		// 类型详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindEntireModel := (&models.KindEntireModel{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
					Preloads: []string{
						"KindCategory",
						"KindSubModels",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindEntireModel, models.KindEntireModel{}, "类型")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_entire_model": kindEntireModel}))
		})
	}
}
