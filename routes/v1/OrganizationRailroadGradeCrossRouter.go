package v1

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrganizationRailroadGradeCrossRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationRailroadGradeCrossRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 道口详情
		r.GET("railroadGradeCross/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationRailroadGradeCross := (&models.OrganizationRailroadGradeCrossModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						clause.Associations,
						"OrganizationWorkshopModel.OrganizationWorkshopTypeModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel.OrganizationRailwayModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationRailroadGradeCross, models.OrganizationRailroadGradeCrossModel{}, "道口")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_railroad_grade_cross": organizationRailroadGradeCross}))
		})
	}
}
