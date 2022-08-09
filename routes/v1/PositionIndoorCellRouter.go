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

// PositionIndoorCellRouter 室内上道位置-柜架格位
type PositionIndoorCellRouter struct{}

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
	ret = models.Init(models.PositionIndoorTierModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorTierUUID}).
		Prepare().
		First(&cls.PositionIndoorTier)
	wrongs.PanicWhenIsEmpty(ret, "所属机柜层")

	return cls
}

// Load 加载路由
//  @receiver PositionIndoorCellRouter
//  @param engine
func (PositionIndoorCellRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorCell", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionIndoorCellModel
			)

			// 表单
			form := (&PositionIndoorCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

			// 新建
			positionIndoorCell := &models.PositionIndoorCellModel{
				BaseModel:          models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:         form.UniqueCode,
				Name:               form.Name,
				PositionIndoorTier: form.PositionIndoorTier,
			}
			if ret = models.Init(models.PositionIndoorCellModel{}).GetSession().Create(&positionIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_indoor_cell": positionIndoorCell}))
		})

		// 删除
		r.DELETE("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				positionIndoorCell models.PositionIndoorCellModel
			)

			// 查询
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			// 删除
			if ret := models.Init(models.PositionIndoorCellModel{}).GetSession().Delete(&positionIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				positionIndoorCell, repeat models.PositionIndoorCellModel
			)

			// 表单
			form := (&PositionIndoorCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

			// 查询
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			// 编辑
			positionIndoorCell.BaseModel.Sort = form.Sort
			positionIndoorCell.UniqueCode = form.UniqueCode
			positionIndoorCell.Name = form.Name
			positionIndoorCell.PositionIndoorTier = form.PositionIndoorTier
			if ret = models.Init(models.PositionIndoorCellModel{}).GetSession().Save(&positionIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_indoor_cell": positionIndoorCell}))
		})

		// 详情
		r.GET("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				positionIndoorCell models.PositionIndoorCellModel
			)
			ret = models.Init(models.PositionIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_cell": positionIndoorCell}))
		})

		// 列表
		r.GET("indoorCell", func(ctx *gin.Context) {
			var positionIndoorCells []models.PositionIndoorCellModel
			models.Init(models.PositionIndoorCellModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionIndoorCells)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_cells": positionIndoorCells}))
		})
	}
}
