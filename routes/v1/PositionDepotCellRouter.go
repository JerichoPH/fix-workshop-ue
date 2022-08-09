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

// PositionDepotCellRouter 仓储柜架格位路由
type PositionDepotCellRouter struct{}

// PositionDepotCellStoreForm 新建仓储柜架格位表单
type PositionDepotCellStoreForm struct {
	Sort                  int64  `form:"sort" json:"sort"`
	UniqueCode            string `form:"unique_code" json:"unique_code"`
	Name                  string `form:"name" json:"name"`
	PositionDepotTierUUID string `form:"position_depot_tier_uuid" json:"position_depot_tier_uuid"`
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
		wrongs.PanicValidate("仓库代码不能必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库名称不能必填")
	}
	if cls.PositionDepotTierUUID == "" {
		wrongs.PanicValidate("所属仓库柜架层必选")
	}
	ret = models.Init(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotTierUUID}).
		Prepare().
		First(&cls.PositionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架层")

	return cls
}

// Load 加载路由
//  @receiver PositionDepotCellRouter
//  @param router
func (PositionDepotCellRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotCell",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotCellModel
			)

			// 表单
			form := (&PositionDepotCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

			// 新建
			positionDepotCell := &models.PositionDepotCellModel{
				BaseModel:         models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:        form.UniqueCode,
				Name:              form.Name,
				PositionDepotTier: form.PositionDepotTier,
			}
			if ret = models.Init(models.PositionDepotCellModel{}).GetSession().Create(&positionDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_depot_cell": positionDepotCell}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionDepotCell models.PositionDepotCellModel
			)

			// 查询
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			// 删除
			if ret := models.Init(models.PositionDepotCellModel{}).GetSession().Delete(&positionDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				positionDepotCell, repeat models.PositionDepotCellModel
			)

			// 表单
			form := (&PositionDepotCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

			// 查询
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			// 编辑
			positionDepotCell.BaseModel.Sort = form.Sort
			positionDepotCell.UniqueCode = form.UniqueCode
			positionDepotCell.Name = form.Name
			positionDepotCell.PositionDepotTier = form.PositionDepotTier
			if ret = models.Init(models.PositionDepotCellModel{}).GetSession().Save(&positionDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_cell": positionDepotCell}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionDepotCell models.PositionDepotCellModel
			)
			ret = models.Init(models.PositionDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_cell": positionDepotCell}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionDepotCells []models.PositionDepotCellModel
			models.Init(models.PositionDepotCellModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionDepotCells)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_cells": positionDepotCells}))
		})
	}
}
