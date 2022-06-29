package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrganizationSectionRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationSectionRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationSection")
	{
		// 区间详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationSection := (&models.OrganizationSectionModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						clause.Associations,
						"OrganizationWorkshopModel.OrganizationWorkshopTypeModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel.OrganizationRailwayModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_section": organizationSection}))
		})
	}
}
