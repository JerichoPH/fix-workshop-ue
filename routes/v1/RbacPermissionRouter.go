package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// RbacPermissionRouter 权限路由
type RbacPermissionRouter struct{}

// RbacPermissionStoreForm 新建权限表单
type RbacPermissionStoreForm struct {
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
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("名称必填"))
	}
	if cls.Uri == "" {
		panic(exceptions.ThrowForbidden("URI必填"))
	}
	if cls.Method == "" {
		panic(exceptions.ThrowForbidden("访问方法必选"))
	}
	if cls.RbacPermissionGroupUUID == "" {
		panic(exceptions.ThrowForbidden("所属权限分组必选"))
	}
	ret = models.Init(&models.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": cls.RbacPermissionGroupUUID}).
		Prepare().
		First(&cls.RbacPermissionGroup)
	tools.ThrowExceptionWhenIsEmptyByDB(ret, "所属权限分组")

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
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Uri == "" {
		panic(exceptions.ThrowForbidden("URI必填"))
	}
	if cls.RbacPermissionGroupUUID == "" {
		panic(exceptions.ThrowForbidden("权限分组必选"))
	}

	return cls
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
			// 表单
			form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

			// 保存
			models.Init(&models.RbacPermissionModel{}).
				DB().
				Create(&models.RbacPermissionModel{
					BaseModel:               models.BaseModel{},
					Name:                    form.Name,
					URI:                     form.Uri,
					Method:                  form.Method,
					RbacPermissionGroupUUID: form.RbacPermissionGroup.UUID,
				})

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 批量添加资源权限
		r.POST("resource", func(ctx *gin.Context) {
			var ret *gorm.DB
			resourceRbacPermission := map[string]string{"列表": "GET", "新建页面": "GET", "新建": "POST", "详情页面": "GET", "编辑页面": "GET", "编辑": "PUT", "删除": "DELETE"}

			// 表单
			form := (&RbacPermissionStoreResourceForm{}).ShouldBind(ctx)

			// 批量新建
			successCount := 0
			for name, method := range resourceRbacPermission {
				// 如果不重复则新建
				var repeat models.RbacPermissionModel
				ret = models.Init(models.RbacPermissionModel{}).
					SetWheres(tools.Map{"name": name, "method": method, "uri": form.Uri}).
					Prepare().
					First(&repeat)
				if !tools.ThrowExceptionWhenIsEmptyByDB(ret, "") {
					if ret = models.Init(models.RbacPermissionModel{}).

						DB().
						Create(&models.RbacPermissionModel{
							Name:                    name,
							URI:                     form.Uri,
							Method:                  method,
							RbacPermissionGroupUUID: form.RbacPermissionGroupUUID,
						}); ret.Error != nil {
						panic(exceptions.ThrowForbidden("批量添加资源权限时错误：" + ret.Error.Error()))
					} else {
						successCount += 1
					}
				}
			}

			ctx.JSON(tools.CorrectIns("成功添加权限：" + strconv.Itoa(successCount) + "个").Created(tools.Map{}))
		})

		// 删除权限
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")
			var rbacPermission models.RbacPermissionModel

			// 查询
			ret = models.Init(models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				Find(&rbacPermission)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限")

			// 删除
			models.Init(&models.RbacPermissionModel{}).
				DB().
				Delete(&rbacPermission)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

			// 查询
			var rbacPermission models.RbacPermissionModel
			ret = models.Init(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacPermission)

			// 修改
			rbacPermission.Name = form.Name
			rbacPermission.URI = form.Uri
			rbacPermission.Method = form.Method
			rbacPermission.RbacPermissionGroupUUID = form.RbacPermissionGroupUUID
			// 保存
			if ret = models.Init(models.RbacPermissionModel{}).
				DB().
				Save(&rbacPermission); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 权限详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var rbacPermission models.RbacPermissionModel
			ret = models.Init(&models.RbacPermissionModel{}).
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
			models.Init(models.RbacPermissionModel{}).
				SetPreloads(tools.Strings{"RbacPermissionGroup"}).
				SetWhereFields("name", "uri", "method", "rbac_permission_group_uuid").
				PrepareQuery(ctx).
				Find(&rbacPermissions)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permissions": rbacPermissions}))
		})
	}

}
