package services

import (
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationWorkAreaService struct{}

// PrepareByAccountOrganizationLevel 根据用户级别设置工区条件
func (OrganizationWorkAreaService) PrepareByAccountOrganizationLevel(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationWorkAreaUuid)
	if !exists || accountOrganizationLevel == "" {
		query = query.Where("false")
	} else {
		if accountOrganizationLevel == tools.OrganizationLevelRailway {
			accountOrganizationRailwayUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)
			query = query.
				Joins("join organization_workshops w on organization_workshop_uuid = w.uuid").
				Joins("join organization_paragraphs w on w.organization_paragraph_uuid = p.uuid").
				Joins("join organization_railways r on p.organization_railway_uuid = r.uuid").
				Where("r.uuid = ?", accountOrganizationRailwayUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelParagraph {
			accountOrganizationParagraphUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)
			query = query.
				Joins("join organization_workshops w on organization_workshop_uuid = w.uuid").
				Joins("join organization_paragraphs w on w.organization_paragraph_uuid = p.uuid").
				Where("p.uuid = ?", accountOrganizationParagraphUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelWorkshop {
			accountOrganizationWorkshopUuid, _ := ctx.Get(tools.AccountOrganizationWorkshopUuid)
			query = query.Where("organization_workshop_uuid = ?", accountOrganizationWorkshopUuid)
		}
	}

	return query
}
