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

// LocationDepotCabinetRouter 仓储仓库柜架路由
type LocationDepotCabinetRouter struct{}

// LocationDepotCabinetStoreForm 新建仓储仓库柜架表单
type LocationDepotCabinetStoreForm struct {
	Sort                   int64  `form:"sort" json:"sort"`
	UniqueCode             string `form:"unique_code" json:"unique_code"`
	Name                   string `form:"name" json:"name"`
	LocationDepotRowUUID   string `form:"location_depot_row_uuid" json:"location_depot_row_uuid"`
	LocationDepotRow       models.LocationDepotRowModel
	LocationDepotTierUUIDs []string `form:"location_depot_tier_uuids" json:"location_depot_tier_uuids"`
	LocationDepotTiers     []models.LocationDepotTierModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotCabinetStoreForm
func (cls LocationDepotCabinetStoreForm) ShouldBind(ctx *gin.Context) LocationDepotCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库柜架代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库柜架名称必填")
	}
	if cls.LocationDepotRowUUID == "" {
		wrongs.PanicValidate("所属仓库排必选")
	}
	ret = models.Init(models.LocationDepotRowModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationDepotRowUUID}).
		Prepare().
		First(&cls.LocationDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库排")
	if len(cls.LocationDepotTierUUIDs) > 0 {
		models.Init(models.LocationDepotTierModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationDepotTierUUIDs).
			Find(&cls.LocationDepotTiers)
	}

	return cls
}

// Load 加载路由
//  @receiver LocationDepotCabinetRouter
//  @param router
func (LocationDepotCabinetRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotCabinet", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotCabinetModel
			)

			// 表单
			form := (&LocationDepotCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

			// 新建
			locationDepotCabinet := &models.LocationDepotCabinetModel{
				BaseModel:          models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:         form.UniqueCode,
				Name:               form.Name,
				LocationDepotRow:   form.LocationDepotRow,
				LocationDepotTiers: form.LocationDepotTiers,
			}
			if ret = models.Init(models.LocationDepotCabinetModel{}).GetSession().Create(&locationDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_cabinet": locationDepotCabinet}))
		})

		// 删除
		r.DELETE("depotCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				locationDepotCabinet models.LocationDepotCabinetModel
			)

			// 查询
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			// 删除
			if ret := models.Init(models.LocationDepotCabinetModel{}).GetSession().Delete(&locationDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                          *gorm.DB
				locationDepotCabinet, repeat models.LocationDepotCabinetModel
			)

			// 表单
			form := (&LocationDepotCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

			// 查询
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			// 编辑
			locationDepotCabinet.BaseModel.Sort = form.Sort
			locationDepotCabinet.UniqueCode = form.UniqueCode
			locationDepotCabinet.Name = form.Name
			locationDepotCabinet.LocationDepotRow = form.LocationDepotRow
			locationDepotCabinet.LocationDepotTiers = form.LocationDepotTiers
			if ret = models.Init(models.LocationDepotCabinetModel{}).GetSession().Save(&locationDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_depot_cabinet": locationDepotCabinet}))
		})

		// 详情
		r.GET("depotCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				locationDepotCabinet models.LocationDepotCabinetModel
			)
			ret = models.Init(models.LocationDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_cabinet": locationDepotCabinet}))
		})

		// 列表
		r.GET("depotCabinet", func(ctx *gin.Context) {
			var locationDepotCabinets []models.LocationDepotCabinetModel
			models.Init(models.LocationDepotCabinetModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotCabinets)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_cabinets": locationDepotCabinets}))
		})
	}
}
