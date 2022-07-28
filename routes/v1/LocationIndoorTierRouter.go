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

// LocationIndoorTierRouter 室内上道位置柜架层路由
type LocationIndoorTierRouter struct{}

// LocationIndoorTierStoreForm 新建室内上道位置柜架层表单
type LocationIndoorTierStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	LocationIndoorCabinetUUID string `form:"location_indoor_cabinet_uuid" json:"location_indoor_cabinet_uuid"`
	LocationIndoorCabinet     models.LocationIndoorCabinetModel
}

// ShouldBind 
//  @receiver cls 
//  @param ctx 
//  @return LocationIndoorTierStoreForm 
func (cls LocationIndoorTierStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorTierStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("柜架层代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("柜架层名称必填")
	}
	if cls.LocationIndoorCabinetUUID == "" {
		wrongs.PanicValidate("所属柜架必选")
	}
	ret = models.Init(models.LocationIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationIndoorCabinetUUID}).
		Prepare().
		First(&cls.LocationIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属柜架")

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorTierRouter
//  @param engine
func (LocationIndoorTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorTier", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationIndoorTierModel
			)

			// 表单
			form := (&LocationIndoorTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层代码")
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层名称")

			// 新建
			locationIndoorTier := &models.LocationIndoorTierModel{
				BaseModel:             models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:            form.UniqueCode,
				Name:                  form.Name,
				LocationIndoorCabinet: form.LocationIndoorCabinet,
			}
			if ret = models.Init(models.LocationIndoorTierModel{}).GetSession().Create(&locationIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_tier": locationIndoorTier}))
		})

		// 删除
		r.DELETE("indoorTier/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorTier models.LocationIndoorTierModel
			)

			// 查询
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			// 删除
			if ret := models.Init(models.LocationIndoorTierModel{}).GetSession().Delete(&locationIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorTier/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationIndoorTier, repeat models.LocationIndoorTierModel
			)

			// 表单
			form := (&LocationIndoorTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层代码")
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层名称")

			// 查询
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			// 编辑
			locationIndoorTier.BaseModel.Sort = form.Sort
			locationIndoorTier.UniqueCode = form.UniqueCode
			locationIndoorTier.Name = form.Name
			locationIndoorTier.LocationIndoorCabinet = form.LocationIndoorCabinet
			if ret = models.Init(models.LocationIndoorTierModel{}).GetSession().Save(&locationIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_tier": locationIndoorTier}))
		})

		// 详情
		r.GET("indoorTier/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationIndoorTier models.LocationIndoorTierModel
			)
			ret = models.Init(models.LocationIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_tier": locationIndoorTier}))
		})

		// 列表
		r.GET("indoorTier", func(ctx *gin.Context) {
			var locationIndoorTier []models.LocationIndoorTierModel
			models.Init(models.LocationIndoorTierModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorTier)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_tier": locationIndoorTier}))
		})
	}
}
