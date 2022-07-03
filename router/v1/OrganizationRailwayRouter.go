package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
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
			middleware.CheckJWT(),
			middleware.CheckPermission(),
			func(ctx *gin.Context) {
				uniqueCode := ctx.Param("unique_code")

				organizationRailway := (&model.OrganizationRailwayModel{
					BaseModel: model.BaseModel{
						Preloads: []string{
							"OrganizationParagraphs",
							"OrganizationParagraphs.OrganizationWorkshops",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationSections",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationWorkAreas",
							"OrganizationParagraphs.OrganizationWorkshops.OrganizationStations",
						},
					},
				}).FindOneByUniqueCode(uniqueCode)
				tool.ThrowErrorWhenIsEmpty(organizationRailway, model.OrganizationRailwayModel{}, "路局")

				ctx.JSON(tool.CorrectIns("").OK(gin.H{"organization_railway": organizationRailway}))
			})
	}
}
