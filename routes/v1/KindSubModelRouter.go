package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

type KindSubModelRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *KindSubModelRouter) Load() {
	r := cls.Router.Group("/api/v1/kindSubModel")
	{
		// 型号详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindSubModel := (&models.KindSubModel{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
					Preloads: []string{
						"KindCategory",
						"KindEntireModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindSubModel, models.KindSubModel{}, "型号")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_sub_model": kindSubModel}))
		})
	}
}
