package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotSectionController struct{}

// PositionDepotSectionStoreForm 新建仓储仓库区域表单
type PositionDepotSectionStoreForm struct {
	Sort                        int64  `form:"sort" json:"sort"`
	UniqueCode                  string `form:"unique_code" json:"unique_code"`
	Name                        string `form:"name" json:"name"`
	PositionDepotStorehouseUuid string `form:"position_depot_storehouse_uuid" json:"position_depot_storehouse_uuid"`
	PositionDepotStorehouse     models.PositionDepotStorehouseModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return PositionDepotSectionStoreForm
func (ins PositionDepotSectionStoreForm) ShouldBind(ctx *gin.Context) PositionDepotSectionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库区域代码不能必填")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库区域名称不能必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("仓库区域名称不能超过64位")
	}
	if ins.PositionDepotStorehouseUuid == "" {
		wrongs.PanicValidate("所属仓库必选")
	}
	ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"uuid": ins.PositionDepotStorehouseUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库")

	return ins
}

// C 新建
func (PositionDepotSectionController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotSectionModel
	)

	// 表单
	form := (&PositionDepotSectionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

	// 新建
	positionDepotSection := &models.PositionDepotSectionModel{
		BaseModel:               models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:              form.UniqueCode,
		Name:                    form.Name,
		PositionDepotStorehouse: form.PositionDepotStorehouse,
	}
	if ret = models.BootByModel(models.PositionDepotSectionModel{}).PrepareByDefaultDbDriver().Create(&positionDepotSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_depot_section": positionDepotSection}))
}

// D 删除
func (PositionDepotSectionController) D(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		positionDepotSection models.PositionDepotSectionModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域")

	// 删除
	if ret := models.BootByModel(models.PositionDepotSectionModel{}).PrepareByDefaultDbDriver().Delete(&positionDepotSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (PositionDepotSectionController) U(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		positionDepotSection, repeat models.PositionDepotSectionModel
	)

	// 表单
	form := (&PositionDepotSectionStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库区域代码")
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库区域名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域")

	// 编辑
	positionDepotSection.BaseModel.Sort = form.Sort
	positionDepotSection.Name = form.Name
	positionDepotSection.PositionDepotStorehouse = form.PositionDepotStorehouse
	if ret = models.BootByModel(models.PositionDepotSectionModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&positionDepotSection); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_section": positionDepotSection}))
}

// S 详情
func (PositionDepotSectionController) S(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		positionDepotSection models.PositionDepotSectionModel
	)
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&positionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_section": positionDepotSection}))
}

// I 列表
func (PositionDepotSectionController) I(ctx *gin.Context) {
	var (
		positionDepotSections []models.PositionDepotSectionModel
		count                 int64
		db                    *gorm.DB
	)
	db = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWhereFields().
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&positionDepotSections)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_sections": positionDepotSections}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&positionDepotSections)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"position_depot_sections": positionDepotSections}, ctx.Query("__page__"), count))
	}
}
