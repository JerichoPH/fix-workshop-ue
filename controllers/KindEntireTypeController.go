package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type KindEntireTypeController struct{}

// KindEntireTypeStoreForm 新建类型表单
type KindEntireTypeStoreForm struct {
	Sort             int64  `form:"sort" json:"sort"`
	UniqueCode       string `form:"unique_code" json:"unique_code"`
	Name             string `form:"name" json:"name"`
	Nickname         string `form:"nickname" json:"nickname"`
	BeEnable         bool   `form:"be_enable" json:"be_enable"`
	KindCategoryUuid string `form:"kind_category_uuid" json:"kind_category_uuid"`
	KindCategory     models.KindCategoryModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return KindEntireTypeStoreForm
func (ins KindEntireTypeStoreForm) ShouldBind(ctx *gin.Context) KindEntireTypeStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("类型代码必填")
	}
	if len(ins.UniqueCode) != 2 {
		wrongs.PanicValidate("类型代码必须是2位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("类型名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("类型名称不能超过64位")
	}
	if ins.KindCategoryUuid == "" {
		wrongs.PanicValidate("所属种类必选")
	}
	ret = models.BootByModel(models.KindCategoryModel{}).
		SetWheres(tools.Map{"uuid": ins.KindCategoryUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.KindCategory)
	wrongs.PanicWhenIsEmpty(ret, "所属种类")

	return ins
}

// N 新建
func (KindEntireTypeController) N(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.KindEntireTypeModel
	)

	// 表单
	form := (&KindEntireTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型代码")
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型名称")

	// 新建
	kindEntireType := &models.KindEntireTypeModel{
		BaseModel:    models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:   form.KindCategory.UniqueCode + form.UniqueCode,
		Name:         form.Name,
		BeEnable:     form.BeEnable,
		KindCategory: form.KindCategory,
	}
	if ret = models.BootByModel(models.KindEntireTypeModel{}).PrepareByDefaultDbDriver().Create(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"kind_entire_type": kindEntireType}))
}

// R 删除
func (KindEntireTypeController) R(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		kindEntireType models.KindEntireTypeModel
	)

	// 查询
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	// 删除
	if ret := models.BootByModel(models.KindEntireTypeModel{}).PrepareByDefaultDbDriver().Delete(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 更新
func (KindEntireTypeController) E(ctx *gin.Context) {
	var (
		ret                    *gorm.DB
		kindEntireType, repeat models.KindEntireTypeModel
	)

	// 表单
	form := (&KindEntireTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型代码")
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型名称")

	// 查询
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	// 编辑
	kindEntireType.BaseModel.Sort = form.Sort
	kindEntireType.Name = form.Name
	kindEntireType.BeEnable = form.BeEnable
	kindEntireType.KindCategory = form.KindCategory
	if ret = models.BootByModel(models.KindEntireTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"kind_entire_type": kindEntireType}))
}

// D 详情
func (KindEntireTypeController) D(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		kindEntireType models.KindEntireTypeModel
	)
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_entire_type": kindEntireType}))
}

// L 列表
func (KindEntireTypeController) L(ctx *gin.Context) {
	var (
		kindEntireTypes []models.KindEntireTypeModel
		count           int64
		db              *gorm.DB
	)
	db = models.BootByModel(models.KindEntireTypeModel{}).
		SetWhereFields("kind_category_uuid").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&kindEntireTypes)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_entire_types": kindEntireTypes}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&kindEntireTypes)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"kind_entire_types": kindEntireTypes}, ctx.Query("__page__"), count))
	}
}
