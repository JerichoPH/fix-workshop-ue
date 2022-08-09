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

// PositionIndoorRoomTypeRouter 机房路由
type PositionIndoorRoomTypeRouter struct{}

// PositionIndoorRoomTypeStoreForm 新建机房表单
type PositionIndoorRoomTypeStoreForm struct {
	Sort                    int64    `form:"sort" json:"sort"`
	UniqueCode              string   `form:"unique_code" json:"unique_code"`
	Name                    string   `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionIndoorRoomTypeStoreForm
func (cls PositionIndoorRoomTypeStoreForm) ShouldBind(ctx *gin.Context) PositionIndoorRoomTypeStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("机房代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("机房名称必填")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param engine
func (PositionIndoorRoomTypeRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("indoorRoomType", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&PositionIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.PositionIndoorRoomTypeModel
			ret = models.Init(models.PositionIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.PositionIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 新建
			positionIndoorRoomType := &models.PositionIndoorRoomTypeModel{
				BaseModel:           models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:          form.UniqueCode,
				Name:                form.Name,
			}
			if ret = models.Init(models.PositionIndoorRoomTypeModel{}).GetSession().Create(&positionIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
		})

		// 删除
		r.DELETE("indoorRoomType/:uuid", func(ctx *gin.Context) {
			// 查询
			locationIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.PositionIndoorRoomTypeModel{}).GetSession().Delete(&locationIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorRoomType/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&PositionIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.PositionIndoorRoomTypeModel
			ret = models.Init(models.PositionIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.PositionIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 查询
			positionIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			positionIndoorRoomType.BaseModel.Sort = form.Sort
			positionIndoorRoomType.UniqueCode = form.UniqueCode
			positionIndoorRoomType.Name = form.Name
			if ret = models.Init(models.PositionIndoorRoomTypeModel{}).GetSession().Save(&positionIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
		})

		// 详情
		r.GET("indoorRoomType/:uuid", func(ctx *gin.Context) {
			positionIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
		})

		// 列表
		r.GET("indoorRoomType", func(ctx *gin.Context) {
			var positionIndoorRoomTypes []models.PositionIndoorRoomTypeModel
			models.Init(models.PositionIndoorRoomTypeModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionIndoorRoomTypes)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_indoor_room_types": positionIndoorRoomTypes}))
		})
	}
}
