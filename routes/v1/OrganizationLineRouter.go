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

// OrganizationLineRouter 线别路由
type OrganizationLineRouter struct{}

// OrganizationLineStoreForm 新建线别表单
type OrganizationLineStoreForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUIDs   []string `form:"organization_railway_uuids" json:"organization_railway_uuids"`
	OrganizationRailways       []*models.OrganizationRailwayModel
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs     []*models.OrganizationParagraphModel
	OrganizationStationUUIDs   []string `form:"organization_station_uuids" json:"organization_station_uuids"`
	OrganizationStations       []*models.OrganizationStationModel
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
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(models.OrganizationRailwayModel{}).GetSession().Where("uuid in ?", cls.OrganizationRailwayUUIDs).Find(&cls.OrganizationRailways)
	}
	// 查询站段
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(models.OrganizationParagraphModel{}).GetSession().Where("uuid in ?", cls.OrganizationParagraphUUIDs).Find(&cls.OrganizationParagraphs)
	}
	// 查询战场
	if len(cls.OrganizationStationUUIDs) > 0 {
		models.Init(models.OrganizationStationModel{}).GetSession().Where("uuid in ?", cls.OrganizationStationUUIDs).Find(&cls.OrganizationStations)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls OrganizationLineRouter) Load(engine *gin.Engine) {
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
				BaseModel:              models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:             form.UniqueCode,
				Name:                   form.Name,
				BeEnable:               form.BeEnable,
				OrganizationRailways:   form.OrganizationRailways,
				OrganizationParagraphs: form.OrganizationParagraphs,
				OrganizationStations:   form.OrganizationStations,
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
			organizationLine.OrganizationStations = form.OrganizationStations
			if ret = (&models.BaseModel{}).SetModel(&models.OrganizationLineModel{}).GetSession().Save(&organizationLine); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_line": organizationLine}))
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
