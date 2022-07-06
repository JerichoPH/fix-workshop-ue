package v1

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type OrganizationWorkAreaRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationWorkAreaRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 工区详情
		r.GET("workArea/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationWorkArea := (&models.OrganizationWorkAreaModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"OrganizationWorkshop",
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationWorkArea, models.OrganizationWorkAreaModel{}, "工区")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_work_area": organizationWorkArea}))
		})
	}
}
