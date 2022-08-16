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

// LocationRailroadGradeCrossRouter 道口路由
type LocationRailroadGradeCrossRouter struct{}

// LocationRailroadGradeCrossStoreForm 新建道口表单
type LocationRailroadGradeCrossStoreForm struct {
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
func (cls LocationRailroadGradeCrossStoreForm) ShouldBind(ctx *gin.Context) LocationRailroadGradeCrossStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(ctx); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("道口代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("道口名称必填")
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

// LocationRailroadGradeCrossBindLocationLines 道口绑定线别表单
type LocationRailroadGradeCrossBindLocationLinesForm struct {
	LocationLineUUIDs []string
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationRailroadGradeCrossBindLocationLinesForm
func (cls LocationRailroadGradeCrossBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationRailroadGradeCrossBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if len(cls.LocationLineUUIDs)>0{
		models.Init(models.LocationLineModel{}).
			Prepare().
			Where("uuid in ?",cls.LocationLineUUIDs).
			Find(&cls.LocationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationRailroadGradeCrossRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationRailroadGradeCross",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationRailroadGradeCrossModel
			)

			// 表单
			form := (&LocationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "道口名称")

			// 新建
			locationRailroadGradeCross := &models.LocationRailroadGradeCrossModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
				BeEnable:   form.BeEnable,
			}
			if ret = models.Init(models.LocationRailroadGradeCrossModel{}).Prepare().Create(&locationRailroadGradeCross); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationRailroadGradeCross models.LocationRailroadGradeCrossModel
			)

			// 查询
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationRailroadGradeCross)
			wrongs.PanicWhenIsEmpty(ret, "道口")

			// 删除
			if ret := models.Init(models.LocationRailroadGradeCrossModel{}).Prepare().Delete(&locationRailroadGradeCross); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                                *gorm.DB
				locationRailroadGradeCross, repeat models.LocationRailroadGradeCrossModel
			)

			// 表单
			form := (&LocationRailroadGradeCrossStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "道口代码")
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "道口名称")

			// 查询
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationRailroadGradeCross)
			wrongs.PanicWhenIsEmpty(ret, "道口")

			// 编辑
			locationRailroadGradeCross.BaseModel.Sort = form.Sort
			locationRailroadGradeCross.UniqueCode = form.UniqueCode
			locationRailroadGradeCross.Name = form.Name
			locationRailroadGradeCross.BeEnable = form.BeEnable
			if ret = models.Init(models.LocationRailroadGradeCrossModel{}).Prepare().Save(&locationRailroadGradeCross); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
		})

		// 道口绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) {
			var (
				ret                                              *gorm.DB
				locationRailroadGradeCross                       models.LocationRailroadGradeCrossModel
				pivotLocationLineAndLocationRailroadGradeCrosses []models.PivotLocationLineAndLocationRailroadGradeCross
			)

			// 表单
			form := (&LocationRailroadGradeCrossBindLocationLinesForm{}).ShouldBind(ctx)

			if ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationRailroadGradeCross); ret.Error != nil {
				wrongs.PanicWhenIsEmpty(ret, "道口")
			}

			// 删除原有绑定关系
			ret = models.Init(models.BaseModel{}).Prepare().Exec("delete from pivot_location_line_and_location_railroad_grade_crosses where location_railroad_grade_crosses_id = ?", locationRailroadGradeCross.ID)

			// 创建绑定关系
			if len(form.LocationLines) > 0 {
				for _, locationLine := range form.LocationLines {
					pivotLocationLineAndLocationRailroadGradeCrosses = append(pivotLocationLineAndLocationRailroadGradeCrosses, models.PivotLocationLineAndLocationRailroadGradeCross{
						LocationLineID:    locationLine.ID,
						LocationRailroadGradeCrossID: locationRailroadGradeCross.ID,
					})
				}
				models.Init(models.PivotLocationLineAndLocationRailroadGradeCross{}).
					Prepare().
					CreateInBatches(&pivotLocationLineAndLocationRailroadGradeCrosses, 100)
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                        *gorm.DB
				locationRailroadGradeCross models.LocationRailroadGradeCrossModel
			)
			ret = models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx).
				First(&locationRailroadGradeCross)
			wrongs.PanicWhenIsEmpty(ret, "道口")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_railroad_grade_cross": locationRailroadGradeCross}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var locationRailroadGradeCrosses []models.LocationRailroadGradeCrossModel
			models.Init(models.LocationRailroadGradeCrossModel{}).
				SetWhereFields("unique_code", "name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
				PrepareQuery(ctx).
				Find(&locationRailroadGradeCrosses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_railroad_grade_crosses": locationRailroadGradeCrosses}))
		})
	}
}
