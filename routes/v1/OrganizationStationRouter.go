package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// OrganizationStationRouter 站场路由
type OrganizationStationRouter struct{}

// OrganizationStationStoreForm 新建站场表单
type OrganizationStationStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	BeEnable                 bool   `form:"be_enable" json:"be_enable"`
	OrganizationWorkshopUUID string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea     models.OrganizationWorkAreaModel
	OrganizationLineUUIDs    []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines        []*models.OrganizationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationStationStoreForm
func (cls OrganizationStationStoreForm) ShouldBind(ctx *gin.Context) OrganizationStationStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		abnormals.PanicValidate("站场代码必填")
	}
	if cls.Name == "" {
		abnormals.PanicValidate("站场名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		abnormals.PanicValidate("所属车间必选")
	}
	cls.OrganizationWorkshop = (&models.OrganizationWorkshopModel{}).FindOneByUUID(cls.OrganizationWorkshopUUID)
	if cls.OrganizationWorkAreaUUID != "" {
		cls.OrganizationWorkArea = (&models.OrganizationWorkAreaModel{}).FindOneByUUID(cls.OrganizationWorkAreaUUID)
	}
	if len(cls.OrganizationLineUUIDs) > 0 {
		models.Init(models.OrganizationLineModel{}).DB().Where("uuid in ?", cls.OrganizationLineUUIDs).Find(&cls.OrganizationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls OrganizationStationRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("station", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationStationStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationStationModel
			ret = models.Init(models.OrganizationStationModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.OrganizationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站场名称")

			// 新建
			organizationStation := &models.OrganizationStationModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				BeEnable:             form.BeEnable,
				OrganizationWorkshop: form.OrganizationWorkshop,
				OrganizationWorkArea: form.OrganizationWorkArea,
				OrganizationLines:    form.OrganizationLines,
			}
			if ret = models.Init(models.OrganizationStationModel{}).DB().Create(&organizationStation); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"organization_station": organizationStation}))
		})

		// 删除
		r.DELETE("station/:uuid", func(ctx *gin.Context) {
			// 查询
			organizationStation := (&models.OrganizationStationModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.OrganizationStationModel{}).DB().Delete(&organizationStation); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("station/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationStationStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationStationModel
			ret = models.Init(models.OrganizationStationModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站场代码")
			ret = models.Init(models.OrganizationStationModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "站场名称")

			// 查询
			organizationStation := (&models.OrganizationStationModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			organizationStation.BaseModel.Sort = form.Sort
			organizationStation.UniqueCode = form.UniqueCode
			organizationStation.Name = form.Name
			organizationStation.BeEnable = form.BeEnable
			organizationStation.OrganizationWorkshop = form.OrganizationWorkshop
			organizationStation.OrganizationWorkArea = form.OrganizationWorkArea
			organizationStation.OrganizationLines = form.OrganizationLines
			if ret = models.Init(models.OrganizationStationModel{}).DB().Save(&organizationStation); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"organization_station": organizationStation}))
		})

		// 详情
		r.GET("station/:uuid", func(ctx *gin.Context) {
			organizationStation := (&models.OrganizationStationModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_station": organizationStation}))
		})

		// 列表
		r.GET("station", func(ctx *gin.Context) {
			var organizationStations []models.OrganizationStationModel
			models.Init(models.OrganizationStationModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&organizationStations)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_stations": organizationStations}))
		})
	}
}
