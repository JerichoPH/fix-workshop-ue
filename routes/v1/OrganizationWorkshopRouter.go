package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationWorkshopRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationWorkshopRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationWorkshop")
	{
		// 车间详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationWorkshop := (&models.OrganizationWorkshopService{
				CTX:       ctx,
				MySqlConn: cls.MySqlConn,
				Preloads: []string{
					clause.Associations,
					"OrganizationWorkshopType",
					"OrganizationParagraph.OrganizationRailway",
				},
				Selects: []string{},
				Omits:   []string{},
			}).FindOneByUniqueCode(uniqueCode)

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_workshop": organizationWorkshop}))
		})
	}
}
