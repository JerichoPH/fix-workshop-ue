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
		Prepare("").
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.Init(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare("").
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	} else {

	}
	if len(cls.LocationLineUUIDs) > 0 {
		models.Init(models.LocationLineModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
	}

	return cls
}

// LocationStationBindLocationLinesForm 站场绑定线别表单
type LocationStationBindLocationLinesForm struct {
	LocationLineUUIDs []string `json:"location_line_uuids"`
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationStationBindLocationLinesForm
func (cls LocationStationBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationStationBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if len(cls.LocationLineUUIDs) > 0 {
		models.Init(models.LocationLineModel{}).
			Prepare("").
			Where("uuid in ?", cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
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
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare("").
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
			if ret = models.Init(models.LocationStationModel{}).Prepare("").Create(&locationStation); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_station": locationStation}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret             *gorm.DB
				locationStation models.LocationStationModel
			)
			// 查询
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			// 删除
			if ret := models.Init(models.LocationStationModel{}).Prepare("").Delete(&locationStation); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				locationStation, repeat models.LocationStationModel
			)

			// 表单
			form := (&LocationStationStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "站场名称")

			// 查询
			ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			// 编辑
			locationStation.BaseModel.Sort = form.Sort
			locationStation.UniqueCode = form.UniqueCode
			locationStation.Name = form.Name
			locationStation.BeEnable = form.BeEnable
			locationStation.OrganizationWorkshop = form.OrganizationWorkshop
			locationStation.OrganizationWorkAreaUUID = form.OrganizationWorkAreaUUID
			locationStation.LocationLines = form.LocationLines
			if ret = models.
				Init(models.LocationStationModel{}).
				Prepare("").
				Where("uuid = ?", ctx.Param("uuid")).
				Updates(map[string]interface{}{
					"sort":                        form.Sort,
					"unique_code":                 form.UniqueCode,
					"name":                        form.Name,
					"be_enable":                   form.BeEnable,
					"organization_workshop_uuid":  form.OrganizationWorkshopUUID,
					"organization_work_area_uuid": form.OrganizationWorkAreaUUID,
				}); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_station": locationStation}))
		})

		// 站场绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) {
			var (
				ret                                  *gorm.DB
				locationStation                      models.LocationStationModel
				pivotLocationLineAndLocationStations []models.PivotLocationLineAndLocationStation
			)

			// 表单
			form := (&LocationStationBindLocationLinesForm{}).ShouldBind(ctx)

			if ret = models.Init(models.LocationStationModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationStation); ret.Error != nil {
				wrongs.PanicWhenIsEmpty(ret, "站场")
			}

			// 删除原有绑定关系
			ret = models.Init(models.BaseModel{}).Prepare("").Exec("delete from pivot_location_line_and_location_stations where location_station_id = ?", locationStation.ID)

			// 创建绑定关系
			if len(form.LocationLines) > 0 {
				for _, locationLine := range form.LocationLines {
					pivotLocationLineAndLocationStations = append(pivotLocationLineAndLocationStations, models.PivotLocationLineAndLocationStation{
						LocationLineID:    locationLine.ID,
						LocationStationID: locationStation.ID,
					})
				}
				models.Init(models.PivotLocationLineAndLocationStation{}).
					Prepare("").
					CreateInBatches(&pivotLocationLineAndLocationStations, 100)
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
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
				SetWhereFields("be_enable").
				PrepareQuery(ctx,"").
				First(&locationStation)
			wrongs.PanicWhenIsEmpty(ret, "站场")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_station": locationStation}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var locationStations []models.LocationStationModel
			models.Init(models.LocationStationModel{}).
				SetWhereFields().
				PrepareQuery(ctx,"").
				Find(&locationStations)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_stations": locationStations}))
		})
	}
}
