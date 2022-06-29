package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationLineRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationLineRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationLine")
	{
		// 线别详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationLine := (&models.OrganizationLineModel{
				BaseModel: models.BaseModel{
					DB:       cls.MySqlConn,
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationLine, models.OrganizationLineModel{}, "线别")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_line": organizationLine}))
		})
	}
}
