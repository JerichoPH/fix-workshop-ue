package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRailroadGradeCrossRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationRailroadGradeCrossRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationRailroadGradeCross")
	{
		// 道口详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationRailroadGradeCross := (&models.OrganizationRailroadGradeCross{
				Preloads: []string{
					clause.Associations,
					"OrganizationWorkshop.OrganizationWorkshopType",
					"OrganizationWorkshop.OrganizationParagraph",
					"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
				},
				Selects: []string{},
				Omits:   []string{},
			}).FindOneByUniqueCode(cls.MySqlConn, uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationRailroadGradeCross, models.OrganizationRailroadGradeCross{}, "道口")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_railroad_grade_cross": organizationRailroadGradeCross}))
		})
	}
}
