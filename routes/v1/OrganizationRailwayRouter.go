package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/OrganizationModels"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// OrganizationRailwayRouter 路局路由
type OrganizationRailwayRouter struct{}

// OrganizationRailwayStoreForm 新建路局表单
type OrganizationRailwayStoreForm struct {
	Sort                       int64    `form:"sort" json:"sort"`
	UniqueCode                 string   `form:"unique_code" json:"unique_code"`
	Name                       string   `form:"name" json:"name"`
	ShortName                  string   `form:"short_name" json:"short_name"`
	BeEnable                   bool     `form:"be_enable" json:"be_enable"`
	OrganizationLineUUIDs      []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines          []*OrganizationModels.OrganizationLineModel
	OrganizationParagraphUUIDs []string `form:"organization_paragraph_uuids" json:"organization_paragraph_uuids"`
	OrganizationParagraphs     []OrganizationModels.OrganizationParagraphModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationRailwayStoreForm
func (cls OrganizationRailwayStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailwayStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("路局代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("路局名称必填")
	}
	if len(cls.OrganizationLineUUIDs) > 0 {
		models.Init(OrganizationModels.OrganizationLineModel{}).DB().Where("uuid in ?", cls.OrganizationLineUUIDs).Find(&cls.OrganizationLines)
	}
	if len(cls.OrganizationParagraphUUIDs) > 0 {
		models.Init(OrganizationModels.OrganizationParagraphModel{}).DB().Where("uuid in ?", cls.OrganizationParagraphUUIDs).Find(&cls.OrganizationParagraphs)
	}

	return cls
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
			var (
				ret    *gorm.DB
				repeat OrganizationModels.OrganizationRailwayModel
			)

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局代码")
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局名称")
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局简称")

			// 新建
			organizationRailway := &OrganizationModels.OrganizationRailwayModel{
				BaseModel:              models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:             form.UniqueCode,
				Name:                   form.Name,
				ShortName:              form.ShortName,
				BeEnable:               form.BeEnable,
				OrganizationLines:      form.OrganizationLines,
				OrganizationParagraphs: form.OrganizationParagraphs,
			}
			if ret = (&models.BaseModel{}).SetModel(OrganizationModels.OrganizationRailwayModel{}).DB().Create(&organizationRailway); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_railway": organizationRailway}))
		})

		// 删除
		r.DELETE("railway/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 查询
			var organizationRailway OrganizationModels.OrganizationRailwayModel
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailway)
			abnormals.PanicWhenIsEmpty(ret, "路局")
			// 删除
			if ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).DB().Delete(&organizationRailway); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("railway/:uuid", func(ctx *gin.Context) {
			var (
				ret                         *gorm.DB
				organizationRailway, repeat OrganizationModels.OrganizationRailwayModel
			)

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局代码")
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局名称")
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "路局简称")

			// 查询
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailway)
			abnormals.PanicWhenIsEmpty(ret, "路局")

			// 修改
			organizationRailway.Sort = form.Sort
			organizationRailway.UniqueCode = form.UniqueCode
			organizationRailway.Name = form.Name
			organizationRailway.ShortName = form.ShortName
			organizationRailway.BeEnable = form.BeEnable
			organizationRailway.OrganizationLines = form.OrganizationLines
			organizationRailway.OrganizationParagraphs = form.OrganizationParagraphs
			if ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).DB().Save(&organizationRailway); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_railway": organizationRailway}))
		})

		// 详情
		r.GET("railway/:uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				organizationRailway OrganizationModels.OrganizationRailwayModel
			)
			ret = models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationRailway)
			abnormals.PanicWhenIsEmpty(ret, "路局")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railway": organizationRailway}))
		})

		// 列表
		r.GET("railway", func(ctx *gin.Context) {
			var organizationRailways []OrganizationModels.OrganizationRailwayModel

			models.Init(OrganizationModels.OrganizationRailwayModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "short_name", "be_enable").
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationRailways)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railways": organizationRailways}))
		})
	}
}
