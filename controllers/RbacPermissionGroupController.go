package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RbacPermissionGroupController struct{}

// RbacPermissionGroupStoreForm 新建权限分组表单
type RbacPermissionGroupStoreForm struct {
	Name string `form:"name" json:"name"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return RbacPermissionGroupStoreForm
func (cls RbacPermissionGroupStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionGroupStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		wrongs.PanicValidate("权限分组名称必填")
	}
	if len(cls.Name) > 64 {
		wrongs.PanicValidate("权限分组名称不能超过64位")
	}

	return cls
}

// C 新建
func (RbacPermissionGroupController) C(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.RbacPermissionGroupModel
	)

	// 表单
	form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "权限分组名称")

	// 保存
	rbacPermissionGroup := &models.RbacPermissionGroupModel{Name: form.Name}
	if ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		PrepareByDefaultDbDriver().
		Create(&rbacPermissionGroup); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}

// D 删除
func (RbacPermissionGroupController) D(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		rbacPermissionGroup models.RbacPermissionGroupModel
	)
	// 查询
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	// 删除权限分组
	if len(rbacPermissionGroup.RbacPermissions) > 0 {
		models.BootByModel(models.RbacPermissionGroupModel{}).PrepareByDefaultDbDriver().Delete(&rbacPermissionGroup.RbacPermissions)
	}

	// 删除
	if ret = models.BootByModel(models.RbacPermissionGroupModel{}).PrepareByDefaultDbDriver().Delete(&rbacPermissionGroup); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// U 编辑
func (RbacPermissionGroupController) U(ctx *gin.Context) {
	var (
		ret                         *gorm.DB
		rbacPermissionGroup, repeat models.RbacPermissionGroupModel
	)

	uuid := ctx.Param("uuid")

	// 表单
	form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"name": form.Name}).
		SetNotWheres(tools.Map{"uuid": uuid}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "权限分组名称")

	// 查询
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	// 修改
	rbacPermissionGroup.Name = form.Name
	models.BootByModel(models.RbacPermissionGroupModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Save(&rbacPermissionGroup)

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}

// S 详情
func (RbacPermissionGroupController) S(ctx *gin.Context) {
	var ret *gorm.DB
	uuid := ctx.Param("uuid")

	// 读取
	var rbacPermissionGroup models.RbacPermissionGroupModel
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": uuid}).
		SetPreloads("RbacPermissions").
		PrepareByDefaultDbDriver().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}

// I 列表
func (RbacPermissionGroupController) I(ctx *gin.Context) {
	var (
		rbacPermissionGroups []models.RbacPermissionGroupModel
		count                int64
		db                   *gorm.DB
	)
	db = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetPreloads("RbacPermissionGroup").
		SetWhereFields("name", "uri", "method", "rbac_permission_group_uuid").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&rbacPermissionGroups)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_permission_groups": rbacPermissionGroups}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&rbacPermissionGroups)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"rbac_permission_groups": rbacPermissionGroups}, ctx.Query("__page__"), count))
	}
}
