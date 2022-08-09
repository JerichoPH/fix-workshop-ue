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

// PositionDepotTierRouter 仓库柜架层路由
type PositionDepotTierRouter struct{}

// PositionDepotTierStoreForm 新建仓库柜架层表单
type PositionDepotTierStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	PositionDepotCabinetUUID string `form:"position_depot_cabinet_uuid" json:"position_depot_cabinet_uuid"`
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
		wrongs.PanicValidate("仓库柜架代码不能必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库柜架名称不能必填")
	}
	if cls.PositionDepotCabinetUUID == "" {
		wrongs.PanicValidate("所属仓库柜架必选")
	}
	ret = models.Init(models.PositionDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotCabinetUUID}).
		Prepare().
		First(&cls.PositionDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架")

	return cls
}

// Load 加载路由
//  @receiver PositionDepotTierRouter
//  @param router
func (PositionDepotTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotTier",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotTierModel
			)

			// 表单
			form := (&PositionDepotTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

			// 新建
			positionDepotTier := &models.PositionDepotTierModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				PositionDepotCabinet: form.PositionDepotCabinet,
			}
			if ret = models.Init(models.PositionDepotTierModel{}).GetSession().Create(&positionDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_tier": positionDepotTier}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionDepotTier models.PositionDepotTierModel
			)

			// 查询
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			// 删除
			if ret := models.Init(models.PositionDepotTierModel{}).GetSession().Delete(&positionDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				positionDepotTier, repeat models.PositionDepotTierModel
			)

			// 表单
			form := (&PositionDepotTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

			// 查询
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			// 编辑
			positionDepotTier.BaseModel.Sort = form.Sort
			positionDepotTier.UniqueCode = form.UniqueCode
			positionDepotTier.Name = form.Name
			positionDepotTier.PositionDepotCabinet = form.PositionDepotCabinet
			if ret = models.Init(models.PositionDepotTierModel{}).GetSession().Save(&positionDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_tier": positionDepotTier}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionDepotTier models.PositionDepotTierModel
			)
			ret = models.Init(models.PositionDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_tier": positionDepotTier}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionDepotTiers []models.PositionDepotTierModel
			models.Init(models.PositionDepotTierModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionDepotTiers)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_tiers": positionDepotTiers}))
		})
	}
}
