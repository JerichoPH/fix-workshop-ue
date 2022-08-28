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

// LocationSectionRouter 区间路由
type LocationSectionRouter struct{}

// LocationSectionStoreForm 新建区间表单
type LocationSectionStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationSectionStoreForm
func (cls LocationSectionStoreForm) ShouldBind(ctx *gin.Context) LocationSectionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("区间代码不能为空")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("区间名称不能为空")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间不能为空")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare("").
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "车间")
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).
			Prepare("").
			First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return cls
}

// LocationSectionBindLocationLinesForm 区间绑定线别表单
type LocationSectionBindLocationLinesForm struct {
	LocationLineUUIDs []string
	LocationLines     []*models.LocationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationSectionBindLocationLinesForm
func (cls LocationSectionBindLocationLinesForm) ShouldBind(ctx *gin.Context) LocationSectionBindLocationLinesForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (LocationSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/locationSection",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationSectionModel
			)

			// 表单
			form := (&LocationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间代码")
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间名称")

			// 新建
			organizationSection := &models.LocationSectionModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				BeEnable:             form.BeEnable,
				OrganizationWorkshop: form.OrganizationWorkshop,
				OrganizationWorkArea: form.OrganizationWorkArea,
			}
			if ret = models.BootByModel(models.LocationSectionModel{}).Prepare("").Create(&organizationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_section": organizationSection}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret             *gorm.DB
				locationSection models.LocationSectionModel
			)
			// 查询
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			// 删除
			if ret := models.BootByModel(models.LocationSectionModel{}).Prepare("").Delete(&locationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				locationSection, repeat models.LocationSectionModel
			)

			// 表单
			form := (&LocationSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间代码")
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "区间名称")

			// 查询
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			// 编辑
			locationSection.BaseModel.Sort = form.Sort
			locationSection.UniqueCode = form.UniqueCode
			locationSection.Name = form.Name
			locationSection.BeEnable = form.BeEnable
			locationSection.OrganizationWorkshop = form.OrganizationWorkshop
			locationSection.OrganizationWorkArea = form.OrganizationWorkArea
			if ret = models.BootByModel(models.LocationSectionModel{}).Prepare("").Save(&locationSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_section": locationSection}))
		})

		// 区间绑定线别
		r.PUT(":uuid/bindLocationLines", func(ctx *gin.Context) {
			var (
				ret                                  *gorm.DB
				locationSection                      models.LocationSectionModel
				pivotLocationLineAndLocationSections []models.PivotLocationLineAndLocationSection
			)

			// 表单
			form := (&LocationSectionBindLocationLinesForm{}).ShouldBind(ctx)

			if ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&locationSection); ret.Error != nil {
				wrongs.PanicWhenIsEmpty(ret, "区间")
			}

			// 删除原有绑定关系
			ret = models.BootByModel(models.BaseModel{}).Prepare("").Exec("delete from pivot_location_line_and_location_sections where location_section_id = ?", locationSection.ID)

			// 创建绑定关系
			if len(form.LocationLines) > 0 {
				for _, locationLine := range form.LocationLines {
					pivotLocationLineAndLocationSections = append(pivotLocationLineAndLocationSections, models.PivotLocationLineAndLocationSection{
						LocationLineID:    locationLine.ID,
						LocationSectionID: locationSection.ID,
					})
				}
				models.BootByModel(models.PivotLocationLineAndLocationSection{}).
					Prepare("").
					CreateInBatches(&pivotLocationLineAndLocationSections, 100)
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				organizationSection models.LocationSectionModel
			)
			ret = models.BootByModel(models.LocationSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetWhereFields("be_enable").
				PrepareQuery(ctx,"").
				First(&organizationSection)
			wrongs.PanicWhenIsEmpty(ret, "区间")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_section": organizationSection}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var locationSections []models.LocationSectionModel
			models.BootByModel(models.LocationSectionModel{}).
				SetWhereFields("unique_code", "Name", "be_enable", "organization_workshop_uuid", "organization_work_area_uuid").
				PrepareQuery(ctx,"").
				Find(&locationSections)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_sections": locationSections}))
		})
	}
}
