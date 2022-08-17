package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// LocationLineRouter 线别路由
type LocationLineRouter struct{}

// LocationLineStoreForm 新建线别表单
type LocationLineStoreForm struct {
	Sort                            int64    `form:"sort" json:"sort"`
	UniqueCode                      string   `form:"unique_code" json:"unique_code"`
	Name                            string   `form:"name" json:"name"`
	BeEnable                        bool     `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUIDs        []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways            []*models.OrganizationRailwayModel
	OrganizationParagraphUUIDs      []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs          []*models.OrganizationParagraphModel
	OrganizationWorkshopUUIDs       []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops           []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUUIDs       []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas           []*models.OrganizationWorkAreaModel
	LocationSectionUUIDs            []string `form:"location_section_uuids" json:"location_section_uuids"`
	LocationSections                []*models.LocationSectionModel
	LocationStationUUIDs            []string `form:"location_station_uuids" json:"location_station_uuids"`
	LocationStations                []*models.LocationStationModel
	LocationRailroadGradeCrossUUIDs []string `form:"location_railroad_grade_cross_uuids" json:"location_railroad_grade_cross_uuids"`
	LocationRailroadGradeCrosses    []*models.LocationRailroadGradeCrossModel
	LocationCenterUUIDs             []string `form:"location_center_uuids" json:"location_center_uuids"`
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
	if cls.Name == "" {
		wrongs.PanicValidate("线别名称必填")
	}
	// 查询路局
	if len(cls.OrganizationRailwayUUIDs) > 0 {
		models.Init(models.OrganizationRailwayModel{}).
			Prepare("").
			Where("uuid in ?", cls.OrganizationRailwayUUIDs).
			Find(&cls.OrganizationRailways)
	}
	// 查询站段
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(models.OrganizationParagraphModel{}).
			Prepare("").
			Where("uuid in ?", cls.OrganizationParagraphUUIDs).
			Find(&cls.OrganizationParagraphs)
	}
	// 查询车间
	if len(cls.OrganizationWorkshopUUIDs) > 0 {
		models.Init(models.OrganizationWorkshopModel{}).
			Prepare("").
			Where("uuid in ?", cls.OrganizationWorkshopUUIDs).
			Find(&cls.OrganizationWorkshops)
	}
	// 查询工区
	if len(cls.OrganizationWorkAreaUUIDs) > 0 {
		models.Init(models.OrganizationWorkAreaModel{}).
			Prepare("").
			Where("uuid in ?", cls.OrganizationWorkAreaUUIDs).
			Find(&cls.OrganizationWorkAreas)
	}
	// 查询区间
	if len(cls.LocationSectionUUIDs) > 0 {
		models.Init(models.LocationSectionModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationSectionUUIDs).
			Find(&cls.LocationSections)
	}
	// 查询站场
	if len(cls.LocationStationUUIDs) > 0 {
		models.Init(models.LocationStationModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationStationUUIDs).
			Find(&cls.LocationStations)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.LocationRailroadGradeCrossModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationRailroadGradeCrossUUIDs).
			Find(&cls.LocationRailroadGradeCrosses)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.LocationRailroadGradeCrossModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationRailroadGradeCrossUUIDs).
			Find(&cls.LocationRailroadGradeCrosses)
	}

	return cls
}

// LocationLineBindForm 线别多对多绑定
type LocationLineBindForm struct {
	OrganizationRailwayUUIDs        []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways            []*models.OrganizationRailwayModel
	OrganizationParagraphUUIDs      []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs          []*models.OrganizationParagraphModel
	OrganizationWorkshopUUIDs       []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops           []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUUIDs       []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas           []*models.OrganizationWorkAreaModel
	LocationSectionUUIDs            []string `form:"location_section_uuids" json:"location_section_uuids"`
	LocationSections                []*models.LocationSectionModel
	LocationStationUUIDs            []string `form:"location_station_uuids" json:"location_station_uuids"`
	LocationStations                []*models.LocationStationModel
	LocationRailroadGradeCrossUUIDs []string `form:"location_railroad_grade_cross_uuids" json:"location_railroad_grade_cross_uuids"`
	LocationRailroadGradeCrosses    []*models.LocationRailroadGradeCrossModel
	LocationCenterUUIDs             []string `form:"location_center_uuids" json:"location_center_uuids"`
	LocationCenters                 []*models.LocationCenterModel
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationLineRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/locationLine",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationLineModel
			)

			// 表单
			form := (&LocationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 新建
			organizationLine := &models.LocationLineModel{
				BaseModel:                    models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:                   form.UniqueCode,
				Name:                         form.Name,
				BeEnable:                     form.BeEnable,
				LocationSections:             form.LocationSections,
				LocationStations:             form.LocationStations,
				LocationRailroadGradeCrosses: form.LocationRailroadGradeCrosses,
				LocationCenters:              form.LocationCenters,
			}
			if ret = models.Init(models.LocationLineModel{}).Prepare("").Create(organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_line": organizationLine}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.LocationLineModel
			)
			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除
			if ret = models.Init(models.LocationLineModel{}).Prepare("").Delete(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				locationLine, repeat models.LocationLineModel
			)

			// 表单
			form := (&LocationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 修改
			locationLine.UniqueCode = form.UniqueCode
			locationLine.Name = form.Name
			locationLine.Sort = form.Sort
			locationLine.BeEnable = form.BeEnable

			if ret = (&models.BaseModel{}).SetModel(&models.LocationLineModel{}).Prepare("").Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			if ret = models.Init(models.LocationLineModel{}).
				Prepare("").
				Where("uuid = ?", ctx.Param("uuid")).
				Updates(map[string]interface{}{
					"sort":        form.Sort,
					"unique_code": form.UniqueCode,
					"name":        form.Name,
					"be_enable":   form.BeEnable,
				}); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_line": locationLine}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.LocationLineModel
			)
			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx,"").
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_line": organizationLine}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var organizationLines []models.LocationLineModel
			models.Init(models.LocationLineModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "sort").
				PrepareQuery(ctx,"").
				Find(&organizationLines)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_lines": organizationLines}))
		})
	}
}
