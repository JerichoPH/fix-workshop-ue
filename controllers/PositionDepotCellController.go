package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotCellController struct{}

// PositionDepotCellStoreForm 新建仓储柜架格位表单
type PositionDepotCellStoreForm struct {
	Sort                  int64  `form:"sort" json:"sort"`
	UniqueCode            string `form:"unique_code" json:"unique_code"`
	Name                  string `form:"name" json:"name"`
	PositionDepotTierUuid string `form:"position_depot_tier_uuid" json:"position_depot_tier_uuid"`
	PositionDepotTier     models.PositionDepotTierModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotCellStoreForm
func (cls PositionDepotCellStoreForm) ShouldBind(ctx *gin.Context) PositionDepotCellStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库格位代码不能必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库格位名称不能必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("仓库格位名称不能超过64位")
	}
	if cls.PositionDepotTierUuid == "" {
		wrongs.PanicValidate("所属仓库柜架层必选")
	}
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotTierUuid}).
		PrepareByDefaultDbDriver().
		First(&cls.PositionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架层")

	return cls
}

// C 新建
func (PositionDepotCellController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotCellModel
	)

	// 表单
	form := (&PositionDepotCellStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

	// 新建
	positionDepotCell := &models.PositionDepotCellModel{
		BaseModel:         models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:        form.UniqueCode,
		Name:              form.Name,
		PositionDepotTier: form.PositionDepotTier,
	}
	if ret = models.BootByModel(models.PositionDepotCellModel{}).PrepareByDefaultDbDriver().Create(&positionDepotCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_depot_cell": positionDepotCell}))
}

// D 删除
func (PositionDepotCellController) D(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotCell models.PositionDepotCellModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCell)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

	// 删除
	if ret := models.BootByModel(models.PositionDepotCellModel{}).PrepareByDefaultDbDriver().Delete(&positionDepotCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (PositionDepotCellController) U(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		positionDepotCell, repeat models.PositionDepotCellModel
	)

	// 表单
	form := (&PositionDepotCellStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCell)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

	// 编辑
	positionDepotCell.BaseModel.Sort = form.Sort
	positionDepotCell.Name = form.Name
	positionDepotCell.PositionDepotTier = form.PositionDepotTier
	if ret = models.BootByModel(models.PositionDepotCellModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionDepotCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_cell": positionDepotCell}))
}

// S 详情
func (PositionDepotCellController) S(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotCell models.PositionDepotCellModel
	)
	ret = models.BootByModel(models.PositionDepotCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotCell)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_cell": positionDepotCell}))
}

// I 列表
func (PositionDepotCellController) I(ctx *gin.Context) {
	var (
		positionDepotCells []models.PositionDepotCellModel
		count           int64
		db              *gorm.DB
	)
	db = models.BootByModel(models.PositionDepotCellModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionDepotCells)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_cells": positionDepotCells}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionDepotCells)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_depot_cells": positionDepotCells}, ctx.Query("__page__"), count))
	}
}
