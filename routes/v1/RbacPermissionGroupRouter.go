package v1

import (
	"fix-workshop-ue/errors"
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

// RbacPermissionGroupUpdateForm 编辑权限分组表单
type RbacPermissionGroupUpdateForm struct {
	Name string `form:"name" json:"name"`
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
			var form RbacPermissionGroupStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(errors.ThrowForbidden(err.Error()))
			}

			// 查重
			var repeat models.RbacPermissionGroupModel
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "权限分组名称")

			// 保存
			var rbacPermissionGroup models.RbacPermissionGroupModel
			rbacPermissionGroup.Name = form.Name
			(&models.BaseModel{}).DB().Create(&rbacPermissionGroup)

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})
	}

	// 删除用户分组
	r.DELETE(":uuid", func(ctx *gin.Context) {
		var ret *gorm.DB
		uuid := ctx.Param("uuid")

		// 查询
		var rbacPermissionGroup models.RbacPermissionGroupModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacPermissionGroupModel{}).
			SetWheres(tools.Map{"uuid": uuid}).
			SetPreloads(tools.Strings{"RbacPermissions"}).
			Prepare().
			First(&rbacPermissionGroup)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "权限分组")

		// 删除权限
		if len(rbacPermissionGroup.RbacPermissions) > 0 {
			(&models.BaseModel{}).SetModel(models.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup.RbacPermissions)
		}

		// 删除
		(&models.BaseModel{}).SetModel(models.RbacPermissionGroupModel{}).DB().Delete(&rbacPermissionGroup)

		ctx.JSON(tools.CorrectIns("").Deleted())
	})

	// 编辑权限分组
	r.PUT(":uuid", func(ctx *gin.Context) {
		var ret *gorm.DB
		uuid := ctx.Param("uuid")

		// 表单
		var form RbacPermissionGroupUpdateForm
		if err := ctx.ShouldBind(&form); err != nil {
			panic(errors.ThrowForbidden(err.Error()))
		}

		// 查重
		var repeat models.RbacPermissionGroupModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacPermissionGroupModel{}).
			SetWheres(tools.Map{"name": form.Name}).
			SetNotWheres(tools.Map{"uuid": uuid}).
			Prepare().
			First(&repeat)
		tools.ThrowErrorWhenIsRepeatByDB(ret, "权限分组名称")

		// 查询
		var rbacPermissionGroup models.RbacPermissionGroupModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacPermissionGroupModel{}).
			SetWheres(tools.Map{"uuid": uuid}).
			Prepare().
			First(&rbacPermissionGroup)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "权限分组")

		// 修改
		if form.Name != "" {
			rbacPermissionGroup.Name = form.Name
		}
		(&models.BaseModel{}).SetModel(models.RbacPermissionGroupModel{}).DB().Save(&rbacPermissionGroup)

		ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
	})

	// 权限分组详情
	r.GET(":uuid", func(ctx *gin.Context) {
		var ret *gorm.DB
		uuid := ctx.Param("uuid")

		// 读取
		var rbacPermissionGroup models.RbacPermissionGroupModel
		ret = (&models.BaseModel{}).
			SetModel(models.RbacPermissionGroupModel{}).
			SetWheres(tools.Map{"uuid": uuid}).
			Prepare().
			First(&rbacPermissionGroup)
		tools.ThrowErrorWhenIsEmptyByDB(ret, "权限分组")

		ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_group": rbacPermissionGroup}))
	})

	// 权限分组列表
	r.GET("", func(ctx *gin.Context) {
		var rbacPermissionGroups models.RbacPermissionGroupModel
		(&models.BaseModel{}).
			SetModel(models.RbacPermissionGroupModel{}).
			PrepareQuery(ctx).
			Find(&rbacPermissionGroups)

		ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_permission_groups": rbacPermissionGroups}))
	})
}
