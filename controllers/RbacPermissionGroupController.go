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
		wrongs.PanicValidate("名称必填")
	}

	return cls
}

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
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "权限分组名称")

	// 保存
	rbacPermissionGroup := &models.RbacPermissionGroupModel{Name: form.Name}
	if ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		PrepareByDefault().
		Create(&rbacPermissionGroup); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}
func (RbacPermissionGroupController) D(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		rbacPermissionGroup models.RbacPermissionGroupModel
	)
	// 查询
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	// 删除权限分组
	if len(rbacPermissionGroup.RbacPermissions) > 0 {
		models.BootByModel(models.RbacPermissionGroupModel{}).PrepareByDefault().Delete(&rbacPermissionGroup.RbacPermissions)
	}

	// 删除
	if ret = models.BootByModel(models.RbacPermissionGroupModel{}).PrepareByDefault().Delete(&rbacPermissionGroup); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
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
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "权限分组名称")

	// 查询
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	// 修改
	rbacPermissionGroup.Name = form.Name
	models.BootByModel(models.RbacPermissionGroupModel{}).SetWheres(tools.Map{"uuid":ctx.Param("uuid")}).PrepareByDefault().Save(&rbacPermissionGroup)

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}
func (RbacPermissionGroupController) S(ctx *gin.Context) {
	var ret *gorm.DB
	uuid := ctx.Param("uuid")

	// 读取
	var rbacPermissionGroup models.RbacPermissionGroupModel
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": uuid}).
		SetPreloads("RbacPermissions").
		PrepareByDefault().
		First(&rbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
}
func (RbacPermissionGroupController) I(ctx *gin.Context) {
	var rbacPermissionGroups []models.RbacPermissionGroupModel
	models.BootByModel(models.RbacPermissionGroupModel{}).
		SetPreloads("RbacPermissions").
		PrepareQuery(ctx, "").
		Find(&rbacPermissionGroups)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_permission_groups": rbacPermissionGroups}))
}
