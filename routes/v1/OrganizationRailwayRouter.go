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

// OrganizationRailwayRouter 路局路由
type OrganizationRailwayRouter struct{}

// OrganizationRailwayStoreForm 新建路局表单
type OrganizationRailwayStoreForm struct {
	Sort              int64    `form:"sort" json:"sort"`
	UniqueCode        string   `form:"unique_code" json:"unique_code"`
	Name              string   `form:"name" json:"name"`
	ShortName         string   `form:"short_name" json:"short_name"`
	BeEnable          bool     `form:"be_enable" json:"be_enable"`
	LocationLineUUIDs []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationRailwayStoreForm
func (cls OrganizationRailwayStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailwayStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("路局代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("路局名称必填")
	}
	if len(cls.LocationLineUUIDs) > 0 {
		models.Init(models.LocationLineModel{}).
			Prepare().
			Where("uuid in ?", cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
	}

	return cls
}

// OrganizationRailwayBindLinesFrom 路局绑定线别
type OrganizationRailwayBindLinesFrom struct {
	LocationLineUUIDs []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
func (cls OrganizationRailwayBindLinesFrom) ShouldBind(ctx *gin.Context) OrganizationRailwayBindLinesFrom {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	models.Init(models.LocationLineModel{}).
		Prepare().
		Where("uuid in ?", cls.LocationLineUUIDs).
		Find(&cls.LocationLines)

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (OrganizationRailwayRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"/api/v1/organizationRailway",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建路局
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.OrganizationRailwayModel
			)

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局代码")
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局名称")
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局简称")

			// 新建
			organizationRailway := &models.OrganizationRailwayModel{
				BaseModel:     models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:    form.UniqueCode,
				Name:          form.Name,
				ShortName:     form.ShortName,
				BeEnable:      form.BeEnable,
			}
			if ret = (&models.BaseModel{}).SetModel(models.OrganizationRailwayModel{}).Prepare().Create(&organizationRailway); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_railway": organizationRailway}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 查询
			var organizationRailway models.OrganizationRailwayModel
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailway)
			wrongs.PanicWhenIsEmpty(ret, "路局")
			// 删除
			if ret = models.Init(models.OrganizationRailwayModel{}).Prepare().Delete(&organizationRailway); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                         *gorm.DB
				organizationRailway, repeat models.OrganizationRailwayModel
			)

			// 表单
			form := (&OrganizationRailwayStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局代码")
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局名称")
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "路局简称")

			// 查询
			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationRailway)
			wrongs.PanicWhenIsEmpty(ret, "路局")

			// 修改
			organizationRailway.Sort = form.Sort
			organizationRailway.UniqueCode = form.UniqueCode
			organizationRailway.Name = form.Name
			organizationRailway.ShortName = form.ShortName
			organizationRailway.BeEnable = form.BeEnable
			if ret = models.Init(models.OrganizationRailwayModel{}).Prepare().Save(&organizationRailway); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_railway": organizationRailway}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				organizationRailway models.OrganizationRailwayModel
			)

			ret = models.Init(models.OrganizationRailwayModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx).
				First(&organizationRailway)
			wrongs.PanicWhenIsEmpty(ret, "路局")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railway": organizationRailway}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var organizationRailways []models.OrganizationRailwayModel

			models.Init(models.OrganizationRailwayModel{}).
				SetWhereFields("uuid", "sort", "unique_code", "name", "short_name", "be_enable").
				PrepareQuery(ctx).
				Find(&organizationRailways)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_railways": organizationRailways}))
		})
	}
}
