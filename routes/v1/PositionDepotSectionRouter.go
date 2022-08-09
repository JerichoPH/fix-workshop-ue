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

// PositionDepotSectionRouter 仓储仓库区域路由
type PositionDepotSectionRouter struct{}

// PositionDepotSectionStoreForm 新建仓储仓库区域表单
type PositionDepotSectionStoreForm struct {
	Sort                        int64  `form:"sort" json:"sort"`
	UniqueCode                  string `form:"unique_code" json:"unique_code"`
	Name                        string `form:"name" json:"name"`
	PositionDepotStorehouseUUID string `form:"position_depot_storehouse_uuid" json:"position_depot_storehouse_uuid"`
	PositionDepotStorehouse     models.PositionDepotStorehouseModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotSectionStoreForm
func (cls PositionDepotSectionStoreForm) ShouldBind(ctx *gin.Context) PositionDepotSectionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库代码不能必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库名称不能必填")
	}
	if cls.PositionDepotStorehouseUUID == "" {
		wrongs.PanicValidate("所属仓库必选")
	}
	ret = models.Init(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotStorehouseUUID}).
		Prepare().
		First(&cls.PositionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库")

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls PositionDepotSectionRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/positionDepotSection",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat models.PositionDepotSectionModel
			)

			// 表单
			form := (&PositionDepotSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

			// 新建
			positionDepotSection := &models.PositionDepotSectionModel{
				BaseModel:               models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				UniqueCode:              form.UniqueCode,
				Name:                    form.Name,
				PositionDepotStorehouse: form.PositionDepotStorehouse,
			}
			if ret = models.Init(models.PositionDepotSectionModel{}).GetSession().Create(&positionDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"position_depot_section": positionDepotSection}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				positionDepotSection models.PositionDepotSectionModel
			)

			// 查询
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			// 删除
			if ret := models.Init(models.PositionDepotSectionModel{}).GetSession().Delete(&positionDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                          *gorm.DB
				positionDepotSection, repeat models.PositionDepotSectionModel
			)

			// 表单
			form := (&PositionDepotSectionStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

			// 查询
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			// 编辑
			positionDepotSection.BaseModel.Sort = form.Sort
			positionDepotSection.UniqueCode = form.UniqueCode
			positionDepotSection.Name = form.Name
			positionDepotSection.PositionDepotStorehouse = form.PositionDepotStorehouse
			if ret = models.Init(models.PositionDepotSectionModel{}).GetSession().Save(&positionDepotSection); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"position_depot_section": positionDepotSection}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret                  *gorm.DB
				positionDepotSection models.PositionDepotSectionModel
			)
			ret = models.Init(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&positionDepotSection)
			wrongs.PanicWhenIsEmpty(ret, "仓库区域")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_section": positionDepotSection}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var positionDepotSections []models.PositionDepotSectionModel
			models.Init(models.PositionDepotSectionModel{}).
				SetWhereFields().
				PrepareQuery(ctx).
				Find(&positionDepotSections)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"position_depot_sections": positionDepotSections}))
		})
	}
}
