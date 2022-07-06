package v1

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrganizationLineRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationLineRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 线别详情
		r.GET("line/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationLine := (&models.OrganizationLineModel{
				BaseModel: models.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationLine, models.OrganizationLineModel{}, "线别")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_line": organizationLine}))
		})
	}
}
