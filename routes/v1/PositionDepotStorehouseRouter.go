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

// PositionDepotStorehouseRouter 仓储仓库路由
type PositionDepotStorehouseRouter struct{}

// PositionDepotStorehouseStoreForm 仓储仓库新建表单
type PositionDepotStorehouseStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	OrganizationWorkshopUUID  string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop      models.OrganizationWorkshopModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotStorehouseStoreForm
func (cls PositionDepotStorehouseStoreForm) ShouldBind(ctx *gin.Context) PositionDepotStorehouseStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库名称必填")
	}
	if cls.OrganizationWorkshopUUID == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.Init(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).
		Prepare().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")

	return cls
}

// Load 加载路由 
//  @receiver cls 
//  @param router 
func (cls PositionDepotStorehouseRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/location",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("depotStorehouse", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotStorehouseModel
			)

			// 表单
			form := (&PositionDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库代码")

			// 新建
			positionDepotStorehouse := &models.PositionDepotStorehouseModel{
				BaseModel:             models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:            form.UniqueCode,
				Name:                  form.Name,
				OrganizationWorkshop:  form.OrganizationWorkshop,
			}
			if ret = models.Init(models.PositionDepotStorehouseModel{}).GetSession().Create(&positionDepotStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
		})

		// 删除
		r.DELETE("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				positionDepotStorehouse models.PositionDepotStorehouseModel
			)

			// 查询
			ret = models.Init(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			// 删除
			if ret := models.Init(models.PositionDepotStorehouseModel{}).GetSession().Delete(&positionDepotStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                             *gorm.DB
				positionDepotStorehouse, repeat models.PositionDepotStorehouseModel
			)

			// 表单
			form := (&PositionDepotStorehouseStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库代码")

			// 查询
			ret = models.Init(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			// 编辑
			positionDepotStorehouse.BaseModel.Sort = form.Sort
			positionDepotStorehouse.UniqueCode = form.UniqueCode
			positionDepotStorehouse.Name = form.Name
			positionDepotStorehouse.OrganizationWorkshop = form.OrganizationWorkshop
			if ret = models.Init(models.PositionDepotStorehouseModel{}).GetSession().Save(&positionDepotStorehouse); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
		})

		// 详情
		r.GET("depotStorehouse/:uuid", func(ctx *gin.Context) {
			var (
				ret                     *gorm.DB
				positionDepotStorehouse models.PositionDepotStorehouseModel
			)
			ret = models.Init(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotStorehouse)
			wrongs.PanicWhenIsEmpty(ret, "仓库")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
		})

		// 列表
		r.GET("depotStorehouse", func(ctx *gin.Context) {
			var positionDepotStorehouses []models.PositionDepotStorehouseModel
			models.Init(models.PositionDepotStorehouseModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionDepotStorehouses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_storehouses": positionDepotStorehouses}))
		})
	}
}
