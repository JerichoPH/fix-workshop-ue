package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionIndoorCabinetController struct{}

// PositionIndoorCabinetStoreForm 新建室内上道位置机柜表单
type PositionIndoorCabinetStoreForm struct {
	Sort                  int64  `form:"sort" json:"sort"`
	UniqueCode            string `form:"unique_code" json:"unique_code"`
	Name                  string `form:"name" json:"name"`
	PositionIndoorRowUuid string `form:"position_indoor_row_uuid" json:"position_indoor_row_uuid"`
	PositionIndoorRow     models.PositionIndoorRowModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionIndoorCabinetStoreForm
func (cls PositionIndoorCabinetStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("柜架代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("柜架名称必填")
	}
	if cls.PositionIndoorRowUuid == "" {
		ret = models.BootByModel(models.PositionIndoorRowModel{}).
			SetWheres(tools.Map{"uuid": cls.PositionIndoorRowUuid}).
			PrepareByDefault().
			First(&cls.PositionIndoorRow)
		wrongs.PanicWhenIsEmpty(ret, "所属排")
	}

	return cls
}

func (PositionIndoorCabinetController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionIndoorCabinetModel
	)

	// 表单
	form := (&PositionIndoorCabinetStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架代码")
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架名称")

	// 新建
	locationIndoorCabinet := &models.PositionIndoorCabinetModel{
		BaseModel:         models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:        form.UniqueCode,
		Name:              form.Name,
		PositionIndoorRow: form.PositionIndoorRow,
	}
	if ret = models.BootByModel(models.PositionIndoorCabinetModel{}).PrepareByDefault().Create(&locationIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_cabinet": locationIndoorCabinet}))
}
func (PositionIndoorCabinetController) D(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		positionIndoorCabinet models.PositionIndoorCabinetModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorCabinetModel{}).PrepareByDefault().Delete(&positionIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionIndoorCabinetController) U(ctx *gin.Context) {
	var (
		ret                           *gorm.DB
		positionIndoorCabinet, repeat models.PositionIndoorCabinetModel
	)

	// 表单
	form := (&PositionIndoorCabinetStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架代码")
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	// 编辑
	positionIndoorCabinet.BaseModel.Sort = form.Sort
	positionIndoorCabinet.Name = form.Name
	positionIndoorCabinet.PositionIndoorRow = form.PositionIndoorRow
	if ret = models.BootByModel(models.PositionIndoorCabinetModel{}).SetWheres(tools.Map{"uuid":ctx.Param("uuid")}).PrepareByDefault().Save(&positionIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
}
func (PositionIndoorCabinetController) S(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		positionIndoorCabinet models.PositionIndoorCabinetModel
	)
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
}
func (PositionIndoorCabinetController) I(ctx *gin.Context) {
	var positionIndoorCabinet []models.PositionIndoorCabinetModel
	models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWhereFields().
		PrepareQuery(ctx, "").
		Find(&positionIndoorCabinet)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
}
