package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotTierController struct{}

// PositionDepotTierStoreForm 新建仓库柜架层表单
type PositionDepotTierStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	PositionDepotCabinetUuid string `form:"position_depot_cabinet_uuid" json:"position_depot_cabinet_uuid"`
	PositionDepotCabinet     models.PositionDepotCabinetModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return PositionDepotTierStoreForm
func (ins PositionDepotTierStoreForm) ShouldBind(ctx *gin.Context) PositionDepotTierStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库柜架层代码不能必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库柜架层名称不能必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("仓库柜架层名称不能超过64位")
	}
	if ins.PositionDepotCabinetUuid == "" {
		wrongs.PanicValidate("所属仓库柜架必选")
	}
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": ins.PositionDepotCabinetUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架")

	return ins
}

// C 新建
func (PositionDepotTierController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotTierModel
	)

	// 表单
	form := (&PositionDepotTierStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

	// 新建
	positionDepotTier := &models.PositionDepotTierModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		PositionDepotCabinet: form.PositionDepotCabinet,
	}
	if ret = models.BootByModel(models.PositionDepotTierModel{}).PrepareByDefaultDbDriver().Create(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_depot_tier": positionDepotTier}))
}

// D 删除
func (PositionDepotTierController) D(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotTier models.PositionDepotTierModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	// 删除
	if ret := models.BootByModel(models.PositionDepotTierModel{}).PrepareByDefaultDbDriver().Delete(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (PositionDepotTierController) U(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		positionDepotTier, repeat models.PositionDepotTierModel
	)

	// 表单
	form := (&PositionDepotTierStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	// 编辑
	positionDepotTier.BaseModel.Sort = form.Sort
	positionDepotTier.Name = form.Name
	positionDepotTier.PositionDepotCabinet = form.PositionDepotCabinet
	if ret = models.BootByModel(models.PositionDepotTierModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_tier": positionDepotTier}))
}

// S 详情
func (PositionDepotTierController) S(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotTier models.PositionDepotTierModel
	)
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_tier": positionDepotTier}))
}

// I 列表
func (PositionDepotTierController) I(ctx *gin.Context) {
	var (
		positionDepotTiers []models.PositionDepotTierModel
		count              int64
		db                 *gorm.DB
	)
	db = models.BootByModel(models.PositionDepotTierModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionDepotTiers)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_tiers": positionDepotTiers}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionDepotTiers)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_depot_tiers": positionDepotTiers}, ctx.Query("__page__"), count))
	}
}
