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

// OrganizationSectionRouter 区间路由
type OrganizationSectionRouter struct{}

// OrganizationSectionStoreForm 新建区间表单
type OrganizationSectionStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationSectionStoreForm
func (cls OrganizationSectionStoreForm) ShouldBind(ctx *gin.Context) OrganizationSectionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("区间代码不能为空")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("区间名称不能为空")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间不能为空")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.Init(models.OrganizationWorkAreaModel{}).
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
func (OrganizationSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("section", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationSectionModel
			)

			// 表单
			form := (&OrganizationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间代码")
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间名称")

			// 新建
			organizationSection := &models.LocationSectionModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				BeEnable:             form.BeEnable,
				OrganizationWorkshop: form.OrganizationWorkshop,
				OrganizationWorkArea: form.OrganizationWorkArea,
			}
			if ret = models.Init(models.LocationSectionModel{}).GetSession().Create(&organizationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_section": organizationSection}))
		})

		// 删除
		r.DELETE("section/:uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				organizationSection models.LocationSectionModel
			)
			// 查询
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			// 删除
			if ret := models.Init(models.LocationSectionModel{}).GetSession().Delete(&organizationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("section/:uuid", func(ctx *gin.Context) {
			var (
				ret                         *gorm.DB
				organizationSection, repeat models.LocationSectionModel
			)

			// 表单
			form := (&OrganizationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间代码")
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间名称")

			// 查询
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			// 编辑
			organizationSection.BaseModel.Sort = form.Sort
			organizationSection.UniqueCode = form.UniqueCode
			organizationSection.Name = form.Name
			organizationSection.BeEnable = form.BeEnable
			organizationSection.OrganizationWorkshop = form.OrganizationWorkshop
			organizationSection.OrganizationWorkArea = form.OrganizationWorkArea
			if ret = models.Init(models.LocationSectionModel{}).GetSession().Save(&organizationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_section": organizationSection}))
		})

		// 详情
		r.GET("section/:uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				organizationSection models.LocationSectionModel
			)
			ret = models.Init(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_section": organizationSection}))
		})

		// 列表
		r.GET("section", func(ctx *gin.Context) {
			var organizationSections []models.LocationSectionModel
			models.Init(models.LocationSectionModel{}).
				SetWhereFields("unique_code", "Name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationSections)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_sections": organizationSections}))
		})
	}
}
