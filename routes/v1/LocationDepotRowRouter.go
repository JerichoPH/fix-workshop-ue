package v1

import (
	"fix-workshop-ue/wrongs"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// LocationDepotRowRouter 仓储仓库排路由
type LocationDepotRowRouter struct{}

// LocationDepotRowStoreForm 新建仓储仓库排表单
type LocationDepotRowStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	LocationDepotSectionUUID  string `form:"location_depot_section_uuid" json:"location_depot_section_uuid"`
	LocationDepotSection      models.LocationDepotSectionModel
	LocationDepotCabinetUUIDs []string `form:"location_depot_cabinet_uuids" json:"location_depot_cabinet_uuids"`
	LocationDepotCabinets     []models.LocationDepotCabinetModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationDepotRowStoreForm
func (cls LocationDepotRowStoreForm) ShouldBind(ctx *gin.Context) LocationDepotRowStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库排代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库排名称必填")
	}
	if cls.LocationDepotSectionUUID == "" {
		wrongs.PanicValidate("所属仓库区域必选")
	}
	ret = models.Init(models.LocationDepotRowModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationDepotSectionUUID}).
		Prepare().
		First(cls.LocationDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库区域")
	if len(cls.LocationDepotCabinetUUIDs) > 0 {
		models.Init(models.LocationDepotCabinetModel{}).
			GetSession().
			Where("uuid in ?", cls.LocationDepotCabinetUUIDs).
			Find(&cls.LocationDepotCabinets)
	}

	return cls
}

func (cls LocationDepotRowRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotRow", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotRowModel
			)

			// 表单
			form := (&LocationDepotRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

			// 新建
			locationDepotRow := &models.LocationDepotRowModel{
				BaseModel:             models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:            form.UniqueCode,
				Name:                  form.Name,
				LocationDepotSection:  form.LocationDepotSection,
				LocationDepotCabinets: form.LocationDepotCabinets,
			}
			if ret = models.Init(models.LocationDepotRowModel{}).GetSession().Create(&locationDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_depot_row": locationDepotRow}))
		})

		// 删除
		r.DELETE("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				locationDepotRow models.LocationDepotRowModel
			)

			// 查询
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			// 删除
			if ret := models.Init(models.LocationDepotRowModel{}).GetSession().Delete(&locationDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret                      *gorm.DB
				locationDepotRow, repeat models.LocationDepotRowModel
			)

			// 表单
			form := (&LocationDepotRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

			// 查询
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			// 编辑
			locationDepotRow.BaseModel.Sort = form.Sort
			locationDepotRow.UniqueCode = form.UniqueCode
			locationDepotRow.Name = form.Name
			locationDepotRow.LocationDepotSection = form.LocationDepotSection
			locationDepotRow.LocationDepotCabinets = form.LocationDepotCabinets
			if ret = models.Init(models.LocationDepotRowModel{}).GetSession().Save(&locationDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_depot_row": locationDepotRow}))
		})

		// 详情
		r.GET("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				locationDepotRow models.LocationDepotRowModel
			)
			ret = models.Init(models.LocationDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_depot_row": locationDepotRow}))
		})

		// 列表
		r.GET("depotRow", func(ctx *gin.Context) {
			var locataionDepotRows []models.LocationDepotRowModel
			models.Init(models.LocationDepotRowModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locataionDepotRows)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"locataion_depot_rows": locataionDepotRows}))
		})
	}
}
