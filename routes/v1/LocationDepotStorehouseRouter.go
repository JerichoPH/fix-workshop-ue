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

// LocationDepotStorehouseRouter 仓储仓库路由 
type LocationDepotStorehouseRouter struct{}

// LocationDepotStorehouseStoreForm 仓储仓库新建表单
type LocationDepotStorehouseStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	OrganizationWorkshopUUID  string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop      models.OrganizationWorkshopModel
	LocationDepotSectionUUIDs []string `form:"location_depot_section_uuids" json:"location_depot_section_uuids"`
	LocationDepotSections     []models.LocationDepotSectionModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotStorehouseStoreForm
func (cls LocationDepotStorehouseStoreForm) ShouldBind(ctx *gin.Context) LocationDepotStorehouseStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")

	if len(cls.LocationDepotSectionUUIDs) > 0 {
		models.Init(models.LocationDepotSectionModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationDepotSectionUUIDs).
			Find(&cls.LocationDepotSections)
	}

	return cls
}

// Load 加载路由 
//  @receiver cls 
//  @param router 
func (cls LocationDepotStorehouseRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotStorehouse", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotStorehouseModel
			)

			// 表单
			form := (&LocationDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库代码")

			// 新建
			locationStorehouse := &models.LocationDepotStorehouseModel{
				BaseModel:             models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:            form.UniqueCode,
				Name:                  form.Name,
				OrganizationWorkshop:  form.OrganizationWorkshop,
				LocationDepotSections: form.LocationDepotSections,
			}
			if ret = models.Init(models.LocationDepotStorehouseModel{}).GetSession().Create(&locationStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_storehouse": locationStorehouse}))
		})

		// 删除
		r.DELETE("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationStorehouse models.LocationDepotStorehouseModel
			)

			// 查询
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			// 删除
			if ret := models.Init(models.LocationDepotStorehouseModel{}).GetSession().Delete(&locationStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationStorehouse, repeat models.LocationDepotStorehouseModel
			)

			// 表单
			form := (&LocationDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库代码")

			// 查询
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			// 编辑
			locationStorehouse.BaseModel.Sort = form.Sort
			locationStorehouse.UniqueCode = form.UniqueCode
			locationStorehouse.Name = form.Name
			locationStorehouse.OrganizationWorkshop = form.OrganizationWorkshop
			locationStorehouse.LocationDepotSections = form.LocationDepotSections
			if ret = models.Init(models.LocationDepotStorehouseModel{}).GetSession().Save(&locationStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_storehouse": locationStorehouse}))
		})

		// 详情
		r.GET("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				locationStorehouse models.LocationDepotStorehouseModel
			)
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_storehouse": locationStorehouse}))
		})

		// 列表
		r.GET("depotStorehouse", func(ctx *gin.Context) {
			var locationDepotStorehouses []models.LocationDepotStorehouseModel
			models.Init(models.LocationDepotStorehouseModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotStorehouses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_storehouses": locationDepotStorehouses}))
		})
	}
}
