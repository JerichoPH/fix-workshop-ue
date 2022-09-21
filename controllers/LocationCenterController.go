package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LocationCenterController struct{}

// LocationCenterStoreForm 新建中心表单
type LocationCenterStoreForm struct {
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
func (cls LocationCenterStoreForm) ShouldBind(ctx *gin.Context) LocationCenterStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("中心代码必填")
	}
	if len(cls.UniqueCode) != 3 {
		wrongs.PanicValidate("中心代码必须是3位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("中心名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("中心名称不能大于64位")
	}
	if cls.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUuid}).
		SetPreloads("OrganizationParagraph", "OrganizationParagraph.OrganizationRailway").
		PrepareByDefault().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")
	if cls.OrganizationWorkAreaUuid != "" {
		models.BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUuid}).
			PrepareByDefault().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "所属工区")
	}
	if len(cls.LocationLineUuids) > 0 {
		models.BootByModel(models.LocationLineModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.LocationLineUuids).
			Find(&cls.LocationLines)
	}

	return cls
}

// LocationCenterBindLocationLinesForm 中心绑定线别表单
type LocationCenterBindLocationLinesForm struct {
	LocationLineUuids []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationCenterBindLocationLinesForm
func (cls LocationCenterBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationCenterBindLocationLinesForm {
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
func (LocationCenterController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.LocationCenterModel
	)

	// 表单
	form := new(LocationCenterStoreForm).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"unique_code": form.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway.UniqueCode + form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "中心代码")
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "中心名称")

	// 新建
	locationCenter := &models.LocationCenterModel{
		BaseModel:                models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:               form.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway.UniqueCode + form.UniqueCode,
		Name:                     form.Name,
		BeEnable:                 form.BeEnable,
		OrganizationWorkshopUuid: form.OrganizationWorkshop.Uuid,
		OrganizationWorkAreaUuid: form.OrganizationWorkAreaUuid,
		LocationLines:            form.LocationLines,
	}
	if ret = models.BootByModel(models.LocationCenterModel{}).PrepareByDefault().Create(&locationCenter); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"location_center": locationCenter}))
}

// D 删除
func (LocationCenterController) D(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		locationCenter models.LocationCenterModel
	)

	// 查询
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationCenter)
	wrongs.PanicWhenIsEmpty(ret, "中心")

	// 删除
	if ret := models.BootByModel(models.LocationCenterModel{}).PrepareByDefault().Delete(&locationCenter); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (LocationCenterController) U(ctx *gin.Context) {
	var (
		ret                    *gorm.DB
		locationCenter, repeat models.LocationCenterModel
	)

	// 表单
	form := (&LocationCenterStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "中心代码")
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "中心名称")

	// 查询
	ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationCenter)
	wrongs.PanicWhenIsEmpty(ret, "中心")

	// 编辑
	locationCenter.BaseModel.Sort = form.Sort
	locationCenter.Name = form.Name
	locationCenter.BeEnable = form.BeEnable
	locationCenter.OrganizationWorkshopUuid = form.OrganizationWorkshop.Uuid
	locationCenter.OrganizationWorkAreaUuid = form.OrganizationWorkAreaUuid
	if ret = models.BootByModel(models.LocationCenterModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefault().Save(&locationCenter); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"location_center": locationCenter}))
}

// PutBindLines 绑定线别
func (LocationCenterController) PutBindLines(ctx *gin.Context) {
	var (
		ret                                 *gorm.DB
		locationCenter                      models.LocationCenterModel
		pivotLocationLineAndLocationCenters []models.PivotLocationLineAndLocationCenter
	)

	// 表单
	form := (&LocationCenterBindLocationLinesForm{}).ShouldBind(ctx)

	if ret = models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationCenter); ret.Error != nil {
		wrongs.PanicWhenIsEmpty(ret, "中心")
	}

	// 删除原有绑定关系
	ret = models.BootByModel(models.BaseModel{}).PrepareByDefault().Exec("delete from pivot_location_line_and_location_centers where location_center_id = ?", locationCenter.Id)

	// 创建绑定关系
	if len(form.LocationLines) > 0 {
		for _, locationLine := range form.LocationLines {
			pivotLocationLineAndLocationCenters = append(pivotLocationLineAndLocationCenters, models.PivotLocationLineAndLocationCenter{
				LocationLineId:   locationLine.Id,
				LocationCenterId: locationCenter.Id,
			})
		}
		models.BootByModel(models.PivotLocationLineAndLocationCenter{}).
			PrepareByDefault().
			CreateInBatches(&pivotLocationLineAndLocationCenters, 100)
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{}))
}

// S 详情
func (LocationCenterController) S(ctx *gin.Context) {
	var locationCenter models.LocationCenterModel
	ret := models.BootByModel(models.LocationCenterModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetWhereFields("be_enable").
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		First(&locationCenter)
	wrongs.PanicWhenIsEmpty(ret, "中心")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_center": locationCenter}))
}

// I 列表
func (LocationCenterController) I(ctx *gin.Context) {
	var locationCenters []models.LocationCenterModel
	models.BootByModel(models.LocationCenterModel{}).
		SetWhereFields().
		SetPreloadsByDefault().
		PrepareUseQueryByDefault(ctx).
		Find(&locationCenters)

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"location_centers": locationCenters}))
}
