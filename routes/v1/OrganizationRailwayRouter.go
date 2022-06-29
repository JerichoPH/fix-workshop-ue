package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRailwayRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationRailwayRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationRailway")
	{
		// 路局详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationRailway := (&models.OrganizationRailwayModel{
				BaseModel: models.BaseModel{
					DB:       cls.MySqlConn,
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationRailway, models.OrganizationRailwayModel{}, "路局")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_railway": organizationRailway}))
		})
	}
}
