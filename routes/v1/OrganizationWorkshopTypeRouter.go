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

// OrganizationWorkshopTypeRouter 车间类型路由
type OrganizationWorkshopTypeRouter struct{}

// OrganizationWorkshopTypeStoreForm 新建车间路由表单
type OrganizationWorkshopTypeStoreForm struct {
	Sort       int64  `gorm:"sort" json:"sort"`
	UniqueCode string `gorm:"unique_code" json:"unique_code"`
	Name       string `gorm:"name" json:"name"`
	Number     string `gorm:"number" json:"number"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkshopTypeStoreForm
func (cls OrganizationWorkshopTypeStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopTypeStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.BombForbidden(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.BombForbidden("车间类型代码必填")
	}
	if cls.Name == "" {
		abnormals.BombForbidden("车间类型名称必填")
	}

	return cls
}

// Load 加载路由
//  @receiver OrganizationWorkshopTypeRouter
//  @param router
func (OrganizationWorkshopTypeRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建车间类型
		r.POST("workshopType", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkshopTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationWorkshopTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&models.OrganizationWorkAreaTypeModel{})
			abnormals.BombWhenIsRepeat(ret, "车间类型代码")
			ret = models.Init(models.OrganizationWorkshopTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&models.OrganizationWorkshopTypeModel{})
			abnormals.BombWhenIsRepeat(ret, "车间类型名称")

			// 新建
			organizationWorkshopType := &models.OrganizationWorkshopTypeModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				Number:     form.Number,
			}
			if ret = models.Init(models.OrganizationWorkshopTypeModel{}).DB().Create(&organizationWorkshopType); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_workshop_type": organizationWorkshopType}))
		})

		// 删除
		r.DELETE("workshopType/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 查询
			organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret = models.Init(models.OrganizationWorkshopTypeModel{}).DB().Delete(&organizationWorkshopType); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("workshopType/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkshopTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationWorkshopTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&models.OrganizationWorkAreaTypeModel{})
			abnormals.BombWhenIsRepeat(ret, "车间类型代码")
			ret = models.Init(models.OrganizationWorkshopTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&models.OrganizationWorkshopTypeModel{})
			abnormals.BombWhenIsRepeat(ret, "车间类型名称")

			// 查询
			organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationWorkshopType.BaseModel.Sort = form.Sort
			organizationWorkshopType.UniqueCode = form.UniqueCode
			organizationWorkshopType.Name = form.Name
			organizationWorkshopType.Number = form.Number
			models.Init(models.OrganizationWorkshopTypeModel{}).DB().Save(&organizationWorkshopType)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_workshop_type": organizationWorkshopType}))
		})

		// 详情
		r.GET("workshopType/:uuid", func(ctx *gin.Context) {
			organizationWorkshopType := (&models.OrganizationWorkshopTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshop_type": organizationWorkshopType}))
		})

		// 列表
		r.GET("workshopType", func(ctx *gin.Context) {
			var organizationWorkshopTypes []models.OrganizationWorkshopTypeModel
			models.Init(models.OrganizationWorkshopTypeModel{}).
				SetWhereFields("sort", "unique_code", "name", "number").
				PrepareQuery(ctx).
				Find(&organizationWorkshopTypes)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_workshop_types": organizationWorkshopTypes}))
		})
	}
}
