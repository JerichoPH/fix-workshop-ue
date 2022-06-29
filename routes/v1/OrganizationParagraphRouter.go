package v1

import (
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationParagraphRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
}

func (cls *OrganizationParagraphRouter) Load() {
	r := cls.Router.Group("/api/v1/organizationParagraph")
	{
		// 站段详情
		r.GET("/:unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			organizationParagraph := (&models.OrganizationParagraph{
				Preloads:  []string{clause.Associations},
				Selects:   []string{},
			}).FindOneByUniqueCode(cls.MySqlConn, uniqueCode)
			tools.ThrowErrorWhenIsEmpty(organizationParagraph, models.OrganizationParagraph{}, "站段")

			ctx.JSON(tools.CorrectIns("").OK(gin.H{"organization_paragraph": organizationParagraph}))
		})
	}
}
