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

// MenuRouter 菜单路由
type MenuRouter struct{}

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
		abnormals.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		abnormals.PanicValidate("菜单名称必填")
	}
	if len(cls.RbacRoleUUIDs) > 0 {
		models.Init(models.RbacRoleModel{}).DB().Where("uuid in ?", cls.RbacRoleUUIDs).Find(&cls.RbacRoles)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *MenuRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/menu",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建菜单
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&MenuStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "菜单名称和URL")

			// 新建
			menu := &models.MenuModel{
				BaseModel:  models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				Name:       form.Name,
				URL:        form.URL,
				URIName:    form.URIName,
				Icon:       form.Icon,
				ParentUUID: form.ParentUUID,
				RbacRoles:  form.RbacRoles,
			}
			if ret = (&models.BaseModel{}).SetModel(models.MenuModel{}).DB().Create(&menu); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"menu": menu}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 查询
			menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret = models.Init(models.MenuModel{}).DB().Delete(&menu); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&MenuStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "菜单名称和URL")

			// 查询
			menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			menu.Name = form.Name
			menu.URL = form.URL
			menu.URIName = form.URIName
			menu.Icon = form.Icon
			menu.ParentUUID = form.ParentUUID
			menu.RbacRoles = form.RbacRoles
			if ret = models.Init(models.MenuModel{}).DB().Save(&menu); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"menu": menu}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			menu := (&models.MenuModel{}).FindOneByUUID(ctx.Param("uuid"))
			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menu": menu}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var menus []models.MenuModel
			models.Init(models.MenuModel{}).
				SetWhereFields("uuid", "name", "url", "parent_uuid").
				SetPreloads("Parent", "Subs", "RbacRoles").
				PrepareQuery(ctx).
				Find(&menus)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menus": menus}))
		})
	}
}
