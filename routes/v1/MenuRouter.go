package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MenuRouter struct{}

type MenuStoreForm struct {
	Name       string `form:"name" json:"name"`
	URL        string `form:"url" json:"url"`
	URIName    string `form:"uri_name" json:"uri_name"`
	ParentUUID string `form:"parent_uuid" json:"parent_uuid"`
}

type MenuUpdateForm struct {
	Name       string `form:"name" json:"name"`
	URL        string `form:"url" json:"url"`
	URIName    string `form:"uri_name" json:"uri_name"`
	ParentUUID string `form:"parent_uuid" json:"parent_uuid"`
}

func (cls *MenuRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/menu",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建菜单
		var form MenuStoreForm
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}

			// 查重
			var repeat models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "菜单名称和URL")

			// 查询

			// 新建
			if ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				DB().
				Create(&models.MenuModel{
					Name:       form.Name,
					URL:        form.URL,
					URIName:    form.URIName,
					ParentUUID: form.ParentUUID,
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
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "菜单")

			// 删除
			if ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
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
			var form MenuUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}

			// 查重
			var repeat models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"name": form.Name, "url": form.URL}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "菜单名称和URL")

			// 查询
			var menu models.MenuModel
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "菜单")

			// 修改
			if form.Name != "" {
				menu.Name = form.Name
			}
			menu.URL = form.URL
			menu.URIName = form.URIName
			menu.ParentUUID = form.ParentUUID
			if ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
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
			ret = (&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetPreloads(tools.Strings{"Parent", "Subs"}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&menu)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "菜单")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menu": menu}))
		})

		// 列表
		r.GET("", func(ctx *gin.Context) {
			var menus []models.MenuModel
			(&models.BaseModel{}).
				SetModel(models.MenuModel{}).
				SetWhereFields(tools.Strings{"uuid", "name", "url", "parent_uuid"}).
				SetPreloads(tools.Strings{"Parent"}).
				PrepareQuery(ctx).
				Find(&menus)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menus": menus}))
		})
	}
}
