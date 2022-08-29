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

// PositionIndoorTierRouter 室内上道位置柜架层路由
type PositionIndoorTierRouter struct{}

// PositionIndoorTierStoreForm 新建室内上道位置柜架层表单
type PositionIndoorTierStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	PositionIndoorCabinetUUID string `form:"position_indoor_cabinet_uuid" json:"position_indoor_cabinet_uuid"`
	PositionIndoorCabinet     models.PositionIndoorCabinetModel
}

// ShouldBind 
//  @receiver cls 
//  @param ctx 
//  @return PositionIndoorTierStoreForm
func (cls PositionIndoorTierStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorTierStoreForm {
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
	if cls.PositionIndoorCabinetUUID == "" {
		wrongs.PanicValidate("所属柜架必选")
	}
	ret = models.BootByModel(models.PositionIndoorCabinetModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorCabinetUUID}).
		PrepareByDefault().
		First(&cls.PositionIndoorCabinet)
	wrongs.PanicWhenIsEmpty(ret, "所属柜架")

	return cls
}

// Load 加载路由
//  @receiver PositionIndoorTierRouter
//  @param engine
func (PositionIndoorTierRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/position",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionIndoorTierModel
			)

			// 表单
			form := (&PositionIndoorTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层代码")
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层名称")

			// 新建
			locationIndoorTier := &models.PositionIndoorTierModel{
				BaseModel:             models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:            form.UniqueCode,
				Name:                  form.Name,
				PositionIndoorCabinet: form.PositionIndoorCabinet,
			}
			if ret = models.BootByModel(models.PositionIndoorTierModel{}).PrepareByDefault().Create(&locationIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_tier": locationIndoorTier}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				positionIndoorTier models.PositionIndoorTierModel
			)

			// 查询
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&positionIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			// 删除
			if ret := models.BootByModel(models.PositionIndoorTierModel{}).PrepareByDefault().Delete(&positionIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				positionIndoorTier, repeat models.PositionIndoorTierModel
			)

			// 表单
			form := (&PositionIndoorTierStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层代码")
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架层名称")

			// 查询
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&positionIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			// 编辑
			positionIndoorTier.BaseModel.Sort = form.Sort
			positionIndoorTier.UniqueCode = form.UniqueCode
			positionIndoorTier.Name = form.Name
			positionIndoorTier.PositionIndoorCabinet = form.PositionIndoorCabinet
			if ret = models.BootByModel(models.PositionIndoorTierModel{}).PrepareByDefault().Save(&positionIndoorTier); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_tier": positionIndoorTier}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				positionIndoorTier models.PositionIndoorTierModel
			)
			ret = models.BootByModel(models.PositionIndoorTierModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&positionIndoorTier)
			wrongs.PanicWhenIsEmpty(ret, "柜架层")

			ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_tier": positionIndoorTier}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionIndoorTier []models.PositionIndoorTierModel
			models.BootByModel(models.PositionIndoorTierModel{}).
				SetWhereFields().
				PrepareQuery(ctx,"").
				Find(&positionIndoorTier)

			ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_tier": positionIndoorTier}))
		})
	}
}
