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

// OrganizationLineRouter 线别路由
type OrganizationLineRouter struct{}

// OrganizationLineStoreForm 新建线别表单
type OrganizationLineStoreForm struct {
	Sort                                int64    `form:"sort" json:"sort"`
	UniqueCode                          string   `form:"unique_code" json:"unique_code"`
	Name                                string   `form:"name" json:"name"`
	BeEnable                            bool     `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUIDs            []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways                []*models.OrganizationRailwayModel
	OrganizationParagraphUUIDs          []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs              []*models.OrganizationParagraphModel
	OrganizationWorkshopUUIDs           []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops               []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUUIDs           []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas               []*models.OrganizationWorkAreaModel
	OrganizationSectionUUIDs            []string `form:"organization_section_uuids" json:"organization_section_uuids"`
	OrganizationSections                []*models.OrganizationSectionModel
	OrganizationStationUUIDs            []string `form:"organization_station_uuids" json:"organization_station_uuids"`
	OrganizationStations                []*models.OrganizationStationModel
	OrganizationRailroadGradeCrossUUIDs []string `form:"organization_railroad_grade_cross_uuids" json:"organization_railroad_grade_cross_uuids"`
	OrganizationRailroadGradeCrosses    []*models.OrganizationRailroadGradeCrossModel
	OrganizationCenterUUIDs             []string `form:"organization_center_uuids" json:"organization_center_uuids"`
	OrganizationCenters                 []*models.OrganizationCenterModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationLineStoreForm
func (cls OrganizationLineStoreForm) ShouldBind(ctx *gin.Context) OrganizationLineStoreForm {
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
	if len(cls.OrganizationSectionUUIDs) > 0 {
		models.Init(models.OrganizationSectionModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationSectionUUIDs).
			Find(&cls.OrganizationSections)
	}
	// 查询站场
	if len(cls.OrganizationStationUUIDs) > 0 {
		models.Init(models.OrganizationStationModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationStationUUIDs).
			Find(&cls.OrganizationStations)
	}
	// 查询道口
	if len(cls.OrganizationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.OrganizationRailroadGradeCrossModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationRailroadGradeCrossUUIDs).
			Find(&cls.OrganizationRailroadGradeCrosses)
	}
	// 查询道口
	if len(cls.OrganizationRailroadGradeCrossUUIDs) > 0 {
		models.Init(models.OrganizationRailroadGradeCrossModel{}).
			GetSession().
			Where("uuid in ?", cls.OrganizationRailroadGradeCrossUUIDs).
			Find(&cls.OrganizationRailroadGradeCrosses)
	}

	return cls
}

// OrganizationLineBindForm 线别多对多绑定
type OrganizationLineBindForm struct {
	OrganizationRailwayUUIDs            []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways                []*models.OrganizationRailwayModel
	OrganizationParagraphUUIDs          []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs              []*models.OrganizationParagraphModel
	OrganizationWorkshopUUIDs           []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops               []*models.OrganizationWorkshopModel
	OrganizationWorkAreaUUIDs           []string `form:"organization_work_area_uuids" json:"organization_work_area_uuids"`
	OrganizationWorkAreas               []*models.OrganizationWorkAreaModel
	OrganizationSectionUUIDs            []string `form:"organization_section_uuids" json:"organization_section_uuids"`
	OrganizationSections                []*models.OrganizationSectionModel
	OrganizationStationUUIDs            []string `form:"organization_station_uuids" json:"organization_station_uuids"`
	OrganizationStations                []*models.OrganizationStationModel
	OrganizationRailroadGradeCrossUUIDs []string `form:"organization_railroad_grade_cross_uuids" json:"organization_railroad_grade_cross_uuids"`
	OrganizationRailroadGradeCrosses    []*models.OrganizationRailroadGradeCrossModel
	OrganizationCenterUUIDs             []string `form:"organization_center_uuids" json:"organization_center_uuids"`
	OrganizationCenters                 []*models.OrganizationCenterModel
}

// ShouldBind 绑定路由
//  @receiver cls
//  @param ctx
//  @return OrganizationLineBindForm
func (cls OrganizationLineBindForm) ShouldBind(ctx *gin.Context) OrganizationLineBindForm {
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
	models.Init(models.OrganizationSectionModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationSectionUUIDs).
		Find(&cls.OrganizationSections)
	// 查询站场
	models.Init(models.OrganizationStationModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationStationUUIDs).
		Find(&cls.OrganizationStations)
	// 查询道口
	models.Init(models.OrganizationRailroadGradeCrossModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationRailroadGradeCrossUUIDs).
		Find(&cls.OrganizationRailroadGradeCrosses)
	// 查询中心
	models.Init(models.OrganizationCenterModel{}).
		GetSession().
		Where("uuid in ?", cls.OrganizationCenterUUIDs).
		Find(&cls.OrganizationCenters)

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationLineRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("line", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 新建
			organizationLine := &models.OrganizationLineModel{
				BaseModel:                        models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:                       form.UniqueCode,
				Name:                             form.Name,
				BeEnable:                         form.BeEnable,
				OrganizationRailways:             form.OrganizationRailways,
				OrganizationParagraphs:           form.OrganizationParagraphs,
				OrganizationWorkshops:            form.OrganizationWorkshops,
				OrganizationWorkAreas:            form.OrganizationWorkAreas,
				OrganizationSections:             form.OrganizationSections,
				OrganizationStations:             form.OrganizationStations,
				OrganizationRailroadGradeCrosses: form.OrganizationRailroadGradeCrosses,
				OrganizationCenters:              form.OrganizationCenters,
			}
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Create(organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_line": organizationLine}))
		})

		// 删除
		r.DELETE("line/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)
			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Delete(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("line/:uuid", func(ctx *gin.Context) {
			var (
				ret                      *gorm.DB
				organizationLine, repeat models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别代码")
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "线别名称")

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 修改
			organizationLine.UniqueCode = form.UniqueCode
			organizationLine.Name = form.Name
			organizationLine.Sort = form.Sort
			organizationLine.BeEnable = form.BeEnable
			organizationLine.OrganizationRailways = form.OrganizationRailways
			organizationLine.OrganizationParagraphs = form.OrganizationParagraphs
			organizationLine.OrganizationWorkshops = form.OrganizationWorkshops
			organizationLine.OrganizationWorkAreas = form.OrganizationWorkAreas
			organizationLine.OrganizationSections = form.OrganizationSections
			organizationLine.OrganizationStations = form.OrganizationStations
			organizationLine.OrganizationRailroadGradeCrosses = form.OrganizationRailroadGradeCrosses
			organizationLine.OrganizationCenters = form.OrganizationCenters
			if ret = (&models.BaseModel{}).SetModel(&models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_line": organizationLine}))
		})

		// 绑定路局
		r.PUT("line/:uuid/bindOrganizationRailways", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_railways where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationRailways = form.OrganizationRailways
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定路局
		r.PUT("line/:uuid/bindOrganizationParagraphs", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_paragraphs where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationParagraphs = form.OrganizationParagraphs
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定车间
		r.PUT("line/:uuid/bindOrganizationWorkshops", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_workshops where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationWorkshops = form.OrganizationWorkshops
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定工区
		r.PUT("line/:uuid/bindOrganizationWorkAreas", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_work_areas where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationWorkAreas = form.OrganizationWorkAreas
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定区间
		r.PUT("line/:uuid/bindOrganizationSections", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_sections where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationSections = form.OrganizationSections
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定站场
		r.PUT("line/:uuid/bindOrganizationStations", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_stations where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationStations = form.OrganizationStations
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定道口
		r.PUT("line/:uuid/bindOrganizationRailroadGradeCrosses", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_railroad_grade_crosses where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationRailroadGradeCrosses = form.OrganizationRailroadGradeCrosses
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 绑定中心
		r.PUT("line/:uuid/bindOrganizationCenters", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)

			// 表单
			form := (&OrganizationLineBindForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			// 删除原有绑定关系
			ret = (&databases.MySql{}).GetConn().Exec("delete from pivot_organization_line_and_organization_centers where organization_line_id = ?", organizationLine.ID)

			// 编辑
			organizationLine.OrganizationCenters = form.OrganizationCenters
			if ret = models.Init(models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("绑定成功").OK(tools.Map{}))
		})

		// 详情
		r.GET("line/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				organizationLine models.OrganizationLineModel
			)
			// 查询
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationLine)
			wrongs.PanicWhenIsEmpty(ret, "线别")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_line": organizationLine}))
		})

		// 列表
		r.GET("line", func(ctx *gin.Context) {
			var organizationLines []models.OrganizationLineModel
			models.Init(models.OrganizationLineModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "sort").
				SetScopes((&models.OrganizationLineModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationLines)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_lines": organizationLines}))
		})
	}
}
