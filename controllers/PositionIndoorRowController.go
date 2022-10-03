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
	PositionIndoorRoomUuid string `form:"position_indoor_room_uuid" json:"position_indoor_room_uuid"`
	PositionIndoorRoom     models.PositionIndoorRoomModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return PositionIndoorRowStoreForm
func (ins PositionIndoorRowStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRowStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("排代码必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("排名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("排名称不能超过64位")
	}
	if ins.PositionIndoorRoomUuid == "" {
		wrongs.PanicValidate("所属机房必选")
	}
	ret = models.BootByModel(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": ins.PositionIndoorRoomUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "所属机房")

	return ins
}

// C 新建
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排代码")
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排名称")

	// 新建
	positionIndoorRow := &models.PositionIndoorRowModel{
		BaseModel:          models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:         form.UniqueCode,
		Name:               form.Name,
		PositionIndoorRoom: form.PositionIndoorRoom,
	}
	if ret = models.BootByModel(models.PositionIndoorRowModel{}).PrepareByDefaultDbDriver().Create(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_row": positionIndoorRow}))
}

// D 删除
func (PositionIndoorRowController) D(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionIndoorRow models.PositionIndoorRowModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorRowModel{}).PrepareByDefaultDbDriver().Delete(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排代码")
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "排名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	// 编辑
	positionIndoorRow.BaseModel.Sort = form.Sort
	positionIndoorRow.Name = form.Name
	positionIndoorRow.PositionIndoorRoom = form.PositionIndoorRoom
	if ret = models.BootByModel(models.PositionIndoorRowModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionIndoorRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_row": positionIndoorRow}))
}

// S 详情
func (PositionIndoorRowController) S(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionIndoorRow models.PositionIndoorRowModel
	)
	ret = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorRow)
	wrongs.PanicWhenIsEmpty(ret, "排")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_row": positionIndoorRow}))
}

// I 列表
func (PositionIndoorRowController) I(ctx *gin.Context) {
	var (
		positionIndoorRows []models.PositionIndoorRowModel
		count              int64
		db                 *gorm.DB
	)
	db = models.BootByModel(models.PositionIndoorRowModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionIndoorRows)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_rows": positionIndoorRows}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionIndoorRows)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_indoor_rows": positionIndoorRows}, ctx.Query("__page__"), count))
	}
}
