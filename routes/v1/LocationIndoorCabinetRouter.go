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

// LocationIndoorCabinetRouter 室内上道位置机柜路由
type LocationIndoorCabinetRouter struct{}

// LocationIndoorCabinetStoreForm 新建室内上道位置机柜表单
type LocationIndoorCabinetStoreForm struct {
	Sort                    int64  `form:"sort" json:"sort"`
	UniqueCode              string `form:"unique_code" json:"unique_code"`
	Name                    string `form:"name" json:"name"`
	LocationIndoorRowUUID   string `form:"location_indoor_row_uuid" json:"location_indoor_row_uuid"`
	LocationIndoorRow       models.LocationIndoorRowModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationIndoorCabinetStoreForm
func (cls LocationIndoorCabinetStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorCabinetStoreForm {
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
	if cls.LocationIndoorRowUUID == "" {
		ret = models.Init(models.LocationIndoorRowModel{}).
			SetWheres(tools.Map{"uuid": cls.LocationIndoorRowUUID}).
			Prepare().
			First(&cls.LocationIndoorRow)
		wrongs.PanicWhenIsEmpty(ret, "所属排")
	}

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorCabinetRouter
//  @param router
func (LocationIndoorCabinetRouter) Load(engine *gin.Engine) {
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
				repeat models.LocationIndoorCabinetModel
			)

			// 表单
			form := (&LocationIndoorCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架代码")
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架名称")

			// 新建
			locationIndoorCabinet := &models.LocationIndoorCabinetModel{
				BaseModel:           models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:          form.UniqueCode,
				Name:                form.Name,
				LocationIndoorRow:   form.LocationIndoorRow,
			}
			if ret = models.Init(models.LocationIndoorCabinetModel{}).GetSession().Create(&locationIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_cabinet": locationIndoorCabinet}))
		})

		// 删除
		r.DELETE("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				locationIndoorCabinet models.LocationIndoorCabinetModel
			)

			// 查询
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			// 删除
			if ret := models.Init(models.LocationIndoorCabinetModel{}).GetSession().Delete(&locationIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                           *gorm.DB
				locationIndoorCabinet, repeat models.LocationIndoorCabinetModel
			)

			// 表单
			form := (&LocationIndoorCabinetStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架代码")
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "柜架名称")

			// 查询
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			// 编辑
			locationIndoorCabinet.BaseModel.Sort = form.Sort
			locationIndoorCabinet.UniqueCode = form.UniqueCode
			locationIndoorCabinet.Name = form.Name
			locationIndoorCabinet.LocationIndoorRow = form.LocationIndoorRow
			if ret = models.Init(models.LocationIndoorCabinetModel{}).GetSession().Save(&locationIndoorCabinet); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_cabinet": locationIndoorCabinet}))
		})

		// 详情
		r.GET("indoorCabinet/:uuid", func(ctx *gin.Context) {
			var (
				ret                   *gorm.DB
				locationIndoorCabinet models.LocationIndoorCabinetModel
			)
			ret = models.Init(models.LocationIndoorCabinetModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorCabinet)
			wrongs.PanicWhenIsEmpty(ret, "柜架")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_cabinet": locationIndoorCabinet}))
		})

		// 列表
		r.GET("indoorCabinet", func(ctx *gin.Context) {
			var locationIndoorCabinet []models.LocationIndoorCabinetModel
			models.Init(models.LocationIndoorCabinetModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorCabinet)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_cabinet": locationIndoorCabinet}))
		})
	}
}
