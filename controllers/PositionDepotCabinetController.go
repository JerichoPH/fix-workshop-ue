package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotCabinetController struct{}

// PositionDepotCabinetStoreForm 新建仓储仓库柜架表单
type PositionDepotCabinetStoreForm struct {
	Sort                 int64  `form:"sort" json:"sort"`
	UniqueCode           string `form:"unique_code" json:"unique_code"`
	Name                 string `form:"name" json:"name"`
	PositionDepotRowUuid string `form:"position_depot_row_uuid" json:"position_depot_row_uuid"`
	PositionDepotRow     models.PositionDepotRowModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return PositionDepotCabinetStoreForm
func (ins PositionDepotCabinetStoreForm) ShouldBind(ctx *gin.Context) PositionDepotCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库柜架代码必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库柜架名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("仓库柜架名称不能超过64位")
	}
	if ins.PositionDepotRowUuid == "" {
		wrongs.PanicValidate("所属仓库排必选")
	}
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"uuid": ins.PositionDepotRowUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库排")

	return ins
}

// C 新建
func (PositionDepotCabinetController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotCabinetModel
	)

	// 表单
	form := (&PositionDepotCabinetStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

	// 新建
	positionDepotCabinet := &models.PositionDepotCabinetModel{
		BaseModel:        models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:       form.UniqueCode,
		Name:             form.Name,
		PositionDepotRow: form.PositionDepotRow,
	}
	if ret = models.BootByModel(models.PositionDepotCabinetModel{}).PrepareByDefaultDbDriver().Create(&positionDepotCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
}

// D 删除
func (PositionDepotCabinetController) D(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		positionDepotCabinet models.PositionDepotCabinetModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

	// 删除
	if ret := models.BootByModel(models.PositionDepotCabinetModel{}).PrepareByDefaultDbDriver().Delete(&positionDepotCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (PositionDepotCabinetController) U(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		positionDepotCabinet, repeat models.PositionDepotCabinetModel
	)

	// 表单
	form := (&PositionDepotCabinetStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

	// 编辑
	positionDepotCabinet.BaseModel.Sort = form.Sort
	positionDepotCabinet.Name = form.Name
	positionDepotCabinet.PositionDepotRow = form.PositionDepotRow
	if ret = models.BootByModel(models.PositionDepotCabinetModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionDepotCabinet); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
}

// S 详情
func (PositionDepotCabinetController) S(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		positionDepotCabinet models.PositionDepotCabinetModel
	)
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
}

// I 列表
func (PositionDepotCabinetController) I(ctx *gin.Context) {
	var (
		positionDepotCabinets []models.PositionDepotCabinetModel
		count                 int64
		db                    *gorm.DB
	)
	db = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionDepotCabinets)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_cabinets": positionDepotCabinets}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionDepotCabinets)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_depot_cabinets": positionDepotCabinets}, ctx.Query("__page__"), count))
	}
}
