package v1

import (
	"fix-workshop-ue/middleware"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
)

type OrganizationStationRouter struct {
	Router *gin.Engine
}

func (cls *OrganizationStationRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/organization",
		middleware.CheckJWT(),
		middleware.CheckPermission(),
	)
	{
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			station := (&model.OrganizationStationModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"OrganizationWorkshop",
						"OrganizationWorkshop.OrganizationWorkshopType",
						"OrganizationWorkshop.OrganizationParagraph",
						"OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
					},
				},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tool.CorrectIns("").OK(gin.H{"station": station}))
		})
	}
}
