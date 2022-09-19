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
	OrganizationWorkAreaTypeUuid       string `form:"organization_work_area_type_uuid" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType           models.OrganizationWorkAreaTypeModel
	OrganizationWorkAreaProfessionUuid string `form:"organization_work_area_profession_uuid" json:"organization_work_area_profession_uuid"`
	OrganizationWorkAreaProfession     models.OrganizationWorkAreaProfessionModel
	OrganizationWorkshopUuid           string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
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
	if len(cls.UniqueCode) != 4 {
		wrongs.PanicValidate("工区代码必须是4位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("工区名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("工区名称不能超过64位")
	}

	if cls.OrganizationWorkAreaTypeUuid == "" {
		wrongs.PanicValidate("工区类型必选")
	}
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).SetWheres(map[string]interface{}{"uuid": cls.OrganizationWorkAreaTypeUuid}).PrepareByDefault().First(&cls.OrganizationWorkAreaType)
	wrongs.PanicWhenIsEmpty(ret, "工区类型")

	if cls.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(map[string]interface{}{"uuid": cls.OrganizationWorkshopUuid}).SetPreloads("OrganizationParagraph", "OrganizationParagraph.OrganizationRailway").PrepareByDefault().First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")

	if cls.OrganizationWorkAreaProfessionUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaProfessionUuid}).PrepareByDefault().First(&cls.OrganizationWorkAreaProfession)
		wrongs.PanicWhenIsEmpty(ret, "工区专业")
	}
	return cls
}

// C 新建
func (OrganizationWorkAreaController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.OrganizationWorkAreaModel
	)

	// 表单
	form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWheres(tools.Map{"unique_code": form.OrganizationWorkshop.OrganizationParagraph.UniqueCode + form.UniqueCode}).
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
			UniqueCode:                         form.OrganizationWorkshop.OrganizationParagraph.UniqueCode + form.UniqueCode,
			Name:                               form.Name,
			OrganizationWorkAreaProfessionUuid: form.OrganizationWorkAreaProfessionUuid,
			OrganizationWorkAreaTypeUuid:       form.OrganizationWorkAreaTypeUuid,
			OrganizationWorkshopUuid:           form.OrganizationWorkshopUuid,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{}))
}

// D 删除
func (OrganizationWorkAreaController) D(ctx *gin.Context) {
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

// U 更新
func (OrganizationWorkAreaController) U(ctx *gin.Context) {
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
			"name":                                   form.Name,
			"organization_work_area_type_uuid":       form.OrganizationWorkAreaType.Uuid,
			"organization_work_area_profession_uuid": form.OrganizationWorkAreaProfessionUuid,
			"be_enable":                              form.BeEnable,
			"organization_workshop_uuid":             form.OrganizationWorkshopUuid,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_work_area": organizationWorkArea}))
}

// S 详情
func (OrganizationWorkAreaController) S(ctx *gin.Context) {
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

// I 列表
func (OrganizationWorkAreaController) I(ctx *gin.Context) {
	var organizationWorkAreas []models.OrganizationWorkAreaModel
	models.BootByModel(models.OrganizationWorkAreaModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "organization_work_area_type_uuid", "organization_workshop_uuid").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&organizationWorkAreas)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_work_areas": organizationWorkAreas}))
}
