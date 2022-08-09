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
	PositionIndoorRoomTypeUUID string `form:"position_indoor_room_type_uuid" json:"position_indoor_room_type_uuid"`
	PositionIndoorRoomType     models.PositionIndoorRoomTypeModel
	LocationStationUUID        string `form:"location_station_uuid" json:"location_station_uuid"`
	LocationStation            models.LocationStationModel
	LocationSectionUUID        string `form:"location_section_uuid" json:"location_section_uuid"`
	LocationSection            models.LocationSectionModel
	LocationCenterUUID         string `form:"location_center_uuid" json:"location_center_uuid"`
	LocationCenter             models.LocationCenterModel
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
	if cls.PositionIndoorRoomTypeUUID == "" {
		wrongs.PanicValidate("机房类型必选")
	}
	ret = models.Init(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorRoomTypeUUID}).
		Prepare().
		First(&cls.PositionIndoorRoomType)
	wrongs.PanicWhenIsEmpty(ret, "机房类型")
	if cls.LocationStationUUID == "" && cls.LocationSectionUUID == "" && cls.LocationCenterUUID == "" {
		wrongs.PanicValidate("归属单位必选")
	}
	if cls.LocationStationUUID != "" {
		ret = models.Init(models.LocationStationModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationStationUUID}).
			Prepare().
			First(&cls.LocationStation)
		wrongs.PanicWhenIsEmpty(ret, "所属战场")
	}
	if cls.LocationSectionUUID != "" {
		ret = models.Init(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationSection}).
			Prepare().
			First(&cls.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属区间")
	}
	if cls.LocationCenterUUID != "" {
		ret = models.Init(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationSection}).
			Prepare().
			First(&cls.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属中心")
	}

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorRoomRouter
//  @param router
func (LocationIndoorRoomRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorRoom",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionIndoorRoomModel
			)

			// 表单
			form := (&LocationIndoorRoomStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 新建
			positionIndoorRoom := &models.PositionIndoorRoomModel{
				BaseModel:              models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:             form.UniqueCode,
				Name:                   form.Name,
				PositionIndoorRoomType: form.PositionIndoorRoomType,
				LocationStation:        form.LocationStation,
			}
			if ret = models.Init(models.PositionIndoorRoomModel{}).GetSession().Create(&positionIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_indoor_room": positionIndoorRoom}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorRoom models.PositionIndoorRoomModel
			)

			// 查询
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			// 删除
			if ret := models.Init(models.PositionIndoorRoomModel{}).GetSession().Delete(&locationIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				positionIndoorRoom, repeat models.PositionIndoorRoomModel
			)

			// 表单
			form := (&LocationIndoorRoomStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 查询
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			// 编辑
			positionIndoorRoom.BaseModel.Sort = form.Sort
			positionIndoorRoom.UniqueCode = form.UniqueCode
			positionIndoorRoom.Name = form.Name
			positionIndoorRoom.PositionIndoorRoomType = form.PositionIndoorRoomType
			positionIndoorRoom.LocationStation = form.LocationStation
			if ret = models.Init(models.PositionIndoorRoomModel{}).GetSession().Save(&positionIndoorRoom); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_indoor_room": positionIndoorRoom}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				positionIndoorRoom models.PositionIndoorRoomModel
			)
			ret = models.Init(models.PositionIndoorRoomModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorRoom)
			wrongs.PanicWhenIsEmpty(ret, "机房")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_room": positionIndoorRoom}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionIndoorRooms []models.PositionIndoorRoomModel
			models.Init(models.PositionIndoorRoomModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionIndoorRooms)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_rooms": positionIndoorRooms}))
		})
	}
}
