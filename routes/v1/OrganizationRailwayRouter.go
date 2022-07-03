package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type OrganizationRailwayRouter struct {
	Router *gin.Engine
}

func (cls *OrganizationRailwayRouter) Load() {
	r := cls.Router.Group("/api/v1/organization")
	{
		// 路局详情
		r.GET(
			"railway/:unique_code",
			middlewares.CheckJWT(),
			middlewares.CheckPermission(),
			func(ctx *gin.Context) {
				uniqueCode := ctx.Param("unique_code")

				organizationRailway := (&models.OrganizationRailwayModel{
					BaseModel: models.BaseModel{
						Preloads: []string{
							"OrganizationParagraphs",
							"OrganizationParagraphs.OrganizationWorkshops",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationSections",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationWorkAreas",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationStations",
						},
					},
				}).FindOneByUniqueCode(uniqueCode)
				tools.ThrowErrorWhenIsEmpty(organizationRailway, models.OrganizationRailwayModel{}, "路局")

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_railway": organizationRailway}))
			})
	}
}
