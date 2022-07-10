package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 权限路由
type RbacPermissionRouter struct{}

// RbacPermissionStoreForm 创建权限表单
type RbacPermissionStoreForm struct {
	Name                    string `form:"name" json:"name"`
	Uri                     string `form:"uri" json:"uri"`
	Method                  string `form:"method" json:"method"`
	RbacPermissionGroupUUID string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
}

// RbacPermissionUpdateForm 编辑权限表单
type RbacPermissionUpdateForm struct {
	Name                    string `form:"name" json:"name"`
	Uri                     string `form:"uri" json:"uri"`
	Method                  string `form:"method" json:"method"`
	RbacPermissionGroupUUID string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
}

func (cls *RbacPermissionRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/rbacPermission",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 创建权限
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			var form RbacPermissionStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}
			if form.Uri == "" {
				panic(exceptions.ThrowForbidden("URI必填"))
			}
			if form.Method == "" {
				panic(exceptions.ThrowForbidden("访问方法必选"))
			}
			if form.RbacPermissionGroupUUID == "" {
				panic(exceptions.ThrowForbidden("所属权限分组必选"))
			}
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": form.RbacPermissionGroupUUID}).
				Prepare().
				First(&models.RbacPermissionGroupModel{})
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "所属权限分组")

			// 查重
			var repeat models.RbacPermissionModel
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uri": form.Uri, "method": form.Method}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "权限Uri和访问方法")

			// 保存
			(&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				DB().
				Create(&models.RbacPermissionModel{
					BaseModel:               models.BaseModel{},
					Name:                    form.Name,
					URI:                     form.Uri,
					Method:                  form.Method,
					RbacPermissionGroupUUID: form.RbacPermissionGroupUUID,
				})

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除权限
		r.DELETE(":id", func(ctx *gin.Context) {
			var ret *gorm.DB
			id := tools.ThrowExceptionWhenIsNotInt(ctx.Param("id"), "权限编号必须是数字")
			var rbacPermission models.RbacPermissionModel

			// 查询
			ret = (&models.BaseModel{}).
				SetModel(models.RbacPermissionModel{}).
				SetWheres(tools.Map{"id": id}).
				Prepare().
				Find(&rbacPermission)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限")

			// 删除
			(&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				DB().
				Delete(&rbacPermission)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form RbacPermissionUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}
			if form.Uri == "" {
				panic(exceptions.ThrowForbidden("URI必填"))
			}
			if form.Method == "" {
				panic(exceptions.ThrowForbidden("访问方法必选"))
			}
			if form.RbacPermissionGroupUUID == "" {
				panic(exceptions.ThrowForbidden("所属权限分组必选"))
			}
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": form.RbacPermissionGroupUUID}).
				Prepare().
				First(&models.RbacPermissionGroupModel{})
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "所属权限分组")

			// 查重
			var repeat models.RbacPermissionModel
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uri": form.Uri, "method": form.Method}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "权限URI")

			// 查询
			var rbacPermission models.RbacPermissionModel
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacPermission)

			// 修改
			if form.Name != "" {
				rbacPermission.Name = form.Name
			}
			if form.Uri != "" {
				rbacPermission.URI = form.Uri
			}
			if form.Method != "" {
				rbacPermission.Method = form.Method
			}
			if form.RbacPermissionGroupUUID != "" {
				rbacPermission.RbacPermissionGroupUUID = form.RbacPermissionGroupUUID
			}

			// 保存
			(&models.BaseModel{}).SetModel(models.RbacPermissionModel{}).DB().Save(&rbacPermission)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 权限详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var rbacPermission models.RbacPermissionModel
			ret = (&models.BaseModel{}).
				SetModel(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				SetPreloads(tools.Strings{"RbacPermissionGroup"}).
				Prepare().
				First(&rbacPermission)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission": rbacPermission}))
		})

		// 权限列表
		r.GET("", func(ctx *gin.Context) {
			var rbacPermissions []models.RbacPermissionModel
			(&models.BaseModel{}).
				SetModel(models.RbacPermissionModel{}).
				SetPreloads(tools.Strings{"RbacPermissionGroup"}).
				SetWhereFields(tools.Strings{"name", "uri", "method", "rbac_permission_group_uuid"}).
				PrepareQuery(ctx).
				Find(&rbacPermissions)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permissions": rbacPermissions}))
		})
	}

}
