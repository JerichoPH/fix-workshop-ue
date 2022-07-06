package v1

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type OrganizationWorkshopRouter struct {
	Router *gin.Engine
}

func (cls *OrganizationWorkshopRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 车间详情
		r.GET("workshop/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationWorkshop := (&models.OrganizationWorkshopModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"OrganizationParagraph",
						"OrganizationParagraph.OrganizationRailway",
						"OrganizationParagraphs",
						"OrganizationSections",
						"OrganizationWorkAreas",
						"OrganizationStations",
						"OrganizationWorkshopTypeModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_workshop": organizationWorkshop}))
		})
	}
}
