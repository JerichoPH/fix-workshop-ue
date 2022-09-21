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
	RbacRoleUuids []string `form:"rbac_role_uuids" json:"rbac_role_uuids"`
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
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("菜单名称不能超过64位")
	}
	if len(cls.RbacRoleUuids) > 0 {
		models.BootByModel(models.RbacRoleModel{}).PrepareByDefaultDbDriver().Where("uuid in ?", cls.RbacRoleUuids).Find(&cls.RbacRoles)
	}

	return cls
}

// C 新建
func (MenuController) C(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&MenuStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.MenuModel
	ret = (&models.BaseModel{}).
		SetModel(models.MenuModel{}).
		SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
		PrepareByDefaultDbDriver().
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
	if ret = (&models.BaseModel{}).SetModel(models.MenuModel{}).PrepareByDefaultDbDriver().Create(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"menu": menu}))
}

// D 删除
func (MenuController) D(ctx *gin.Context) {
	var ret *gorm.DB

	// 查询
	menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))

	// 删除
	if ret = models.BootByModel(models.MenuModel{}).PrepareByDefaultDbDriver().Delete(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 更新
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
		PrepareByDefaultDbDriver().
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
	if ret = models.BootByModel(models.MenuModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&menu); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"menu": menu}))
}

// S 详情
func (MenuController) S(ctx *gin.Context) {
	menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))
	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"menu": menu}))
}

// I 列表
func (MenuController) I(ctx *gin.Context) {
	var (
		menus []models.MenuModel
		count           int64
		db              *gorm.DB
	)
	db = models.BootByModel(models.MenuModel{}).
		SetWhereFields("uuid", "name", "url", "parent_uuid").
		SetPreloads("Parent", "Subs", "RbacRoles").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&menus)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"menus": menus}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&menus)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"menus": menus}, ctx.Query("__page__"), count))
	}
}
