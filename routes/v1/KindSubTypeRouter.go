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

// KindSubTypeRouter 型号路由
type KindSubTypeRouter struct{}

// KindSubTypeStoreForm 新建型号表单
type KindSubTypeStoreForm struct {
	Sort               int64  `form:"sort" json:"sort"`
	UniqueCode         string `form:"unique_code" json:"unique_code"`
	Name               string `form:"name" json:"name"`
	Nickname           string `form:"nickname" json:"nickname"`
	BeEnable           bool   `form:"be_enable" json:"be_enable"`
	KindEntireTypeUUID string `form:"kind_entire_type_uuid" json:"kind_entire_type_uuid"`
	KindEntireType     models.KindEntireTypeModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return KindSubTypeStoreForm
func (cls KindSubTypeStoreForm) ShouldBind(ctx *gin.Context) KindSubTypeStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("型号代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("型号名称必填")
	}
	if cls.KindEntireTypeUUID == "" {
		wrongs.PanicValidate("所属类型必选")
	}
	ret = models.Init(models.KindEntireTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.KindEntireTypeUUID}).
		Prepare().
		First(&cls.KindEntireType)
	wrongs.PanicWhenIsEmpty(ret, "所属类型")

	return cls
}

// Load 加载路由
//  @receiver KindSubTypeRouter
//  @param engine
func (KindSubTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/kind",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("subType", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.KindSubTypeModel
			)

			// 表单
			form := (&KindSubTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "型号代码")
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "型号名称")

			// 新建
			kindSubType := &models.KindSubTypeModel{
				BaseModel:      models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:     form.UniqueCode,
				Name:           form.Name,
				BeEnable:       form.BeEnable,
				Nickname:       form.Nickname,
				KindEntireType: form.KindEntireType,
			}
			if ret = models.Init(models.KindSubTypeModel{}).GetSession().Create(&kindSubType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"kind_sub_type": kindSubType}))
		})

		// 删除
		r.DELETE("subType/:uuid", func(ctx *gin.Context) {
			var (
				ret         *gorm.DB
				kindSubType models.KindSubTypeModel
			)

			// 查询
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindSubType)
			wrongs.PanicWhenIsEmpty(ret, "型号")

			// 删除
			if ret := models.Init(models.KindSubTypeModel{}).GetSession().Delete(&kindSubType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("subType/:uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				kindSubType, repeat models.KindSubTypeModel
			)

			// 表单
			form := (&KindSubTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "型号代码")
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "型号名称")

			// 查询
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindSubType)
			wrongs.PanicWhenIsEmpty(ret, "型号")

			// 编辑
			kindSubType.BaseModel.Sort = form.Sort
			kindSubType.UniqueCode = form.UniqueCode
			kindSubType.Name = form.Name
			kindSubType.BeEnable = form.BeEnable
			kindSubType.Nickname = form.Nickname
			kindSubType.KindEntireType = form.KindEntireType
			if ret = models.Init(models.KindSubTypeModel{}).GetSession().Save(&kindSubType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"kind_sub_type": kindSubType}))
		})

		// 详情
		r.GET("subType/:uuid", func(ctx *gin.Context) {
			var (
				ret         *gorm.DB
				kindSubType models.KindSubTypeModel
			)
			ret = models.Init(models.KindSubTypeModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&kindSubType)
			wrongs.PanicWhenIsEmpty(ret, "型号")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_sub_type": kindSubType}))
		})

		// 列表
		r.GET("subType", func(ctx *gin.Context) {
			var kindSubTypes []models.KindSubTypeModel
			models.Init(models.KindSubTypeModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&kindSubTypes)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"kind_sub_types": kindSubTypes}))
		})
	}
}
