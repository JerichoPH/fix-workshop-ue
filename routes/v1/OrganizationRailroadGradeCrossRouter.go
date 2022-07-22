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

// OrganizationRailroadGradeCrossRouter 道口路由
type OrganizationRailroadGradeCrossRouter struct{}

// OrganizationRailroadGradeCrossStoreForm 新建道口表单
type OrganizationRailroadGradeCrossStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"" json:""`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationModels.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationModels.OrganizationWorkAreaModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return OrganizationCenterStoreForm
func (cls OrganizationRailroadGradeCrossStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailroadGradeCrossStoreForm {
	var ret *gorm.DB

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
	ret = models.Init(OrganizationModels.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	abnormals.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.Init(OrganizationModels.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare().
			First(&cls.OrganizationWorkArea)
		abnormals.PanicWhenIsEmpty(ret, "工区")
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
			var (
				ret    *gorm.DB
				repeat OrganizationModels.OrganizationRailroadGradeCrossModel
			)

			// 表单
			form := (&OrganizationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口名称")

			// 新建
			organizationRailroadGradeCross := &OrganizationModels.OrganizationRailroadGradeCrossModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).DB().Create(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 删除
		r.DELETE("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			var (
				ret                            *gorm.DB
				organizationRailroadGradeCross OrganizationModels.OrganizationRailroadGradeCrossModel
			)

			// 查询
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailroadGradeCross)
			abnormals.PanicWhenIsEmpty(ret, "道口")

			// 删除
			if ret := models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).DB().Delete(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			var (
				ret                                    *gorm.DB
				organizationRailroadGradeCross, repeat OrganizationModels.OrganizationRailroadGradeCrossModel
			)

			// 表单
			form := (&OrganizationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "道口名称")

			// 查询
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailroadGradeCross)
			abnormals.PanicWhenIsEmpty(ret, "道口")

			// 编辑
			organizationRailroadGradeCross.BaseModel.Sort = form.Sort
			organizationRailroadGradeCross.UniqueCode = form.UniqueCode
			organizationRailroadGradeCross.Name = form.Name
			organizationRailroadGradeCross.BeEnable = form.BeEnable
			if ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).DB().Save(&organizationRailroadGradeCross); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 详情
		r.GET("railroadGradeCross/:uuid", func(ctx *gin.Context) {
			var (
				ret                            *gorm.DB
				organizationRailroadGradeCross OrganizationModels.OrganizationRailroadGradeCrossModel
			)
			ret = models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&organizationRailroadGradeCross)
			abnormals.PanicWhenIsEmpty(ret, "道口")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationRailroadGradeCross": organizationRailroadGradeCross}))
		})

		// 列表
		r.GET("railroadGradeCross", func(ctx *gin.Context) {
			var organizationRailroadGradeCrosses []OrganizationModels.OrganizationRailroadGradeCrossModel
			models.Init(OrganizationModels.OrganizationRailroadGradeCrossModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&organizationRailroadGradeCrosses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organizationRailroadGradeCrosses": organizationRailroadGradeCrosses}))
		})
	}
}
