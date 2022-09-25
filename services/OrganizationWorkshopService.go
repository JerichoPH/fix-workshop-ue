package services

import (
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationWorkshopService struct{}

func (OrganizationWorkshopService) PrepareByAccountOrganizationLevel(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationLevel)

	if !exists || accountOrganizationLevel == "" {
		query = query.Where("false")
	} else {
		if accountOrganizationLevel == tools.OrganizationLevelRailway {
			accountOrganizationRailwayUuid, _ := ctx.Get(tools.AccountOrganizationRailwayUuid)

			query = query.
				Joins("join organization_paragraphs p on organization_paragraph_uuid = p.uuid").
				Joins("join organization_railways r on p.organization_railway_uuid = r.uuid").
				Where("r.uuid = ?", accountOrganizationRailwayUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelParagraph {
			accountOrganizationParagraphUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)

			query = query.Where("organization_paragraph_uuid = ?", accountOrganizationParagraphUuid)
		}
	}

	return query
}
