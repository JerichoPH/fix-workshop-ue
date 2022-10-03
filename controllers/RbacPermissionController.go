package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
)

type RbacPermissionController struct{}

// RbacPermissionStoreForm 新建权限表单
type RbacPermissionStoreForm struct {
	Sort                    int64  `form:"sort" json:"sort"`
	Name                    string `form:"name" json:"name"`
	Uri                     string `form:"uri" json:"uri"`
	Method                  string `form:"method" json:"method"`
	RbacPermissionGroupUuid string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     models.RbacPermissionGroupModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return RbacPermissionStoreForm
func (ins RbacPermissionStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.Name == "" {
		wrongs.PanicValidate("权限名称必填")
	}
	if len(ins.Name) > 64 {
		wrongs.PanicValidate("权限名称不能超过64位")
	}
	if ins.Uri == "" {
		wrongs.PanicValidate("URI必填")
	}
	if ins.Method == "" {
		wrongs.PanicValidate("访问方法必选")
	}
	if ins.RbacPermissionGroupUuid == "" {
		wrongs.PanicValidate("所属权限分组必选")
	}
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ins.RbacPermissionGroupUuid}).
		PrepareByDefaultDbDriver().
		First(&ins.RbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	return ins
}

// RbacPermissionStoreResourceForm 批量添加资源权限
type RbacPermissionStoreResourceForm struct {
	Uri                     string `form:"uri" json:"uri"`
	RbacPermissionGroupUuid string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return RbacPermissionStoreResourceForm
func (ins RbacPermissionStoreResourceForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreResourceForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicForbidden(err.Error())
	}
	if ins.Uri == "" {
		wrongs.PanicForbidden("URI必填")
	}
	if ins.RbacPermissionGroupUuid == "" {
		wrongs.PanicForbidden("权限分组必选")
	}

	return ins
}

func (RbacPermissionController) N(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

	// 新建
	rbacPermission := &models.RbacPermissionModel{
		BaseModel:               models.BaseModel{Sort: form.Sort, Uuid: uuid.NewV4().String()},
		Name:                    form.Name,
		URI:                     form.Uri,
		Method:                  form.Method,
		RbacPermissionGroupUuid: form.RbacPermissionGroup.Uuid,
	}
	if ret = models.BootByModel(&models.RbacPermissionModel{}).
		PrepareByDefaultDbDriver().
		Create(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_permission": rbacPermission}))
}
func (RbacPermissionController) PostResource(ctx *gin.Context) {
	var ret *gorm.DB
	resourceRbacPermission := map[string]string{"列表": "GET", "新建页面": "GET", "新建": "POST", "详情页面": "GET", "编辑页面": "GET", "编辑": "PUT", "删除": "DELETE"}

	// 表单
	form := (&RbacPermissionStoreResourceForm{}).ShouldBind(ctx)

	// 批量新建
	successCount := 0
	for name, method := range resourceRbacPermission {
		// 如果不重复则新建
		var repeat models.RbacPermissionModel
		ret = models.BootByModel(models.RbacPermissionModel{}).
			SetWheres(tools.Map{"name": name, "method": method, "uri": form.Uri}).
			PrepareByDefaultDbDriver().
			First(&repeat)
		if !wrongs.PanicWhenIsEmpty(ret, "") {
			if ret = models.BootByModel(models.RbacPermissionModel{}).
				PrepareByDefaultDbDriver().
				Create(&models.RbacPermissionModel{
					BaseModel:               models.BaseModel{Uuid: uuid.NewV4().String()},
					Name:                    name,
					URI:                     form.Uri,
					Method:                  method,
					RbacPermissionGroupUuid: form.RbacPermissionGroupUuid,
				}); ret.Error != nil {
				wrongs.PanicForbidden("批量添加资源权限时错误：" + ret.Error.Error())
			} else {
				successCount += 1
			}
		}
	}

	ctx.JSON(tools.CorrectBoot("成功添加权限：" + strconv.Itoa(successCount) + "个").Created(tools.Map{}))
}
func (RbacPermissionController) R(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)
	// 查询
	ret = models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	// 删除
	if ret = models.BootByModel(&models.RbacPermissionModel{}).
		PrepareByDefaultDbDriver().
		Delete(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func (RbacPermissionController) E(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)

	// 表单
	form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	// 修改
	rbacPermission.Name = form.Name
	rbacPermission.URI = form.Uri
	rbacPermission.Method = form.Method
	rbacPermission.RbacPermissionGroupUuid = form.RbacPermissionGroupUuid
	if ret = models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		Save(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_permission": rbacPermission}))
}

// D 详情
func (RbacPermissionController) D(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)

	ret = models.BootByModel(&models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetPreloads("RbacPermissionGroup").
		PrepareByDefaultDbDriver().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_permission": rbacPermission}))
}

// L 列表
func (RbacPermissionController) L(ctx *gin.Context) {
	var (
		rbacPermissions []models.RbacPermissionModel
		count           int64
		db              *gorm.DB
	)
	db = models.BootByModel(models.RbacPermissionModel{}).
		SetPreloads("RbacPermissionGroup").
		SetWhereFields("name", "uri", "method", "rbac_permission_group_uuid").
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&rbacPermissions)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"rbac_permissions": rbacPermissions}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&rbacPermissions)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"rbac_permissions": rbacPermissions}, ctx.Query("__page__"), count))
	}
}
