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

// LocationDepotStorehouseRouter 仓储仓库路由 
type LocationDepotStorehouseRouter struct{}

func(cls *LocationDepotStorehouseStoreForm) ShouldBind(ctx *gin.Context) LocationDepotStorehouseStoreForm{

}

type LocationDepotStorehouseStoreForm struct{}

// Load 加载路由 
//  @receiver cls 
//  @param router 
func (cls LocationDepotStorehouseRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotStorehouse", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.LocationDepotStorehouseModel
			)

			// 表单
			form := (&LocationDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "仓库代码")
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "仓库名称")

			// 新建
			locationDepotStorehouse := &models.LocationDepotStorehouseModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
			}
			if ret = models.Init(models.LocationDepotStorehouseModel{}).DB().Create(&locationDepotStorehouse); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"locationDepotStorehouse": locationDepotStorehouse}))
		})

		// 删除
		r.DELETE("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				locationDepotStorehouse models.LocationDepotStorehouseModel
			)

			// 查询
			ret := models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotStorehouse)
			abnormals.PanicWhenIsEmpty(ret, "仓库")

			// 删除
			if ret := models.Init(models.LocationDepotStorehouseModel{}).DB().Delete(&locationDepotStorehouse); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                             *gorm.DB
				locationDepotStorehouse, repeat models.LocationDepotStorehouseModel
			)

			// 表单
			form := (&LocationDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "仓库代码")
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "仓库名称")

			// 查询
			ret := models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotStorehouse)
			abnormals.PanicWhenIsEmpty(ret, "仓库")

			// 编辑
			locationDepotStorehouse.BaseModel.Sort = form.Sort
			locationDepotStorehouse.UniqueCode = form.UniqueCode
			locationDepotStorehouse.Name = form.Name
			locationDepotStorehouse.BeEnable = form.BeEnable
			if ret = models.Init(models.LocationDepotStorehouseModel{}).DB().Save(&locationDepotStorehouse); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"locationDepotStorehouse": locationDepotStorehouse}))
		})

		// 详情
		r.GET("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				locationDepotStorehouse models.LocationDepotStorehouseModel
			)
			ret = models.Init(models.LocationDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&locationDepotStorehouse)
			abnormals.PanicWhenIsEmpty(ret, "仓库")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"locationDepotStorehouse": locationDepotStorehouse}))
		})

		// 列表
		r.GET("depotStorehouse", func(ctx *gin.Context) {
			var locationDepotStorehouses []models.LocationDepotStorehouseModel
			models.Init(models.LocationDepotStorehouseModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&locationDepotStorehouses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"locationDepotStorehouses": locationDepotStorehouses}))
		})
	}
}
