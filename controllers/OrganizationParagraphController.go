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
//  @receiver ins
//  @param ctx
//  @return OrganizationParagraphStoreForm
func (ins OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("站段代码必填")
	}
	if len(ins.UniqueCode) != 4 {
		wrongs.PanicValidate("站段代码必须是4位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("站段名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("站段名称不能超过64位")
	}
	if ins.OrganizationRailwayUuid == "" {
		wrongs.PanicValidate("所属路局必选")
	}
	ret = models.BootByModel(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": ins.OrganizationRailwayUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.OrganizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")
	if len(ins.OrganizationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.OrganizationLineUuids).
			Find(&ins.OrganizationLines)
	}

	return ins
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段代码")
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
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
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).PrepareByDefaultDbDriver().Create(&organizationParagraph); ret.Error != nil {
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
		PrepareByDefaultDbDriver().
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	// 删除
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).PrepareByDefaultDbDriver().Delete(&organizationParagraph); ret.Error != nil {
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段代码")
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站段名称")

	// 查询
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	// 编辑
	organizationParagraph.BaseModel.Sort = form.Sort
	organizationParagraph.Name = form.Name
	organizationParagraph.BeEnable = form.BeEnable
	organizationParagraph.OrganizationRailway = form.OrganizationRailway
	if ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&organizationParagraph); ret.Error != nil {
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
		PrepareUseQueryByDefaultDbDriver(ctx).
		First(&organizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_paragraph": organizationParagraph}))
}

// I 列表
func (OrganizationParagraphController) I(ctx *gin.Context) {
	var (
		organizationParagraphs []models.OrganizationParagraphModel
		count                  int64
		db                     *gorm.DB
	)
	db = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWhereFields("uuid", "sort", "unique_code", "name", "shot_name", "be_enable", "organization_railway_uuid").
		SetPreloadsByDefault().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&organizationParagraphs)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"organization_paragraphs": organizationParagraphs}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&organizationParagraphs)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"organization_paragraphs": organizationParagraphs}, ctx.Query("__page__"), count))
	}
}
