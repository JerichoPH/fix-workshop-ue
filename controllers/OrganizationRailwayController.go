package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationRailwayController struct{}

// OrganizationRailwayStoreForm 新建路局表单
type OrganizationRailwayStoreForm struct {
	Sort              int64    `form:"sort" json:"sort"`
	UniqueCode        string   `form:"unique_code" json:"unique_code"`
	Name              string   `form:"name" json:"name"`
	ShortName         string   `form:"short_name" json:"short_name"`
	BeEnable          bool     `form:"be_enable" json:"be_enable"`
	LocationLineUuids []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationRailwayStoreForm
func (cls OrganizationRailwayStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailwayStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("路局代码必填")
	}
	if len(cls.UniqueCode) != 3 {
		wrongs.PanicValidate("路局代码必须是3位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("路局名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("路局名称不能超过64位")
	}
	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// OrganizationRailwayBindLinesFrom 路局绑定线别
type OrganizationRailwayBindLinesFrom struct {
	LocationLineUUIDs []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
func (cls OrganizationRailwayBindLinesFrom) ShouldBind(ctx *gin.Context) OrganizationRailwayBindLinesFrom {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	models.BootByModel(models.LocationLineModel{}).
		PrepareByDefaultDbDriver().
		Where("uuid in ?", cls.LocationLineUUIDs).
		Find(&cls.LocationLines)

	return cls
}

// C 新建
func (OrganizationRailwayController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.OrganizationRailwayModel
	)

	// 表单
	form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局代码")
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局名称")
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"short_name": form.ShortName}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局简称")

	// 新建
	organizationRailway := &models.OrganizationRailwayModel{
		BaseModel:  models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
		ShortName:  form.ShortName,
		BeEnable:   form.BeEnable,
	}
	if ret = (&models.BaseModel{}).SetModel(models.OrganizationRailwayModel{}).PrepareByDefaultDbDriver().Create(&organizationRailway); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_railway": organizationRailway}))
}

// D 删除
func (OrganizationRailwayController) D(ctx *gin.Context) {
	var ret *gorm.DB

	// 查询
	var organizationRailway models.OrganizationRailwayModel
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")
	// 删除
	if ret = models.BootByModel(models.OrganizationRailwayModel{}).PrepareByDefaultDbDriver().Delete(&organizationRailway); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (OrganizationRailwayController) U(ctx *gin.Context) {
	var (
		ret                         *gorm.DB
		organizationRailway, repeat models.OrganizationRailwayModel
	)

	// 表单
	form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局代码")
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局名称")
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"short_name": form.ShortName}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "路局简称")

	// 查询
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")

	// 修改
	organizationRailway.Sort = form.Sort
	organizationRailway.Name = form.Name
	organizationRailway.ShortName = form.ShortName
	organizationRailway.BeEnable = form.BeEnable
	if ret = models.BootByModel(models.OrganizationRailwayModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&organizationRailway); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_railway": organizationRailway}))
}

// S 详情
func (OrganizationRailwayController) S(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationRailway models.OrganizationRailwayModel
	)

	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		PrepareUseQueryByDefaultDbDriver(ctx).
		First(&organizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_railway": organizationRailway}))
}

// I 列表
func (OrganizationRailwayController) I(ctx *gin.Context) {
	var (
		organizationRailways []models.OrganizationRailwayModel
		count                int64
		db                   *gorm.DB
	)
	db = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWhereFields("uuid", "sort", "unique_code", "name", "short_name", "be_enable").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&organizationRailways)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_railways": organizationRailways}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&organizationRailways)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"organization_railways": organizationRailways}, ctx.Query("__page__"), count))
	}
}
