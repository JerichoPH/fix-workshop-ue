package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationRailwayRouter struct{}

// OrganizationRailwayStoreForm 路局新建表单
type OrganizationRailwayStoreForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	ShortName                  string   `form:"short_name" json:"short_name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationLineUUIDs      []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines          []*models.OrganizationLineModel
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs     []models.OrganizationParagraphModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationRailwayStoreForm
func (cls OrganizationRailwayStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailwayStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.UniqueCode == "" {
		panic(exceptions.ThrowForbidden("路局代码必填"))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("路局名称必填"))
	}
	if len(cls.OrganizationLineUUIDs) > 0 {
		models.Init(models.OrganizationLineModel{}).
			DB().
			Where("uuid in ?", cls.OrganizationLineUUIDs).
			Find(&cls.OrganizationLines)
	}
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(models.OrganizationParagraphModel{}).
			DB().
			Where("uuid in ?", cls.OrganizationParagraphUUIDs).
			Find(&cls.OrganizationParagraphs)
	}

	return cls
}

// OrganizationRailwayUpdateForm 编辑表单
type OrganizationRailwayUpdateForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	ShortName                  string   `form:"short_name" json:"short_name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationLineUUIDs      []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *OrganizationRailwayRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建路局
		r.POST("railway", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 新建
			if ret = (&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				DB().
				Create(&models.OrganizationRailwayModel{
					UniqueCode:             form.UniqueCode,
					Name:                   form.Name,
					ShortName:              form.ShortName,
					BeEnable:               form.BeEnable,
					OrganizationLines:      form.OrganizationLines,
					OrganizationParagraphs: form.OrganizationParagraphs,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除
		r.DELETE("railway/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 查询
			var organizationRailway models.OrganizationRailwayModel
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationRailway)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局")

			// 删除
			models.Init(models.OrganizationRailwayModel{}).
				DB().
				Delete(&organizationRailway)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("railway/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 查询
			var organizationRailway models.OrganizationRailwayModel
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationRailway)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局")

			// 修改
			organizationRailway.Sort = form.Sort
			organizationRailway.UniqueCode = form.UniqueCode
			organizationRailway.Name = form.Name
			organizationRailway.ShortName = form.ShortName
			organizationRailway.BeEnable = form.BeEnable
			organizationRailway.OrganizationLines = form.OrganizationLines
			organizationRailway.OrganizationParagraphs = form.OrganizationParagraphs
			if ret = models.Init(models.OrganizationRailwayModel{}).
				DB().
				Save(&organizationRailway); ret.Error != nil {
				panic(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET("railway/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var organizationRailway models.OrganizationRailwayModel
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationRailway)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railway": organizationRailway}))
		})

		// 列表
		r.GET("railway", func(ctx *gin.Context) {
			var organizationRailways []models.OrganizationRailwayModel

			models.Init(models.OrganizationRailwayModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "short_name", "be_enable").
				PrepareQuery(ctx).
				Find(&organizationRailways)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railways": organizationRailways}))
		})
	}
}
