package v1

import (
	"fix-workshop-ue/wrongs"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
		wrongs.PanicValidate(err.Error())
	}
	if cls.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if cls.Nickname == "" {
		wrongs.PanicValidate("昵称必填")
	}
	if cls.Password != cls.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
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
		wrongs.PanicValidate(err.Error())
	}
	if cls.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if cls.Nickname == "" {
		wrongs.PanicValidate("昵称必填")
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
		wrongs.PanicValidate(err.Error())
	}
	if cls.OldPassword == "" {
		wrongs.PanicValidate("原始密码必填")
	}
	if cls.NewPassword == "" {
		wrongs.PanicValidate("新密码必填")
	}
	if cls.PasswordConfirmation == "" {
		wrongs.PanicValidate("确认密码必填")
	}
	if cls.NewPassword != cls.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls AccountRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/account",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", func(ctx *gin.Context) {
			// 表单
			form := (&AccountStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "用户名")
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

			account := &models.AccountModel{
				BaseModel: models.BaseModel{UUID: uuid.NewV4().String()},
				Username:  form.Username,
				Password:  string(bytes),
				Nickname:  form.Nickname,
			}
			if ret = models.Init(models.AccountModel{}).SetOmits(clause.Associations).GetSession().Create(&account); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("新建成功").Created(tools.Map{"account": account}))
		})

		// 编辑
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&AccountUpdateForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.AccountModel
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "用户账号")
			ret = models.Init(models.AccountModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
				Prepare().
				First(&repeat)
			wrongs.PanicWhenIsRepeat(ret, "用户昵称")

			// 查询
			account := (&models.AccountModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 编辑
			account.Username = form.Username
			account.Nickname = form.Nickname
			if ret = models.Init(models.AccountModel{}).GetSession().Save(&account); ret.Error != nil {
				wrongs.PanicForbidden(ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 修改密码
		r.PUT(":uuid/updatePassword", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&AccountUpdatePasswordForm{}).ShouldBind(ctx)

			// 查询
			account := (&models.AccountModel{}).FindOneByUUID(ctx.Param("uuid"))

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.OldPassword)); err != nil {
				wrongs.PanicUnAuth("旧密码输入错误")
			}

			// 修改密码
			bytes, _ := bcrypt.GenerateFromPassword([]byte(form.NewPassword), 14)
			account.Password = string(bytes)

			if ret = models.Init(models.AccountModel{}).GetSession().Save(&account); ret.Error != nil {
				wrongs.PanicForbidden("编辑失败：" + ret.Error.Error())
			}

			ctx.JSON(tools.CorrectIns("密码修改成功").Updated(tools.Map{}))
		})

		// 用户详情
		r.GET(":uuid", func(ctx *gin.Context) {
			account := (&models.AccountModel{}).FindOneByUUID(ctx.Param("uuid"))
			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account": account}))
		})

		// 用户列表
		r.GET("", func(ctx *gin.Context) {
			var accounts []models.AccountModel
			models.Init(models.AccountModel{}).
				SetPreloads("AccountStatus").
				PrepareQuery(ctx).
				Find(&accounts)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"accounts": accounts}))
		})
	}
}
