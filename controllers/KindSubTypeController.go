package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type KindSubTypeController struct{}

// KindSubTypeStoreForm 新建型号表单
type KindSubTypeStoreForm struct {
	Sort               int64  `form:"sort" json:"sort"`
	UniqueCode         string `form:"unique_code" json:"unique_code"`
	Name               string `form:"name" json:"name"`
	Nickname           string `form:"nickname" json:"nickname"`
	BeEnable           bool   `form:"be_enable" json:"be_enable"`
	KindEntireTypeUuid string `form:"kind_entire_type_uuid" json:"kind_entire_type_uuid"`
	KindEntireType     models.KindEntireTypeModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return KindSubTypeStoreForm
func (cls KindSubTypeStoreForm) ShouldBind(ctx *gin.Context) KindSubTypeStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("型号代码必填")
	}
	if len(cls.UniqueCode) != 2 {
		wrongs.PanicValidate("型号代码必须是2位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("型号名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("型号名称不能超过64位")
	}
	if cls.KindEntireTypeUuid == "" {
		wrongs.PanicValidate("所属类型必选")
	}
	ret = models.BootByModel(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.KindEntireTypeUuid}).
		PrepareByDefaultDbDriver().
		First(&cls.KindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "所属类型")

	return cls
}

// N 新建
func (KindSubTypeController) N(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.KindSubTypeModel
	)

	// 表单
	form := (&KindSubTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "型号代码")
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "型号名称")

	// 新建
	kindSubType := &models.KindSubTypeModel{
		BaseModel:      models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:     form.KindEntireType.UniqueCode + form.UniqueCode,
		Name:           form.Name,
		BeEnable:       form.BeEnable,
		Nickname:       form.Nickname,
		KindEntireType: form.KindEntireType,
	}
	if ret = models.BootByModel(models.KindSubTypeModel{}).PrepareByDefaultDbDriver().Create(&kindSubType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"kind_sub_type": kindSubType}))
}

// R 删除
func (KindSubTypeController) R(ctx *gin.Context) {
	var (
		ret         *gorm.DB
		kindSubType models.KindSubTypeModel
	)

	// 查询
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindSubType)
	wrongs.PanicWhenIsEmpty(ret, "型号")

	// 删除
	if ret := models.BootByModel(models.KindSubTypeModel{}).PrepareByDefaultDbDriver().Delete(&kindSubType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 编辑
func (KindSubTypeController) E(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		kindSubType, repeat models.KindSubTypeModel
	)

	// 表单
	form := (&KindSubTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "型号代码")
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "型号名称")

	// 查询
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindSubType)
	wrongs.PanicWhenIsEmpty(ret, "型号")

	// 编辑
	kindSubType.BaseModel.Sort = form.Sort
	kindSubType.Name = form.Name
	kindSubType.BeEnable = form.BeEnable
	kindSubType.Nickname = form.Nickname
	kindSubType.KindEntireType = form.KindEntireType
	if ret = models.BootByModel(models.KindSubTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&kindSubType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"kind_sub_type": kindSubType}))
}

// D 详情
func (KindSubTypeController) D(ctx *gin.Context) {
	var (
		ret         *gorm.DB
		kindSubType models.KindSubTypeModel
	)
	ret = models.BootByModel(models.KindSubTypeModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&kindSubType)
	wrongs.PanicWhenIsEmpty(ret, "型号")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_sub_type": kindSubType}))
}

// L 列表
func (KindSubTypeController) L(ctx *gin.Context) {
	var (
		kindSubTypes []models.KindSubTypeModel
		count        int64
		db           *gorm.DB
	)
	db = models.BootByModel(models.KindSubTypeModel{}).
		SetWhereFields("kind_entire_type_uuid").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&kindSubTypes)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"kind_sub_types": kindSubTypes}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&kindSubTypes)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"kind_sub_types": kindSubTypes}, ctx.Query("__page__"), count))
	}
}
