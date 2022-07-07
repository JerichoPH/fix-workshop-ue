package v1

import (
	"fix-workshop-ue/errors"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RbacRoleRouter struct {
}

func (cls *RbacRoleRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/rbacRole",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// 新建角色
		r.POST("", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			var form models.RbacRoleStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(err)
			}

			// 查重
			var repeat models.RbacRoleStoreForm
			ret = (&models.BaseModel{}).
				SetModel(models.RbacRoleModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "角色名称")

			// 保存
			ret = (&models.BaseModel{}).DB().Create(&models.RbacRoleModel{Name: form.Name})
			if ret.Error != nil {
				panic(errors.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})
	}

	// 编辑角色
	r.POST(":id", func(ctx *gin.Context) {
		var ret *gorm.DB
		id:=tools.ThrowErrorWhenIsNotInt(ctx.Param("id"))

		// 表单
		var form models.RbacRoleStoreForm
	})

	// 角色列表
	r.GET("", func(ctx *gin.Context) {
		var rbacRoles models.RbacRoleModel
		(&models.BaseModel{}).
			SetModel(models.RbacRoleModel{}).
			PrepareQuery(ctx).
			Find(&rbacRoles)

		ctx.JSON(tools.CorrectIns("").OK(tools.Map{"rbac_roles": rbacRoles}))
	})
}
