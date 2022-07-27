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

// LocationDepotTierRouter 仓库柜架层路由
type LocationDepotTierRouter struct{}

// LocationDepotTierStoreForm 新建仓库柜架层表单
type LocationDepotTierStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	LocationDepotCabinetUUID string `form:"location_depot_cabinet_uuid" json:"location_depot_cabinet_uuid"`
	LocationDepotCabinet     models.LocationDepotCabinetModel
	LocationDepotCellUUIDs   []string `form:"location_depot_cell_uuids" json:"location_depot_cell_uuids"`
	LocationDepotCells       []models.LocationDepotCellModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotTierStoreForm
func (cls LocationDepotTierStoreForm) ShouldBind(ctx *gin.Context) LocationDepotTierStoreForm {
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
	if cls.LocationDepotCabinetUUID == "" {
		wrongs.PanicValidate("所属仓库柜架必选")
	}
	ret = models.Init(models.LocationDepotCabinetModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationDepotCabinetUUID}).
		Prepare().
		First(&cls.LocationDepotCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库柜架")
	if len(cls.LocationDepotCellUUIDs) > 0 {
		models.Init(models.LocationDepotCellModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationDepotCellUUIDs).
			Find(&cls.LocationDepotCells)
	}

	return cls
}

// Load 加载路由
//  @receiver LocationDepotTierRouter
//  @param router
func (LocationDepotTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotTier", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotTierModel
			)

			// 表单
			form := (&LocationDepotTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

			// 新建
			locationDepotTier := &models.LocationDepotTierModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				LocationDepotCabinet: form.LocationDepotCabinet,
				LocationDepotCells:   form.LocationDepotCells,
			}
			if ret = models.Init(models.LocationDepotTierModel{}).GetSession().Create(&locationDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_tier": locationDepotTier}))
		})

		// 删除
		r.DELETE("depotTier/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationDepotTier models.LocationDepotTierModel
			)

			// 查询
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			// 删除
			if ret := models.Init(models.LocationDepotTierModel{}).GetSession().Delete(&locationDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotTier/:uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				locationDepotTier, repeat models.LocationDepotTierModel
			)

			// 表单
			form := (&LocationDepotTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层代码")
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架层名称")

			// 查询
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			// 编辑
			locationDepotTier.BaseModel.Sort = form.Sort
			locationDepotTier.UniqueCode = form.UniqueCode
			locationDepotTier.Name = form.Name
			locationDepotTier.LocationDepotCabinet = form.LocationDepotCabinet
			locationDepotTier.LocationDepotCells = form.LocationDepotCells
			if ret = models.Init(models.LocationDepotTierModel{}).GetSession().Save(&locationDepotTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_depot_tier": locationDepotTier}))
		})

		// 详情
		r.GET("depotTier/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationDepotTier models.LocationDepotTierModel
			)
			ret = models.Init(models.LocationDepotTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotTier)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架层")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_tier": locationDepotTier}))
		})

		// 列表
		r.GET("depotTier", func(ctx *gin.Context) {
			var locationDepotTiers []models.LocationDepotTierModel
			models.Init(models.LocationDepotTierModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotTiers)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_tiers": locationDepotTiers}))
		})
	}
}
