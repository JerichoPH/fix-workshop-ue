package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationSectionController struct{}

// LocationSectionStoreForm 新建区间表单
type LocationSectionStoreForm struct {
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

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationSectionStoreForm
func (cls LocationSectionStoreForm) ShouldBind(ctx *gin.Context) LocationSectionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("区间代码不能为空")
	}
	if len(cls.UniqueCode) != 6 {
		wrongs.PanicValidate("区间代码必须是6位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("区间名称不能为空")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("区间名称不能64位")
	}
	if cls.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间不能为空")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUuid}).
		PrepareByDefaultDbDriver().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUuid}).
			PrepareByDefaultDbDriver().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}
	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// LocationSectionBindLocationLinesForm 区间绑定线别表单
type LocationSectionBindLocationLinesForm struct {
	LocationLineUuids []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationSectionBindLocationLinesForm
func (cls LocationSectionBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationSectionBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// C 新建
func (LocationSectionController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationSectionModel
	)

	// 表单
	form := (&LocationSectionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "区间代码")
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "区间名称")

	// 新建
	organizationSection := &models.LocationSectionModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		BeEnable:             form.BeEnable,
		OrganizationWorkshop: form.OrganizationWorkshop,
		OrganizationWorkArea: form.OrganizationWorkArea,
		LocationLines:        form.LocationLines,
	}
	if ret = models.BootByModel(models.LocationSectionModel{}).PrepareByDefaultDbDriver().Create(&organizationSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_section": organizationSection}))
}

func (LocationSectionController) D(ctx *gin.Context) {
	var (
		ret             *gorm.DB
		locationSection models.LocationSectionModel
	)
	// 查询
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationSection)
	wrongs.PanicWhenIsEmpty(ret, "区间")

	// 删除
	if ret := models.BootByModel(models.LocationSectionModel{}).PrepareByDefaultDbDriver().Delete(&locationSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (LocationSectionController) U(ctx *gin.Context) {
	var (
		ret                     *gorm.DB
		locationSection, repeat models.LocationSectionModel
	)

	// 表单
	form := (&LocationSectionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "区间代码")
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "区间名称")

	// 查询
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationSection)
	wrongs.PanicWhenIsEmpty(ret, "区间")

	// 编辑
	locationSection.BaseModel.Sort = form.Sort
	locationSection.Name = form.Name
	locationSection.BeEnable = form.BeEnable
	locationSection.OrganizationWorkshopUuid = form.OrganizationWorkshop.Uuid
	locationSection.OrganizationWorkAreaUuid = form.OrganizationWorkAreaUuid
	if ret = models.BootByModel(models.LocationSectionModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&locationSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_section": locationSection}))
}

// PutBindLines 绑定线别
func (LocationSectionController) PutBindLines(ctx *gin.Context) {
	var (
		ret                                  *gorm.DB
		locationSection                      models.LocationSectionModel
		pivotLocationLineAndLocationSections []models.PivotLocationLineAndLocationSectionModel
	)

	// 表单
	form := new(LocationSectionBindLocationLinesForm).ShouldBind(ctx)

	if ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&locationSection); ret.Error != nil {
		wrongs.PanicWhenIsEmpty(ret, "区间")
	}

	// 删除原有绑定关系
	ret = models.BootByModel(models.BaseModel{}).PrepareByDefaultDbDriver().Exec("delete from pivot_location_line_and_location_sections where location_section_id = ?", locationSection.Id)

	// 创建绑定关系
	if len(form.LocationLines) > 0 {
		for _, locationLine := range form.LocationLines {
			pivotLocationLineAndLocationSections = append(pivotLocationLineAndLocationSections, models.PivotLocationLineAndLocationSectionModel{
				LocationLineId:    locationLine.Id,
				LocationSectionId: locationSection.Id,
			})
		}
		models.BootByModel(models.PivotLocationLineAndLocationSectionModel{}).
			PrepareByDefaultDbDriver().
			CreateInBatches(&pivotLocationLineAndLocationSections, 100)
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{}))
}

// S 详情
func (LocationSectionController) S(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationSection models.LocationSectionModel
	)
	ret = models.BootByModel(models.LocationSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloadsByDefault().
		PrepareUseQueryByDefaultDbDriver(ctx).
		First(&organizationSection)
	wrongs.PanicWhenIsEmpty(ret, "区间")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_section": organizationSection}))
}

// I 列表
func (LocationSectionController) I(ctx *gin.Context) {
	var (
		locationSections []models.LocationSectionModel
		count            int64
		db               *gorm.DB
	)
	db = models.BootByModel(models.LocationSectionModel{}).
		SetWhereFields("unique_code", "Name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
		SetPreloadsByDefault().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&locationSections)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_sections": locationSections}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&locationSections)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"location_sections": locationSections}, ctx.Query("__page__"), count))
	}
}
