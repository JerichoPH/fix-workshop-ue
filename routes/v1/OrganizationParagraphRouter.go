package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrganizationParagraphRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationParagraphRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 站段详情
		r.GET("paragraph/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationParagraph := (&models.OrganizationParagraphModel{
				BaseModel: models.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationParagraph, models.OrganizationParagraphModel{}, "站段")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_paragraph": organizationParagraph}))
		})
	}
}
