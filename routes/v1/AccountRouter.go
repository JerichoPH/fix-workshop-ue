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

type AccountRouter struct{}

// AccountStoreForm 新建用户表单
type AccountStoreForm struct {
	Username             string `form:"username" json:"username" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
	Nickname             string `form:"nickname" json:"nickname" binding:"required"`
}

// AccountUpdateForm 编辑用户表单
type AccountUpdateForm struct {
	Username                string `form:"username" json:"string" uri:"username"`
	Nickname                string `form:"nickname" json:"nickname" uri:"nickname"`
	AccountStatusUniqueCode string `form:"account_status_unique_code" json:"account_status_unique_code" uri:"account_status_unique_code"`
}

// AccountUpdatePasswordForm 修改密码表单
type AccountUpdatePasswordForm struct {
	OldPassword          string `form:"old_password" binding:"required"`
	NewPassword          string `form:"new_password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}

// Load 加载路由
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
			var form AccountStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}

			if form.Password != form.PasswordConfirmation {
				panic(exceptions.ThrowForbidden("两次密码输入不一致"))
			}

			// 检查重复项（用户名）
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

			if ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetOmits(tools.Strings{clause.Associations}).
				DB().
				Create(&models.AccountModel{
					Username:                form.Username,
					Password:                string(bytes),
					Nickname:                form.Nickname,
					AccountStatusUniqueCode: "DEFAULT",
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
			var form AccountUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}

			// 查重
			var repeat models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "用户账号")
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "用户昵称")

			// 查询
			var account models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "用户")

			// 编辑
			if form.Username != "" {
				account.Username = form.Username
			}
			if form.Nickname != "" {
				account.Nickname = form.Nickname
			}

			(&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				DB().
				Save(&account)

			ctx.JSON(tools.CorrectIns("").Updated(tools.Map{}))
		})

		// 修改密码
		r.PUT(":uuid/updatePassword", func(ctx *gin.Context) {
			var ret *gorm.DB
			var account models.AccountModel
			uuid := ctx.Param("uuid")

			// 表单
			var form AccountUpdatePasswordForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.NewPassword != form.PasswordConfirmation {
				panic(exceptions.ThrowForbidden("两次密码输入不一致"))
			}

			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
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

			if ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
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
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
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
			(&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetPreloads(tools.Strings{"AccountStatus"}).
				PrepareQuery(ctx).
				Find(&accounts)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"accounts": accounts}))
		})
	}
}
