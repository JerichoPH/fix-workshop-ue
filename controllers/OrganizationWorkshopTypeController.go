package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationWorkshopTypeController struct{}

// OrganizationWorkshopTypeStoreForm 新建车间路由表单
type OrganizationWorkshopTypeStoreForm struct {
	Sort       int64  `gorm:"sort" json:"sort"`
	UniqueCode string `gorm:"unique_code" json:"unique_code"`
	Name       string `gorm:"name" json:"name"`
	NumberCode string `gorm:"number_code" json:"number_code"`
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return OrganizationWorkshopTypeStoreForm
func (ins OrganizationWorkshopTypeStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopTypeStoreForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("车间类型代码必填")
	}
	if len(ins.UniqueCode) > 64 {
		wrongs.PanicValidate("车间类型代码不能超过64位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("车间类型名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("车间类型名称不能超过64位")
	}

	return ins
}

// C 新建
func (OrganizationWorkshopTypeController) C(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&OrganizationWorkshopTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&models.OrganizationWorkAreaTypeModel{})
	wrongs.PanicWhenIsRepeat(ret, "车间类型代码")
	ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&models.OrganizationWorkshopTypeModel{})
	wrongs.PanicWhenIsRepeat(ret, "车间类型名称")

	// 新建
	organizationWorkshopType := &models.OrganizationWorkshopTypeModel{
		BaseModel:  models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
		NumberCode: form.NumberCode,
	}
	if ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).PrepareByDefaultDbDriver().Create(&organizationWorkshopType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_workshop_type": organizationWorkshopType}))
}

// D 删除
func (OrganizationWorkshopTypeController) D(ctx *gin.Context) {
	var ret *gorm.DB

	// 查询
	organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 删除
	if ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).PrepareByDefaultDbDriver().Delete(&organizationWorkshopType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (OrganizationWorkshopTypeController) U(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&OrganizationWorkshopTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&models.OrganizationWorkAreaTypeModel{})
	wrongs.PanicWhenIsRepeat(ret, "车间类型代码")
	ret = models.BootByModel(models.OrganizationWorkshopTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&models.OrganizationWorkshopTypeModel{})
	wrongs.PanicWhenIsRepeat(ret, "车间类型名称")

	// 查询
	organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 编辑
	organizationWorkshopType.BaseModel.Sort = form.Sort
	organizationWorkshopType.UniqueCode = form.UniqueCode
	organizationWorkshopType.Name = form.Name
	organizationWorkshopType.NumberCode = form.NumberCode
	models.BootByModel(models.OrganizationWorkshopTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&organizationWorkshopType)

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_workshop_type": organizationWorkshopType}))
}

// S 详情
func (OrganizationWorkshopTypeController) S(ctx *gin.Context) {
	organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_workshop_type": organizationWorkshopType}))
}

// I 列表
func (OrganizationWorkshopTypeController) I(ctx *gin.Context) {
	var (
		organizationWorkshopTypes []models.OrganizationWorkshopTypeModel
		count                     int64
		db                        *gorm.DB
	)
	db = models.BootByModel(models.OrganizationWorkshopTypeModel{}).
		SetWhereFields("sort", "unique_code", "name", "number").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&organizationWorkshopTypes)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_workshop_types": organizationWorkshopTypes}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&organizationWorkshopTypes)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"organization_workshop_types": organizationWorkshopTypes}, ctx.Query("__page__"), count))
	}
}
