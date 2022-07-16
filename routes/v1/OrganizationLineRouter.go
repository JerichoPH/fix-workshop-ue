package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationLineRouter struct{}

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
func (cls OrganizationLineStoreForm) ShouldBind(ctx *gin.Context) OrganizationLineStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.UniqueCode == "" {
		panic(exceptions.ThrowForbidden("线别代码必填"))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("线别名称必填"))
	}
	// 查询路局
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		(&models.BaseModel{}).
			SetModel(models.OrganizationRailwayModel{}).
			SetScopes((&models.OrganizationRailwayModel{}).ScopeBeEnable).
			DB().
			Where("uuid in ?", cls.OrganizationRailwayUUIDs).
			Find(&cls.OrganizationRailways)
	}
	// 查询站段
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		(&models.BaseModel{}).
			SetModel(models.OrganizationParagraphModel{}).
			SetScopes((&models.OrganizationParagraphModel{}).ScopeBeEnable).
			DB().
			Where("uuid in ?", cls.OrganizationParagraphUUIDs).
			Find(&cls.OrganizationParagraphs)
	}
	// 查询战场
	if len(cls.OrganizationStationUUIDs) > 0 {
		(&models.BaseModel{}).
			SetModel(models.OrganizationStationModel{}).
			SetScopes((&models.OrganizationStationModel{}).ScopeBeEnable).
			DB().
			Where("uuid in ?", cls.OrganizationStationUUIDs).
			Find(&cls.OrganizationStations)
	}

	return cls
}

// 加载路由
func (cls *OrganizationLineRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("line", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationLineModel
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别代码")
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别名称")

			// 新建
			if ret = models.Init(models.OrganizationLineModel{}).
				DB().
				Create(&models.OrganizationLineModel{
					BaseModel:              models.BaseModel{Sort: form.Sort},
					UniqueCode:             form.UniqueCode,
					Name:                   form.Name,
					BeEnable:               form.BeEnable,
					OrganizationRailways:   form.OrganizationRailways,
					OrganizationParagraphs: form.OrganizationParagraphs,
					OrganizationStations:   form.OrganizationStations,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除
		r.DELETE("line/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var organizationLine models.OrganizationLineModel
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationLine)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "线别")

			if ret = models.Init(models.OrganizationLineModel{}).
				DB().
				Delete(&organizationLine); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("line/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (&OrganizationLineStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationLineModel
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别代码")
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别名称")

			// 查询
			var organizationLine models.OrganizationLineModel
			ret = models.Init(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationLine)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "线别")

			// 修改
			organizationLine.UniqueCode = form.UniqueCode
			organizationLine.Name = form.Name
			organizationLine.Sort = form.Sort
			organizationLine.BeEnable = form.BeEnable
			organizationLine.OrganizationRailways = form.OrganizationRailways
			organizationLine.OrganizationParagraphs = form.OrganizationParagraphs
			organizationLine.OrganizationStations = form.OrganizationStations
			if ret = (&models.BaseModel{}).
				SetModel(&models.OrganizationLineModel{}).
				DB().
				Save(&organizationLine); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET("line/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var organizationLine models.OrganizationLineModel
			ret = models.Init(models.OrganizationLineModel{}).
				SetScopes((&models.OrganizationLineModel{}).ScopeBeEnable).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationLine)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "线别")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_line": organizationLine}))
		})

		// 列表
		r.GET("line", func(ctx *gin.Context) {
			var organizationLines []models.OrganizationLineModel
			models.Init(models.OrganizationLineModel{}).
				SetScopes((&models.OrganizationLineModel{}).ScopeBeEnable).
				SetWhereFields("unique_code", "name", "be_enable", "sort").
				PrepareQuery(ctx).
				Find(&organizationLines)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_lines": organizationLines}))
		})
	}
}
