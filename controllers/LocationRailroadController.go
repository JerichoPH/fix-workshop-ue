package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationRailroadController struct{}

// LocationRailroadStoreForm 新建道口表单
type LocationRailroadStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUuid string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUuid string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
	LocationLineUuids        []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines            []*models.LocationLineModel
}

// ShouldBind 表单绑定
//  @receiver ins
//  @param ctx
//  @return LocationCenterStoreForm
func (ins LocationRailroadStoreForm) ShouldBind(ctx *gin.Context) LocationRailroadStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("道口代码必填")
	}
	if len(ins.UniqueCode) != 5 {
		wrongs.PanicValidate("道口代码必须是5位")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("道口名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("道口名称不能超过64位")
	}
	if ins.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": ins.OrganizationWorkshopUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if ins.OrganizationWorkAreaUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": ins.OrganizationWorkAreaUuid}).
			PrepareByDefaultDbDriver().
			First(&ins.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}
	if len(ins.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationLineUuids).
			Find(&ins.LocationLines)
	}

	return ins
}

// LocationRailroadBindLocationLinesForm 道口绑定线别表单
type LocationRailroadBindLocationLinesForm struct {
	LocationLineUuids []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return LocationRailroadBindLocationLinesForm
func (ins LocationRailroadBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationRailroadBindLocationLinesForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if len(ins.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.LocationLineUuids).
			Find(&ins.LocationLines)
	}

	return ins
}

// N 新建
func (LocationRailroadController) N(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationRailroadModel
	)

	// 表单
	form := (&LocationRailroadStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口代码")
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口名称")

	// 新建
	locationRailroadGradeCross := &models.LocationRailroadModel{
		BaseModel:                models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		BeEnable:                 form.BeEnable,
		OrganizationWorkshopUuid: form.OrganizationWorkshop.Uuid,
		OrganizationWorkAreaUuid: form.OrganizationWorkAreaUuid,
		LocationLines:            form.LocationLines,
	}
	if ret = models.BootByModel(models.LocationRailroadModel{}).PrepareByDefaultDbDriver().Create(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_railroad": locationRailroadGradeCross}))
}

// R 删除
func (LocationRailroadController) R(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		locationRailroadGradeCross models.LocationRailroadModel
	)

	// 查询
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	// 删除
	if ret := models.BootByModel(models.LocationRailroadModel{}).PrepareByDefaultDbDriver().Delete(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 编辑
func (LocationRailroadController) E(ctx *gin.Context) {
	var (
		ret                                *gorm.DB
		locationRailroadGradeCross, repeat models.LocationRailroadModel
	)

	// 表单
	form := (&LocationRailroadStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口代码")
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口名称")

	// 查询
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	// 编辑
	locationRailroadGradeCross.BaseModel.Sort = form.Sort
	locationRailroadGradeCross.Name = form.Name
	locationRailroadGradeCross.BeEnable = form.BeEnable
	locationRailroadGradeCross.OrganizationWorkshopUuid = form.OrganizationWorkshop.Uuid
	locationRailroadGradeCross.OrganizationWorkAreaUuid = form.OrganizationWorkAreaUuid
	locationRailroadGradeCross.LocationLines = form.LocationLines
	if ret = models.BootByModel(models.LocationRailroadModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_railroad": locationRailroadGradeCross}))
}

// PutBindLines 道口绑定线别
func (LocationRailroadController) PutBindLines(ctx *gin.Context) {
	var (
		ret                                              *gorm.DB
		locationRailroadGradeCross                       models.LocationRailroadModel
		pivotLocationLineAndLocationRailroadGradeCrosses []models.PivotLocationLineAndLocationRailroadGradeCrossModel
	)

	// 表单
	form := (&LocationRailroadBindLocationLinesForm{}).ShouldBind(ctx)

	if ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicWhenIsEmpty(ret, "道口")
	}

	// 删除原有绑定关系
	ret = models.BootByModel(models.BaseModel{}).PrepareByDefaultDbDriver().Exec("delete from pivot_location_line_and_location_railroad_grade_crosses where location_railroad_grade_cross_id = ?", locationRailroadGradeCross.Id)

	// 创建绑定关系
	if len(form.LocationLines) > 0 {
		for _, locationLine := range form.LocationLines {
			pivotLocationLineAndLocationRailroadGradeCrosses = append(pivotLocationLineAndLocationRailroadGradeCrosses, models.PivotLocationLineAndLocationRailroadGradeCrossModel{
				LocationLineId:     locationLine.Id,
				LocationRailroadId: locationRailroadGradeCross.Id,
			})
		}
		models.BootByModel(models.PivotLocationLineAndLocationRailroadGradeCrossModel{}).
			PrepareByDefaultDbDriver().
			CreateInBatches(&pivotLocationLineAndLocationRailroadGradeCrosses, 100)
	}

	ctx.JSON(tools.CorrectBoot("绑定成功").Updated(tools.Map{}))
}

// D 详情
func (LocationRailroadController) D(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		locationRailroadGradeCross models.LocationRailroadModel
	)
	ret = models.BootByModel(models.LocationRailroadModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloads("OrganizationWorkshop", "OrganizationWorkArea", "LocationLine").
		SetPreloadsByDefault().
		PrepareUseQueryByDefaultDbDriver(ctx).
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_railroad": locationRailroadGradeCross}))
}

// L 列表
func (LocationRailroadController) L(ctx *gin.Context) {
	var (
		locationRailroadGradeCrosses []models.LocationRailroadModel
		count                        int64
		db                           *gorm.DB
	)
	db = models.BootByModel(models.LocationRailroadModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
		SetPreloads("OrganizationWorkshop", "OrganizationWorkArea", "LocationLine").
		SetPreloadsByDefault().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&locationRailroadGradeCrosses)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_railroads": locationRailroadGradeCrosses}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&locationRailroadGradeCrosses)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"location_railroads": locationRailroadGradeCrosses}, ctx.Query("__page__"), count))
	}
}
