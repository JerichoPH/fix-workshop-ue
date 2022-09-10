package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionIndoorRowController struct{}

// PositionIndoorRowStoreForm 新建上道位置排表单
type PositionIndoorRowStoreForm struct {
	Sort                   int64  `form:"sort" json:"sort"`
	UniqueCode             string `form:"unique_code" json:"unique_code"`
	Name                   string `form:"name" json:"name"`
	PositionIndoorRoomUUID string `form:"position_indoor_room_uuid" json:"position_indoor_room_uuid"`
	PositionIndoorRoom     models.PositionIndoorRoomModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionIndoorRowStoreForm
func (cls PositionIndoorRowStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRowStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("排代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("排名称必填")
	}
	if cls.PositionIndoorRoomUUID == "" {
		wrongs.PanicValidate("所属机房必选")
	}
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorRoomUUID}).
		PrepareByDefault().
		First(&cls.PositionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "所属机房")

	return cls
}

func (PositionIndoorRowController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionIndoorRowModel
	)

	// 表单
	form := (&PositionIndoorRowStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排代码")
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排名称")

	// 新建
	positionIndoorRow := &models.PositionIndoorRowModel{
		BaseModel:          models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:         form.UniqueCode,
		Name:               form.Name,
		PositionIndoorRoom: form.PositionIndoorRoom,
	}
	if ret = models.BootByModel(models.PositionIndoorRowModel{}).PrepareByDefault().Create(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_row": positionIndoorRow}))
}
func (PositionIndoorRowController) D(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionIndoorRow models.PositionIndoorRowModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorRowModel{}).PrepareByDefault().Delete(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionIndoorRowController) U(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		positionIndoorRow, repeat models.PositionIndoorRowModel
	)

	// 表单
	form := (&PositionIndoorRowStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排代码")
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	// 编辑
	positionIndoorRow.BaseModel.Sort = form.Sort
	positionIndoorRow.UniqueCode = form.UniqueCode
	positionIndoorRow.Name = form.Name
	positionIndoorRow.PositionIndoorRoom = form.PositionIndoorRoom
	if ret = models.BootByModel(models.PositionIndoorRowModel{}).PrepareByDefault().Save(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_row": positionIndoorRow}))
}
func (PositionIndoorRowController) S(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionIndoorRow models.PositionIndoorRowModel
	)
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_row": positionIndoorRow}))
}
func (PositionIndoorRowController) I(ctx *gin.Context) {
	var positionIndoorRows []models.PositionIndoorRowModel
	models.BootByModel(models.PositionIndoorRowModel{}).
		SetWhereFields().
		PrepareUseQueryByDefault(ctx).
		Find(&positionIndoorRows)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_rows": positionIndoorRows}))
}
