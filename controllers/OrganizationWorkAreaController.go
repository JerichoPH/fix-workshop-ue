package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationWorkAreaController struct{}

// OrganizationWorkAreaStoreForm 新建工区表单
type OrganizationWorkAreaStoreForm struct {
	Sort                               int64  `form:"sort" json:"sort"`
	UniqueCode                         string `form:"unique_code" json:"unique_code"`
	Name                               string `form:"name" json:"name"`
	BeEnable                           bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkAreaTypeUUID       string `form:"organization_work_area_type_uuid" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType           models.OrganizationWorkAreaTypeModel
	OrganizationWorkAreaProfessionUUID string `form:"organization_work_area_profession_uuid" json:"organization_work_area_profession_uuid"`
	OrganizationWorkAreaProfession     models.OrganizationWorkAreaProfessionModel
	OrganizationWorkshopUUID           string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop               models.OrganizationWorkshopModel
}

// ShouldBind 绑定表单
//  @receiver cl
//  @param ctx
//  @return OrganizationWorkAreaStoreForm
func (cls OrganizationWorkAreaStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("工区代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("工区名称必填")
	}

	if cls.OrganizationWorkAreaTypeUUID == "" {
		wrongs.PanicValidate("工区类型必选")
	}
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).SetWheres(map[string]interface{}{"uuid": cls.OrganizationWorkAreaTypeUUID}).PrepareByDefault().First(&cls.OrganizationWorkAreaType)
	wrongs.PanicWhenIsEmpty(ret, "工区类型")

	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(map[string]interface{}{"uuid": cls.OrganizationWorkshopUUID}).PrepareByDefault().First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")

	if cls.OrganizationWorkAreaProfessionUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaProfessionUUID}).PrepareByDefault().First(&cls.OrganizationWorkAreaProfession)
		wrongs.PanicWhenIsEmpty(ret, "工区专业")
	}
	return cls
}

// Store 新建
func (OrganizationWorkAreaController) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.OrganizationWorkAreaModel
	)

	// 表单
	form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区代码")
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区名称")

	// 新建
	if ret = models.BootByModel(models.OrganizationWorkAreaModel{}).PrepareByDefault().Create(
		&models.OrganizationWorkAreaModel{
			BaseModel:                          models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
			UniqueCode:                         form.UniqueCode,
			Name:                               form.Name,
			OrganizationWorkAreaProfessionUuid: form.OrganizationWorkAreaProfessionUUID,
			OrganizationWorkAreaTypeUuid:       form.OrganizationWorkAreaTypeUUID,
			OrganizationWorkshopUuid:           form.OrganizationWorkshopUUID,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{}))
}

// Destroy 删除
func (OrganizationWorkAreaController) Destroy(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkArea models.OrganizationWorkAreaModel
	)
	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationWorkArea)
	wrongs.PanicWhenIsEmpty(ret, "工区")

	// 删除
	if ret := models.BootByModel(models.OrganizationWorkAreaModel{}).PrepareByDefault().Delete(&organizationWorkArea); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// Update 更新
func (OrganizationWorkAreaController) Update(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		organizationWorkArea, repeat models.OrganizationWorkAreaModel
	)

	// 表单
	form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区代码")
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区名称")

	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationWorkArea)
	wrongs.PanicWhenIsEmpty(ret, "工区")

	// 编辑
	if ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		Updates(map[string]interface{}{
			"unique_code":                            form.UniqueCode,
			"name":                                   form.Name,
			"organization_work_area_type_uuid":       form.OrganizationWorkAreaType.Uuid,
			"organization_work_area_profession_uuid": form.OrganizationWorkAreaProfessionUUID,
			"be_enable":                              form.BeEnable,
			"organization_workshop_uuid":             form.OrganizationWorkshopUUID,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_work_area": organizationWorkArea}))
}

// Show 详情
func (OrganizationWorkAreaController) Show(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkArea models.OrganizationWorkAreaModel
	)
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		First(&organizationWorkArea)
	wrongs.PanicWhenIsEmpty(ret, "工区")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_work_area": organizationWorkArea}))
}

// Index 列表
func (OrganizationWorkAreaController) Index(ctx *gin.Context) {
	var organizationWorkAreas []models.OrganizationWorkAreaModel
	models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "organization_work_area_type_uuid", "organization_workshop_uuid").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&organizationWorkAreas)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_work_areas": organizationWorkAreas}))
}
