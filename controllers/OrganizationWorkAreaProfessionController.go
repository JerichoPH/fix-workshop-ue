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
//  @receiver ins
//  @param ctx
//  @return OrganizationWorkAreaProfessionStoreForm
func (ins OrganizationWorkAreaProfessionStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaProfessionStoreForm {

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("工区专业代码必填")
	}
	if len(ins.UniqueCode) > 64 {
		wrongs.PanicValidate("工区专业代码不能超过64位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("工区专业名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("工区专业名称不能超过64位")
	}

	return ins
}

// C 新建
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业代码")
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业名称")

	// 新建
	organizationWorkAreaProfession := &models.OrganizationWorkAreaProfessionModel{
		BaseModel:  models.BaseModel{Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).PrepareByDefaultDbDriver().Create(&organizationWorkAreaProfession)
	if ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}

// D 删除
func (OrganizationWorkAreaProfessionController) D(ctx *gin.Context) {
	var (
		ret                            *gorm.DB
		organizationWorkAreaProfession models.OrganizationWorkAreaProfessionModel
	)

	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	// 删除
	if ret := models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).PrepareByDefaultDbDriver().Delete(&organizationWorkAreaProfession); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业代码")
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区专业名称")

	// 查询
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	// 编辑
	organizationWorkAreaProfession.UniqueCode = form.UniqueCode
	organizationWorkAreaProfession.Name = form.Name
	if ret = models.
		BootByModel(models.OrganizationWorkAreaProfessionModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().
		Updates(map[string]interface{}{
			"unique_code": form.UniqueCode,
			"name":        form.Name,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}

// S 详情
func (OrganizationWorkAreaProfessionController) S(ctx *gin.Context) {
	var (
		ret                            *gorm.DB
		organizationWorkAreaProfession models.OrganizationWorkAreaProfessionModel
	)
	ret = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationWorkAreaProfession)
	wrongs.PanicWhenIsEmpty(ret, "工区专业")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_work_area_profession": organizationWorkAreaProfession}))
}

// I 列表
func (OrganizationWorkAreaProfessionController) I(ctx *gin.Context) {
	var (
		organizationWorkAreaProfessions []models.OrganizationWorkAreaProfessionModel
		count                           int64
		db                              *gorm.DB
	)
	db = models.BootByModel(models.OrganizationWorkAreaProfessionModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&organizationWorkAreaProfessions)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_work_area_professions": organizationWorkAreaProfessions}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&organizationWorkAreaProfessions)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"organization_work_area_professions": organizationWorkAreaProfessions}, ctx.Query("__page__"), count))
	}
}
