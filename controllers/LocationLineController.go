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
//  @receiver ins
//  @param ctx
//  @return LocationLineStoreForm
func (ins LocationLineStoreForm) ShouldBind(ctx *gin.Context) LocationLineStoreForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("线别代码必填")
	}
	if len(ins.UniqueCode) != 5 {
		wrongs.PanicValidate("线别代码必须是5位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("线别名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("线别名称不能超过64位")
	}
	// 查询路局
	if len(ins.OrganizationRailwayUuids) > 0 {
		models.BootByModel(models.OrganizationRailwayModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.OrganizationRailwayUuids).
			Find(&ins.OrganizationRailways)
	}
	// 查询站段
	if len(ins.OrganizationParagraphUuids) > 0 {
		models.BootByModel(models.OrganizationParagraphModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.OrganizationParagraphUuids).
			Find(&ins.OrganizationParagraphs)
	}
	// 查询车间
	if len(ins.OrganizationWorkshopUuids) > 0 {
		models.BootByModel(models.OrganizationWorkshopModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.OrganizationWorkshopUuids).
			Find(&ins.OrganizationWorkshops)
	}
	// 查询工区
	if len(ins.OrganizationWorkAreaUuids) > 0 {
		models.BootByModel(models.OrganizationWorkAreaModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.OrganizationWorkAreaUuids).
			Find(&ins.OrganizationWorkAreas)
	}
	// 查询区间
	if len(ins.LocationSectionUuids) > 0 {
		models.BootByModel(models.LocationSectionModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationSectionUuids).
			Find(&ins.LocationSections)
	}
	// 查询站场
	if len(ins.LocationStationUuids) > 0 {
		models.BootByModel(models.LocationStationModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationStationUuids).
			Find(&ins.LocationStations)
	}
	// 查询道口
	if len(ins.LocationRailroadGradeCrossUuids) > 0 {
		models.BootByModel(models.LocationRailroadGradeCrossModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationRailroadGradeCrossUuids).
			Find(&ins.LocationRailroadGradeCrosses)
	}
	// 查询道口
	if len(ins.LocationRailroadGradeCrossUuids) > 0 {
		models.BootByModel(models.LocationRailroadGradeCrossModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationRailroadGradeCrossUuids).
			Find(&ins.LocationRailroadGradeCrosses)
	}

	return ins
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

// N 新建
func (LocationLineController) N(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationLineModel
	)

	// 表单
	form := (&LocationLineStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别代码")
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
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
	if ret = models.BootByModel(models.LocationLineModel{}).PrepareByDefaultDbDriver().Create(organizationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_line": organizationLine}))
}

// R 删除
func (LocationLineController) R(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		organizationLine models.LocationLineModel
	)
	// 查询
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&organizationLine)
	wrongs.PanicWhenIsEmpty(ret, "线别")

	// 删除
	if ret = models.BootByModel(models.LocationLineModel{}).PrepareByDefaultDbDriver().Delete(&organizationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 编辑
func (LocationLineController) E(ctx *gin.Context) {
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别代码")
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "线别名称")

	// 查询
	ret = models.BootByModel(models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationLine)
	wrongs.PanicWhenIsEmpty(ret, "线别")

	// 修改
	locationLine.Name = form.Name
	locationLine.Sort = form.Sort
	locationLine.BeEnable = form.BeEnable
	if ret = models.BootByModel(&models.LocationLineModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		Save(&locationLine); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_line": locationLine}))
}

// D 详情
func (LocationLineController) D(ctx *gin.Context) {
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

// L 列表
func (LocationLineController) L(ctx *gin.Context) {
	var (
		locationLines []models.LocationLineModel
		count         int64
		db            *gorm.DB
	)
	db = models.BootByModel(models.LocationLineModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "sort").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&locationLines)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_lines": locationLines}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&locationLines)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"location_lines": locationLines}, ctx.Query("__page__"), count))
	}
}
