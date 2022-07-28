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

// LocationDepotSectionRouter 仓储仓库区域路由
type LocationDepotSectionRouter struct{}

// LocationDepotSectionStoreForm 新建仓储仓库区域表单
type LocationDepotSectionStoreForm struct {
	Sort                        int64  `form:"sort" json:"sort"`
	UniqueCode                  string `form:"unique_code" json:"unique_code"`
	Name                        string `form:"name" json:"name"`
	LocationDepotStorehouseUUID string `form:"location_depot_storehouse_uuid" json:"location_depot_storehouse_uuid"`
	LocationDepotStorehouse     models.LocationDepotStorehouseModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotSectionStoreForm
func (cls LocationDepotSectionStoreForm) ShouldBind(ctx *gin.Context) LocationDepotSectionStoreForm {
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
	if cls.LocationDepotStorehouseUUID == "" {
		wrongs.PanicValidate("所属仓库必选")
	}
	ret = models.Init(models.LocationDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationDepotStorehouseUUID}).
		Prepare().
		First(&cls.LocationDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库")

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls LocationDepotSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotSection", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotSectionModel
			)

			// 表单
			form := (&LocationDepotSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

			// 新建
			locationDepotSection := &models.LocationDepotSectionModel{
				BaseModel:               models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:              form.UniqueCode,
				Name:                    form.Name,
				LocationDepotStorehouse: form.LocationDepotStorehouse,
			}
			if ret = models.Init(models.LocationDepotSectionModel{}).GetSession().Create(&locationDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_section": locationDepotSection}))
		})

		// 删除
		r.DELETE("depotSection/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				locationDepotSection models.LocationDepotSectionModel
			)

			// 查询
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			// 删除
			if ret := models.Init(models.LocationDepotSectionModel{}).GetSession().Delete(&locationDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotSection/:uuid", func(ctx *gin.Context) {
			var (
				ret                          *gorm.DB
				locationDepotSection, repeat models.LocationDepotSectionModel
			)

			// 表单
			form := (&LocationDepotSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

			// 查询
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			// 编辑
			locationDepotSection.BaseModel.Sort = form.Sort
			locationDepotSection.UniqueCode = form.UniqueCode
			locationDepotSection.Name = form.Name
			locationDepotSection.LocationDepotStorehouse = form.LocationDepotStorehouse
			if ret = models.Init(models.LocationDepotSectionModel{}).GetSession().Save(&locationDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_depot_section": locationDepotSection}))
		})

		// 详情
		r.GET("depotSection/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				locationDepotSection models.LocationDepotSectionModel
			)
			ret = models.Init(models.LocationDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_section": locationDepotSection}))
		})

		// 列表
		r.GET("depotSection", func(ctx *gin.Context) {
			var locationDepotSections []models.LocationDepotSectionModel
			models.Init(models.LocationDepotSectionModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotSections)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_sections": locationDepotSections}))
		})
	}
}
