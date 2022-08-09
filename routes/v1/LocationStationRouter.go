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

// LocationStationRouter 站场路由
type LocationStationRouter struct{}

// LocationStationStoreForm 新建站场表单
type LocationStationStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
	LocationLineUUIDs        []string `form:"location_line_uuids" json:"location_line_uuids"`
	LocationLines            []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationStationStoreForm
func (cls LocationStationStoreForm) ShouldBind(ctx *gin.Context) LocationStationStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("站场代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("站场名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.Init(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationStationRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationStation",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationStationModel
			)

			// 表单
			form := (&LocationStationStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场名称")

			// 新建
			locationStation := &models.LocationStationModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				BeEnable:             form.BeEnable,
				OrganizationWorkshop: form.OrganizationWorkshop,
				OrganizationWorkArea: form.OrganizationWorkArea,
				LocationLines:        form.LocationLines,
			}
			if ret = models.Init(models.LocationStationModel{}).GetSession().Create(&locationStation); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_station": locationStation}))
		})

		// 删除
		r.DELETE("station/:uuid", func(ctx *gin.Context) {
			var (
				ret             *gorm.DB
				locationStation models.LocationStationModel
			)
			// 查询
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			// 删除
			if ret := models.Init(models.LocationStationModel{}).GetSession().Delete(&locationStation); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                         *gorm.DB
				organizationStation, repeat models.LocationStationModel
			)

			// 表单
			form := (&LocationStationStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场名称")

			// 查询
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&organizationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			// 编辑
			organizationStation.BaseModel.Sort = form.Sort
			organizationStation.UniqueCode = form.UniqueCode
			organizationStation.Name = form.Name
			organizationStation.BeEnable = form.BeEnable
			organizationStation.OrganizationWorkshop = form.OrganizationWorkshop
			organizationStation.OrganizationWorkArea = form.OrganizationWorkArea
			organizationStation.LocationLines = form.LocationLines
			if ret = models.Init(models.LocationStationModel{}).GetSession().Save(&organizationStation); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_station": organizationStation}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret             *gorm.DB
				locationStation models.LocationStationModel
			)
			// 查询
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				Prepare().
				First(&locationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_station": locationStation}))
		})

		// 列表
		r.GET("station", func(ctx *gin.Context) {
			var locationStations []models.LocationStationModel
			models.Init(models.LocationStationModel{}).
				SetWhereFields().
				SetScopes((&models.BaseModel{}).ScopeBeEnable).
				PrepareQuery(ctx).
				Find(&locationStations)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_stations": locationStations}))
		})
	}
}
