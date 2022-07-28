package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// LocationIndoorRoomRouter 室内上道位置机房路由
type LocationIndoorRoomRouter struct{}

// LocationIndoorRoomStoreForm 新建室内上道位置机房表单
type LocationIndoorRoomStoreForm struct {
	Sort                       int64  `form:"sort" json:"sort"`
	UniqueCode                 string `form:"unique_code" json:"unique_code"`
	Name                       string `form:"name" json:"name"`
	LocationIndoorRoomTypeUUID string `form:"location_indoor_room_type_uuid" json:"location_indoor_room_type_uuid"`
	LocationIndoorRoomType     models.LocationIndoorRoomTypeModel
	OrganizationStationUUID    string `form:"organization_station_uuid" json:"organization_station_uuid"`
	OrganizationStation        models.OrganizationStationModel
	OrganizationSectionUUID    string `form:"organization_section_uuid" json:"organization_section_uuid"`
	OrganizationSection        models.OrganizationSectionModel
	OrganizationCenterUUID     string `form:"organization_center_uuid" json:"organization_center_uuid"`
	OrganizationCenter         models.OrganizationCenterModel
}

// ShouldBind 绑定表单
//  @receiver LocationIndoorRoomStoreForm
//  @param ctx
//  @return LocationIndoorRoomStoreForm
func (cls LocationIndoorRoomStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorRoomStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("机房代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("机房名称必填")
	}
	if cls.LocationIndoorRoomTypeUUID == "" {
		wrongs.PanicValidate("机房类型必选")
	}
	ret = models.Init(models.LocationIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationIndoorRoomTypeUUID}).
		Prepare().
		First(&cls.LocationIndoorRoomType)
	wrongs.PanicWhenIsEmpty(ret, "机房类型")
	if cls.OrganizationStationUUID == "" && cls.OrganizationSectionUUID == "" && cls.OrganizationCenterUUID == "" {
		wrongs.PanicValidate("归属单位必选")
	}
	if cls.OrganizationStationUUID != "" {
		ret = models.Init(models.OrganizationStationModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationStationUUID}).
			Prepare().
			First(&cls.OrganizationStation)
		wrongs.PanicWhenIsEmpty(ret, "所属战场")
	}
	if cls.OrganizationSectionUUID != "" {
		ret = models.Init(models.OrganizationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationSection}).
			Prepare().
			First(&cls.OrganizationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属区间")
	}
	if cls.OrganizationCenterUUID != "" {
		ret = models.Init(models.OrganizationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationSection}).
			Prepare().
			First(&cls.OrganizationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属中心")
	}

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorRoomRouter
//  @param router
func (LocationIndoorRoomRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorRoom", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationIndoorRoomModel
			)

			// 表单
			form := (&LocationIndoorRoomStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 新建
			locationIndoorRoom := &models.LocationIndoorRoomModel{
				BaseModel:              models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:             form.UniqueCode,
				Name:                   form.Name,
				LocationIndoorRoomType: form.LocationIndoorRoomType,
				OrganizationStation:    form.OrganizationStation,
			}
			if ret = models.Init(models.LocationIndoorRoomModel{}).GetSession().Create(&locationIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_room": locationIndoorRoom}))
		})

		// 删除
		r.DELETE("indoorRoom/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorRoom models.LocationIndoorRoomModel
			)

			// 查询
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			// 删除
			if ret := models.Init(models.LocationIndoorRoomModel{}).GetSession().Delete(&locationIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorRoom/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationIndoorRoom, repeat models.LocationIndoorRoomModel
			)

			// 表单
			form := (&LocationIndoorRoomStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 查询
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			// 编辑
			locationIndoorRoom.BaseModel.Sort = form.Sort
			locationIndoorRoom.UniqueCode = form.UniqueCode
			locationIndoorRoom.Name = form.Name
			locationIndoorRoom.LocationIndoorRoomType = form.LocationIndoorRoomType
			locationIndoorRoom.OrganizationStation = form.OrganizationStation
			if ret = models.Init(models.LocationIndoorRoomModel{}).GetSession().Save(&locationIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_room": locationIndoorRoom}))
		})

		// 详情
		r.GET("indoorRoom/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorRoom models.LocationIndoorRoomModel
			)
			ret = models.Init(models.LocationIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_room": locationIndoorRoom}))
		})

		// 列表
		r.GET("indoorRoom", func(ctx *gin.Context) {
			var locationIndoorRooms []models.LocationIndoorRoomModel
			models.Init(models.LocationIndoorRoomModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorRooms)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_rooms": locationIndoorRooms}))
		})
	}
}
