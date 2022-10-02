package v1

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) Load(engine *gin.Engine) {
	// 用户与权鉴
	new(AuthorizationRouter).Load(engine)       // 权鉴
	new(AccountRouter).Load(engine)             // 用户                                                                                                                                                          // 用户
	new(RbacRoleRouter).Load(engine)            // 角色
	new(RbacPermissionGroupRouter).Load(engine) // 权限分组
	new(RbacPermissionRouter).Load(engine)      // 权限
	new(MenuRouter).Load(engine)                // 菜单

	// 组织机构
	new(OrganizationRailwayRouter).Load(engine)            // 路局
	new(OrganizationParagraphRouter).Load(engine)          // 站段
	new(OrganizationWorkshopTypeRouter).Load(engine)       // 车间类型
	new(OrganizationWorkshopRouter).Load(engine)           // 车间
	new(OrganizationWorkAreaTypeRouter).Load(engine)       // 工区类型
	new(OrganizationWorkAreaProfessionRouter).Load(engine) //  工区专业
	new(OrganizationWorkAreaRouter).Load(engine)           // 工区

	// 使用处所
	new(LocationLineRouter).Load(engine)               // 线别
	new(LocationStationRouter).Load(engine)            // 站场
	new(LocationSectionRouter).Load(engine)            // 区间
	new(LocationCenterRouter).Load(engine)             // 中心
	new(LocationRailroadGradeCrossRouter).Load(engine) // 道口

	// 存放位置-仓储
	new(PositionDepotStorehouseRouter).Load(engine) // 仓库
	new(PositionDepotSectionRouter).Load(engine)    // 仓库区域
	new(PositionDepotRowRouter).Load(engine)        // 仓库排
	new(PositionDepotCabinetRouter).Load(engine)    // 仓库柜架
	new(PositionDepotTierRouter).Load(engine)       // 仓库柜架层
	new(PositionDepotCellRouter).Load(engine)       // 仓库柜架格位

	// 使用位置-室内上道位置
	new(PositionIndoorRoomTypeRouter).Load(engine) // 室内上道位置机房类型
	new(PositionIndoorRoomRouter).Load(engine)     // 室内上道位置机房
	new(PositionIndoorRowRouter).Load(engine)      // 室内上道位置排
	new(PositionIndoorCabinetRouter).Load(engine)  // 室内上道位置柜架
	new(PositionIndoorTierRouter).Load(engine)     // 室内上道位置柜架层
	new(PositionIndoorCellRouter).Load(engine)     // 室内上道位置柜架格位

	// 种类型
	new(KindCategoryRouter).Load(engine)   // 种类
	new(KindEntireTypeRouter).Load(engine) // 类型
	new(KindSubTypeRouter).Load(engine)    // 型号

	// 同步
	new(SyncRouter).Load(engine) // 同步
}
