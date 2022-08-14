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

// OrganizationParagraphRouter 站段路由
type OrganizationParagraphRouter struct{}

// OrganizationParagraphStoreForm 新建站段表单
type OrganizationParagraphStoreForm struct {
	Sort                    int64  `form:"sort" json:"sort"`
	UniqueCode              string `form:"unique_code" json:"unique_code"`
	Name                    string `form:"name" json:"name"`
	ShortName               string `form:"short_name" json:"short_name"`
	BeEnable                bool   `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUID string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway     models.OrganizationRailwayModel
	OrganizationLineUUIDs   []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines       []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationParagraphStoreForm
func (cls OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("站段代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("站段名称必填")
	}
	if cls.OrganizationRailwayUUID == "" {
		wrongs.PanicValidate("所属路局必选")
	}
	ret = models.Init(models.OrganizationRailwayModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationRailwayUUID}).
		Prepare().
		First(&cls.OrganizationRailway)
	wrongs.PanicWhenIsEmpty(ret, "路局")
	if len(cls.OrganizationLineUUIDs) > 0 {
		models.Init(models.LocationLineModel{}).
			Prepare().
			Where("uuid in ?", cls.OrganizationLineUUIDs).
			Find(&cls.OrganizationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationParagraphRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/organizationParagraph",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationParagraphModel
			)

			// 表单
			form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站段代码")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站段名称")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站段简称")

			// 新建
			organizationParagraph := &models.OrganizationParagraphModel{
				BaseModel:           models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:          form.UniqueCode,
				Name:                form.Name,
				ShortName:           form.ShortName,
				BeEnable:            form.BeEnable,
				OrganizationRailway: form.OrganizationRailway,
				LocationLines:       form.OrganizationLines,
			}
			if ret = models.Init(models.OrganizationParagraphModel{}).Prepare().Create(&organizationParagraph); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				organizationParagraph models.OrganizationParagraphModel
			)

			// 查询
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationParagraph)
			wrongs.PanicWhenIsEmpty(ret, "站段")

			// 删除
			if ret = models.Init(models.OrganizationParagraphModel{}).Prepare().Delete(&organizationParagraph); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				organizationParagraph models.OrganizationParagraphModel
			)

			// 表单
			form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationParagraphModel
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站段代码")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站段名称")

			// 查询
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationParagraph)
			wrongs.PanicWhenIsEmpty(ret, "站段")

			// 编辑
			organizationParagraph.BaseModel.Sort = form.Sort
			organizationParagraph.UniqueCode = form.UniqueCode
			organizationParagraph.Name = form.Name
			organizationParagraph.ShortName = form.ShortName
			organizationParagraph.BeEnable = form.BeEnable
			organizationParagraph.OrganizationRailway = form.OrganizationRailway
			organizationParagraph.LocationLines = form.OrganizationLines
			if ret = models.Init(models.OrganizationParagraphModel{}).Prepare().Save(&organizationParagraph); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				organizationParagraph models.OrganizationParagraphModel
			)

			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx).
				First(&organizationParagraph)
			wrongs.PanicWhenIsEmpty(ret, "站段")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var organizationParagraphs []models.OrganizationParagraphModel
			models.Init(models.OrganizationParagraphModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "shot_name", "be_enable", "organization_railway_uuid").
				PrepareQuery(ctx).
				Find(&organizationParagraphs)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_paragraphs": organizationParagraphs}))
		})
	}
}
