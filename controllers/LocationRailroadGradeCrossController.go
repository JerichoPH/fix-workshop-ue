package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationRailroadGradeCrossController struct{}

// LocationRailroadGradeCrossStoreForm 新建道口表单
type LocationRailroadGradeCrossStoreForm struct {
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
//  @receiver cls
//  @param ctx
//  @return LocationCenterStoreForm
func (cls LocationRailroadGradeCrossStoreForm) ShouldBind(ctx *gin.Context) LocationRailroadGradeCrossStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("道口代码必填")
	}
	if len(cls.UniqueCode) != 5 {
		wrongs.PanicValidate("道口代码必须是5位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("道口名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("道口名称不能超过64位")
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
	}
	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// LocationRailroadGradeCrossBindLocationLinesForm 道口绑定线别表单
type LocationRailroadGradeCrossBindLocationLinesForm struct {
	LocationLineUuids []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationRailroadGradeCrossBindLocationLinesForm
func (cls LocationRailroadGradeCrossBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationRailroadGradeCrossBindLocationLinesForm {
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
func (LocationRailroadGradeCrossController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationRailroadGradeCrossModel
	)

	// 表单
	form := (&LocationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口代码")
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口名称")

	// 新建
	locationRailroadGradeCross := &models.LocationRailroadGradeCrossModel{
		BaseModel:                models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		BeEnable:                 form.BeEnable,
		OrganizationWorkshopUuid: form.OrganizationWorkshop.Uuid,
		OrganizationWorkAreaUuid: form.OrganizationWorkAreaUuid,
		LocationLines:            form.LocationLines,
	}
	if ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).PrepareByDefault().Create(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
}

// D 删除
func (LocationRailroadGradeCrossController) D(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		locationRailroadGradeCross models.LocationRailroadGradeCrossModel
	)

	// 查询
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	// 删除
	if ret := models.BootByModel(models.LocationRailroadGradeCrossModel{}).PrepareByDefault().Delete(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (LocationRailroadGradeCrossController) U(ctx *gin.Context) {
	var (
		ret                                *gorm.DB
		locationRailroadGradeCross, repeat models.LocationRailroadGradeCrossModel
	)

	// 表单
	form := (&LocationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口代码")
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "道口名称")

	// 查询
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	// 编辑
	locationRailroadGradeCross.BaseModel.Sort = form.Sort
	locationRailroadGradeCross.Name = form.Name
	locationRailroadGradeCross.BeEnable = form.BeEnable
	locationRailroadGradeCross.OrganizationWorkshopUuid = form.OrganizationWorkshop.Uuid
	locationRailroadGradeCross.OrganizationWorkAreaUuid = form.OrganizationWorkAreaUuid
	locationRailroadGradeCross.LocationLines = form.LocationLines
	if ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefault().Save(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
}

// PutBindLines 道口绑定线别
func (LocationRailroadGradeCrossController) PutBindLines(ctx *gin.Context) {
	var (
		ret                                              *gorm.DB
		locationRailroadGradeCross                       models.LocationRailroadGradeCrossModel
		pivotLocationLineAndLocationRailroadGradeCrosses []models.PivotLocationLineAndLocationRailroadGradeCross
	)

	// 表单
	form := (&LocationRailroadGradeCrossBindLocationLinesForm{}).ShouldBind(ctx)

	if ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationRailroadGradeCross); ret.Error != nil {
		wrongs.PanicWhenIsEmpty(ret, "道口")
	}

	// 删除原有绑定关系
	ret = models.BootByModel(models.BaseModel{}).PrepareByDefault().Exec("delete from pivot_location_line_and_location_railroad_grade_crosses where location_railroad_grade_cross_id = ?", locationRailroadGradeCross.Id)

	// 创建绑定关系
	if len(form.LocationLines) > 0 {
		for _, locationLine := range form.LocationLines {
			pivotLocationLineAndLocationRailroadGradeCrosses = append(pivotLocationLineAndLocationRailroadGradeCrosses, models.PivotLocationLineAndLocationRailroadGradeCross{
				LocationLineId:               locationLine.Id,
				LocationRailroadGradeCrossId: locationRailroadGradeCross.Id,
			})
		}
		models.BootByModel(models.PivotLocationLineAndLocationRailroadGradeCross{}).
			PrepareByDefault().
			CreateInBatches(&pivotLocationLineAndLocationRailroadGradeCrosses, 100)
	}

	ctx.JSON(tools.CorrectBoot("绑定成功").Updated(tools.Map{}))
}

func (LocationRailroadGradeCrossController) S(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		locationRailroadGradeCross models.LocationRailroadGradeCrossModel
	)
	ret = models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloads("OrganizationWorkshop", "OrganizationWorkArea", "LocationLines").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		First(&locationRailroadGradeCross)
	wrongs.PanicWhenIsEmpty(ret, "道口")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
}

func (LocationRailroadGradeCrossController) I(ctx *gin.Context) {
	var locationRailroadGradeCrosses []models.LocationRailroadGradeCrossModel
	models.BootByModel(models.LocationRailroadGradeCrossModel{}).
		SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
		SetPreloads("OrganizationWorkshop", "OrganizationWorkArea", "LocationLines").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&locationRailroadGradeCrosses)

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_railroad_grade_crosses": locationRailroadGradeCrosses}))
}
