package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MenuController struct{}

// MenuStoreForm 新建菜单表单
type MenuStoreForm struct {
	Sort          int64    `form:"sort" json:"sort"`
	Name          string   `form:"name" json:"name"`
	URL           string   `form:"url" json:"url"`
	URIName       string   `form:"uri_name" json:"uri_name"`
	ParentUUID    string   `form:"parent_uuid" json:"parent_uuid"`
	Icon          string   `form:"icon" json:"icon"`
	RbacRoleUUIDs []string `form:"rbac_role_uuids" json:"rbac_role_uuids"`
	RbacRoles     []*models.RbacRoleModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return MenuStoreForm
func (cls MenuStoreForm) ShouldBind(ctx *gin.Context) MenuStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		wrongs.PanicValidate("菜单名称必填")
	}
	if len(cls.RbacRoleUUIDs) > 0 {
		models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Where("uuid in ?", cls.RbacRoleUUIDs).Find(&cls.RbacRoles)
	}

	return cls
}

func (MenuController) C(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&MenuStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.MenuModel
	ret = (&models.BaseModel{}).
		SetModel(models.MenuModel{}).
		SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "菜单名称和URL")

	// 新建
	menu := &models.MenuModel{
		BaseModel:  models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		Name:       form.Name,
		URL:        form.URL,
		URIName:    form.URIName,
		Icon:       form.Icon,
		ParentUuid: form.ParentUUID,
		RbacRoles:  form.RbacRoles,
	}
	if ret = (&models.BaseModel{}).SetModel(models.MenuModel{}).PrepareByDefault().Create(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"menu": menu}))
}
func (MenuController) D(ctx *gin.Context) {
	var ret *gorm.DB

	// 查询
	menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 删除
	if ret = models.BootByModel(models.MenuModel{}).PrepareByDefault().Delete(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (MenuController) U(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&MenuStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.MenuModel
	ret = (&models.BaseModel{}).
		SetModel(models.MenuModel{}).
		SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "菜单名称和URL")

	// 查询
	menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 编辑
	menu.Name = form.Name
	menu.URL = form.URL
	menu.URIName = form.URIName
	menu.Icon = form.Icon
	menu.ParentUuid = form.ParentUUID
	menu.RbacRoles = form.RbacRoles
	if ret = models.BootByModel(models.MenuModel{}).PrepareByDefault().Save(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"menu": menu}))
}
func (MenuController) S(ctx *gin.Context) {
	menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))
	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"menu": menu}))
}
func (MenuController) I(ctx *gin.Context) {
	var menus []models.MenuModel
	models.BootByModel(models.MenuModel{}).
		SetWhereFields("uuid", "name", "url", "parent_uuid").
		SetPreloads("Parent", "Subs", "RbacRoles").
		PrepareQuery(ctx, "").
		Find(&menus)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"menus": menus}))
}
