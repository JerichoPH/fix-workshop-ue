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
	RbacPermissionGroupUUID string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     models.RbacPermissionGroupModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return RbacPermissionStoreForm
func (cls RbacPermissionStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		wrongs.PanicValidate("名称必填")
	}
	if cls.Uri == "" {
		wrongs.PanicValidate("URI必填")
	}
	if cls.Method == "" {
		wrongs.PanicValidate("访问方法必选")
	}
	if cls.RbacPermissionGroupUUID == "" {
		wrongs.PanicValidate("所属权限分组必选")
	}
	ret = models.BootByModel(models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&cls.RbacPermissionGroup)
	wrongs.PanicWhenIsEmpty(ret, "权限分组")

	return cls
}

// RbacPermissionStoreResourceForm 批量添加资源权限
type RbacPermissionStoreResourceForm struct {
	Uri                     string `form:"uri" json:"uri"`
	RbacPermissionGroupUUID string `form:"rbac_permission_group_uuid" json:"rbac_permission_group_uuid"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return RbacPermissionStoreResourceForm
func (cls RbacPermissionStoreResourceForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreResourceForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicForbidden(err.Error())
	}
	if cls.Uri == "" {
		wrongs.PanicForbidden("URI必填")
	}
	if cls.RbacPermissionGroupUUID == "" {
		wrongs.PanicForbidden("权限分组必选")
	}

	return cls
}

func(RbacPermissionController) C(ctx *gin.Context){
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
		PrepareByDefault().
		Create(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Created(tools.Map{"rbac_permission": rbacPermission}))
}
func(RbacPermissionController) PostResource(ctx *gin.Context){
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
			PrepareByDefault().
			First(&repeat)
		if !wrongs.PanicWhenIsEmpty(ret, "") {
			if ret = models.BootByModel(models.RbacPermissionModel{}).
				PrepareByDefault().
				Create(&models.RbacPermissionModel{
					BaseModel:               models.BaseModel{Uuid: uuid.NewV4().String()},
					Name:                    name,
					URI:                     form.Uri,
					Method:                  method,
					RbacPermissionGroupUuid: form.RbacPermissionGroupUUID,
				}); ret.Error != nil {
				wrongs.PanicForbidden("批量添加资源权限时错误：" + ret.Error.Error())
			} else {
				successCount += 1
			}
		}
	}

	ctx.JSON(tools.CorrectBoot("成功添加权限：" + strconv.Itoa(successCount) + "个").Created(tools.Map{}))
}
func(RbacPermissionController) D(ctx *gin.Context){
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)
	// 查询
	ret = models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	// 删除
	if ret = models.BootByModel(&models.RbacPermissionModel{}).
		PrepareByDefault().
		Delete(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}
func(RbacPermissionController) U(ctx *gin.Context){
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)

	// 表单
	form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

	// 查询
	ret = models.BootByModel(models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	// 修改
	rbacPermission.Name = form.Name
	rbacPermission.URI = form.Uri
	rbacPermission.Method = form.Method
	rbacPermission.RbacPermissionGroupUuid = form.RbacPermissionGroupUUID
	if ret = models.BootByModel(models.RbacPermissionModel{}).
		PrepareByDefault().
		Save(&rbacPermission); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{"rbac_permission": rbacPermission}))
}
func(RbacPermissionController) S(ctx *gin.Context){
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionModel
	)

	ret = models.BootByModel(&models.RbacPermissionModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetPreloads("RbacPermissionGroup").
		PrepareByDefault().
		First(&rbacPermission)
	wrongs.PanicWhenIsEmpty(ret, "权限")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_permission": rbacPermission}))
}
func(RbacPermissionController) I(ctx *gin.Context){
	var rbacPermissions []models.RbacPermissionModel
	models.BootByModel(models.RbacPermissionModel{}).
		SetPreloads("RbacPermissionGroup").
		SetWhereFields("name", "uri", "method", "rbac_permission_group_uuid").
		PrepareQuery(ctx,"").
		Find(&rbacPermissions)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"rbac_permissions": rbacPermissions}))
}