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

// KindCategoryRouter 种类路由
type KindCategoryRouter struct{}

// KindCategoryStoreForm 新建种类表单
type KindCategoryStoreForm struct {
	Sort                int64    `form:"sort" json:"sort"`
	UniqueCode          string   `form:"unique_code" json:"unique_code"`
	Name                string   `form:"name" json:"name"`
	Nickname            string   `form:"nickname" json:"nickname"`
	BeEnable            bool     `form:"be_enable" json:"be_enable"`
	Race                string `form:"race" json:"race"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return KindCategoryStoreForm
func (cls KindCategoryStoreForm) ShouldBind(ctx *gin.Context) KindCategoryStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("种类代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("种类名称必填")
	}

	return cls
}

func (KindCategoryRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kindCategory",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.KindCategoryModel
			)

			// 表单
			form := (&KindCategoryStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "种类代码")
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "种类名称")

			// 新建
			kindCategory := &models.KindCategoryModel{
				BaseModel:       models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:      form.UniqueCode,
				Name:            form.Name,
				BeEnable:        form.BeEnable,
				Nickname:        form.Nickname,
				Race:            form.Race,
			}
			if ret = models.BootByModel(models.KindCategoryModel{}).Prepare("").Create(&kindCategory); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"kind_category": kindCategory}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				kindCategory models.KindCategoryModel
			)

			// 查询
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&kindCategory)
			wrongs.PanicWhenIsEmpty(ret, "种类")

			// 删除
			if ret := models.BootByModel(models.KindCategoryModel{}).Prepare("").Delete(&kindCategory); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				kindCategory, repeat models.KindCategoryModel
			)

			// 表单
			form := (&KindCategoryStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "种类代码")
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "种类名称")

			// 查询
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&kindCategory)
			wrongs.PanicWhenIsEmpty(ret, "种类")

			// 编辑
			kindCategory.BaseModel.Sort = form.Sort
			kindCategory.UniqueCode = form.UniqueCode
			kindCategory.Name = form.Name
			kindCategory.Nickname = form.Nickname
			kindCategory.BeEnable = form.BeEnable
			kindCategory.Race = form.Race
			if ret = models.BootByModel(models.KindCategoryModel{}).Prepare("").Save(&kindCategory); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"kind_category": kindCategory}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret          *gorm.DB
				kindCategory models.KindCategoryModel
			)
			ret = models.BootByModel(models.KindCategoryModel{}).
				SetScopes((&models.BaseModel{}).ScopeBeEnableTrue).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&kindCategory)
			wrongs.PanicWhenIsEmpty(ret, "种类")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_category": kindCategory}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var kindCategories []models.KindCategoryModel
			models.BootByModel(models.KindCategoryModel{}).
				SetScopes((&models.BaseModel{}).ScopeBeEnableTrue).
				SetWhereFields().
				PrepareQuery(ctx,"").
				Find(&kindCategories)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_categories": kindCategories}))
		})
	}
}
