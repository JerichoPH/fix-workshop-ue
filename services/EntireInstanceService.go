package services

import (
	"fix-workshop-ue/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EntireInstanceService struct{}

// GetAccountOrganizationLevel 获取用户级别
func (EntireInstanceService) GetAccountOrganizationLevel(ctx *gin.Context) {
	fmt.Println(ctx.Get(tools.AccountOrganizationLevel))
	fmt.Println(ctx.Get(tools.AccountOrganizationWorkAreaProfessionUniqueCode))
}

// PrepareByAccountOrganizationLevel 根据用户所属单位设置搜索级别
func (EntireInstanceService) PrepareByAccountOrganizationLevel(ctx *gin.Context, dbSession *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationLevel)
	if !exists || accountOrganizationLevel == "" {
		dbSession.Where("ture = false")
	} else {
		if accountOrganizationLevel == tools.OrganizationLevelRailway {
			accountOrganizationRailwayUuid, _ := ctx.Get(tools.AccountOrganizationRailwayUuid)
			dbSession.Where("belong_to_organization_railway_uuid", accountOrganizationRailwayUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelParagraph {
			accountOrganizationParagraphUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)
			dbSession.Where("belong_to_organization_paragraph_uuid", accountOrganizationParagraphUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelWorkshop {
			accountOrganizationWorkshopUuid, _ := ctx.Get(tools.AccountOrganizationWorkshopUuid)
			dbSession.Where("belong_to_organization_workshop_uuid", accountOrganizationWorkshopUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelWorkArea {
			accountOrganizationWorkAreaUuid, _ := ctx.Get(tools.AccountOrganizationWorkAreaUuid)
			dbSession = dbSession.Where("belong_to_organization_work_area_uuid", accountOrganizationWorkAreaUuid)
		}
	}

	return dbSession
}
