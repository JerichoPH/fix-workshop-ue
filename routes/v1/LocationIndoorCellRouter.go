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

type LocationIndoorCellRouter struct{}

type LocationIndoorCellStoreForm struct {
	Sort                   int64  `form:"sort" json:"sort"`
	UniqueCode             string `form:"unique_code" json:"unique_code"`
	Name                   string `form:"name" json:"name"`
	LocationIndoorTierUUID string `form:"location_indoor_tier_uuid" json:"location_indoor_tier_uuid"`
	LocationIndoorTier     models.LocationIndoorTierModel
}

// ShouldBind
//  @receiver cls
//  @param ctx
//  @return LocationIndoorCellStoreForm
func (cls LocationIndoorCellStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorCellStoreForm {
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
	if cls.LocationIndoorTierUUID == "" {
		wrongs.PanicValidate("所属机柜层必选")
	}
	ret = models.Init(models.LocationIndoorTierModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationIndoorTierUUID}).
		Prepare().
		First(&cls.LocationIndoorTier)
	wrongs.PanicWhenIsEmpty(ret, "所属机柜层")

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorCellRouter
//  @param engine
func (LocationIndoorCellRouter) Load(engine *gin.Engine) {
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
				repeat models.LocationIndoorCellModel
			)

			// 表单
			form := (&LocationIndoorCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

			// 新建
			locationIndoorCell := &models.LocationIndoorCellModel{
				BaseModel:          models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:         form.UniqueCode,
				Name:               form.Name,
				LocationIndoorTier: form.LocationIndoorTier,
			}
			if ret = models.Init(models.LocationIndoorCellModel{}).GetSession().Create(&locationIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_cell": locationIndoorCell}))
		})

		// 删除
		r.DELETE("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorCell models.LocationIndoorCellModel
			)

			// 查询
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			// 删除
			if ret := models.Init(models.LocationIndoorCellModel{}).GetSession().Delete(&locationIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationIndoorCell, repeat models.LocationIndoorCellModel
			)

			// 表单
			form := (&LocationIndoorCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位代码")
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机柜格位名称")

			// 查询
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			// 编辑
			locationIndoorCell.BaseModel.Sort = form.Sort
			locationIndoorCell.UniqueCode = form.UniqueCode
			locationIndoorCell.Name = form.Name
			locationIndoorCell.LocationIndoorTier = form.LocationIndoorTier
			if ret = models.Init(models.LocationIndoorCellModel{}).GetSession().Save(&locationIndoorCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_cell": locationIndoorCell}))
		})

		// 详情
		r.GET("indoorCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorCell models.LocationIndoorCellModel
			)
			ret = models.Init(models.LocationIndoorCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCell)
			wrongs.PanicWhenIsEmpty(ret, "机柜格位")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_cell": locationIndoorCell}))
		})

		// 列表
		r.GET("indoorCell", func(ctx *gin.Context) {
			var locationIndoorCells []models.LocationIndoorCellModel
			models.Init(models.LocationIndoorCellModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorCells)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_cells": locationIndoorCells}))
		})
	}
}
