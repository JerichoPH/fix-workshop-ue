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
//  @receiver cls
//  @param ctx
//  @return PositionDepotTierStoreForm
func (cls PositionDepotTierStoreForm) ShouldBind(ctx *gin.Context) PositionDepotTierStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库柜架层代码不能必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库柜架层名称不能必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("仓库柜架层名称不能超过64位")
	}
	if cls.PositionDepotCabinetUuid == "" {
		wrongs.PanicValidate("所属仓库柜架必选")
	}
	ret = models.BootByModel(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotCabinetUuid}).
		PrepareByDefault().
		First(&cls.PositionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架")

	return cls
}

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
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

	// 新建
	positionDepotTier := &models.PositionDepotTierModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		PositionDepotCabinet: form.PositionDepotCabinet,
	}
	if ret = models.BootByModel(models.PositionDepotTierModel{}).PrepareByDefault().Create(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_depot_tier": positionDepotTier}))
}
func (PositionDepotTierController) D(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotTier models.PositionDepotTierModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	// 删除
	if ret := models.BootByModel(models.PositionDepotTierModel{}).PrepareByDefault().Delete(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
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
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	// 编辑
	positionDepotTier.BaseModel.Sort = form.Sort
	positionDepotTier.Name = form.Name
	positionDepotTier.PositionDepotCabinet = form.PositionDepotCabinet
	if ret = models.BootByModel(models.PositionDepotTierModel{}).SetWheres(tools.Map{"uuid":ctx.Param("uuid")}).PrepareByDefault().Save(&positionDepotTier); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_tier": positionDepotTier}))
}
func (PositionDepotTierController) S(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		positionDepotTier models.PositionDepotTierModel
	)
	ret = models.BootByModel(models.PositionDepotTierModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionDepotTier)
	wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_depot_tier": positionDepotTier}))
}
func (PositionDepotTierController) I(ctx *gin.Context) {
	var positionDepotTiers []models.PositionDepotTierModel
	models.BootByModel(models.PositionDepotTierModel{}).
		SetWhereFields().
		PrepareQuery(ctx,"").
		Find(&positionDepotTiers)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_depot_tiers": positionDepotTiers}))
}
