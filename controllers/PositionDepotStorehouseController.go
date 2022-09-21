package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotStorehouseController struct{}

// PositionDepotStorehouseStoreForm 仓储仓库新建表单
type PositionDepotStorehouseStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	OrganizationWorkshopUuid string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop     models.OrganizationWorkshopModel
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
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("仓库名称不能超过64位")
	}
	if cls.OrganizationWorkshopUuid == "" {
		wrongs.PanicValidate("所属车间必选")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUuid}).
		PrepareByDefaultDbDriver().
		First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "所属车间")

	return cls
}

// C 新建
func (PositionDepotStorehouseController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotStorehouseModel
	)

	// 表单
	form := (&PositionDepotStorehouseStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库代码")

	// 新建
	positionDepotStorehouse := &models.PositionDepotStorehouseModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		OrganizationWorkshop: form.OrganizationWorkshop,
	}
	if ret = models.BootByModel(models.PositionDepotStorehouseModel{}).PrepareByDefaultDbDriver().Create(&positionDepotStorehouse); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
}

// D 删除
func (PositionDepotStorehouseController) D(ctx *gin.Context) {
	var (
		ret                     *gorm.DB
		positionDepotStorehouse models.PositionDepotStorehouseModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "仓库")

	// 删除
	if ret := models.BootByModel(models.PositionDepotStorehouseModel{}).PrepareByDefaultDbDriver().Delete(&positionDepotStorehouse); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (PositionDepotStorehouseController) U(ctx *gin.Context) {
	var (
		ret                             *gorm.DB
		positionDepotStorehouse, repeat models.PositionDepotStorehouseModel
	)

	// 表单
	form := (&PositionDepotStorehouseStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库代码")

	// 查询
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "仓库")

	// 编辑
	positionDepotStorehouse.BaseModel.Sort = form.Sort
	positionDepotStorehouse.Name = form.Name
	positionDepotStorehouse.OrganizationWorkshop = form.OrganizationWorkshop
	if ret = models.BootByModel(models.PositionDepotStorehouseModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionDepotStorehouse); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
}

// S 详情
func (PositionDepotStorehouseController) S(ctx *gin.Context) {
	var (
		ret                     *gorm.DB
		positionDepotStorehouse models.PositionDepotStorehouseModel
	)
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "仓库")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_storehouse": positionDepotStorehouse}))
}

// I 列表
func (PositionDepotStorehouseController) I(ctx *gin.Context) {

	var (
		positionDepotStorehouses []models.PositionDepotStorehouseModel
		count                    int64
		db                       *gorm.DB
	)
	db = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionDepotStorehouses)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_storehouses": positionDepotStorehouses}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionDepotStorehouses)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_depot_storehouses": positionDepotStorehouses}, ctx.Query("__page__"), count))
	}
}
