package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationLineController struct{}

// LocationLineStoreForm 新建线别表单
type LocationLineStoreForm struct {
	Sort                            int64    `form:"sort" json:"sort"`
	UniqueCode                      string   `form:"unique_code" json:"unique_code"`
	Name                            string   `form:"name" json:"name"`
	BeEnable                        bool     `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUuids        []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways            []*models.OrganizationRailwayModel
	OrganizationParagraphUuids      []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs          []*models.OrganizationParagraphModel
	OrganizationWorkshopUuids       []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops           []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUuids       []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas           []*models.OrganizationWorkAreaModel
	LocationSectionUuids            []string `form:"location_section_uuids" json:"location_section_uuids"`
	LocationSections                []*models.LocationSectionModel
	LocationStationUuids            []string `form:"location_station_uuids" json:"location_station_uuids"`
	LocationStations                []*models.LocationStationModel
	LocationRailroadGradeCrossUuids []string `form:"location_railroad_grade_cross_uuids" json:"location_railroad_grade_cross_uuids"`
	LocationRailroadGradeCrosses    []*models.LocationRailroadGradeCrossModel
	LocationCenterUuids             []string `form:"location_center_uuids" json:"location_center_uuids"`
	LocationCenters                 []*models.LocationCenterModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationLineStoreForm
func (cls LocationLineStoreForm) ShouldBind(ctx *gin.Context) LocationLineStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("线别代码必填")
	}
	if len(cls.UniqueCode) != 5 {
		wrongs.PanicValidate("线别代码必须是5位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("线别名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("线别名称不能超过64位")
	}
	// 查询路局
	if len(cls.OrganizationRailwayUuids) > 0 {
		models.BootByModel(models.OrganizationRailwayModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.OrganizationRailwayUuids).
			Find(&cls.OrganizationRailways)
	}
	// 查询站段
	if len(cls.OrganizationParagraphUuids) > 0 {
		models.BootByModel(models.OrganizationParagraphModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.OrganizationParagraphUuids).
			Find(&cls.OrganizationParagraphs)
	}
	// 查询车间
	if len(cls.OrganizationWorkshopUuids) > 0 {
		models.BootByModel(models.OrganizationWorkshopModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.OrganizationWorkshopUuids).
			Find(&cls.OrganizationWorkshops)
	}
	// 查询工区
	if len(cls.OrganizationWorkAreaUuids) > 0 {
		models.BootByModel(models.OrganizationWorkAreaModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.OrganizationWorkAreaUuids).
			Find(&cls.OrganizationWorkAreas)
	}
	// 查询区间
	if len(cls.LocationSectionUuids) > 0 {
		models.BootByModel(models.LocationSectionModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationSectionUuids).
			Find(&cls.LocationSections)
	}
	// 查询站场
	if len(cls.LocationStationUuids) > 0 {
		models.BootByModel(models.LocationStationModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationStationUuids).
			Find(&cls.LocationStations)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUuids) > 0 {
		models.BootByModel(models.LocationRailroadGradeCrossModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationRailroadGradeCrossUuids).
			Find(&cls.LocationRailroadGradeCrosses)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUuids) > 0 {
		models.BootByModel(models.LocationRailroadGradeCrossModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationRailroadGradeCrossUuids).
			Find(&cls.LocationRailroadGradeCrosses)
	}

	return cls
}

// LocationLineBindForm 线别多对多绑定
type LocationLineBindForm struct {
	OrganizationRailwayUuids        []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways            []*models.OrganizationRailwayModel
	OrganizationParagraphUuids      []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs          []*models.OrganizationParagraphModel
	OrganizationWorkshopUuids       []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops           []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUuids       []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas           []*models.OrganizationWorkAreaModel
	LocationSectionUuids            []string `form:"location_section_uuids" json:"location_section_uuids"`
	LocationSections                []*models.LocationSectionModel
	LocationStationUuids            []string `form:"location_station_uuids" json:"location_station_uuids"`
	LocationStations                []*models.LocationStationModel
	LocationRailroadGradeCrossUuids []string `form:"location_railroad_grade_cross_uuids" json:"location_railroad_grade_cross_uuids"`
	LocationRailroadGradeCrosses    []*models.LocationRailroadGradeCrossModel
	LocationCenterUuids             []string `form:"location_center_uuids" json:"location_center_uuids"`
	LocationCenters                 []*models.LocationCenterModel
}

// C 新建
func (LocationLineController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationLineModel
	)

	// 表单
	form := (&LocationLineStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别代码")
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别名称")

	// 新建
	organizationLine := &models.LocationLineModel{
		BaseModel:                    models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:                   form.UniqueCode,
		Name:                         form.Name,
		BeEnable:                     form.BeEnable,
		LocationSections:             form.LocationSections,
		LocationStations:             form.LocationStations,
		LocationRailroadGradeCrosses: form.LocationRailroadGradeCrosses,
		LocationCenters:              form.LocationCenters,
	}
	if ret = models.BootByModel(models.LocationLineModel{}).PrepareByDefault().Create(organizationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_line": organizationLine}))
}

// D 删除
func (LocationLineController) D(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		organizationLine models.LocationLineModel
	)
	// 查询
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&organizationLine)
	wrongs.PanicWhenIsEmpty(ret, "线别")

	// 删除
	if ret = models.BootByModel(models.LocationLineModel{}).PrepareByDefault().Delete(&organizationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (LocationLineController) U(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		locationLine, repeat models.LocationLineModel
	)

	// 表单
	form := (&LocationLineStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别代码")
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别名称")

	// 查询
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationLine)
	wrongs.PanicWhenIsEmpty(ret, "线别")

	// 修改
	locationLine.Name = form.Name
	locationLine.Sort = form.Sort
	locationLine.BeEnable = form.BeEnable
	if ret = models.BootByModel(&models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		Save(&locationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_line": locationLine}))
}

// S 详情
func (LocationLineController) S(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		organizationLine models.LocationLineModel
	)
	// 查询
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		PrepareUseQuery(ctx, "").
		First(&organizationLine)
	wrongs.PanicWhenIsEmpty(ret, "线别")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_line": organizationLine}))
}

// I 列表
func (LocationLineController) I(ctx *gin.Context) {
	var organizationLines []models.LocationLineModel
	models.BootByModel(models.LocationLineModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "sort").
		PrepareUseQuery(ctx, "").
		Find(&organizationLines)

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_lines": organizationLines}))
}
