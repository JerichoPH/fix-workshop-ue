package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
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

			organizationWorkArea := (&model.OrganizationWorkAreaModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"OrganizationWorkshop",
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(organizationWorkArea, model.OrganizationWorkAreaModel{}, "工区")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"organization_work_area": organizationWorkArea}))
		})
	}
}
