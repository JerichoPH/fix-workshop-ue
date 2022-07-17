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

// OrganizationSectionRouter 区间路由
type OrganizationSectionRouter struct{}

// OrganizationSectionStoreForm 新建区间表单
type OrganizationSectionStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationSectionStoreForm
func (cls OrganizationSectionStoreForm) ShouldBind(ctx *gin.Context) OrganizationSectionStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.BombForbidden(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.BombForbidden("区间代码不能为空")
	}
	if cls.Name == "" {
		abnormals.BombForbidden("区间名称不能为空")
	}
	if cls.OrganizationWorkshopUUID == "" {
		abnormals.BombForbidden("所属车间不能为空")
	}
	cls.OrganizationWorkshop = (&models.OrganizationWorkshopModel{}).FindOneByUUID(cls.OrganizationWorkshopUUID)

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls OrganizationSectionRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("section", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationSectionModel
			ret = models.Init(models.OrganizationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.BombWhenIsRepeat(ret, "区间代码")
			ret = models.Init(models.OrganizationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.BombWhenIsRepeat(ret, "区间名称")

			// 新建
			organizationSection := &models.OrganizationSectionModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				BeEnable:             form.BeEnable,
				OrganizationWorkshop: form.OrganizationWorkshop,
			}
			if ret = models.Init(models.OrganizationSectionModel{}).DB().Create(&organizationSection); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_section": organizationSection}))
		})

		// 删除
		r.DELETE("section/:uuid", func(ctx *gin.Context) {
			// 查询
			organizationSection := (&models.OrganizationSectionModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.OrganizationSectionModel{}).DB().Delete(&organizationSection); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("section/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationSectionModel
			ret = models.Init(models.OrganizationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.BombWhenIsRepeat(ret, "区间代码")
			ret = models.Init(models.OrganizationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.BombWhenIsRepeat(ret, "区间名称")

			// 查询
			organizationSection := (&models.OrganizationSectionModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationSection.BaseModel.Sort = form.Sort
			organizationSection.UniqueCode = form.UniqueCode
			organizationSection.Name = form.Name
			organizationSection.BeEnable = form.BeEnable
			organizationSection.OrganizationWorkshop = form.OrganizationWorkshop
			if ret = models.Init(models.OrganizationSectionModel{}).DB().Save(&organizationSection); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_section": organizationSection}))
		})

		// 详情
		r.GET("section/:uuid", func(ctx *gin.Context) {
			organizationSection := (&models.OrganizationSectionModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_section": organizationSection}))
		})

		// 列表
		r.GET("section", func(ctx *gin.Context) {
			var organizationSections []models.OrganizationSectionModel
			models.Init(models.OrganizationSectionModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&organizationSections)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_sections": organizationSections}))
		})
	}
}
