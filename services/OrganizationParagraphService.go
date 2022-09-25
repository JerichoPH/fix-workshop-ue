package services

import (
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationParagraphService struct{}

// PrepareByAccountOrganizationLevel 根据用户级别设置站段条件
func (OrganizationParagraphService) PrepareByAccountOrganizationLevel(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationLevel)

	if !exists || accountOrganizationLevel == "" {
		query = query.Where("false")
	} else {
		if accountOrganizationLevel == tools.OrganizationLevelRailway {
			accountOrganizationRailwayUuid, _ := ctx.Get(tools.AccountOrganizationRailwayUuid)
			query = query.Where("organization_railway_uuid = ?", accountOrganizationRailwayUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelParagraph {
			accountOrganizationParagraphUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)
			query = query.Where("uuid = ?", accountOrganizationParagraphUuid)
		}
	}

	return query
}
