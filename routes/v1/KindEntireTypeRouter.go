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

// KindEntireTypeRouter 类型路由
type KindEntireTypeRouter struct{}

// KindEntireTypeStoreForm 新建类型表单
type KindEntireTypeStoreForm struct {
	Sort             int64  `form:"sort" json:"sort"`
	UniqueCode       string `form:"unique_code" json:"unique_code"`
	Name             string `form:"name" json:"name"`
	Nickname         string `form:"nickname" json:"nickname"`
	BeEnable         bool   `form:"be_enable" json:"be_enable"`
	KindCategoryUUID string `form:"kind_category_uuid" json:"kind_category_uuid"`
	KindCategory     models.KindCategoryModel
	KindSubTypeUUIDs []string `form:"kind_sub_type_uuids" json:"kind_sub_type_uuids"`
	KindSubTypes     []models.KindSubTypeModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return KindEntireTypeStoreForm
func (cls KindEntireTypeStoreForm) ShouldBind(ctx *gin.Context) KindEntireTypeStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("类型代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("类型名称必填")
	}
	if cls.KindCategoryUUID == "" {
		wrongs.PanicValidate("所属种类必选")
	}
	ret = models.Init(models.KindCategoryModel{}).
		SetWheres(tools.Map{"uuid": cls.KindCategoryUUID}).
		Prepare().
		First(&cls.KindCategory)
	wrongs.PanicWhenIsEmpty(ret, "所属种类")
	if len(cls.KindSubTypeUUIDs) > 0 {
		models.Init(models.KindSubTypeModel{}).
			GetSession().
			Where("uuid in ?", cls.KindSubTypeUUIDs).
			Find(&cls.KindSubTypes)
	}

	return cls
}

// Load 加载路由
//  @receiver KindEntireTypeRouter
//  @param engine
func (KindEntireTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kind",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("entireType", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.KindEntireTypeModel
			)

			// 表单
			form := (&KindEntireTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "类型代码")
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "类型名称")

			// 新建
			kindEntireType := &models.KindEntireTypeModel{
				BaseModel:    models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:   form.UniqueCode,
				Name:         form.Name,
				BeEnable:     form.BeEnable,
				KindCategory: form.KindCategory,
				KindSubTypes: form.KindSubTypes,
			}
			if ret = models.Init(models.KindEntireTypeModel{}).GetSession().Create(&kindEntireType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"kind_entire_type": kindEntireType}))
		})

		// 删除
		r.DELETE("entireType/:uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				kindEntireType models.KindEntireTypeModel
			)

			// 查询
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindEntireType)
			wrongs.PanicWhenIsEmpty(ret, "类型")

			// 删除
			if ret := models.Init(models.KindEntireTypeModel{}).GetSession().Delete(&kindEntireType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("entireType/:uuid", func(ctx *gin.Context) {
			var (
				ret                    *gorm.DB
				kindEntireType, repeat models.KindEntireTypeModel
			)

			// 表单
			form := (&KindEntireTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "类型代码")
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "类型名称")

			// 查询
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindEntireType)
			wrongs.PanicWhenIsEmpty(ret, "类型")

			// 编辑
			kindEntireType.BaseModel.Sort = form.Sort
			kindEntireType.UniqueCode = form.UniqueCode
			kindEntireType.Name = form.Name
			kindEntireType.BeEnable = form.BeEnable
			kindEntireType.KindCategory = form.KindCategory
			kindEntireType.KindSubTypes = form.KindSubTypes
			if ret = models.Init(models.KindEntireTypeModel{}).GetSession().Save(&kindEntireType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"kind_entire_type": kindEntireType}))
		})

		// 详情
		r.GET("entireType/:uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				kindEntireType models.KindEntireTypeModel
			)
			ret = models.Init(models.KindEntireTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindEntireType)
			wrongs.PanicWhenIsEmpty(ret, "类型")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_entire_type": kindEntireType}))
		})

		// 列表
		r.GET("entireType", func(ctx *gin.Context) {
			var kindEntireTypes []models.KindEntireTypeModel
			models.Init(models.KindEntireTypeModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&kindEntireTypes)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_entire_types": kindEntireTypes}))
		})
	}
}
