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

// OrganizationCenterRouter 中心路由
type OrganizationCenterRouter struct{}

// OrganizationCenterStoreForm 新建中心表单
type OrganizationCenterStoreForm struct {
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
//  @return OrganizationCenterStoreForm
func (cls OrganizationCenterStoreForm) ShouldBind(ctx *gin.Context) OrganizationCenterStoreForm {
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
func (cls OrganizationCenterRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("center", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationCenterModel
			)

			// 表单
			form := (&OrganizationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 新建
			organizationCenter := &models.OrganizationCenterModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(models.OrganizationCenterModel{}).GetSession().Create(&organizationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organizationCenter": organizationCenter}))
		})

		// 删除
		r.DELETE("center/:uuid", func(ctx *gin.Context) {
			var (
				ret                *gorm.DB
				organizationCenter models.OrganizationCenterModel
			)

			// 查询
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 删除
			if ret := models.Init(models.OrganizationCenterModel{}).GetSession().Delete(&organizationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("center/:uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				organizationCenter, repeat models.OrganizationCenterModel
			)

			// 表单
			form := (&OrganizationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 查询
			ret = models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 编辑
			organizationCenter.BaseModel.Sort = form.Sort
			organizationCenter.UniqueCode = form.UniqueCode
			organizationCenter.Name = form.Name
			organizationCenter.BeEnable = form.BeEnable
			if ret = models.Init(models.OrganizationCenterModel{}).GetSession().Save(&organizationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organizationCenter": organizationCenter}))
		})

		// 详情
		r.GET("center/:uuid", func(ctx *gin.Context) {
			var organizationCenter models.OrganizationCenterModel
			ret := models.Init(models.OrganizationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationCenter": organizationCenter}))
		})

		// 列表
		r.GET("center", func(ctx *gin.Context) {
			var organizationCenters []models.OrganizationCenterModel
			models.Init(models.OrganizationCenterModel{}).
				SetWhereFields().
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationCenters)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationCenters": organizationCenters}))
		})
	}
}
