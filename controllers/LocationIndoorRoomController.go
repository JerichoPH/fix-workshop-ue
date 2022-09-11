package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionIndoorRoomController struct{}

// PositionIndoorRoomStoreForm 新建室内上道位置机房表单
type PositionIndoorRoomStoreForm struct {
	Sort                       int64  `form:"sort" json:"sort"`
	UniqueCode                 string `form:"unique_code" json:"unique_code"`
	Name                       string `form:"name" json:"name"`
	PositionIndoorRoomTypeUuid string `form:"position_indoor_room_type_uuid" json:"position_indoor_room_type_uuid"`
	PositionIndoorRoomType     models.PositionIndoorRoomTypeModel
	LocationStationUuid        string `form:"location_station_uuid" json:"location_station_uuid"`
	LocationStation            models.LocationStationModel
	LocationSectionUuid        string `form:"location_section_uuid" json:"location_section_uuid"`
	LocationSection            models.LocationSectionModel
	LocationCenterUuid         string `form:"location_center_uuid" json:"location_center_uuid"`
	LocationCenter             models.LocationCenterModel
}

// ShouldBind 绑定表单
//  @receiver PositionIndoorRoomStoreForm
//  @param ctx
//  @return PositionIndoorRoomStoreForm
func (cls PositionIndoorRoomStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRoomStoreForm {
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
	if cls.PositionIndoorRoomTypeUuid == "" {
		wrongs.PanicValidate("机房类型必选")
	}
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorRoomTypeUuid}).
		PrepareByDefault().
		First(&cls.PositionIndoorRoomType)
	wrongs.PanicWhenIsEmpty(ret, "机房类型")
	if cls.LocationStationUuid == "" && cls.LocationSectionUuid == "" && cls.LocationCenterUuid == "" {
		wrongs.PanicValidate("归属单位必选")
	}
	if cls.LocationStationUuid != "" {
		ret = models.BootByModel(models.LocationStationModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationStationUuid}).
			PrepareByDefault().
			First(&cls.LocationStation)
		wrongs.PanicWhenIsEmpty(ret, "所属战场")
	}
	if cls.LocationSectionUuid != "" {
		ret = models.BootByModel(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationSection}).
			PrepareByDefault().
			First(&cls.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属区间")
	}
	if cls.LocationCenterUuid != "" {
		ret = models.BootByModel(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationSection}).
			PrepareByDefault().
			First(&cls.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属中心")
	}

	return cls
}

func (PositionIndoorRoomController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionIndoorRoomModel
	)

	// 表单
	form := (&PositionIndoorRoomStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房名称")

	// 新建
	positionIndoorRoom := &models.PositionIndoorRoomModel{
		BaseModel:              models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:             form.UniqueCode,
		Name:                   form.Name,
		PositionIndoorRoomType: form.PositionIndoorRoomType,
		LocationStation:        form.LocationStation,
	}
	if ret = models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefault().Create(&positionIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_room": positionIndoorRoom}))
}
func (PositionIndoorRoomController) D(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		locationIndoorRoom models.PositionIndoorRoomModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefault().Delete(&locationIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionIndoorRoomController) U(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		positionIndoorRoom, repeat models.PositionIndoorRoomModel
	)

	// 表单
	form := (&PositionIndoorRoomStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	// 编辑
	positionIndoorRoom.BaseModel.Sort = form.Sort
	positionIndoorRoom.Name = form.Name
	positionIndoorRoom.PositionIndoorRoomType = form.PositionIndoorRoomType
	positionIndoorRoom.LocationStation = form.LocationStation
	if ret = models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefault().Save(&positionIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_room": positionIndoorRoom}))
}
func (PositionIndoorRoomController) S(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		positionIndoorRoom models.PositionIndoorRoomModel
	)
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_room": positionIndoorRoom}))
}
func (PositionIndoorRoomController) I(ctx *gin.Context) {
	var positionIndoorRooms []models.PositionIndoorRoomModel
	models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWhereFields().
		PrepareQuery(ctx, "").
		Find(&positionIndoorRooms)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_rooms": positionIndoorRooms}))
}
