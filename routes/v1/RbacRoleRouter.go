package v1

import (
	"fix-workshop-ue/errors"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
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
	RbacPermissionIDs []string `form:"permission_ids[]" json:"permission_ids"`
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
				panic(err)
			}

			// 查重
			var repeat RbacRoleStoreForm
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "角色名称")

			// 保存
			ret = (&models.BaseModel{}).DB().Create(&models.RbacRoleModel{Name: form.Name})
			if ret.Error != nil {
				panic(errors.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})
	}

	// 编辑角色
	r.PUT(":id", func(ctx *gin.Context) {
		var ret *gorm.DB
		id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "角色编号必须是数字")

		// 表单
		var form RbacRoleUpdateForm
		if err := ctx.ShouldBind(&form); err != nil {
			panic(err)
		}

		// 查重
		var repeat models.RbacRoleModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			SetWheres(tools.Map{"name": form.Name}).
			SetNotWheres(tools.Map{"id": id}).
			Prepare().
			First(&repeat)
		tools.ThrowErrorWhenIsRepeatByDB(ret, "角色名称")

		// 查询
		var rbacRole models.RbacRoleModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			SetWheres(tools.Map{"id": id}).
			Prepare().
			First(&rbacRole)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "角色")

		// 修改
		if form.Name != "" {
			rbacRole.Name = form.Name
		}
		(&models.BaseModel{}).SetModel(models.RbacRoleModel{}).DB().Save(&rbacRole)

		ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
	})

	// 绑定用户
	r.PUT(":id/bindAccounts", func(ctx *gin.Context) {
		var ret *gorm.DB
		id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "角色编号必须是数字")

		// 表单
		var rbacRoleBindAccountsForm RbacRoleBindAccountsForm
		if err := ctx.ShouldBind(&rbacRoleBindAccountsForm); err != nil {
			panic(err)
		}

		// 获取角色
		var rbacRole models.RbacRoleModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			SetWheres(tools.Map{"id": id}).
			Prepare().
			First(&rbacRole)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "角色")

		// 获取用户
		var accounts []*models.AccountModel
		ret = (&models.BaseModel{}).
			SetModel(models.AccountModel{}).
			Prepare().
			Where("uuid in ?", rbacRoleBindAccountsForm.AccountUUIDs).
			Find(&accounts)
		if len(accounts) == 0 {
			panic(errors.ThrowEmpty("用户不存在"))
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
	r.PUT(":id/bindPermissions", func(ctx *gin.Context) {
		var ret *gorm.DB
		id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "角色编号必须是数字")

		// 表单
		var form RbacRoleBindPermissionsForm
		if err := ctx.ShouldBind(&form); err != nil {
			panic(err)
		}

		// 查询角色
		var rbacRole models.RbacRoleModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			SetWheres(tools.Map{"id": id}).
			Prepare().
			First(&rbacRole)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "角色")

		// 查询权限
		var rbacPermissions []*models.RbacPermissionModel
		(&models.BaseModel{}).
			SetModel(models.RbacPermissionModel{}).
			DB().
			Where("id IN ?", form.RbacPermissionIDs).
			Find(&rbacPermissions)
		if len(rbacPermissions) == 0 {
			panic(errors.ThrowForbidden("没有找到权限"))
		}

		// 绑定
		rbacRole.RbacPermissions = rbacPermissions
		(&models.BaseModel{}).SetModel(models.RbacRoleModel{}).DB().Save(&rbacRole)

		ctx.JSON(tools.CorrectIns("绑定成功").Updated(tools.Map{}))
	})

	// 角色详情
	r.GET(":id", func(ctx *gin.Context) {
		var ret *gorm.DB
		var rbacRole models.RbacRoleModel
		id := tools.ThrowErrorWhenIsNotInt(ctx.Param("id"), "角色编号必须是数字")

		ret = (&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			SetWheres(tools.Map{"id": id}).
			SetPreloads(tools.Strings{"RbacPermissions", "Accounts"}).
			Prepare().
			First(&rbacRole)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "角色")

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
