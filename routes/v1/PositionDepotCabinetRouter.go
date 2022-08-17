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

// PositionDepotCabinetRouter 仓储仓库柜架路由
type PositionDepotCabinetRouter struct{}

// PositionDepotCabinetStoreForm 新建仓储仓库柜架表单
type PositionDepotCabinetStoreForm struct {
	Sort                 int64  `form:"sort" json:"sort"`
	UniqueCode           string `form:"unique_code" json:"unique_code"`
	Name                 string `form:"name" json:"name"`
	PositionDepotRowUUID string `form:"position_depot_row_uuid" json:"position_depot_row_uuid"`
	PositionDepotRow     models.PositionDepotRowModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotCabinetStoreForm
func (cls PositionDepotCabinetStoreForm) ShouldBind(ctx *gin.Context) PositionDepotCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库柜架代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库柜架名称必填")
	}
	if cls.PositionDepotRowUUID == "" {
		wrongs.PanicValidate("所属仓库排必选")
	}
	ret = models.Init(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotRowUUID}).
		Prepare("").
		First(&cls.PositionDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库排")

	return cls
}

// Load 加载路由
//  @receiver PositionDepotCabinetRouter
//  @param router
func (PositionDepotCabinetRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotCabinet",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotCabinetModel
			)

			// 表单
			form := (&PositionDepotCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

			// 新建
			positionDepotCabinet := &models.PositionDepotCabinetModel{
				BaseModel:        models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:       form.UniqueCode,
				Name:             form.Name,
				PositionDepotRow: form.PositionDepotRow,
			}
			if ret = models.Init(models.PositionDepotCabinetModel{}).Prepare("").Create(&positionDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				positionDepotCabinet models.PositionDepotCabinetModel
			)

			// 查询
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&positionDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			// 删除
			if ret := models.Init(models.PositionDepotCabinetModel{}).Prepare("").Delete(&positionDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                          *gorm.DB
				positionDepotCabinet, repeat models.PositionDepotCabinetModel
			)

			// 表单
			form := (&PositionDepotCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架代码")
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库柜架名称")

			// 查询
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&positionDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			// 编辑
			positionDepotCabinet.BaseModel.Sort = form.Sort
			positionDepotCabinet.UniqueCode = form.UniqueCode
			positionDepotCabinet.Name = form.Name
			positionDepotCabinet.PositionDepotRow = form.PositionDepotRow
			if ret = models.Init(models.PositionDepotCabinetModel{}).Prepare("").Save(&positionDepotCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				positionDepotCabinet models.PositionDepotCabinetModel
			)
			ret = models.Init(models.PositionDepotCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare("").
				First(&positionDepotCabinet)
			wrongs.PanicWhenIsEmpty(ret, "仓库柜架")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_cabinet": positionDepotCabinet}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionDepotCabinets []models.PositionDepotCabinetModel
			models.Init(models.PositionDepotCabinetModel{}).
				SetWhereFields().
				PrepareQuery(ctx,"").
				Find(&positionDepotCabinets)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_cabinets": positionDepotCabinets}))
		})
	}
}
