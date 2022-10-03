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
//  @receiver ins
//  @param ctx
//  @return PositionIndoorCabinetStoreForm
func (ins PositionIndoorCabinetStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("柜架代码必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("柜架名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("柜架名称不能超过64位")
	}
	if ins.PositionIndoorRowUuid == "" {
		ret = models.BootByModel(models.PositionIndoorRowModel{}).
			SetWheres(tools.Map{"uuid": ins.PositionIndoorRowUuid}).
			PrepareByDefaultDbDriver().
			First(&ins.PositionIndoorRow)
		wrongs.PanicWhenIsEmpty(ret, "所属排")
	}

	return ins
}

// C 新建
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架代码")
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架名称")

	// 新建
	locationIndoorCabinet := &models.PositionIndoorCabinetModel{
		BaseModel:         models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:        form.UniqueCode,
		Name:              form.Name,
		PositionIndoorRow: form.PositionIndoorRow,
	}
	if ret = models.BootByModel(models.PositionIndoorCabinetModel{}).PrepareByDefaultDbDriver().Create(&locationIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_cabinet": locationIndoorCabinet}))
}

// D 删除
func (PositionIndoorCabinetController) D(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		positionIndoorCabinet models.PositionIndoorCabinetModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorCabinetModel{}).PrepareByDefaultDbDriver().Delete(&positionIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架代码")
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "柜架名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	// 编辑
	positionIndoorCabinet.BaseModel.Sort = form.Sort
	positionIndoorCabinet.Name = form.Name
	positionIndoorCabinet.PositionIndoorRow = form.PositionIndoorRow
	if ret = models.BootByModel(models.PositionIndoorCabinetModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionIndoorCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
}

// S 详情
func (PositionIndoorCabinetController) S(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		positionIndoorCabinet models.PositionIndoorCabinetModel
	)
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "柜架")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
}

// I 列表
func (PositionIndoorCabinetController) I(ctx *gin.Context) {
	var (
		positionIndoorCabinet []models.PositionIndoorCabinetModel
		count                 int64
		db                    *gorm.DB
	)
	db = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionIndoorCabinet)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionIndoorCabinet)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}, ctx.Query("__page__"), count))
	}
}
