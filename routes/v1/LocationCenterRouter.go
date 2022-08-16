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

// LocationCenterRouter 中心路由
type LocationCenterRouter struct{}

// LocationCenterStoreForm 新建中心表单
type LocationCenterStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"" json:""`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return LocationCenterStoreForm
func (cls LocationCenterStoreForm) ShouldBind(ctx *gin.Context) LocationCenterStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(ctx); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("中心代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("中心名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")
	if cls.OrganizationWorkAreaUUID != "" {
		models.Init(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare().
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return cls
}

// LocationCenterBindLocationLinesForm 中心绑定线别表单
type LocationCenterBindLocationLinesForm struct {
	LocationLineUUIDs []string
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationCenterBindLocationLinesForm
func (cls LocationCenterBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationCenterBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if len(cls.LocationLineUUIDs) > 0 {
		models.Init(models.LocationLineModel{}).
			Prepare().
			Where("uuid in ?", cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationCenterRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationCenter",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationCenterModel
			)

			// 表单
			form := (&LocationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 新建
			locationCenter := &models.LocationCenterModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(models.LocationCenterModel{}).Prepare().Create(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_center": locationCenter}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				locationCenter models.LocationCenterModel
			)

			// 查询
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 删除
			if ret := models.Init(models.LocationCenterModel{}).Prepare().Delete(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                    *gorm.DB
				locationCenter, repeat models.LocationCenterModel
			)

			// 表单
			form := (&LocationCenterStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心代码")
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "中心名称")

			// 查询
			ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			// 编辑
			locationCenter.BaseModel.Sort = form.Sort
			locationCenter.UniqueCode = form.UniqueCode
			locationCenter.Name = form.Name
			locationCenter.BeEnable = form.BeEnable
			if ret = models.Init(models.LocationCenterModel{}).Prepare().Save(&locationCenter); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_center": locationCenter}))
		})

		// 中心绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) {
			var (
				ret                                 *gorm.DB
				locationCenter                      models.LocationCenterModel
				pivotLocationLineAndLocationCenters []models.PivotLocationLineAndLocationCenter
			)

			// 表单
			form := (&LocationCenterBindLocationLinesForm{}).ShouldBind(ctx)

			if ret = models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationCenter); ret.Error != nil {
				wrongs.PanicWhenIsEmpty(ret, "中心")
			}

			// 删除原有绑定关系
			ret = models.Init(models.BaseModel{}).Prepare().Exec("delete from pivot_location_line_and_location_centers where location_center_id = ?", locationCenter.ID)

			// 创建绑定关系
			if len(form.LocationLines) > 0 {
				for _, locationLine := range form.LocationLines {
					pivotLocationLineAndLocationCenters = append(pivotLocationLineAndLocationCenters, models.PivotLocationLineAndLocationCenter{
						LocationLineID:   locationLine.ID,
						LocationCenterID: locationCenter.ID,
					})
				}
				models.Init(models.PivotLocationLineAndLocationCenter{}).
					Prepare().
					CreateInBatches(&pivotLocationLineAndLocationCenters, 100)
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var locationCenter models.LocationCenterModel
			ret := models.Init(models.LocationCenterModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx).
				First(&locationCenter)
			wrongs.PanicWhenIsEmpty(ret, "中心")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_center": locationCenter}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var locationCenters []models.LocationCenterModel
			models.Init(models.LocationCenterModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationCenters)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_centers": locationCenters}))
		})
	}
}
