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

// OrganizationWorkshopRouter 车间路由
type OrganizationWorkshopRouter struct{}

// OrganizationWorkshopStoreForm 新建车间表单
type OrganizationWorkshopStoreForm struct {
	Sort                         int64  `form:"sort" json:"sort"`
	UniqueCode                   string `form:"unique_code" json:"unique_code"`
	Name                         string `form:"name" json:"name"`
	BeEnable                     bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopTypeUUID string `form:"organization_workshop_type_uuid" json:"organization_workshop_type_uuid"`
	OrganizationWorkshopType     models.OrganizationWorkshopTypeModel
	OrganizationParagraphUUID    string `form:"organization_paragraph_uuid" json:"organization_paragraph_uuid"`
	OrganizationParagraph        models.OrganizationParagraphModel
	OrganizationSectionUUIDs     []string `form:"organization_section_uuids" json:"organization_section_uuids"`
	OrganizationSections         []models.OrganizationSectionModel
	OrganizationWorkAreaUUIDs    []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas        []models.OrganizationWorkAreaModel
	OrganizationStationUUIDs     []string `form:"organization_station_uuids" json:"organization_station_uuids"`
	OrganizationStations         []models.OrganizationStationModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkshopStoreForm
func (cls OrganizationWorkshopStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("车间代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("车间名称必填")
	}
	if cls.OrganizationWorkshopTypeUUID == "" {
		wrongs.PanicValidate("车间类型必选")
	}
	cls.OrganizationWorkshopType = (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(cls.OrganizationWorkshopTypeUUID)
	if cls.OrganizationParagraphUUID == "" {
		wrongs.PanicValidate("所属站段必选")
	}
	ret = models.Init(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationParagraphUUID}).
		Prepare().
		First(&cls.OrganizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")
	if len(cls.OrganizationSectionUUIDs) > 0 {
		models.Init(models.OrganizationSectionModel{}).GetSession().Where("uuid in ?", cls.OrganizationSectionUUIDs).Find(&cls.OrganizationSections)
	}
	if len(cls.OrganizationWorkAreaUUIDs) > 0 {
		models.Init(models.OrganizationWorkAreaModel{}).GetSession().Where("uuid in ?", cls.OrganizationWorkAreaUUIDs).Find(&cls.OrganizationWorkAreas)
	}
	if len(cls.OrganizationStationUUIDs) > 0 {
		models.Init(models.OrganizationStationModel{}).GetSession().Where("uuid in ?", cls.OrganizationStationUUIDs).Find(&cls.OrganizationStations)
	}

	return cls
}

// Load 加载路由
//  @receiver OrganizationWorkshopRouter
//  @param router
func (OrganizationWorkshopRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("workshop", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationWorkshopModel
			)

			// 表单
			form := (&OrganizationWorkshopStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间代码")
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间名称")

			// 新建
			organizationWorkshop := &models.OrganizationWorkshopModel{
				BaseModel:                models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:               form.UniqueCode,
				Name:                     form.Name,
				OrganizationWorkshopType: form.OrganizationWorkshopType,
				OrganizationParagraph:    form.OrganizationParagraph,
				OrganizationSections:     form.OrganizationSections,
				OrganizationWorkAreas:    form.OrganizationWorkAreas,
				OrganizationStations:     form.OrganizationStations,
			}
			if ret = models.Init(models.OrganizationWorkshopModel{}).GetSession().Create(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 删除
		r.DELETE("workshop/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			// 查询
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			// 删除
			if ret = models.Init(models.OrganizationWorkshopModel{}).GetSession().Delete(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("workshop/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			// 表单
			form := (&OrganizationWorkshopStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkshopModel
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间代码")
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间名称")

			// 查询
			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			// 编辑
			organizationWorkshop.UniqueCode = form.UniqueCode
			organizationWorkshop.Name = form.Name
			organizationWorkshop.BeEnable = form.BeEnable
			organizationWorkshop.OrganizationWorkshopType = form.OrganizationWorkshopType
			organizationWorkshop.OrganizationParagraph = form.OrganizationParagraph
			organizationWorkshop.OrganizationSections = form.OrganizationSections
			organizationWorkshop.OrganizationWorkAreas = form.OrganizationWorkAreas
			organizationWorkshop.OrganizationStations = form.OrganizationStations
			if ret = models.Init(models.OrganizationWorkshopModel{}).GetSession().Save(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 详情
		r.GET("workshop/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			ret = models.Init(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 列表
		r.GET("workshop", func(ctx *gin.Context) {
			var organizationWorkshops []models.OrganizationWorkshopModel

			models.Init(models.OrganizationWorkshopModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_type_uuid", "organization_paragraph_uuid").
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationWorkshops)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshops": organizationWorkshops}))
		})
	}
}
