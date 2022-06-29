package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrganizationRailwayRouter struct {
	Router    *gin.Engine
}

func (cls *OrganizationRailwayRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationRailway")
	{
		// 路局详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationRailway := (&models.OrganizationRailwayModel{
				BaseModel: models.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUniqueCode(uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationRailway, models.OrganizationRailwayModel{}, "路局")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_railway": organizationRailway}))
		})
	}
}
