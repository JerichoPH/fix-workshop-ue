package v1

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) Load(engine *gin.Engine) {
	// 用户与权鉴
	(&AuthorizationRouter{}).Load(engine)       // 权鉴
	(&AccountRouter{}).Load(engine)             // 用户                                                                                                                                                          // 用户
	(&RbacRoleRouter{}).Load(engine)            // 角色
	(&RbacPermissionGroupRouter{}).Load(engine) // 权限分组
	(&RbacPermissionRouter{}).Load(engine)      // 权限
	(&MenuRouter{}).Load(engine)                // 菜单

	// 组织机构
	(&OrganizationLineRouter{}).Load(engine)               // 线别
	(&OrganizationRailwayRouter{}).Load(engine)            // 路局
	(&OrganizationParagraphRouter{}).Load(engine)          // 站段
	(&OrganizationWorkshopTypeRouter{}).Load(engine)       // 车间类型
	(&OrganizationWorkshopRouter{}).Load(engine)           // 车间
	(&OrganizationWorkAreaTypeRouter{}).Load(engine)       // 工区类型
	(&OrganizationWorkAreaRouter{}).Load(engine)           // 工区
	(&OrganizationSectionRouter{}).Load(engine)            // 区间
	(&OrganizationCenterRouter{}).Load(engine)             // 中心
	(&OrganizationRailroadGradeCrossRouter{}).Load(engine) // 道口
	(&OrganizationStationRouter{}).Load(engine)            // 站场

	// 仓储位置
	(&LocationDepotStorehouseRouter{}).Load(engine) // 仓库
	(&LocationDepotSectionRouter{}).Load(engine)    // 仓库区域
	(&LocationDepotRowRouter{}).Load(engine)        // 仓库排
	(&LocationDepotCabinetRouter{}).Load(engine)    // 仓库柜架
	(&LocationDepotTierRouter{}).Load(engine)       // 仓库柜架层
	(&LocationDepotCellRouter{}).Load(engine)       // 仓库柜架格位

	// 室内上道位置
	(&LocationIndoorRoomTypeRouter{}).Load(engine) // 室内上道位置机房类型
	(&LocationIndoorRoomRouter{}).Load(engine)     // 室内上道位置机房
	(&LocationIndoorRowRouter{}).Load(engine)      // 室内上道位置排
	(&LocationIndoorCabinetRouter{}).Load(engine)  // 室内上道位置柜架
	(&LocationIndoorTierRouter{}).Load(engine)     // 室内上道位置柜架层
	(&LocationIndoorCellRouter{}).Load(engine)     // 室内上道位置柜架格位
}
