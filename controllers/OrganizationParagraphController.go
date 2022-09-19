package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrganizationParagraphController struct{}

// OrganizationParagraphStoreForm 新建站段表单
type OrganizationParagraphStoreForm struct {
	Sort                    int64  `form:"sort" json:"sort"`
	UniqueCode              string `form:"unique_code" json:"unique_code"`
	Name                    string `form:"name" json:"name"`
	BeEnable                bool   `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUuid string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway     models.OrganizationRailwayModel
	OrganizationLineUuids   []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines       []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationParagraphStoreForm
func (cls OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("站段代码必填")
	}
	if len(cls.UniqueCode) != 4{
		wrongs.PanicValidate("站段代码必须是4位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("站段名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("站段名称不能超过64位")
	}
	if cls.OrganizationRailwayUuid == "" {
		wrongs.PanicValidate("所属路局必选")
	}
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationRailwayUuid}).
		PrepareByDefault().
		First(&cls.OrganizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")
	if len(cls.OrganizationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.OrganizationLineUuids).
			Find(&cls.OrganizationLines)
	}

	return cls
}

// C 新建
func (OrganizationParagraphController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.OrganizationParagraphModel
	)

	// 表单
	form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段代码")
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段名称")

	// 新建
	organizationParagraph := &models.OrganizationParagraphModel{
		BaseModel:           models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:          form.UniqueCode,
		Name:                form.Name,
		BeEnable:            form.BeEnable,
		OrganizationRailway: form.OrganizationRailway,
	}
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).PrepareByDefault().Create(&organizationParagraph); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_paragraph": organizationParagraph}))
}

// D 删除
func (OrganizationParagraphController) D(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationParagraph models.OrganizationParagraphModel
	)

	// 查询
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	// 删除
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).PrepareByDefault().Delete(&organizationParagraph); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (OrganizationParagraphController) U(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationParagraph models.OrganizationParagraphModel
	)

	// 表单
	form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.OrganizationParagraphModel
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段代码")
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段名称")

	// 查询
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	// 编辑
	organizationParagraph.BaseModel.Sort = form.Sort
	organizationParagraph.Name = form.Name
	organizationParagraph.BeEnable = form.BeEnable
	organizationParagraph.OrganizationRailway = form.OrganizationRailway
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefault().Save(&organizationParagraph); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_paragraph": organizationParagraph}))
}

// S 详情
func (OrganizationParagraphController) S(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationParagraph models.OrganizationParagraphModel
	)

	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_paragraph": organizationParagraph}))
}

// I 列表
func (OrganizationParagraphController) I(ctx *gin.Context) {
	var organizationParagraphs []models.OrganizationParagraphModel
	models.BootByModel(models.OrganizationParagraphModel{}).
		SetWhereFields("uuid", "sort", "unique_code", "name", "shot_name", "be_enable", "organization_railway_uuid").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&organizationParagraphs)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_paragraphs": organizationParagraphs}))
}
