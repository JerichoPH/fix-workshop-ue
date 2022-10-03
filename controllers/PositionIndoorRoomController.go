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
func (ins PositionIndoorRoomStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRoomStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("机房代码必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("机房名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("机房名称不能超过64位")
	}
	if ins.PositionIndoorRoomTypeUuid == "" {
		wrongs.PanicValidate("机房类型必选")
	}
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"uuid": ins.PositionIndoorRoomTypeUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionIndoorRoomType)
	wrongs.PanicWhenIsEmpty(ret, "机房类型")
	if ins.LocationStationUuid == "" && ins.LocationSectionUuid == "" && ins.LocationCenterUuid == "" {
		wrongs.PanicValidate("归属单位必选")
	}
	if ins.LocationStationUuid != "" {
		ret = models.BootByModel(models.LocationStationModel{}).
			SetWheres(tools.Map{"uuid": ins.LocationStationUuid}).
			PrepareByDefaultDbDriver().
			First(&ins.LocationStation)
		wrongs.PanicWhenIsEmpty(ret, "所属战场")
	}
	if ins.LocationSectionUuid != "" {
		ret = models.BootByModel(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": ins.LocationSection}).
			PrepareByDefaultDbDriver().
			First(&ins.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属区间")
	}
	if ins.LocationCenterUuid != "" {
		ret = models.BootByModel(models.LocationSectionModel{}).
			SetWheres(tools.Map{"uuid": ins.LocationSection}).
			PrepareByDefaultDbDriver().
			First(&ins.LocationSection)
		wrongs.PanicWhenIsEmpty(ret, "所属中心")
	}

	return ins
}

// C 新建
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
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
	if ret = models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefaultDbDriver().Create(&positionIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_room": positionIndoorRoom}))
}

// D 删除
func (PositionIndoorRoomController) D(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		locationIndoorRoom models.PositionIndoorRoomModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefaultDbDriver().Delete(&locationIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	// 编辑
	positionIndoorRoom.BaseModel.Sort = form.Sort
	positionIndoorRoom.Name = form.Name
	positionIndoorRoom.PositionIndoorRoomType = form.PositionIndoorRoomType
	positionIndoorRoom.LocationStation = form.LocationStation
	if ret = models.BootByModel(models.PositionIndoorRoomModel{}).PrepareByDefaultDbDriver().Save(&positionIndoorRoom); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_room": positionIndoorRoom}))
}

// S 详情
func (PositionIndoorRoomController) S(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		positionIndoorRoom models.PositionIndoorRoomModel
	)
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "机房")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_room": positionIndoorRoom}))
}

// I 列表
func (PositionIndoorRoomController) I(ctx *gin.Context) {
	var (
		positionIndoorRooms []models.PositionIndoorRoomModel
		count               int64
		db                  *gorm.DB
	)
	db = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionIndoorRooms)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_rooms": positionIndoorRooms}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionIndoorRooms)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_indoor_rooms": positionIndoorRooms}, ctx.Query("__page__"), count))
	}
}
