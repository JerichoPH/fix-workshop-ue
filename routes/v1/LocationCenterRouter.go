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

// LocationCenterRouter 中心路由
type LocationCenterRouter struct{}

// LocationCenterStoreForm 新建中心表单
type LocationCenterStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"" json:""`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return LocationCenterStoreForm
func (cls LocationCenterStoreForm) ShouldBind(ctx *gin.Context) LocationCenterStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(ctx); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("中心代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("中心名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")
	if cls.OrganizationWorkAreaUUID != "" {
		models.Init(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationCenterRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("center", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationCenterModel
			)

			// 表单
			form := (&LocationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 新建
			locationCenter := &models.LocationCenterModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(models.LocationCenterModel{}).GetSession().Create(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_center": locationCenter}))
		})

		// 删除
		r.DELETE("center/:uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				locationCenter models.LocationCenterModel
			)

			// 查询
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 删除
			if ret := models.Init(models.LocationCenterModel{}).GetSession().Delete(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("center/:uuid", func(ctx *gin.Context) {
			var (
				ret                    *gorm.DB
				locationCenter, repeat models.LocationCenterModel
			)

			// 表单
			form := (&LocationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 查询
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 编辑
			locationCenter.BaseModel.Sort = form.Sort
			locationCenter.UniqueCode = form.UniqueCode
			locationCenter.Name = form.Name
			locationCenter.BeEnable = form.BeEnable
			if ret = models.Init(models.LocationCenterModel{}).GetSession().Save(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_center": locationCenter}))
		})

		// 详情
		r.GET("center/:uuid", func(ctx *gin.Context) {
			var locationCenter models.LocationCenterModel
			ret := models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_center": locationCenter}))
		})

		// 列表
		r.GET("center", func(ctx *gin.Context) {
			var locationCenters []models.LocationCenterModel
			models.Init(models.LocationCenterModel{}).
				SetWhereFields().
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&locationCenters)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_centers": locationCenters}))
		})
	}
}
