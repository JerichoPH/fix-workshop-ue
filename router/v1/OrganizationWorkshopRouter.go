package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
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

			organizationWorkshop := (&model.OrganizationWorkshopModel{
				BaseModel: model.BaseModel{
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

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"organization_workshop": organizationWorkshop}))
		})
	}
}
