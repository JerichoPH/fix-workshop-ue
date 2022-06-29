package v1

import (
	"github.com/gin-gonic/gin"
)

type V1Router struct {
	Router     *gin.Engine
}

func (cls *V1Router) Load() {
	(&AuthorizationRouter{Router: cls.Router}).Load() // 权鉴
	(&AccountRouter{Router: cls.Router}).Load()                                                                                                                                                           // 用户
	//(&AccountStatusRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load() // 用户状态
	//
	//// 种类型
	//(&KindCategoryRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()    // 种类
	//(&KindEntireModelRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load() // 类型
	//(&KindSubModelRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()    // 型号
	//
	//// 组织机构
	//(&OrganizationLineRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()               // 站场
	//(&OrganizationRailwayRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()            // 路局
	//(&OrganizationParagraphRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()          // 站段
	//(&OrganizationWorkshopRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()           // 车间
	//(&OrganizationSectionRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()            // 区间
	//(&OrganizationRailroadGradeCrossRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load() // 道口
	//(&OrganizationWorkAreaRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()           // 工区
	//(&OrganizationStationRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()            // 站场
	//
	//// 器材
	//(&EntireInstanceRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load() //器材
	//
	//// 上道位置
	//(&LocationInstallRoomRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load()     // 机房
	//(&LocationInstallRoomTypeRouter{Router: cls.Router, MySqlConn: cls.MySqlConn, MsSqlConn: cls.MsSqlConn, AppConfig: cls.AppConfig, DBConfig: cls.DBConfig}).Load() // 机房类型
}
