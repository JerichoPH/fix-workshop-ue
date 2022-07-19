package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// OrganizationWorkAreaRouter 工区路由
type OrganizationWorkAreaRouter struct{}

// OrganizationWorkAreaStoreForm 新建工区表单
type OrganizationWorkAreaStoreForm struct {
	Sort                         int64  `form:"sort" json:"sort"`
	UniqueCode                   string `form:"unique_code" json:"unique_code"`
	Name                         string `form:"name" json:"name"`
	BeEnable                     bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkAreaTypeUUID string `form:"organization_work_area_type_uuid" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType     models.OrganizationWorkAreaTypeModel
	OrganizationWorkshopUUID     string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop         models.OrganizationWorkshopModel
	OrganizationSectionUUIDs     []string `form:"organization_section_uuids" json:"organization_section_uuids"`
	OrganizationSections         []models.OrganizationSectionModel
	OrganizationStationUUIDs     []string `form:"organization_station_uuids" json:"organization_station_uuids"`
	OrganizationStations         []models.OrganizationStationModel
}

// ShouldBind 绑定表单
//  @receiver cl
//  @param ctx
//  @return OrganizationWorkAreaStoreForm
func (cls OrganizationWorkAreaStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("工区代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("工区名称必填")
	}
	if cls.OrganizationWorkAreaTypeUUID == "" {
		abnormals.PanicValidate("工区类型必选")
	}
	cls.OrganizationWorkAreaType = (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(cls.OrganizationWorkAreaTypeUUID)
	if cls.OrganizationWorkshopUUID == "" {
		abnormals.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	abnormals.PanicWhenIsEmpty(ret, "车间")
	if len(cls.OrganizationSections) > 0 {
		models.Init(models.OrganizationSectionModel{}).DB().Where("uuid in ?", cls.OrganizationSectionUUIDs).Find(&cls.OrganizationSections)
	}
	if len(cls.OrganizationStationUUIDs) > 0 {
		models.Init(models.OrganizationStationModel{}).DB().Where("uuid in ?", cls.OrganizationStationUUIDs).Find(&cls.OrganizationStations)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls OrganizationWorkAreaRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("workArea", func(ctx *gin.Context) {
			var (
				ret *gorm.DB
				repeat models.OrganizationWorkAreaModel
			)

			// 表单
			form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区代码")
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区名称")

			// 新建
			organizationWorkArea := &models.OrganizationWorkAreaModel{
				BaseModel:                models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:               form.UniqueCode,
				Name:                     form.Name,
				BeEnable:                 form.BeEnable,
				OrganizationWorkAreaType: form.OrganizationWorkAreaType,
				OrganizationWorkshop:     form.OrganizationWorkshop,
				OrganizationSections:     form.OrganizationSections,
				OrganizationStations:     form.OrganizationStations,
			}
			if ret = models.Init(models.OrganizationWorkAreaModel{}).DB().Create(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 删除
		r.DELETE("workArea/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkArea models.OrganizationWorkAreaModel
			)
			// 查询
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationWorkArea)
			abnormals.PanicWhenIsEmpty(ret, "工区")

			// 删除
			if ret := models.Init(models.OrganizationWorkAreaModel{}).DB().Delete(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("workArea/:uuid", func(ctx *gin.Context) {
			var (
				ret                          *gorm.DB
				organizationWorkArea, repeat models.OrganizationWorkAreaModel
			)

			// 表单
			form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区代码")
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区名称")

			// 查询
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationWorkArea)
			abnormals.PanicWhenIsEmpty(ret, "工区")

			// 编辑
			organizationWorkArea.BaseModel.Sort = form.Sort
			organizationWorkArea.UniqueCode = form.UniqueCode
			organizationWorkArea.Name = form.Name
			organizationWorkArea.BeEnable = form.BeEnable
			organizationWorkArea.OrganizationWorkAreaType = form.OrganizationWorkAreaType
			organizationWorkArea.OrganizationWorkshop = form.OrganizationWorkshop
			organizationWorkArea.OrganizationSections = form.OrganizationSections
			organizationWorkArea.OrganizationStations = form.OrganizationStations
			if ret = models.Init(models.OrganizationWorkAreaModel{}).DB().Save(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 详情
		r.GET("workArea/:uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkArea models.OrganizationWorkAreaModel
			)
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationWorkArea)
			abnormals.PanicWhenIsEmpty(ret, "工区")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 列表
		r.GET("workArea", func(ctx *gin.Context) {
			var organizationWorkAreas []models.OrganizationWorkAreaModel
			models.Init(models.OrganizationWorkAreaModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "organization_work_area_type_uuid", "organization_workshop_uuid").
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationWorkAreas)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_areas": organizationWorkAreas}))
		})
	}
}
