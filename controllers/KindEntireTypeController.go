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
//  @receiver cls
//  @param ctx
//  @return KindEntireTypeStoreForm
func (cls KindEntireTypeStoreForm) ShouldBind(ctx *gin.Context) KindEntireTypeStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("类型代码必填")
	}
	if len(cls.UniqueCode) != 2 {
		wrongs.PanicValidate("类型代码必须是2位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("类型名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("类型名称不能超过64位")
	}
	if cls.KindCategoryUuid == "" {
		wrongs.PanicValidate("所属种类必选")
	}
	ret = models.BootByModel(models.KindCategoryModel{}).
		SetWheres(tools.Map{"uuid": cls.KindCategoryUuid}).
		PrepareByDefault().
		First(&cls.KindCategory)
	wrongs.PanicWhenIsEmpty(ret, "所属种类")

	return cls
}

// C 新建
func (KindEntireTypeController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.KindEntireTypeModel
	)

	// 表单
	form := (&KindEntireTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型代码")
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
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
	if ret = models.BootByModel(models.KindEntireTypeModel{}).PrepareByDefault().Create(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"kind_entire_type": kindEntireType}))
}

// D 删除
func (KindEntireTypeController) D(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		kindEntireType models.KindEntireTypeModel
	)

	// 查询
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	// 删除
	if ret := models.BootByModel(models.KindEntireTypeModel{}).PrepareByDefault().Delete(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 更新
func (KindEntireTypeController) U(ctx *gin.Context) {
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
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型代码")
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "类型名称")

	// 查询
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	// 编辑
	kindEntireType.BaseModel.Sort = form.Sort
	kindEntireType.Name = form.Name
	kindEntireType.BeEnable = form.BeEnable
	kindEntireType.KindCategory = form.KindCategory
	if ret = models.BootByModel(models.KindEntireTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefault().Save(&kindEntireType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"kind_entire_type": kindEntireType}))
}

// S 详情
func (KindEntireTypeController) S(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		kindEntireType models.KindEntireTypeModel
	)
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&kindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "类型")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_entire_type": kindEntireType}))
}

// I 列表
func (KindEntireTypeController) I(ctx *gin.Context) {
	var kindEntireTypes []models.KindEntireTypeModel
	models.BootByModel(models.KindEntireTypeModel{}).
		SetWhereFields().
		PrepareUseQuery(ctx, "").
		Find(&kindEntireTypes)

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_entire_types": kindEntireTypes}))
}
