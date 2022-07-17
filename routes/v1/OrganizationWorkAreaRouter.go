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

// OrganizationWorkAreaRouter 工区路由
type OrganizationWorkAreaRouter struct{}

// OrganizationWorkAreaStoreForm 新建工区表单
type OrganizationWorkAreaStoreForm struct {
	Sort                         int64  `form:"sort" json:"sort"`
	UniqueCode                   string `form:"unique_code" json:"unique_code"`
	Name                         string `form:"name" json:"name"`
	BeEnable                     bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkAreaTypeUUID string `form:"organization_work_area_type_uuid" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType     models.OrganizationWorkAreaTypeModel
	OrganizationWorkshopUUID     string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop         models.OrganizationWorkshopModel
	OrganizationSections         []models.OrganizationSectionModel
	OrganizationStations         []models.OrganizationStationModel
}

// ShouldBind 绑定表单
//  @receiver cl
//  @param ctx
//  @return OrganizationWorkAreaStoreForm
func (cls OrganizationWorkAreaStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("工区代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("工区名称必填")
	}
	if cls.OrganizationWorkAreaTypeUUID == "" {
		abnormals.PanicValidate("工区类型必选")
	}

	return cls
}

func (cls OrganizationWorkAreaRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("workArea", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkAreaModel
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区代码")
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区名称")

			// 新建
			organizationWorkArea := &models.OrganizationWorkAreaModel{
				BaseModel:                models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:               form.UniqueCode,
				Name:                     form.Name,
				BeEnable:                 form.BeEnable,
				OrganizationWorkAreaType: form.OrganizationWorkAreaType,
				OrganizationWorkshop:     form.OrganizationWorkshop,
				OrganizationSections:     form.OrganizationSections,
				OrganizationStations:     form.OrganizationStations,
			}
			if ret = models.Init(models.OrganizationWorkAreaModel{}).DB().Create(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 删除
		r.DELETE("workArea/:uuid", func(ctx *gin.Context) {
			// 查询
			organizationWorkArea := (&models.OrganizationWorkAreaModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.OrganizationWorkAreaModel{}).DB().Delete(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("workArea/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkAreaStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkAreaModel
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区代码")
			ret = models.Init(models.OrganizationWorkAreaModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "工区名称")

			// 查询
			organizationWorkArea := (&models.OrganizationWorkAreaModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationWorkArea.BaseModel.Sort = form.Sort
			organizationWorkArea.UniqueCode = form.UniqueCode
			organizationWorkArea.Name = form.Name
			organizationWorkArea.BeEnable = form.BeEnable
			organizationWorkArea.OrganizationWorkAreaType = form.OrganizationWorkAreaType
			organizationWorkArea.OrganizationWorkshop = form.OrganizationWorkshop
			organizationWorkArea.OrganizationSections = form.OrganizationSections
			organizationWorkArea.OrganizationStations = form.OrganizationStations
			if ret = models.Init(models.OrganizationWorkAreaModel{}).DB().Save(&organizationWorkArea); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 详情
		r.GET("workArea/:uuid", func(ctx *gin.Context) {
			organizationWorkArea := (&models.OrganizationWorkAreaModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_area": organizationWorkArea}))
		})

		// 列表
		r.GET("workArea", func(ctx *gin.Context) {
			var organizationWorkAreas []models.OrganizationWorkAreaModel
			models.Init(models.OrganizationWorkAreaModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&organizationWorkAreas)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_areas": organizationWorkAreas}))
		})
	}
}
