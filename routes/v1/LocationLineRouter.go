package v1

import (
	"fix-workshop-ue/databases"
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
	Sort                                int64    `form:"sort" json:"sort"`
	UniqueCode                          string   `form:"unique_code" json:"unique_code"`
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
			GetSession().
			Where("uuid in ?", cls.OrganizationRailwayUUIDs).
			Find(&cls.OrganizationRailways)
	}
	// 查询站段
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(models.OrganizationParagraphModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationParagraphUUIDs).
			Find(&cls.OrganizationParagraphs)
	}
	// 查询车间
	if len(cls.OrganizationWorkshopUUIDs) > 0 {
		models.Init(models.OrganizationWorkshopModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationWorkshopUUIDs).
			Find(&cls.OrganizationWorkshops)
	}
	// 查询工区
	if len(cls.OrganizationWorkAreaUUIDs) > 0 {
		models.Init(models.OrganizationWorkAreaModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationWorkAreaUUIDs).
			Find(&cls.OrganizationWorkAreas)
	}
	// 查询区间
	if len(cls.LocationSectionUUIDs) > 0 {
		models.Init(models.LocationSectionModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationSectionUUIDs).
			Find(&cls.LocationSections)
	}
	// 查询站场
	if len(cls.LocationStationUUIDs) > 0 {
		models.Init(models.LocationStationModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationStationUUIDs).
			Find(&cls.LocationStations)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.LocationRailroadGradeCrossModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationRailroadGradeCrossUUIDs).
			Find(&cls.LocationRailroadGradeCrosses)
	}
	// 查询道口
	if len(cls.LocationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.LocationRailroadGradeCrossModel{}).
			GetSession().
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

// ShouldBind 绑定路由
//  @receiver cls
//  @param ctx
//  @return LocationLineBindForm
func (cls LocationLineBindForm) ShouldBind(ctx *gin.Context) LocationLineBindForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	// 查询路局
	models.Init(models.OrganizationRailwayModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationRailwayUUIDs).
		Find(&cls.OrganizationRailways)
	// 查询站段
	models.Init(models.OrganizationParagraphModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationParagraphUUIDs).
		Find(&cls.OrganizationParagraphs)
	// 查询车间
	models.Init(models.OrganizationWorkshopModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationWorkshopUUIDs).
		Find(&cls.OrganizationWorkshops)
	// 查询工区
	models.Init(models.OrganizationWorkAreaModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationWorkAreaUUIDs).
		Find(&cls.OrganizationWorkAreas)
	// 查询区间
	models.Init(models.LocationSectionModel{}).
		GetSession().
		Where("uuid in ?", cls.LocationSectionUUIDs).
		Find(&cls.LocationSections)
	// 查询站场
	models.Init(models.LocationStationModel{}).
		GetSession().
		Where("uuid in ?", cls.LocationStationUUIDs).
		Find(&cls.LocationStations)
	// 查询道口
	models.Init(models.LocationRailroadGradeCrossModel{}).
		GetSession().
		Where("uuid in ?", cls.LocationRailroadGradeCrossUUIDs).
		Find(&cls.LocationRailroadGradeCrosses)
	// 查询中心
	models.Init(models.LocationCenterModel{}).
		GetSession().
		Where("uuid in ?", cls.LocationCenterUUIDs).
		Find(&cls.LocationCenters)

	return cls
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
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 新建
			organizationLine := &models.LocationLineModel{
				BaseModel:                    models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:                   form.UniqueCode,
				Name:                         form.Name,
				BeEnable:                     form.BeEnable,
				OrganizationRailways:         form.OrganizationRailways,
				OrganizationParagraphs:       form.OrganizationParagraphs,
				OrganizationWorkshops:        form.OrganizationWorkshops,
				OrganizationWorkAreas:        form.OrganizationWorkAreas,
				LocationSections:             form.LocationSections,
				LocationStations:             form.LocationStations,
				LocationRailroadGradeCrosses: form.LocationRailroadGradeCrosses,
				LocationCenters:              form.LocationCenters,
			}
			if ret = models.Init(models.LocationLineModel{}).GetSession().Create(organizationLine); ret.Error != nil {
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
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除
			if ret = models.Init(models.LocationLineModel{}).GetSession().Delete(&organizationLine); ret.Error != nil {
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
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 修改
			locationLine.UniqueCode = form.UniqueCode
			locationLine.Name = form.Name
			locationLine.Sort = form.Sort
			locationLine.BeEnable = form.BeEnable
			locationLine.OrganizationRailways = form.OrganizationRailways
			locationLine.OrganizationParagraphs = form.OrganizationParagraphs
			locationLine.OrganizationWorkshops = form.OrganizationWorkshops
			locationLine.OrganizationWorkAreas = form.OrganizationWorkAreas
			locationLine.LocationSections = form.LocationSections
			locationLine.LocationStations = form.LocationStations
			locationLine.LocationRailroadGradeCrosses = form.LocationRailroadGradeCrosses
			locationLine.LocationCenters = form.LocationCenters
			if ret = (&models.BaseModel{}).SetModel(&models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_line": locationLine}))
		})

		// 绑定路局
		r.PUT(":uuid/bindOrganizationRailways", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_organization_railways where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.OrganizationRailways = form.OrganizationRailways
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定路局
		r.PUT(":uuid/bindOrganizationParagraphs", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_organization_paragraphs where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.OrganizationParagraphs = form.OrganizationParagraphs
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定车间
		r.PUT(":uuid/bindOrganizationWorkshops", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_organization_workshops where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.OrganizationWorkshops = form.OrganizationWorkshops
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定工区
		r.PUT(":uuid/bindOrganizationWorkAreas", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_organization_work_areas where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.OrganizationWorkAreas = form.OrganizationWorkAreas
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定区间
		r.PUT(":uuid/bindOrganizationSections", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_location_sections where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.LocationSections = form.LocationSections
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定站场
		r.PUT(":uuid/bindOrganizationStations", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_location_stations where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.LocationStations = form.LocationStations
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定道口
		r.PUT(":uuid/bindOrganizationRailroadGradeCrosses", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_location_railroad_grade_crosses where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.LocationRailroadGradeCrosses = form.LocationRailroadGradeCrosses
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定中心
		r.PUT(":uuid/bindOrganizationCenters", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				locationLine models.LocationLineModel
			)

			// 表单
			form := (&LocationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.LocationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_location_line_and_location_centers where location_line_id = ?", locationLine.ID)

			// 编辑
			locationLine.LocationCenters = form.LocationCenters
			if ret = models.Init(models.LocationLineModel{}).GetSession().Save(&locationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
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
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_line": organizationLine}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var organizationLines []models.LocationLineModel
			models.Init(models.LocationLineModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "sort").
				SetScopes((&models.LocationLineModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationLines)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_lines": organizationLines}))
		})
	}
}
