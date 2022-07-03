package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
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

			organizationRailroadGradeCross := (&model.OrganizationRailroadGradeCrossModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						clause.Associations,
						"OrganizationWorkshopModel.OrganizationWorkshopTypeModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel",
						"OrganizationWorkshopModel.OrganizationParagraphModel.OrganizationRailwayModel",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(organizationRailroadGradeCross, model.OrganizationRailroadGradeCrossModel{}, "道口")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"organization_railroad_grade_cross": organizationRailroadGradeCross}))
		})
	}
}
