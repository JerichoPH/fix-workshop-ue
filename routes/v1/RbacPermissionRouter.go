package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
)

// RbacPermissionRouter 权限路由
type RbacPermissionRouter struct{}

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
	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.BombForbidden(err.Error())
	}
	if cls.Name == "" {
		abnormals.BombForbidden("名称必填")
	}
	if cls.Uri == "" {
		abnormals.BombForbidden("URI必填")
	}
	if cls.Method == "" {
		abnormals.BombForbidden("访问方法必选")
	}
	if cls.RbacPermissionGroupUUID == "" {
		abnormals.BombForbidden("所属权限分组必选")
	}
	cls.RbacPermissionGroup = (&models.RbacPermissionGroupModel{}).FindOneByUUID(cls.RbacPermissionGroupUUID)

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
		abnormals.BombForbidden(err.Error())
	}
	if cls.Uri == "" {
		abnormals.BombForbidden("URI必填")
	}
	if cls.RbacPermissionGroupUUID == "" {
		abnormals.BombForbidden("权限分组必选")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *RbacPermissionRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/rbacPermission",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

			// 新建
			rbacPermission := &models.RbacPermissionModel{
				BaseModel:               models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				Name:                    form.Name,
				URI:                     form.Uri,
				Method:                  form.Method,
				RbacPermissionGroupUUID: form.RbacPermissionGroup.UUID,
			}
			if ret = models.Init(&models.RbacPermissionModel{}).
				DB().
				Create(&rbacPermission); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"rbac_permission": rbacPermission}))
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
				if !abnormals.BombWhenIsEmptyByDB(ret, "") {
					if ret = models.Init(models.RbacPermissionModel{}).
						DB().
						Create(&models.RbacPermissionModel{
							BaseModel:               models.BaseModel{UUID: uuid.NewV4().String()},
							Name:                    name,
							URI:                     form.Uri,
							Method:                  method,
							RbacPermissionGroupUUID: form.RbacPermissionGroupUUID,
						}); ret.Error != nil {
						abnormals.BombForbidden("批量添加资源权限时错误：" + ret.Error.Error())
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

			// 查询
			rbacPermission := (&models.RbacPermissionModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 删除
			if ret = models.Init(&models.RbacPermissionModel{}).
				DB().
				Delete(&rbacPermission); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

			// 查询
			rbacPermission := (&models.RbacPermissionModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 修改
			rbacPermission.Name = form.Name
			rbacPermission.URI = form.Uri
			rbacPermission.Method = form.Method
			rbacPermission.RbacPermissionGroupUUID = form.RbacPermissionGroupUUID
			if ret = models.Init(models.RbacPermissionModel{}).
				DB().
				Save(&rbacPermission); ret.Error != nil {
				abnormals.BombForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"rbac_permission": rbacPermission}))
		})

		// 权限详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			var rbacPermission models.RbacPermissionModel
			ret = models.Init(&models.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetPreloads(tools.Strings{"RbacPermissionGroup"}).
				Prepare().
				First(&rbacPermission)
			abnormals.BombWhenIsEmptyByDB(ret, "权限")

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
