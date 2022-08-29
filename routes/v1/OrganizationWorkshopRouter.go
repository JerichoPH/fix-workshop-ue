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

// OrganizationWorkshopRouter 车间路由
type OrganizationWorkshopRouter struct{}

// OrganizationWorkshopStoreForm 新建车间表单
type OrganizationWorkshopStoreForm struct {
	Sort                         int64  `form:"sort" json:"sort"`
	UniqueCode                   string `form:"unique_code" json:"unique_code"`
	Name                         string `form:"name" json:"name"`
	BeEnable                     bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopTypeUUID string `form:"organization_workshop_type_uuid" json:"organization_workshop_type_uuid"`
	OrganizationWorkshopType     models.OrganizationWorkshopTypeModel
	OrganizationParagraphUUID    string `form:"organization_paragraph_uuid" json:"organization_paragraph_uuid"`
	OrganizationParagraph        models.OrganizationParagraphModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkshopStoreForm
func (cls OrganizationWorkshopStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("车间代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("车间名称必填")
	}
	if cls.OrganizationWorkshopTypeUUID == "" {
		wrongs.PanicValidate("车间类型必选")
	}
	cls.OrganizationWorkshopType = (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(cls.OrganizationWorkshopTypeUUID)
	if cls.OrganizationParagraphUUID == "" {
		wrongs.PanicValidate("所属站段必选")
	}
	ret = models.BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationParagraphUUID}).
		PrepareByDefault().
		First(&cls.OrganizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "站段")

	return cls
}

// Load 加载路由
//  @receiver OrganizationWorkshopRouter
//  @param router
func (OrganizationWorkshopRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkshop",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationWorkshopModel
			)

			// 表单
			form := (&OrganizationWorkshopStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间代码")
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间名称")

			// 新建
			organizationWorkshop := &models.OrganizationWorkshopModel{
				BaseModel:                models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:               form.UniqueCode,
				Name:                     form.Name,
				OrganizationWorkshopType: form.OrganizationWorkshopType,
				OrganizationParagraph:    form.OrganizationParagraph,
			}
			if ret = models.BootByModel(models.OrganizationWorkshopModel{}).PrepareByDefault().Create(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			// 查询
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			// 删除
			if ret = models.BootByModel(models.OrganizationWorkshopModel{}).PrepareByDefault().Delete(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			// 表单
			form := (&OrganizationWorkshopStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkshopModel
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间代码")
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "车间名称")

			// 查询
			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				PrepareByDefault().
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			// 编辑
			organizationWorkshop.UniqueCode = form.UniqueCode
			organizationWorkshop.Name = form.Name
			organizationWorkshop.BeEnable = form.BeEnable
			organizationWorkshop.OrganizationWorkshopType = form.OrganizationWorkshopType
			organizationWorkshop.OrganizationParagraph = form.OrganizationParagraph
			if ret = models.BootByModel(models.OrganizationWorkshopModel{}).PrepareByDefault().Save(&organizationWorkshop); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				organizationWorkshop models.OrganizationWorkshopModel
			)

			ret = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareUseQueryByDefault(ctx).
				First(&organizationWorkshop)
			wrongs.PanicWhenIsEmpty(ret, "车间")

			ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_workshop": organizationWorkshop}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var (
				organizationWorkshops []models.OrganizationWorkshopModel
				dbSession             *gorm.DB
			)

			dbSession = models.BootByModel(models.OrganizationWorkshopModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_type_uuid", "organization_paragraph_uuid").
				PrepareQuery(ctx, "")

			organizationWorkshopTypeUniqueCodes, exists := ctx.GetQueryArray("organization_workshop_type_unique_codes")
			if exists {
				dbSession = dbSession.Joins("join organization_workshop_types owt on organization_workshops.organization_workshop_type_uuid = owt.uuid").
					Where("owt.unique_code in ?", organizationWorkshopTypeUniqueCodes)
			}

			dbSession.Find(&organizationWorkshops)

			ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"organization_workshops": organizationWorkshops}))
		})
	}
}
