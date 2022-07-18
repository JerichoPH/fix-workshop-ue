package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// OrganizationParagraphRouter 站段路由
type OrganizationParagraphRouter struct{}

// OrganizationParagraphStoreForm 新建站段表单
type OrganizationParagraphStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	ShortName                 string `form:"short_name" json:"short_name"`
	BeEnable                  bool   `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUID   string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway       models.OrganizationRailwayModel
	OrganizationWorkshopUUIDs []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops     []models.OrganizationWorkshopModel
	OrganizationLineUUIDs     []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines         []*models.OrganizationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationParagraphStoreForm
func (cls OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("站段代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("站段名称必填")
	}
	if cls.OrganizationRailwayUUID == "" {
		abnormals.PanicValidate("所属路局必选")
	}
	cls.OrganizationRailway = (&models.OrganizationRailwayModel{}).FindOneByUUID(cls.OrganizationRailwayUUID)
	if len(cls.OrganizationWorkshopUUIDs) > 0 {
		models.Init(models.OrganizationWorkshopModel{}).DB().Where("uuid in ?", cls.OrganizationWorkshopUUIDs).Find(&cls.OrganizationWorkshops)
	}
	if len(cls.OrganizationLineUUIDs) > 0 {
		models.Init(models.OrganizationLineModel{}).DB().Where("uuid in ?", cls.OrganizationLineUUIDs).Find(&cls.OrganizationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *OrganizationParagraphRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("paragraph", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationParagraphModel
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站段代码")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站段名称")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站段简称")

			// 新建
			organizationParagraph := &models.OrganizationParagraphModel{
				BaseModel:           models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:          form.UniqueCode,
				Name:                form.Name,
				ShortName:           form.ShortName,
				BeEnable:            form.BeEnable,
				OrganizationRailway: form.OrganizationRailway,
				OrganizationWorkshops: form.OrganizationWorkshops,
				OrganizationLines:     form.OrganizationLines,
			}
			if ret = models.Init(models.OrganizationParagraphModel{}).DB().Create(&organizationParagraph); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 删除
		r.DELETE("paragraph/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 查询
			organizationParagraph := (&models.OrganizationParagraphModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret = models.Init(models.OrganizationParagraphModel{}).DB().Delete(&organizationParagraph); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("paragraph/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationParagraphModel
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站段代码")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站段名称")

			// 查询
			organizationParagraph := (&models.OrganizationParagraphModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationParagraph.BaseModel.Sort = form.Sort
			organizationParagraph.UniqueCode = form.UniqueCode
			organizationParagraph.Name = form.Name
			organizationParagraph.ShortName = form.ShortName
			organizationParagraph.BeEnable = form.BeEnable
			organizationParagraph.OrganizationRailway = form.OrganizationRailway
			organizationParagraph.OrganizationWorkshops = form.OrganizationWorkshops
			organizationParagraph.OrganizationLines = form.OrganizationLines
			if ret = models.Init(models.OrganizationParagraphModel{}).DB().Save(&organizationParagraph); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 详情
		r.GET("paragraph/:uuid", func(ctx *gin.Context) {
			organizationParagraph := (&models.OrganizationParagraphModel{}).FindOneByUUID(ctx.Param("uuid"))
			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_paragraph": organizationParagraph}))
		})

		// 列表
		r.GET("paragraph", func(ctx *gin.Context) {
			var organizationParagraphs []models.OrganizationParagraphModel
			models.Init(models.OrganizationParagraphModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "shot_name", "be_enable", "organization_railway_uuid").
				PrepareQuery(ctx).
				Find(&organizationParagraphs)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_paragraphs": organizationParagraphs}))
		})
	}
}
