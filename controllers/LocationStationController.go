package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationStationController struct{}

// LocationStationStoreForm 新建站场表单
type LocationStationStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUuid string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUuid string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
	LocationLineUUIDs        []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines            []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationStationStoreForm
func (cls LocationStationStoreForm) ShouldBind(ctx *gin.Context) LocationStationStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("站场代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("站场名称必填")
	}
	if cls.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUuid}).
		PrepareByDefault().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUuid}).
			PrepareByDefault().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	} else {

	}
	if len(cls.LocationLineUUIDs) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
	}

	return cls
}

// LocationStationBindLocationLinesForm 站场绑定线别表单
type LocationStationBindLocationLinesForm struct {
	LocationLineUuids []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationStationBindLocationLinesForm
func (cls LocationStationBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationStationBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// C 新建
func (LocationStationController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationStationModel
	)

	// 表单
	form := (&LocationStationStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站场代码")
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站场名称")

	// 新建
	locationStation := &models.LocationStationModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		BeEnable:             form.BeEnable,
		OrganizationWorkshop: form.OrganizationWorkshop,
		OrganizationWorkArea: form.OrganizationWorkArea,
		LocationLines:        form.LocationLines,
	}
	if ret = models.BootByModel(models.LocationStationModel{}).PrepareByDefault().Create(&locationStation); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_station": locationStation}))
}

// D 删除
func (LocationStationController) D(ctx *gin.Context) {
	var (
		ret             *gorm.DB
		locationStation models.LocationStationModel
	)
	// 查询
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationStation)
	wrongs.PanicWhenIsEmpty(ret, "站场")

	// 删除
	if ret := models.BootByModel(models.LocationStationModel{}).PrepareByDefault().Delete(&locationStation); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (LocationStationController) U(ctx *gin.Context) {
	var (
		ret                     *gorm.DB
		locationStation, repeat models.LocationStationModel
	)

	// 表单
	form := (&LocationStationStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站场代码")
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "站场名称")

	// 查询
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationStation)
	wrongs.PanicWhenIsEmpty(ret, "站场")

	// 编辑
	if ret = models.
		BootByModel(models.LocationStationModel{}).
		PrepareByDefault().
		Where("uuid = ?", ctx.Param("uuid")).
		Updates(map[string]interface{}{
			"sort":                        form.Sort,
			"name":                        form.Name,
			"be_enable":                   form.BeEnable,
			"organization_workshop_uuid":  form.OrganizationWorkshopUuid,
			"organization_work_area_uuid": form.OrganizationWorkAreaUuid,
		}); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_station": locationStation}))
}

// PutBindLocationLines 编辑站场到线别
func (LocationStationController) PutBindLocationLines(ctx *gin.Context) {
	var (
		ret                                  *gorm.DB
		locationStation                      models.LocationStationModel
		pivotLocationLineAndLocationStations []models.PivotLocationLineAndLocationStation
	)

	// 表单
	form := (&LocationStationBindLocationLinesForm{}).ShouldBind(ctx)

	if ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationStation); ret.Error != nil {
		wrongs.PanicWhenIsEmpty(ret, "站场")
	}

	// 删除原有绑定关系
	ret = models.BootByModel(models.BaseModel{}).PrepareByDefault().Exec("delete from pivot_location_line_and_location_stations where location_station_id = ?", locationStation.Id)

	// 创建绑定关系
	if len(form.LocationLines) > 0 {
		for _, locationLine := range form.LocationLines {
			pivotLocationLineAndLocationStations = append(pivotLocationLineAndLocationStations, models.PivotLocationLineAndLocationStation{
				LocationLineId:    locationLine.Id,
				LocationStationId: locationStation.Id,
			})
		}
		models.BootByModel(models.PivotLocationLineAndLocationStation{}).
			PrepareByDefault().
			CreateInBatches(&pivotLocationLineAndLocationStations, 100)
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{}))
}

// S 详情
func (LocationStationController) S(ctx *gin.Context) {
	var (
		ret             *gorm.DB
		locationStation models.LocationStationModel
	)
	// 查询
	ret = models.BootByModel(models.LocationStationModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		First(&locationStation)
	wrongs.PanicWhenIsEmpty(ret, "站场")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"location_station": locationStation}))
}

// I 列表
func (LocationStationController) I(ctx *gin.Context) {
	var locationStations []models.LocationStationModel
	models.BootByModel(models.LocationStationModel{}).
		SetWhereFields().
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&locationStations)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"location_stations": locationStations}))
}
