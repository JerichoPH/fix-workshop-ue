package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
)

type OrganizationStationRouter struct {
	Router *gin.Engine
}

func (cls *OrganizationStationRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/organization",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			station := (&models.OrganizationStationModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"OrganizationWorkshop",
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"station": station}))
		})
	}
}
