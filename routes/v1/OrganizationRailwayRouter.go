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

type OrganizationRailwayStoreForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	ShortName                  string   `form:"short_name" json:"short_name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationLineUUIDs      []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
}

type OrganizationRailwayUpdateForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	ShortName                  string   `form:"short_name" json:"short_name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationLineUUIDs      []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
}

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
			var form OrganizationRailwayStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.UniqueCode == "" {
				panic(exceptions.ThrowForbidden("路局代码必填"))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("路局名称必填"))
			}
			var organizationLines []*models.OrganizationLineModel
			if len(form.OrganizationLineUUIDs) > 0 {
				(&models.BaseModel{}).
					SetModel(models.OrganizationLineModel{}).
					DB().
					Where("uuid in ?", form.OrganizationLineUUIDs).
					Find(&organizationLines)
			}
			var organizationParagraphs []models.OrganizationParagraphModel
			if len(form.OrganizationParagraphUUIDs) > 0 {
				(&models.BaseModel{}).
					SetModel(models.OrganizationParagraphModel{}).
					DB().
					Where("uuid in ?", form.OrganizationParagraphUUIDs).
					Find(&organizationParagraphs)
			}

			// 新建
			if ret = (&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				DB().
				Create(&models.OrganizationRailwayModel{
					UniqueCode:             form.UniqueCode,
					Name:                   form.Name,
					ShortName:              form.ShortName,
					BeEnable:               form.BeEnable,
					OrganizationLines:      organizationLines,
					OrganizationParagraphs: organizationParagraphs,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除
		r.DELETE("railway/:unique_code", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 查询
			var organizationRailway models.OrganizationRailwayModel
			(&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationRailway)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局")

			// 删除
			(&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				DB().
				Delete(&organizationRailway)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("railway/:unique_code", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form OrganizationRailwayUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.UniqueCode == "" {
				panic(exceptions.ThrowForbidden("路局代码必填"))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("路局名称必填"))
			}
			var organizationLines []*models.OrganizationLineModel
			if len(form.OrganizationLineUUIDs) > 0 {
				(&models.BaseModel{}).
					SetModel(models.OrganizationLineModel{}).
					DB().
					Where("uuid in ?", form.OrganizationLineUUIDs).
					Find(&organizationLines)
			}
			var organizationParagraphs []models.OrganizationParagraphModel
			if len(form.OrganizationParagraphUUIDs) > 0 {
				(&models.BaseModel{}).
					SetModel(models.OrganizationParagraphModel{}).
					DB().
					Where("uuid in ?", form.OrganizationParagraphUUIDs).
					Find(&organizationParagraphs)
			}

			// 查询
			var organizationRailway models.OrganizationRailwayModel
			ret = (&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
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
			organizationRailway.OrganizationLines = organizationLines
			organizationRailway.OrganizationParagraphs = organizationParagraphs
			if ret = (&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				DB().
				Save(&organizationRailway); ret.Error != nil {
				panic(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET("railway/:unique_code", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var organizationRailway models.OrganizationRailwayModel
			ret = (&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationRailway)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railway": organizationRailway}))
		})

		// 列表
		r.GET("railway", func(ctx *gin.Context) {
			var organizationRailways []models.OrganizationRailwayModel
			{
			}
			(&models.BaseModel{}).
				SetModel(models.OrganizationRailwayModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "short_name", "be_enable").
				PrepareQuery(ctx).
				Find(&organizationRailways)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railways": organizationRailways}))
		})
	}
}
