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

// LocationDepotCellRouter 仓储柜架格位路由
type LocationDepotCellRouter struct{}

// LocationDepotCellStoreForm 新建仓储柜架格位表单
type LocationDepotCellStoreForm struct {
	Sort                  int64  `form:"sort" json:"sort"`
	UniqueCode            string `form:"unique_code" json:"unique_code"`
	Name                  string `form:"name" json:"name"`
	LocationDepotTierUUID string `form:"location_depot_tier_uuid" json:"location_depot_tier_uuid"`
	LocationDepotTier     models.LocationDepotTierModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotCellStoreForm
func (cls LocationDepotCellStoreForm) ShouldBind(ctx *gin.Context) LocationDepotCellStoreForm {
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
	if cls.LocationDepotTierUUID == "" {
		wrongs.PanicValidate("所属仓库柜架层必选")
	}
	ret = models.Init(models.LocationDepotTierModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationDepotTierUUID}).
		Prepare().
		First(&cls.LocationDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架层")

	return cls
}

// Load 加载路由
//  @receiver LocationDepotCellRouter
//  @param router
func (LocationDepotCellRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotCell", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotCellModel
			)

			// 表单
			form := (&LocationDepotCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

			// 新建
			locationDepotCell := &models.LocationDepotCellModel{
				BaseModel:         models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:        form.UniqueCode,
				Name:              form.Name,
				LocationDepotTier: form.LocationDepotTier,
			}
			if ret = models.Init(models.LocationDepotCellModel{}).GetSession().Create(&locationDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_cell": locationDepotCell}))
		})

		// 删除
		r.DELETE("depotCell/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationDepotCell models.LocationDepotCellModel
			)

			// 查询
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			// 删除
			if ret := models.Init(models.LocationDepotCellModel{}).GetSession().Delete(&locationDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotCell/:uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				locationDepotCell, repeat models.LocationDepotCellModel
			)

			// 表单
			form := (&LocationDepotCellStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位代码")
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架格位名称")

			// 查询
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			// 编辑
			locationDepotCell.BaseModel.Sort = form.Sort
			locationDepotCell.UniqueCode = form.UniqueCode
			locationDepotCell.Name = form.Name
			locationDepotCell.LocationDepotTier = form.LocationDepotTier
			if ret = models.Init(models.LocationDepotCellModel{}).GetSession().Save(&locationDepotCell); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_depot_cell": locationDepotCell}))
		})

		// 详情
		r.GET("depotCell/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationDepotCell models.LocationDepotCellModel
			)
			ret = models.Init(models.LocationDepotCellModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCell)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架格位")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_cell": locationDepotCell}))
		})

		// 列表
		r.GET("depotCell", func(ctx *gin.Context) {
			var locationDepotCells []models.LocationDepotCellModel
			models.Init(models.LocationDepotCellModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotCells)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_cells": locationDepotCells}))
		})
	}
}
