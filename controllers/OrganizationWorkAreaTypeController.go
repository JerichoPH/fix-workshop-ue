package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationWorkAreaTypeController struct{}

// OrganizationWorkAreaTypeStoreForm 新建工区表单
type OrganizationWorkAreaTypeStoreForm struct {
	Sort       int64  `form:"sort" json:"sort"`
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkAreaTypeStoreForm
func (cls OrganizationWorkAreaTypeStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaTypeStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("工区类型代码必填")
	}
	if len(cls.UniqueCode) > 64{
		wrongs.PanicValidate("工区类型代码不能超过64位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("工区类型名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("工区类型名称不能超过64位")
	}

	return cls
}

// C 新建
func (OrganizationWorkAreaTypeController) C(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&OrganizationWorkAreaTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.OrganizationWorkAreaTypeModel
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区类型代码")
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区类型名称")

	// 新建
	organizationWorkAreaType := &models.OrganizationWorkAreaTypeModel{
		BaseModel:  models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	if ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).PrepareByDefaultDbDriver().Create(&organizationWorkAreaType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
}

// D 删除
func (OrganizationWorkAreaTypeController) D(ctx *gin.Context) {
	// 查询
	organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 删除
	if ret := models.BootByModel(models.OrganizationWorkAreaTypeModel{}).PrepareByDefaultDbDriver().Delete(&organizationWorkAreaType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (OrganizationWorkAreaTypeController) U(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&OrganizationWorkAreaTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.OrganizationWorkAreaTypeModel
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区类型代码")
	ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "工区类型名称")

	// 查询
	organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 编辑
	organizationWorkAreaType.BaseModel.Sort = form.Sort
	organizationWorkAreaType.UniqueCode = form.UniqueCode
	organizationWorkAreaType.Name = form.Name
	if ret = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&organizationWorkAreaType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
}

// S 详情
func (OrganizationWorkAreaTypeController) S(ctx *gin.Context) {
	organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
}

// I 列表
func (OrganizationWorkAreaTypeController) I(ctx *gin.Context) {
	var (
		organizationWorkAreaTypes []models.OrganizationWorkAreaTypeModel
		count                     int64
		db              *gorm.DB
	)
	db = models.BootByModel(models.OrganizationWorkAreaTypeModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&organizationWorkAreaTypes)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_work_area_types": organizationWorkAreaTypes}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&organizationWorkAreaTypes)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"organization_work_area_types": organizationWorkAreaTypes}, ctx.Query("__page__"), count))
	}
}
