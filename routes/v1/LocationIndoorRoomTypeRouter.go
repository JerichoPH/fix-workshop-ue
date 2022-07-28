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

// LocationIndoorRoomTypeRouter 机房路由
type LocationIndoorRoomTypeRouter struct{}

// LocationIndoorRoomTypeStoreForm 新建机房表单
type LocationIndoorRoomTypeStoreForm struct {
	Sort                    int64    `form:"sort" json:"sort"`
	UniqueCode              string   `form:"unique_code" json:"unique_code"`
	Name                    string   `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return LocationIndoorRoomTypeStoreForm
func (cls LocationIndoorRoomTypeStoreForm) ShouldBind(ctx *gin.Context) LocationIndoorRoomTypeStoreForm {
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
func (LocationIndoorRoomTypeRouter) Load(engine *gin.Engine) {
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
			form := (&LocationIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.LocationIndoorRoomTypeModel
			ret = models.Init(models.LocationIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.LocationIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 新建
			locationIndoorRoomType := &models.LocationIndoorRoomTypeModel{
				BaseModel:           models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:          form.UniqueCode,
				Name:                form.Name,
			}
			if ret = models.Init(models.LocationIndoorRoomTypeModel{}).GetSession().Create(&locationIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"location_indoor_room_type": locationIndoorRoomType}))
		})

		// 删除
		r.DELETE("indoorRoomType/:uuid", func(ctx *gin.Context) {
			// 查询
			locationIndoorRoomType := (&models.LocationIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret := models.Init(models.LocationIndoorRoomTypeModel{}).GetSession().Delete(&locationIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("indoorRoomType/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&LocationIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.LocationIndoorRoomTypeModel
			ret = models.Init(models.LocationIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房代码")
			ret = models.Init(models.LocationIndoorRoomTypeModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "机房名称")

			// 查询
			locationIndoorRoomType := (&models.LocationIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			locationIndoorRoomType.BaseModel.Sort = form.Sort
			locationIndoorRoomType.UniqueCode = form.UniqueCode
			locationIndoorRoomType.Name = form.Name
			if ret = models.Init(models.LocationIndoorRoomTypeModel{}).GetSession().Save(&locationIndoorRoomType); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"location_indoor_room_type": locationIndoorRoomType}))
		})

		// 详情
		r.GET("indoorRoomType/:uuid", func(ctx *gin.Context) {
			locationIndoorRoomType := (&models.LocationIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_room_type": locationIndoorRoomType}))
		})

		// 列表
		r.GET("indoorRoomType", func(ctx *gin.Context) {
			var locationIndoorRoomTypes []models.LocationIndoorRoomTypeModel
			models.Init(models.LocationIndoorRoomTypeModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationIndoorRoomTypes)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"location_indoor_room_types": locationIndoorRoomTypes}))
		})
	}
}
