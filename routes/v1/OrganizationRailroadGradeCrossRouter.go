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

// OrganizationRailroadGradeCrossRouter 道口路由
type OrganizationRailroadGradeCrossRouter struct{}

// OrganizationRailroadGradeCrossStoreForm 新建道口表单
type OrganizationRailroadGradeCrossStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"" json:""`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return OrganizationCenterStoreForm
func (cls OrganizationRailroadGradeCrossStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailroadGradeCrossStoreForm {
	if err := ctx.ShouldBind(ctx); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("道口代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("道口名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		abnormals.PanicValidate("所属车间必选")
	}
	cls.OrganizationWorkshop = (&models.OrganizationWorkshopModel{}).FindOneByUUID(cls.OrganizationWorkshopUUID)
	if cls.OrganizationWorkAreaUUID != "" {
		cls.OrganizationWorkArea = (&models.OrganizationWorkAreaModel{}).FindOneByUUID(cls.OrganizationWorkAreaUUID)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls OrganizationRailroadGradeCrossRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("railroadGradeCross", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationRailroadGradeCrossModel
			ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口名称")

			// 新建
			organizationRailroadGradeCross := &models.OrganizationRailroadGradeCrossModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).DB().Create(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 删除
		r.DELETE("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			// 查询
			organizationRailroadGradeCross := (&models.OrganizationRailroadGradeCrossModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.OrganizationRailroadGradeCrossModel{}).DB().Delete(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationRailroadGradeCrossModel
			ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口名称")

			// 查询
			organizationRailroadGradeCross := (&models.OrganizationRailroadGradeCrossModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationRailroadGradeCross.BaseModel.Sort = form.Sort
			organizationRailroadGradeCross.UniqueCode = form.UniqueCode
			organizationRailroadGradeCross.Name = form.Name
			organizationRailroadGradeCross.BeEnable = form.BeEnable
			if ret = models.Init(models.OrganizationRailroadGradeCrossModel{}).DB().Save(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 详情
		r.GET("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			organizationRailroadGradeCross := (&models.OrganizationRailroadGradeCrossModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 列表
		r.GET("railroadGradeCross", func(ctx *gin.Context) {
			var organizationRailroadGradeCrosses []models.OrganizationRailroadGradeCrossModel
			models.Init(models.OrganizationRailroadGradeCrossModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&organizationRailroadGradeCrosses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationRailroadGradeCrosses": organizationRailroadGradeCrosses}))
		})
	}
}
