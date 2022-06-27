package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

type KindCategoryRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

// Load 加载路由
func (cls *KindCategoryRouter) Load() {
	r := cls.Router.Group("/api/v1/kindCategory")
	{
		// 种类详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			kindCategory := (&models.KindCategoryService{
				CTX:       ctx,
				MySqlConn: cls.MySqlConn,
				Preloads: []string{
					"KindEntireModels",
					"KindEntireModels.KindSubModels",
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(kindCategory, models.KindCategory{}, "种类")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"kind_category": kindCategory}))
		})
	}
}
