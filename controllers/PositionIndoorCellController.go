package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionIndoorCellController struct{}

// PositionIndoorCellStoreForm 新建室内上道位置-柜架格位表单
type PositionIndoorCellStoreForm struct {
	Sort                   int64  `form:"sort" json:"sort"`
	UniqueCode             string `form:"unique_code" json:"unique_code"`
	Name                   string `form:"name" json:"name"`
	PositionIndoorTierUUID string `form:"position_indoor_tier_uuid" json:"position_indoor_tier_uuid"`
	PositionIndoorTier     models.PositionIndoorTierModel
}

// ShouldBind
//  @receiver cls
//  @param ctx
//  @return PositionIndoorCellStoreForm
func (cls PositionIndoorCellStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorCellStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("机柜格位代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("机柜格位名称必填")
	}
	if cls.PositionIndoorTierUUID == "" {
		wrongs.PanicValidate("所属机柜层必选")
	}
	ret = models.BootByModel(models.PositionIndoorTierModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorTierUUID}).
		PrepareByDefault().
		First(&cls.PositionIndoorTier)
	wrongs.PanicWhenIsEmpty(ret, "所属机柜层")

	return cls
}

func (PositionIndoorCellController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionIndoorCellModel
	)

	// 表单
	form := (&PositionIndoorCellStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

	// 新建
	positionIndoorCell := &models.PositionIndoorCellModel{
		BaseModel:          models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:         form.UniqueCode,
		Name:               form.Name,
		PositionIndoorTier: form.PositionIndoorTier,
	}
	if ret = models.BootByModel(models.PositionIndoorCellModel{}).PrepareByDefault().Create(&positionIndoorCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_cell": positionIndoorCell}))
}
func (PositionIndoorCellController) D(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		positionIndoorCell models.PositionIndoorCellModel
	)

	// 查询
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCell)
	wrongs.PanicWhenIsEmpty(ret, "机柜格位")

	// 删除
	if ret := models.BootByModel(models.PositionIndoorCellModel{}).PrepareByDefault().Delete(&positionIndoorCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionIndoorCellController) U(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		positionIndoorCell, repeat models.PositionIndoorCellModel
	)

	// 表单
	form := (&PositionIndoorCellStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

	// 查询
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCell)
	wrongs.PanicWhenIsEmpty(ret, "机柜格位")

	// 编辑
	positionIndoorCell.BaseModel.Sort = form.Sort
	positionIndoorCell.UniqueCode = form.UniqueCode
	positionIndoorCell.Name = form.Name
	positionIndoorCell.PositionIndoorTier = form.PositionIndoorTier
	if ret = models.BootByModel(models.PositionIndoorCellModel{}).PrepareByDefault().Save(&positionIndoorCell); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_cell": positionIndoorCell}))
}
func (PositionIndoorCellController) S(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		positionIndoorCell models.PositionIndoorCellModel
	)
	ret = models.BootByModel(models.PositionIndoorCellModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionIndoorCell)
	wrongs.PanicWhenIsEmpty(ret, "机柜格位")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_cell": positionIndoorCell}))
}
func (PositionIndoorCellController) I(ctx *gin.Context) {
	var positionIndoorCells []models.PositionIndoorCellModel
	models.BootByModel(models.PositionIndoorCellModel{}).
		SetWhereFields().
		PrepareQuery(ctx,"").
		Find(&positionIndoorCells)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_cells": positionIndoorCells}))
}
