package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/RbacModels"
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
	RbacPermissionGroup     RbacModels.RbacPermissionGroupModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return RbacPermissionStoreForm
func (cls RbacPermissionStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		abnormals.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		abnormals.PanicValidate("名称必填")
	}
	if cls.Uri == "" {
		abnormals.PanicValidate("URI必填")
	}
	if cls.Method == "" {
		abnormals.PanicValidate("访问方法必选")
	}
	if cls.RbacPermissionGroupUUID == "" {
		abnormals.PanicValidate("所属权限分组必选")
	}
	ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		Prepare().
		First(&cls.RbacPermissionGroup)
	abnormals.PanicWhenIsEmpty(ret, "权限分组")

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
		abnormals.PanicForbidden(err.Error())
	}
	if cls.Uri == "" {
		abnormals.PanicForbidden("URI必填")
	}
	if cls.RbacPermissionGroupUUID == "" {
		abnormals.PanicForbidden("权限分组必选")
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
			rbacPermission := &RbacModels.RbacPermissionModel{
				BaseModel:               models.BaseModel{Sort: form.Sort, UUID: uuid.NewV4().String()},
				Name:                    form.Name,
				URI:                     form.Uri,
				Method:                  form.Method,
				RbacPermissionGroupUUID: form.RbacPermissionGroup.UUID,
			}
			if ret = models.Init(&RbacModels.RbacPermissionModel{}).
				DB().
				Create(&rbacPermission); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
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
				var repeat RbacModels.RbacPermissionModel
				ret = models.Init(RbacModels.RbacPermissionModel{}).
					SetWheres(tools.Map{"name": name, "method": method, "uri": form.Uri}).
					Prepare().
					First(&repeat)
				if !abnormals.PanicWhenIsEmpty(ret, "") {
					if ret = models.Init(RbacModels.RbacPermissionModel{}).
						DB().
						Create(&RbacModels.RbacPermissionModel{
							BaseModel:               models.BaseModel{UUID: uuid.NewV4().String()},
							Name:                    name,
							URI:                     form.Uri,
							Method:                  method,
							RbacPermissionGroupUUID: form.RbacPermissionGroupUUID,
						}); ret.Error != nil {
						abnormals.PanicForbidden("批量添加资源权限时错误：" + ret.Error.Error())
					} else {
						successCount += 1
					}
				}
			}

			ctx.JSON(tools.CorrectIns("成功添加权限：" + strconv.Itoa(successCount) + "个").Created(tools.Map{}))
		})

		// 删除权限
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				rbacPermission RbacModels.RbacPermissionModel
			)
			// 查询
			ret = models.Init(RbacModels.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&rbacPermission)
			abnormals.PanicWhenIsEmpty(ret, "权限")

			// 删除
			if ret = models.Init(&RbacModels.RbacPermissionModel{}).
				DB().
				Delete(&rbacPermission); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				rbacPermission RbacModels.RbacPermissionModel
			)

			// 表单
			form := (&RbacPermissionStoreForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(RbacModels.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&rbacPermission)
			abnormals.PanicWhenIsEmpty(ret, "权限")

			// 修改
			rbacPermission.Name = form.Name
			rbacPermission.URI = form.Uri
			rbacPermission.Method = form.Method
			rbacPermission.RbacPermissionGroupUUID = form.RbacPermissionGroupUUID
			if ret = models.Init(RbacModels.RbacPermissionModel{}).
				DB().
				Save(&rbacPermission); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"rbac_permission": rbacPermission}))
		})

		// 权限详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var (
				ret            *gorm.DB
				rbacPermission RbacModels.RbacPermissionModel
			)

			ret = models.Init(&RbacModels.RbacPermissionModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				SetPreloads("RbacPermissionGroup").
				Prepare().
				First(&rbacPermission)
			abnormals.PanicWhenIsEmpty(ret, "权限")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission": rbacPermission}))
		})

		// 权限列表
		r.GET("", func(ctx *gin.Context) {
			var rbacPermissions []RbacModels.RbacPermissionModel
			models.Init(RbacModels.RbacPermissionModel{}).
				SetPreloads("RbacPermissionGroup").
				SetWhereFields("name", "uri", "method", "rbac_permission_group_uuid").
				PrepareQuery(ctx).
				Find(&rbacPermissions)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permissions": rbacPermissions}))
		})
	}

}
