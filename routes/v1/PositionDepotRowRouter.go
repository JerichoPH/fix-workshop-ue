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

// PositionDepotRowRouter 仓储仓库排路由
type PositionDepotRowRouter struct{}

// PositionDepotRowStoreForm 新建仓储仓库排表单
type PositionDepotRowStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	PositionDepotRowTypeUUID string `form:"position_depot_row_type_uuid" json:"position_depot_row_type_uuid"`
	PositionDepotRowType     models.PositionDepotRowTypeModel
	PositionDepotSectionUUID string `form:"position_depot_section_uuid" json:"position_depot_section_uuid"`
	PositionDepotSection     models.PositionDepotSectionModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotRowStoreForm
func (cls PositionDepotRowStoreForm) ShouldBind(ctx *gin.Context) PositionDepotRowStoreForm {
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
	if cls.PositionDepotRowTypeUUID == "" {
		wrongs.PanicValidate("所属排类型必选")
	}
	models.Init(models.PositionDepotRowTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotRowTypeUUID}).
		Prepare().
		First(&cls.PositionDepotRowType)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库排类型")
	if cls.PositionDepotSectionUUID == "" {
		wrongs.PanicValidate("所属仓库区域必选")
	}
	ret = models.Init(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotSectionUUID}).
		Prepare().
		First(&cls.PositionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库区域")

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param engine
func (PositionDepotRowRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotRow", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotRowModel
			)

			// 表单
			form := (&PositionDepotRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

			// 新建
			positionDepotRow := &models.PositionDepotRowModel{
				BaseModel:            models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:           form.UniqueCode,
				Name:                 form.Name,
				PositionDepotRowType: form.PositionDepotRowType,
				PositionDepotSection: form.PositionDepotSection,
			}
			if ret = models.Init(models.PositionDepotRowModel{}).GetSession().Create(&positionDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_depot_row": positionDepotRow}))
		})

		// 删除
		r.DELETE("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				positionDepotRow models.PositionDepotRowModel
			)

			// 查询
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			// 删除
			if ret := models.Init(models.PositionDepotRowModel{}).GetSession().Delete(&positionDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret                      *gorm.DB
				positionDepotRow, repeat models.PositionDepotRowModel
			)

			// 表单
			form := (&PositionDepotRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

			// 查询
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			// 编辑
			positionDepotRow.BaseModel.Sort = form.Sort
			positionDepotRow.UniqueCode = form.UniqueCode
			positionDepotRow.Name = form.Name
			positionDepotRow.PositionDepotRowType = form.PositionDepotRowType
			positionDepotRow.PositionDepotSection = form.PositionDepotSection
			if ret = models.Init(models.PositionDepotRowModel{}).GetSession().Save(&positionDepotRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_row": positionDepotRow}))
		})

		// 详情
		r.GET("depotRow/:uuid", func(ctx *gin.Context) {
			var (
				ret              *gorm.DB
				locationDepotRow models.PositionDepotRowModel
			)
			ret = models.Init(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotRow)
			wrongs.PanicWhenIsEmpty(ret, "仓库排")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_row": locationDepotRow}))
		})

		// 列表
		r.GET("depotRow", func(ctx *gin.Context) {
			var locationDepotRows []models.PositionDepotRowModel
			models.Init(models.PositionDepotRowModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotRows)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_rows": locationDepotRows}))
		})
	}
}
