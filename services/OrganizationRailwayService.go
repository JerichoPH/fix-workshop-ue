package services

import (
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationRailwayService struct{}

// PrepareByAccountOrganizationLevel 根据用户级别设置路局条件
func (OrganizationRailwayService) PrepareByAccountOrganizationLevel(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	accountOrganizationLevel, exists := ctx.Get(tools.AccountOrganizationLevel)
	if !exists || accountOrganizationLevel != tools.OrganizationLevelRailway {
		query = query.Where("false")
	} else {
		query = query.Where("uuid = ?", tools.AccountOrganizationRailwayUuid)
	}

	return query
}
