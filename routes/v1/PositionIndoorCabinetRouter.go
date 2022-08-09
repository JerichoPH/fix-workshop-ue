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

// PositionIndoorCabinetRouter 室内上道位置机柜路由
type PositionIndoorCabinetRouter struct{}

// PositionIndoorCabinetStoreForm 新建室内上道位置机柜表单
type PositionIndoorCabinetStoreForm struct {
	Sort                  int64  `form:"sort" json:"sort"`
	UniqueCode            string `form:"unique_code" json:"unique_code"`
	Name                  string `form:"name" json:"name"`
	PositionIndoorRowUUID string `form:"position_indoor_row_uuid" json:"position_indoor_row_uuid"`
	PositionIndoorRow     models.PositionIndoorRowModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionIndoorCabinetStoreForm
func (cls PositionIndoorCabinetStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorCabinetStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("柜架代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("柜架名称必填")
	}
	if cls.PositionIndoorRowUUID == "" {
		ret = models.Init(models.PositionIndoorRowModel{}).
			SetWheres(tools.Map{"uuid": cls.PositionIndoorRowUUID}).
			Prepare().
			First(&cls.PositionIndoorRow)
		wrongs.PanicWhenIsEmpty(ret, "所属排")
	}

	return cls
}

// Load 加载路由
//  @receiver PositionIndoorCabinetRouter
//  @param router
func (PositionIndoorCabinetRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorCabinet", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionIndoorCabinetModel
			)

			// 表单
			form := (&PositionIndoorCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架代码")
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架名称")

			// 新建
			locationIndoorCabinet := &models.PositionIndoorCabinetModel{
				BaseModel:         models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:        form.UniqueCode,
				Name:              form.Name,
				PositionIndoorRow: form.PositionIndoorRow,
			}
			if ret = models.Init(models.PositionIndoorCabinetModel{}).GetSession().Create(&locationIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_indoor_cabinet": locationIndoorCabinet}))
		})

		// 删除
		r.DELETE("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				positionIndoorCabinet models.PositionIndoorCabinetModel
			)

			// 查询
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			// 删除
			if ret := models.Init(models.PositionIndoorCabinetModel{}).GetSession().Delete(&positionIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                           *gorm.DB
				positionIndoorCabinet, repeat models.PositionIndoorCabinetModel
			)

			// 表单
			form := (&PositionIndoorCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架代码")
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架名称")

			// 查询
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			// 编辑
			positionIndoorCabinet.BaseModel.Sort = form.Sort
			positionIndoorCabinet.UniqueCode = form.UniqueCode
			positionIndoorCabinet.Name = form.Name
			positionIndoorCabinet.PositionIndoorRow = form.PositionIndoorRow
			if ret = models.Init(models.PositionIndoorCabinetModel{}).GetSession().Save(&positionIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
		})

		// 详情
		r.GET("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				positionIndoorCabinet models.PositionIndoorCabinetModel
			)
			ret = models.Init(models.PositionIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
		})

		// 列表
		r.GET("indoorCabinet", func(ctx *gin.Context) {
			var positionIndoorCabinet []models.PositionIndoorCabinetModel
			models.Init(models.PositionIndoorCabinetModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionIndoorCabinet)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_cabinet": positionIndoorCabinet}))
		})
	}
}