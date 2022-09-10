package controllers

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RbacRoleController struct{}

// RbacRoleStoreForm 创建角色表单
type RbacRoleStoreForm struct {
	Sort int64  `form:"sort" json:"sort"`
	Name string `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return RbacRoleStoreForm
func (cls RbacRoleStoreForm) ShouldBind(ctx *gin.Context) RbacRoleStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		wrongs.PanicValidate("名称必填")
	}

	return cls
}

// RbacRoleUpdateForm 编辑角色表单
type RbacRoleUpdateForm struct {
	Name string `form:"name" json:"name"`
}

// RbacRoleBindAccountsForm 角色绑定用户表单
type RbacRoleBindAccountsForm struct {
	AccountUUIDs []string `form:"account_uuids[]" json:"account_uuids"`
	Accounts     []*models.AccountModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return RbacRoleBindAccountsForm
func (cls RbacRoleBindAccountsForm) ShouldBind(ctx *gin.Context) RbacRoleBindAccountsForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicForbidden(err.Error())
	}

	if len(cls.AccountUUIDs) > 0 {
		models.BootByModel(models.AccountModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.AccountUUIDs).
			Find(&cls.Accounts)
	}

	return cls
}

// RbacRoleBindPermissionsForm 角色绑定权限表单
type RbacRoleBindPermissionsForm struct {
	RbacPermissionUUIDs []string `form:"rbac_permission_uuids" json:"rbac_permission_uuids"`
	RbacPermissions     []*models.RbacPermissionModel
}

// ShouldBind 表单绑定
//  @receiver cls
//  @param ctx
//  @return RbacRoleBindPermissionsForm
func (cls RbacRoleBindPermissionsForm) ShouldBind(ctx *gin.Context) RbacRoleBindPermissionsForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicForbidden(err.Error())
	}

	if len(cls.RbacPermissionUUIDs) > 0 {
		models.BootByModel(models.RbacPermissionModel{}).
			PrepareByDefault().
			Where("uuid in ?", cls.RbacPermissionUUIDs).
			Find(&cls.RbacPermissions)
	}

	return cls
}

func (RbacRoleController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat RbacRoleStoreForm
	)

	// 表单
	form := (&RbacRoleStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "角色名称")

	// 新建
	rbacRole := &models.RbacRoleModel{
		BaseModel: models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		Name:      form.Name,
	}
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Create(rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_role": rbacRole}))
}
func (RbacRoleController) D(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleModel
	)
	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Delete(&rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (RbacRoleController) U(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		rbacRole, repeat models.RbacRoleModel
	)

	// 表单
	form := (&RbacRoleStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "角色名称")

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 修改
	rbacRole.Name = form.Name
	models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Save(&rbacRole)

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_role": rbacRole}))
}
func (RbacRoleController) PutBindAccounts(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleModel
	)

	// 表单
	form := (&RbacRoleBindAccountsForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除原有绑定关系
	new(databases.Launcher).GetDatabaseConn().Exec("delete from pivot_rbac_role_and_rbac_permissions where rbac_role_id = ?", rbacRole.Id)

	// 绑定
	rbacRole.Accounts = form.Accounts
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Save(&rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("绑定成功").Updated(tools.Map{}))
}
func (RbacRoleController) PutBindRbacPermissions(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleModel
	)

	// 表单
	form := (&RbacRoleBindPermissionsForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除原有绑定关系
	new(databases.Launcher).GetDatabaseConn().Exec("delete from pivot_rbac_role_and_rbac_permissions where rbac_role_id = ?", rbacRole.Id)

	// 绑定
	rbacRole.RbacPermissions = form.RbacPermissions
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefault().Save(&rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("绑定成功").Updated(tools.Map{}))
}
func (RbacRoleController) S(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleModel
	)

	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetPreloads(
			"RbacPermissions",
			"RbacPermissions.RbacPermissionGroup",
			"Accounts",
			"Menus",
		).
		PrepareByDefault().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_role": rbacRole}))
}
func (RbacRoleController) I(ctx *gin.Context) {
	var rbacRoles []models.RbacRoleModel
	models.BootByModel(models.RbacRoleModel{}).
		PrepareUseQueryByDefault(ctx).
		Find(&rbacRoles)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_roles": rbacRoles}))
}
