package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RbacRoleRouter struct{}

// RbacRoleStoreForm 创建角色表单
type RbacRoleStoreForm struct {
	Name string `form:"name" json:"name"`
}

// RbacRoleUpdateForm 编辑角色表单
type RbacRoleUpdateForm struct {
	Name string `form:"name" json:"name"`
}

// RbacRoleBindAccountsForm 角色绑定用户表单
type RbacRoleBindAccountsForm struct {
	AccountUUIDs []string `form:"account_uuids[]" json:"account_uuids"`
}

// RbacRoleBindPermissionsForm 角色绑定权限表单
type RbacRoleBindPermissionsForm struct {
	RbacPermissionUUIDs []string `form:"rbac_permission_uuids" json:"rbac_permission_uuids"`
}

func (cls *RbacRoleRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/rbacRole",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建角色
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			var form RbacRoleStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}

			// 查重
			var repeat RbacRoleStoreForm
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "角色名称")

			// 保存
			ret = (&models.BaseModel{}).DB().Create(&models.RbacRoleModel{Name: form.Name})
			if ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除角色
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")
			var rbacRole models.RbacRoleModel

			// 查询
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacRole)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "角色")

			// 删除
			if ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				DB().
				Delete(&rbacRole); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑角色
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form RbacRoleUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("名称必填"))
			}

			// 查重
			var repeat models.RbacRoleModel
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "角色名称")

			// 查询
			var rbacRole models.RbacRoleModel
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacRole)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "角色")

			// 修改
			if form.Name != "" {
				rbacRole.Name = form.Name
			}
			(&models.BaseModel{}).SetModel(models.RbacRoleModel{}).DB().Save(&rbacRole)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 绑定用户
		r.PUT(":uuid/bindAccounts", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var rbacRoleBindAccountsForm RbacRoleBindAccountsForm
			if err := ctx.ShouldBind(&rbacRoleBindAccountsForm); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}

			// 获取角色
			var rbacRole models.RbacRoleModel
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacRole)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "角色")

			// 获取用户
			var accounts []*models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				Prepare().
				Where("uuid in ?", rbacRoleBindAccountsForm.AccountUUIDs).
				Find(&accounts)
			if len(accounts) == 0 {
				panic(exceptions.ThrowEmpty("用户不存在"))
			}

			// 添加绑定关系
			rbacRole.Accounts = accounts
			(&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				DB().
				Save(&rbacRole)

			ctx.JSON(tools.CorrectIns("绑定成功").Updated(tools.Map{}))
		})

		// 绑定权限
		r.PUT(":uuid/bindPermissions", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form RbacRoleBindPermissionsForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}

			// 查询角色
			var rbacRole models.RbacRoleModel
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacRole)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "角色")

			// 查询权限
			var rbacPermissions []*models.RbacPermissionModel
			fmt.Println(form.RbacPermissionUUIDs)
			(&models.BaseModel{}).
				SetModel(models.RbacPermissionModel{}).
				DB().
				Where("uuid IN ?", form.RbacPermissionUUIDs).
				Find(&rbacPermissions)
			if len(rbacPermissions) == 0 {
				panic(exceptions.ThrowForbidden("没有找到权限"))
			}

			// 绑定
			rbacRole.RbacPermissions = rbacPermissions
			(&models.BaseModel{}).SetModel(models.RbacRoleModel{}).DB().Save(&rbacRole)

			ctx.JSON(tools.CorrectIns("绑定成功").Updated(tools.Map{}))
		})

		// 角色详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			var rbacRole models.RbacRoleModel
			uuid := ctx.Param("uuid")

			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				SetPreloads(tools.Strings{
					"RbacPermissions",
					"RbacPermissions.RbacPermissionGroup",
					"Accounts",
					"Accounts.AccountStatus",
					"Menus",
				}).
				Prepare().
				First(&rbacRole)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "角色")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_role": rbacRole}))
		})

		// 角色列表
		r.GET("", func(ctx *gin.Context) {
			var rbacRoles []models.RbacRoleModel
			(&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				PrepareQuery(ctx).
				Find(&rbacRoles)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_roles": rbacRoles}))
		})
	}
}
