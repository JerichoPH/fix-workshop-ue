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

// PositionIndoorRowRouter 上道位置排路由
type PositionIndoorRowRouter struct{}

// PositionIndoorRowStoreForm 新建上道位置排表单
type PositionIndoorRowStoreForm struct {
	Sort                   int64  `form:"sort" json:"sort"`
	UniqueCode             string `form:"unique_code" json:"unique_code"`
	Name                   string `form:"name" json:"name"`
	PositionIndoorRoomUUID string `form:"position_indoor_room_uuid" json:"position_indoor_room_uuid"`
	PositionIndoorRoom     models.PositionIndoorRoomModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionIndoorRowStoreForm
func (cls PositionIndoorRowStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRowStoreForm {
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
	if cls.PositionIndoorRoomUUID == "" {
		wrongs.PanicValidate("所属机房必选")
	}
	ret = models.Init(models.PositionIndoorRoomModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionIndoorRoomUUID}).
		Prepare().
		First(&cls.PositionIndoorRoom)
	wrongs.PanicWhenIsEmpty(ret, "所属机房")

	return cls
}

// Load 加载路由
//  @receiver PositionIndoorRowRouter
//  @param router
func (PositionIndoorRowRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionIndoorRow",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionIndoorRowModel
			)

			// 表单
			form := (&PositionIndoorRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排代码")
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排名称")

			// 新建
			positionIndoorRow := &models.PositionIndoorRowModel{
				BaseModel:          models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:         form.UniqueCode,
				Name:               form.Name,
				PositionIndoorRoom: form.PositionIndoorRoom,
			}
			if ret = models.Init(models.PositionIndoorRowModel{}).Prepare().Create(&positionIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_indoor_row": positionIndoorRow}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionIndoorRow models.PositionIndoorRowModel
			)

			// 查询
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			// 删除
			if ret := models.Init(models.PositionIndoorRowModel{}).Prepare().Delete(&positionIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                       *gorm.DB
				positionIndoorRow, repeat models.PositionIndoorRowModel
			)

			// 表单
			form := (&PositionIndoorRowStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排代码")
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "排名称")

			// 查询
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			// 编辑
			positionIndoorRow.BaseModel.Sort = form.Sort
			positionIndoorRow.UniqueCode = form.UniqueCode
			positionIndoorRow.Name = form.Name
			positionIndoorRow.PositionIndoorRoom = form.PositionIndoorRoom
			if ret = models.Init(models.PositionIndoorRowModel{}).Prepare().Save(&positionIndoorRow); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_indoor_row": positionIndoorRow}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret               *gorm.DB
				positionIndoorRow models.PositionIndoorRowModel
			)
			ret = models.Init(models.PositionIndoorRowModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionIndoorRow)
			wrongs.PanicWhenIsEmpty(ret, "排")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_row": positionIndoorRow}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionIndoorRows []models.PositionIndoorRowModel
			models.Init(models.PositionIndoorRowModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionIndoorRows)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_rows": positionIndoorRows}))
		})
	}
}
