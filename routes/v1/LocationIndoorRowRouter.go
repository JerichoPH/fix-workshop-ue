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

// LocationIndoorRowRouter 上道位置排路由
type LocationIndoorRowRouter struct{}

// LocationIndoorRowStoreForm 新建上道位置排表单
type LocationIndoorRowStoreForm struct {
	Sort                       int64  `form:"sort" json:"sort"`
	UniqueCode                 string `form:"unique_code" json:"unique_code"`
	Name                       string `form:"name" json:"name"`
	LocationIndoorRoomUUID     string `form:"location_indoor_room_uuid" json:"location_indoor_room_uuid"`
	LocationIndoorRoom         models.LocationIndoorRoomModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationIndoorRowStoreForm
func (cls LocationIndoorRowStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorRowStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("排代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("排名称必填")
	}
	if cls.LocationIndoorRoomUUID == "" {
		wrongs.PanicValidate("所属机房必选")
	}
	ret = models.Init(models.LocationIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": cls.LocationIndoorRoomUUID}).
		Prepare().
		First(&cls.LocationIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "所属机房")

	return cls
}

// Load 加载路由
//  @receiver LocationIndoorRowRouter
//  @param router
func (LocationIndoorRowRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorRow", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationIndoorRowModel
			)

			// 表单
			form := (&LocationIndoorRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排代码")
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排名称")

			// 新建
			locationIndoorRow := &models.LocationIndoorRowModel{
				BaseModel:              models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:             form.UniqueCode,
				Name:                   form.Name,
				LocationIndoorRoom:     form.LocationIndoorRoom,
			}
			if ret = models.Init(models.LocationIndoorRowModel{}).GetSession().Create(&locationIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_row": locationIndoorRow}))
		})

		// 删除
		r.DELETE("indoorRow/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationIndoorRow models.LocationIndoorRowModel
			)

			// 查询
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			// 删除
			if ret := models.Init(models.LocationIndoorRowModel{}).GetSession().Delete(&locationIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorRow/:uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				locationIndoorRow, repeat models.LocationIndoorRowModel
			)

			// 表单
			form := (&LocationIndoorRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排代码")
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排名称")

			// 查询
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			// 编辑
			locationIndoorRow.BaseModel.Sort = form.Sort
			locationIndoorRow.UniqueCode = form.UniqueCode
			locationIndoorRow.Name = form.Name
			locationIndoorRow.LocationIndoorRoom = form.LocationIndoorRoom
			if ret = models.Init(models.LocationIndoorRowModel{}).GetSession().Save(&locationIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_row": locationIndoorRow}))
		})

		// 详情
		r.GET("indoorRow/:uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				locationIndoorRow models.LocationIndoorRowModel
			)
			ret = models.Init(models.LocationIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_row": locationIndoorRow}))
		})

		// 列表
		r.GET("indoorRow", func(ctx *gin.Context) {
			var locationIndoorRows []models.LocationIndoorRowModel
			models.Init(models.LocationIndoorRowModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorRows)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_rows": locationIndoorRows}))
		})
	}
}
