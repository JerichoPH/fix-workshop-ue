package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationWorkAreaRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationWorkAreaRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationWorkArea")
	{
		// 工区详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationWorkArea := (&models.OrganizationWorkArea{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
					Preloads: []string{
						clause.Associations,
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationWorkArea, models.OrganizationWorkArea{}, "工区")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_work_area": organizationWorkArea}))
		})
	}
}
