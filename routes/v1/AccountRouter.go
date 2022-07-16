package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AccountRouter 用户路由
type AccountRouter struct{}

// AccountStoreForm 新建用户表单
type AccountStoreForm struct {
	Username             string `form:"username" json:"username" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
	Nickname             string `form:"nickname" json:"nickname" binding:"required"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AccountStoreForm
func (cls AccountStoreForm) ShouldBind(ctx *gin.Context) AccountStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Username == "" {
		panic(exceptions.ThrowForbidden("账号必填"))
	}
	if cls.Nickname == "" {
		panic(exceptions.ThrowForbidden("昵称必填"))
	}
	if cls.Password != cls.PasswordConfirmation {
		panic(exceptions.ThrowForbidden("两次密码输入不一致"))
	}

	return cls
}

// AccountUpdateForm 编辑用户表单
type AccountUpdateForm struct {
	Username                string `form:"username" json:"string" uri:"username"`
	Nickname                string `form:"nickname" json:"nickname" uri:"nickname"`
	AccountStatusUniqueCode string `form:"account_status_unique_code" json:"account_status_unique_code" uri:"account_status_unique_code"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AccountUpdateForm
func (cls AccountUpdateForm) ShouldBind(ctx *gin.Context) AccountUpdateForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Username == "" {
		panic(exceptions.ThrowForbidden("账号必填"))
	}
	if cls.Nickname == "" {
		panic(exceptions.ThrowForbidden("昵称必填"))
	}

	return cls
}

// AccountUpdatePasswordForm 修改密码表单
type AccountUpdatePasswordForm struct {
	OldPassword          string `form:"old_password" binding:"required"`
	NewPassword          string `form:"new_password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AccountUpdatePasswordForm
func (cls AccountUpdatePasswordForm) ShouldBind(ctx *gin.Context) AccountUpdatePasswordForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.OldPassword == "" {
		panic(exceptions.ThrowForbidden("原始密码必填"))
	}
	if cls.NewPassword == "" {
		panic(exceptions.ThrowForbidden("新密码必填"))
	}
	if cls.PasswordConfirmation == "" {
		panic(exceptions.ThrowForbidden("确认密码必填"))
	}
	if cls.NewPassword != cls.PasswordConfirmation {
		panic(exceptions.ThrowForbidden("两次密码输入不一致"))
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *AccountRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/account",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建用户
		r.POST("", func(ctx *gin.Context) {
			// 表单验证
			form := (&AccountStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "用户名")
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

			if ret = models.Init(models.AccountModel{}).
				SetOmits(clause.Associations).
				DB().
				Create(&models.AccountModel{
					Username: form.Username,
					Password: string(bytes),
					Nickname: form.Nickname,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("新建成功").Created(tools.Map{}))
		})

		// 编辑用户
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			form := (&AccountUpdateForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.AccountModel
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "用户账号")
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "用户昵称")

			// 查询
			var account models.AccountModel
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "用户")

			// 编辑
			account.Username = form.Username
			account.Nickname = form.Nickname
			if ret = models.Init(models.AccountModel{}).
				DB().
				Save(&account); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 修改密码
		r.PUT(":uuid/updatePassword", func(ctx *gin.Context) {
			var ret *gorm.DB
			var account models.AccountModel
			uuid := ctx.Param("uuid")

			// 表单
			form := (&AccountUpdatePasswordForm{}).ShouldBind(ctx)

			// 查询
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.OldPassword)); err != nil {
				panic(exceptions.ThrowUnAuthorization("旧密码输入错误"))
			}

			// 修改密码
			bytes, _ := bcrypt.GenerateFromPassword([]byte(form.NewPassword), 14)
			account.Password = string(bytes)

			if ret = models.Init(models.AccountModel{}).
				DB().
				Save(&account); ret.Error != nil {
				panic(exceptions.ThrowForbidden("编辑失败：" + ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("密码修改成功").Updated(tools.Map{}))
		})

		// 用户详情
		r.GET(":uuid", func(ctx *gin.Context) {
			uuid := ctx.Param("uuid")

			var account models.AccountModel
			var ret *gorm.DB
			ret = models.Init(models.AccountModel{}).
				SetPreloads(tools.Strings{"AccountStatus"}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "用户")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account": account}))
		})

		// 用户列表
		r.GET("", func(ctx *gin.Context) {
			var accounts []models.AccountModel
			models.Init(models.AccountModel{}).
				SetPreloads(tools.Strings{"AccountStatus"}).
				PrepareQuery(ctx).
				Find(&accounts)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"accounts": accounts}))
		})
	}
}
