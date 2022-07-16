package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MenuRouter 菜单路由
type MenuRouter struct{}

// MenuStoreForm 新建菜单表单
type MenuStoreForm struct {
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
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("名称必填"))
	}
	if len(cls.RbacRoleUUIDs) == 0 {
		panic(exceptions.ThrowEmpty("所属角色必选"))
	}
	// 查询角色
	models.Init(models.RbacRoleModel{}).
		DB().
		Where("uuid in ?", cls.RbacRoleUUIDs).
		Find(&cls.RbacRoles)
	if len(cls.RbacRoles) == 0 {
		panic(exceptions.ThrowEmpty("所选角色不存在"))
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
			exceptions.ThrowWhenIsRepeatByDB(ret, "菜单名称和URL")

			// 新建
			if ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				DB().
				Create(&models.MenuModel{
					Name:       form.Name,
					URL:        form.URL,
					URIName:    form.URIName,
					Icon:       form.Icon,
					ParentUUID: form.ParentUUID,
					RbacRoles:  form.RbacRoles,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 查询
			var menu models.MenuModel
			ret = models.Init(models.MenuModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			exceptions.ThrowWhenIsEmptyByDB(ret, "菜单")

			// 删除
			if ret = models.Init(models.MenuModel{}).
				DB().
				Delete(&menu); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (&MenuStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			exceptions.ThrowWhenIsRepeatByDB(ret, "菜单名称和URL")

			// 查询
			var menu models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			exceptions.ThrowWhenIsEmptyByDB(ret, "菜单")

			// 修改
			menu.Name = form.Name
			menu.URL = form.URL
			menu.URIName = form.URIName
			menu.Icon = form.Icon
			menu.ParentUUID = form.ParentUUID
			menu.RbacRoles = form.RbacRoles
			if ret = models.Init(models.MenuModel{}).
				DB().
				Save(&menu); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var menu models.MenuModel
			ret = models.Init(models.MenuModel{}).
				SetPreloads(tools.Strings{"Parent", "Subs", "RbacRoles"}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			exceptions.ThrowWhenIsEmptyByDB(ret, "菜单")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menu": menu}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var menus []models.MenuModel
			models.Init(models.MenuModel{}).
				SetWhereFields("uuid", "name", "url", "parent_uuid").
				SetPreloads(tools.Strings{"Parent", "Subs", "RbacRoles"}).
				PrepareQuery(ctx).
				Find(&menus)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menus": menus}))
		})
	}
}
