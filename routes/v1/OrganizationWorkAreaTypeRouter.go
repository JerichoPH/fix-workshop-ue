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

// OrganizationWorkAreaTypeRouter 工区类型路由
type OrganizationWorkAreaTypeRouter struct{}

// OrganizationWorkAreaTypeStoreForm 新建工区表单
type OrganizationWorkAreaTypeStoreForm struct {
	Sort                      int64    `form:"sort" json:"sort"`
	UniqueCode                string   `form:"unique_code" json:"unique_code"`
	Name                      string   `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationWorkAreaTypeStoreForm
func (cls OrganizationWorkAreaTypeStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaTypeStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("工区代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("工区名称必填")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationWorkAreaTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/organizationWorkAreaType",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkAreaTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkAreaTypeModel
			ret = models.Init(models.OrganizationWorkAreaTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "工区类型代码")
			ret = models.Init(models.OrganizationWorkAreaTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "工区类型名称")

			// 新建
			organizationWorkAreaType := &models.OrganizationWorkAreaTypeModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
			}
			if ret = models.Init(models.OrganizationWorkAreaTypeModel{}).GetSession().Create(&organizationWorkAreaType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			// 查询
			organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.OrganizationWorkAreaTypeModel{}).GetSession().Delete(&organizationWorkAreaType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationWorkAreaTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationWorkAreaTypeModel
			ret = models.Init(models.OrganizationWorkAreaTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "工区类型代码")
			ret = models.Init(models.OrganizationWorkAreaTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "工区类型名称")

			// 查询
			organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationWorkAreaType.BaseModel.Sort = form.Sort
			organizationWorkAreaType.UniqueCode = form.UniqueCode
			organizationWorkAreaType.Name = form.Name
			if ret = models.Init(models.OrganizationWorkAreaTypeModel{}).GetSession().Save(&organizationWorkAreaType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			organizationWorkAreaType := (&models.OrganizationWorkAreaTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_area_type": organizationWorkAreaType}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var organizationWorkAreaType []models.OrganizationWorkAreaTypeModel
			models.Init(models.OrganizationWorkAreaTypeModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&organizationWorkAreaType)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_work_area_types": organizationWorkAreaType}))
		})
	}
}
