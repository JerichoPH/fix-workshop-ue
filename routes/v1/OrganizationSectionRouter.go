package v1

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type OrganizationSectionRouter struct {
	Router *gin.Engine
}

func (cls *OrganizationSectionRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 区间详情
		r.GET("section/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationSection := (&models.OrganizationSectionModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"OrganizationWorkshop",
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_section": organizationSection}))
		})
	}
}
