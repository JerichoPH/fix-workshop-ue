package v1

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/RbacModels"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RbacPermissionGroupRouter 权限分组路由
type RbacPermissionGroupRouter struct{}

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
		abnormals.PanicValidate(err.Error())
	}
	if cls.Name == "" {
		abnormals.PanicValidate("名称必填")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *RbacPermissionGroupRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/rbacPermissionGroup",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 创建权限分组
		r.POST("", func(ctx *gin.Context) {
			var (
				ret    *gorm.DB
				repeat RbacModels.RbacPermissionGroupModel
			)

			// 表单
			form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "权限分组名称")

			// 保存
			rbacPermissionGroup := &RbacModels.RbacPermissionGroupModel{Name: form.Name}
			if ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				DB().
				Create(&rbacPermissionGroup); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
		})

		// 删除用户分组
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var (
				ret                 *gorm.DB
				rbacPermissionGroup RbacModels.RbacPermissionGroupModel
			)
			// 查询
			ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&rbacPermissionGroup)
			abnormals.PanicWhenIsEmpty(ret, "权限分组")

			// 删除权限分组
			if len(rbacPermissionGroup.RbacPermissions) > 0 {
				models.Init(RbacModels.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup.RbacPermissions)
			}

			// 删除
			if ret = models.Init(RbacModels.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup); ret.Error != nil {
				abnormals.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限分组
		r.PUT(":uuid", func(ctx *gin.Context) {
			var (
				ret                         *gorm.DB
				rbacPermissionGroup, repeat RbacModels.RbacPermissionGroupModel
			)

			uuid := ctx.Param("uuid")

			// 表单
			form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

			// 查重
			ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			abnormals.PanicWhenIsRepeat(ret, "权限分组名称")

			// 查询
			ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&rbacPermissionGroup)
			abnormals.PanicWhenIsEmpty(ret, "权限分组")

			// 修改
			rbacPermissionGroup.Name = form.Name
			models.Init(RbacModels.RbacPermissionGroupModel{}).DB().Save(&rbacPermissionGroup)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
		})

		// 权限分组详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 读取
			var rbacPermissionGroup RbacModels.RbacPermissionGroupModel
			ret = models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				SetPreloads("RbacPermissions").
				Prepare().
				First(&rbacPermissionGroup)
			abnormals.PanicWhenIsEmpty(ret, "权限分组")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
		})

		// 权限分组列表
		r.GET("", func(ctx *gin.Context) {
			var rbacPermissionGroups []RbacModels.RbacPermissionGroupModel
			models.Init(RbacModels.RbacPermissionGroupModel{}).
				SetPreloads("RbacPermissions").
				PrepareQuery(ctx).
				Find(&rbacPermissionGroups)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_groups": rbacPermissionGroups}))
		})
	}
}
