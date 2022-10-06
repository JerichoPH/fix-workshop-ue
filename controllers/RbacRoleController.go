package controllers

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
)

type RbacRoleController struct{}

// RbacRoleStoreForm 创建角色表单
type RbacRoleStoreForm struct {
	Sort int64  `form:"sort" json:"sort"`
	Name string `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return RbacRoleStoreForm
func (ins RbacRoleStoreForm) ShouldBind(ctx *gin.Context) RbacRoleStoreForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.Name == "" {
		wrongs.PanicValidate("角色名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("角色名称不能超过64位")
	}

	return ins
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
//  @receiver ins
//  @param ctx
//  @return RbacRoleBindAccountsForm
func (ins RbacRoleBindAccountsForm) ShouldBind(ctx *gin.Context) RbacRoleBindAccountsForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicForbidden(err.Error())
	}

	if len(ins.AccountUUIDs) > 0 {
		models.BootByModel(models.AccountModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.AccountUUIDs).
			Find(&ins.Accounts)
	}

	return ins
}

// RbacRoleBindPermissionsForm 角色绑定权限表单
type RbacRoleBindPermissionsForm struct {
	RbacPermissionUuids []string `form:"rbac_permission_uuids" json:"rbac_permission_uuids"`
	RbacPermissions     []*models.RbacPermissionModel
}

// ShouldBind 表单绑定
//  @receiver ins
//  @param ctx
//  @return RbacRoleBindPermissionsForm
func (ins RbacRoleBindPermissionsForm) ShouldBind(ctx *gin.Context) RbacRoleBindPermissionsForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicForbidden(err.Error())
	}

	if len(ins.RbacPermissionUuids) > 0 {
		models.BootByModel(models.RbacPermissionModel{}).
			PrepareByDefaultDbDriver().
			Where("uuid in ?", ins.RbacPermissionUuids).
			Find(&ins.RbacPermissions)
	}

	return ins
}

// RbacRoleBindPermissionsByRbacPermissionGroupForm 角色绑定权限表单（根据权限分组）
type RbacRoleBindPermissionsByRbacPermissionGroupForm struct {
	RbacPermissionUuids     []string `form:"rbac_permission_uuids" json:"rbac_permission_uuids"`
	RbacPermissions         []*models.RbacPermissionModel
	RbacPermissionGroupUuid string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     models.RbacPermissionGroupModel
}

// ShouldBind 绑定表单
func (ins RbacRoleBindPermissionsByRbacPermissionGroupForm) ShouldBind(ctx *gin.Context) RbacRoleBindPermissionsByRbacPermissionGroupForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}

	if ins.RbacPermissionGroupUuid == "" {
		wrongs.PanicValidate("所属权限分组必选")
	}
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).SetWheres(tools.Map{"uuid": ins.RbacPermissionGroupUuid}).PrepareByDefaultDbDriver().First(&ins.RbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	if len(ins.RbacPermissionUuids) > 0 {
		models.BootByModel(models.RbacPermissionModel{}).PrepareByDefaultDbDriver().Where("uuid in ?", ins.RbacPermissionUuids).Find(&ins.RbacPermissions)
	}

	return ins
}

// N 新建
func (RbacRoleController) N(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat RbacRoleStoreForm
	)

	// 表单
	form := (&RbacRoleStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "角色名称")

	// 新建
	rbacRole := &models.RbacRoleModel{
		BaseModel: models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		Name:      form.Name,
	}
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefaultDbDriver().Create(rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_role": rbacRole}))
}

// R 删除
func (RbacRoleController) R(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleModel
	)
	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除
	if ret = models.BootByModel(models.RbacRoleModel{}).PrepareByDefaultDbDriver().Delete(&rbacRole); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 编辑
func (RbacRoleController) E(ctx *gin.Context) {
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "角色名称")

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 修改
	rbacRole.Name = form.Name
	models.BootByModel(models.RbacRoleModel{}).PrepareByDefaultDbDriver().Save(&rbacRole)

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_role": rbacRole}))
}

// PutBindAccounts 角色绑定用户
func (RbacRoleController) PutBindAccounts(ctx *gin.Context) {
	var (
		ret          *gorm.DB
		rbacRole     models.RbacRoleModel
		successCount uint64
	)

	// 表单
	form := (&RbacRoleBindAccountsForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除原有绑定关系
	new(databases.Launcher).GetDatabaseConn().Exec("delete from pivot_rbac_role_and_accounts where rbac_role_id = ?", rbacRole.Id)

	// 绑定角色与用户
	if len(form.Accounts) > 0 {
		for _, account := range form.Accounts {
			pivotRbacRoleAndAccount := models.PivotRbacRoleAndAccountModel{
				RbacRoleUuid: rbacRole.Uuid,
				AccountUuid:  account.Uuid,
			}

			models.BootByModel(models.PivotRbacRoleAndAccountModel{}).
				PrepareByDefaultDbDriver().
				Create(&pivotRbacRoleAndAccount)
			successCount++
		}
	}

	ctx.JSON(tools.CorrectBoot("绑定成功" + strconv.Itoa(int(successCount)) + "项").Updated(tools.Map{}))
}

// PutBindRbacPermissionsByRbacPermissionGroup 角色绑定权限（根据权限分组）
func (RbacRoleController) PutBindRbacPermissionsByRbacPermissionGroup(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		rbacRole          models.RbacRoleModel
		successCount      uint64
		rbacPermissions   []models.RbacPermissionModel
		rbacPermissionIds []uint64
	)

	// 表单
	form := (&RbacRoleBindPermissionsByRbacPermissionGroupForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除原有绑定关系
	// 1、查找现有权限分组下所有权限id
	models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"rbac_permission_group_uuid": form.RbacPermissionGroup.Uuid}).
		PrepareByDefaultDbDriver().
		Find(&rbacPermissions)
	if len(rbacPermissions) > 0 {
		for _, rbacPermission := range rbacPermissions {
			rbacPermissionIds = append(rbacPermissionIds, rbacPermission.Id)
		}
	}

	new(databases.Launcher).GetDatabaseConn().Exec(`
delete from pivot_rbac_role_and_rbac_permissions
		where rbac_role_id = ?
		and rbac_permission_id in ?
`, rbacRole.Id, rbacPermissionIds)

	// 绑定角色与权限关系
	if len(form.RbacPermissions) > 0 {
		for _, rbacPermission := range form.RbacPermissions {
			pivotRbacRoleAndPermission := models.PivotRbacRoleAndPermissionModel{
				RbacRoleUuid:       rbacRole.Uuid,
				RbacPermissionUuid: rbacPermission.Uuid,
			}

			models.BootByModel(models.PivotRbacRoleAndPermissionModel{}).
				PrepareByDefaultDbDriver().
				Create(&pivotRbacRoleAndPermission)
			successCount++
		}
	}

	ctx.JSON(tools.CorrectBoot("绑定成功" + strconv.Itoa(int(successCount)) + "项").Updated(tools.Map{}))
}

// PutBindRbacPermissions 角色绑定权限
func (RbacRoleController) PutBindRbacPermissions(ctx *gin.Context) {
	var (
		ret          *gorm.DB
		rbacRole     models.RbacRoleModel
		successCount uint64
	)

	// 表单
	form := (&RbacRoleBindPermissionsForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacRoleModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	// 删除原有绑定关系
	new(databases.Launcher).GetDatabaseConn().Exec("delete from pivot_rbac_role_and_rbac_permissions where rbac_role_id = ?", rbacRole.Id)

	// 绑定角色与权限关系
	if len(form.RbacPermissions) > 0 {
		for _, rbacPermission := range form.RbacPermissions {
			pivotRbacRoleAndPermission := models.PivotRbacRoleAndPermissionModel{
				RbacRoleUuid:       rbacRole.Uuid,
				RbacPermissionUuid: rbacPermission.Uuid,
			}

			models.BootByModel(models.PivotRbacRoleAndPermissionModel{}).
				PrepareByDefaultDbDriver().
				Create(&pivotRbacRoleAndPermission)
			successCount++
		}
	}

	ctx.JSON(tools.CorrectBoot("绑定成功" + strconv.Itoa(int(successCount)) + "项").Updated(tools.Map{}))
}

// D 详情
func (RbacRoleController) D(ctx *gin.Context) {
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
		PrepareByDefaultDbDriver().
		First(&rbacRole)
	wrongs.PanicWhenIsEmpty(ret, "角色")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_role": rbacRole}))
}

// L 列表
func (RbacRoleController) L(ctx *gin.Context) {
	var (
		rbacRoles []models.RbacRoleModel
		count     int64
		db        *gorm.DB
	)
	db = models.BootByModel(models.RbacRoleModel{}).
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&rbacRoles)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_roles": rbacRoles}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&rbacRoles)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"rbac_roles": rbacRoles}, ctx.Query("__page__"), count))
	}
}
