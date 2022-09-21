package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotRowController struct{}

// PositionDepotRowStoreForm 新建仓储仓库排表单
type PositionDepotRowStoreForm struct {
	Sort                     int64  `form:"sort" json:"sort"`
	UniqueCode               string `form:"unique_code" json:"unique_code"`
	Name                     string `form:"name" json:"name"`
	PositionDepotRowTypeUuid string `form:"position_depot_row_type_uuid" json:"position_depot_row_type_uuid"`
	PositionDepotRowType     models.PositionDepotRowTypeModel
	PositionDepotSectionUuid string `form:"position_depot_section_uuid" json:"position_depot_section_uuid"`
	PositionDepotSection     models.PositionDepotSectionModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return PositionDepotRowStoreForm
func (cls PositionDepotRowStoreForm) ShouldBind(ctx *gin.Context) PositionDepotRowStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库排代码必填")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库排名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("仓库排名称不能超过64位")
	}
	if cls.PositionDepotRowTypeUuid == "" {
		wrongs.PanicValidate("所属排类型必选")
	}
	models.BootByModel(models.PositionDepotRowTypeModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotRowTypeUuid}).
		PrepareByDefault().
		First(&cls.PositionDepotRowType)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库排类型")
	if cls.PositionDepotSectionUuid == "" {
		wrongs.PanicValidate("所属仓库区域必选")
	}
	ret = models.BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"uuid": cls.PositionDepotSectionUuid}).
		PrepareByDefault().
		First(&cls.PositionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "所属仓库区域")

	return cls
}

func (PositionDepotRowController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.PositionDepotRowModel
	)

	// 表单
	form := (&PositionDepotRowStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

	// 新建
	positionDepotRow := &models.PositionDepotRowModel{
		BaseModel:            models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode:           form.UniqueCode,
		Name:                 form.Name,
		PositionDepotRowType: form.PositionDepotRowType,
		PositionDepotSection: form.PositionDepotSection,
	}
	if ret = models.BootByModel(models.PositionDepotRowModel{}).PrepareByDefault().Create(&positionDepotRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_depot_row": positionDepotRow}))
}
func (PositionDepotRowController) D(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		positionDepotRow models.PositionDepotRowModel
	)

	// 查询
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "仓库排")

	// 删除
	if ret := models.BootByModel(models.PositionDepotRowModel{}).PrepareByDefault().Delete(&positionDepotRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionDepotRowController) U(ctx *gin.Context) {
	var (
		ret                      *gorm.DB
		positionDepotRow, repeat models.PositionDepotRowModel
	)

	// 表单
	form := (&PositionDepotRowStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库排代码")
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "仓库排名称")

	// 查询
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&positionDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "仓库排")

	// 编辑
	positionDepotRow.BaseModel.Sort = form.Sort
	positionDepotRow.Name = form.Name
	positionDepotRow.PositionDepotRowType = form.PositionDepotRowType
	positionDepotRow.PositionDepotSection = form.PositionDepotSection
	if ret = models.BootByModel(models.PositionDepotRowModel{}).SetWheres(tools.Map{"uuid":ctx.Param("uuid")}).PrepareByDefault().Save(&positionDepotRow); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_depot_row": positionDepotRow}))
}
func (PositionDepotRowController) S(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		locationDepotRow models.PositionDepotRowModel
	)
	ret = models.BootByModel(models.PositionDepotRowModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&locationDepotRow)
	wrongs.PanicWhenIsEmpty(ret, "仓库排")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_row": locationDepotRow}))
}
func (PositionDepotRowController) I(ctx *gin.Context) {
	var locationDepotRows []models.PositionDepotRowModel
	models.BootByModel(models.PositionDepotRowModel{}).
		SetWhereFields().
		PrepareUseQuery(ctx, "").
		Find(&locationDepotRows)

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"position_depot_rows": locationDepotRows}))
}
