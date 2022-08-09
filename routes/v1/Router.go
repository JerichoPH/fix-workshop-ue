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
	(&LocationLineRouter{}).Load(engine)                   // 线别
	(&OrganizationRailwayRouter{}).Load(engine)            // 路局
	(&OrganizationParagraphRouter{}).Load(engine)          // 站段
	(&OrganizationWorkshopTypeRouter{}).Load(engine)       // 车间类型
	(&OrganizationWorkshopRouter{}).Load(engine)           // 车间
	(&OrganizationWorkAreaTypeRouter{}).Load(engine)       // 工区类型
	(&OrganizationWorkAreaRouter{}).Load(engine)           // 工区
	(&OrganizationSectionRouter{}).Load(engine)            // 区间
	(&LocationCenterRouter{}).Load(engine)                 // 中心
	(&OrganizationRailroadGradeCrossRouter{}).Load(engine) // 道口
	(&OrganizationStationRouter{}).Load(engine)            // 站场

	// 仓储位置
	(&PositionDepotStorehouseRouter{}).Load(engine) // 仓库
	(&PositionDepotSectionRouter{}).Load(engine)    // 仓库区域
	(&PositionDepotRowRouter{}).Load(engine)        // 仓库排
	(&PositionDepotCabinetRouter{}).Load(engine)    // 仓库柜架
	(&PositionDepotTierRouter{}).Load(engine)       // 仓库柜架层
	(&PositionDepotCellRouter{}).Load(engine)       // 仓库柜架格位

	// 室内上道位置
	(&PositionIndoorRoomTypeRouter{}).Load(engine) // 室内上道位置机房类型
	(&LocationIndoorRoomRouter{}).Load(engine)     // 室内上道位置机房
	(&PositionIndoorRowRouter{}).Load(engine)      // 室内上道位置排
	(&PositionIndoorCabinetRouter{}).Load(engine)  // 室内上道位置柜架
	(&PositionIndoorTierRouter{}).Load(engine)     // 室内上道位置柜架层
	(&PositionIndoorCellRouter{}).Load(engine)     // 室内上道位置柜架格位

	// 种类型
	(&KindCategoryRouter{}).Load(engine)   // 种类
	(&KindEntireTypeRouter{}).Load(engine) // 类型
	(&KindSubTypeRouter{}).Load(engine)    // 型号
}
