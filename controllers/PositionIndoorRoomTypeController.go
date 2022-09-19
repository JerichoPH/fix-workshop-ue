package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionIndoorRoomTypeController struct{}

// PositionIndoorRoomTypeStoreForm 新建机房表单
type PositionIndoorRoomTypeStoreForm struct {
	Sort       int64  `form:"sort" json:"sort"`
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
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
		wrongs.PanicValidate("机房类型代码必填")
	}
	if len(cls.UniqueCode) > 64 {
		wrongs.PanicValidate("机柜代码不能超过64位")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("机房类型名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("机房类型名称不能超过64位")
	}

	return cls
}

func (PositionIndoorRoomTypeController) C(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&PositionIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.PositionIndoorRoomTypeModel
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房名称")

	// 新建
	positionIndoorRoomType := &models.PositionIndoorRoomTypeModel{
		BaseModel:  models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	if ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).PrepareByDefault().Create(&positionIndoorRoomType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
}
func (PositionIndoorRoomTypeController) D(ctx *gin.Context) {
	// 查询
	locationIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 删除
	if ret := models.BootByModel(models.PositionIndoorRoomTypeModel{}).PrepareByDefault().Delete(&locationIndoorRoomType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (PositionIndoorRoomTypeController) U(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&PositionIndoorRoomTypeStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.PositionIndoorRoomTypeModel
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"unique_code": form.UniqueCode}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房代码")
	ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "机房名称")

	// 查询
	positionIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 编辑
	positionIndoorRoomType.BaseModel.Sort = form.Sort
	positionIndoorRoomType.Name = form.Name
	if ret = models.BootByModel(models.PositionIndoorRoomTypeModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefault().Save(&positionIndoorRoomType); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
}
func (PositionIndoorRoomTypeController) S(ctx *gin.Context) {
	positionIndoorRoomType := (&models.PositionIndoorRoomTypeModel{}).FindOneByUUID(ctx.Param("uuid"))

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_room_type": positionIndoorRoomType}))
}
func (PositionIndoorRoomTypeController) I(ctx *gin.Context) {
	var positionIndoorRoomTypes []models.PositionIndoorRoomTypeModel
	models.BootByModel(models.PositionIndoorRoomTypeModel{}).
		SetWhereFields().
		PrepareQuery(ctx, "").
		Find(&positionIndoorRoomTypes)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"position_indoor_room_types": positionIndoorRoomTypes}))
}
