package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationWorkAreaProfessionController struct{}

// OrganizationWorkAreaProfessionStoreForm 新建工区专业表单
type OrganizationWorkAreaProfessionStoreForm struct {
	UniqueCode string `json:"unique_code"`
	Name       string `json:"name"`
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkAreaProfessionStoreForm
func (cls OrganizationWorkAreaProfessionStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaProfessionStoreForm {

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("工区专业代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("工区专业名称必填")
	}

	return cls
}

func (OrganizationWorkAreaProfessionController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.OrganizationWorkAreaProfessionModel
	)

	// 表单
	form := (&OrganizationWorkAreaProfessionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业代码")
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业名称")

	// 新建
	organizationWorkAreaProfession := &models.OrganizationWorkAreaProfessionModel{
		BaseModel:  models.BaseModel{Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).PrepareByDefault().Create(&organizationWorkAreaProfession)
	if ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}
func (OrganizationWorkAreaProfessionController) D(ctx *gin.Context) {
	var (
		ret                            *gorm.DB
		organizationWorkAreaProfession models.OrganizationWorkAreaProfessionModel
	)

	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	// 删除
	if ret := models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).PrepareByDefault().Delete(&organizationWorkAreaProfession); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (OrganizationWorkAreaProfessionController) U(ctx *gin.Context) {
	var (
		ret                                    *gorm.DB
		organizationWorkAreaProfession, repeat models.OrganizationWorkAreaProfessionModel
	)

	// 表单
	form := (&OrganizationWorkAreaProfessionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业代码")
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业名称")

	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	// 编辑
	organizationWorkAreaProfession.UniqueCode = form.UniqueCode
	organizationWorkAreaProfession.Name = form.Name
	if ret = models.
		BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		PrepareByDefault().
		Where("uuid = ?", ctx.Param("uuid")).
		Updates(map[string]interface{}{
			"unique_code": form.UniqueCode,
			"name":        form.Name,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}
func (OrganizationWorkAreaProfessionController) S(ctx *gin.Context) {
	var (
		ret                            *gorm.DB
		organizationWorkAreaProfession models.OrganizationWorkAreaProfessionModel
	)
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}
func (OrganizationWorkAreaProfessionController) I(ctx *gin.Context) {
	var organizationWorkAreaProfessions []models.OrganizationWorkAreaProfessionModel
	models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWhereFields().
		PrepareUseQueryByDefault(ctx).
		Find(&organizationWorkAreaProfessions)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_work_area_professions": organizationWorkAreaProfessions}))
}
