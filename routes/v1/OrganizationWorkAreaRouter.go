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

			organizationWorkArea := (&models.OrganizationWorkAreaModel{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
					Preloads: []string{
						clause.Associations,
						"OrganizationWorkshopModel.OrganizationWorkshopTypeModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel.OrganizationRailwayModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationWorkArea, models.OrganizationWorkAreaModel{}, "工区")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_work_area": organizationWorkArea}))
		})
	}
}
