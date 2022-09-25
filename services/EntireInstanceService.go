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

// PrepareByAccountOrganizationLevel 根据用户级别设置器材条件
func (EntireInstanceService) PrepareByAccountOrganizationLevel(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationLevel)
	if !exists || accountOrganizationLevel == "" {
		query = query.Where("false")
	} else {
		if accountOrganizationLevel == tools.OrganizationLevelRailway {
			accountOrganizationRailwayUuid, _ := ctx.Get(tools.AccountOrganizationRailwayUuid)
			query = query.Where("belong_to_organization_railway_uuid", accountOrganizationRailwayUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelParagraph {
			accountOrganizationParagraphUuid, _ := ctx.Get(tools.AccountOrganizationParagraphUuid)
			query = query.Where("belong_to_organization_paragraph_uuid", accountOrganizationParagraphUuid)
		}
		if accountOrganizationLevel == tools.OrganizationLevelWorkshop {
			accountOrganizationWorkshopUuid, _ := ctx.Get(tools.AccountOrganizationWorkshopUuid)
			accountOrganizationWorkshopTypeUniqueCode, _ := ctx.Get(tools.AccountOrganizationWorkshopTypeUniqueCode)
			if accountOrganizationWorkshopTypeUniqueCode == tools.OrganizationLevelSceneWorkshop {
				query = query.Where("use_place_current_organization_workshop_uuid", accountOrganizationWorkshopUuid)
			} else {
				query = query.Where("belong_to_organization_workshop_uuid = ?", accountOrganizationWorkshopUuid)
			}
		}
		if accountOrganizationLevel == tools.OrganizationLevelWorkArea {
			accountOrganizationWorkAreaUuid, _ := ctx.Get(tools.AccountOrganizationWorkAreaUuid)
			accountOrganizationWorkAreaTypeUniqueCode, _ := ctx.Get(tools.AccountOrganizationWorkAreaTypeUniqueCode)
			if accountOrganizationWorkAreaTypeUniqueCode == tools.OrganizationLevelSceneWorkArea {
				query = query.Where("use_place_current_organization_work_area_uuid", accountOrganizationWorkAreaUuid)
			} else {
				query = query.Where("belong_to_organization_work_area_uuid = ?", accountOrganizationWorkAreaUuid)
			}
		}
	}

	return query
}
