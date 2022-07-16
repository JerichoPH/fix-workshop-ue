package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RbacPermissionGroupRouter 权限分组路由
type RbacPermissionGroupRouter struct{}

// RbacPermissionGroupStoreForm 创建权限分组表单
type RbacPermissionGroupStoreForm struct {
	Name string `form:"name" json:"name"`
}

func (cls RbacPermissionGroupStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionGroupStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("名称必填"))
	}

	return cls
}

func (cls *RbacPermissionGroupRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/rbacPermissionGroup",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 创建权限分组
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.RbacPermissionGroupModel
			ret = models.Init(models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "权限分组名称")

			// 保存
			var rbacPermissionGroup models.RbacPermissionGroupModel
			rbacPermissionGroup.Name = form.Name
			if ret = models.Init(models.RbacPermissionGroupModel{}).
				DB().
				Create(&rbacPermissionGroup); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除用户分组
		r.DELETE(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 查询
			var rbacPermissionGroup models.RbacPermissionGroupModel
			ret = models.Init(models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				SetPreloads(tools.Strings{"RbacPermissions"}).
				Prepare().
				First(&rbacPermissionGroup)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限分组")

			// 删除权限
			if len(rbacPermissionGroup.RbacPermissions) > 0 {
				models.Init(models.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup.RbacPermissions)
			}

			// 删除
			models.Init(models.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 编辑权限分组
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (RbacPermissionGroupStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.RbacPermissionGroupModel
			ret = models.Init(models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "权限分组名称")

			// 查询
			var rbacPermissionGroup models.RbacPermissionGroupModel
			ret = models.Init(models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&rbacPermissionGroup)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限分组")

			// 修改
			rbacPermissionGroup.Name = form.Name
			models.Init(models.RbacPermissionGroupModel{}).DB().Save(&rbacPermissionGroup)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 权限分组详情
		r.GET(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 读取
			var rbacPermissionGroup models.RbacPermissionGroupModel
			ret = models.Init(models.RbacPermissionGroupModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				SetPreloads(tools.Strings{"RbacPermissions"}).
				Prepare().
				First(&rbacPermissionGroup)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "权限分组")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
		})

		// 权限分组列表
		r.GET("", func(ctx *gin.Context) {
			var rbacPermissionGroups []models.RbacPermissionGroupModel
			models.Init(models.RbacPermissionGroupModel{}).
				SetPreloads(tools.Strings{"RbacPermissions"}).
				PrepareQuery(ctx).
				Find(&rbacPermissionGroups)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_groups": rbacPermissionGroups}))
		})
	}
}
