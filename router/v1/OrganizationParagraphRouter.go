package v1

import (
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
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

			organizationParagraph := (&model.OrganizationParagraphModel{
				BaseModel: model.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tool.ThrowErrorWhenIsEmpty(organizationParagraph, model.OrganizationParagraphModel{}, "站段")

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"organization_paragraph": organizationParagraph}))
		})
	}
}
