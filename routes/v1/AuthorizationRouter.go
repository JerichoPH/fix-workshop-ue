package v1

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AuthorizationRegisterForm 注册表单
type AuthorizationRegisterForm struct {
	Username             string `form:"username" json:"username" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
	Nickname             string `form:"nickname" json:"nickname" binding:"required"`
}

// AuthorizationLoginForm 登录表单
type AuthorizationLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthorizationRouter struct {
}

func (cls *AuthorizationRouter) Load(router *gin.Engine) {
	r := router.Group("/api/v1/authorization")
	{
		// 注册
		r.POST("/register", func(ctx *gin.Context) {
			// 表单验证
			var authorizationRegisterForm AuthorizationRegisterForm
			if err := ctx.ShouldBind(&authorizationRegisterForm); err != nil {
				panic(err)
			}

			if authorizationRegisterForm.Password != authorizationRegisterForm.PasswordConfirmation {
				panic(errors.ThrowForbidden("两次密码输入不一致"))
			}

			// 检查重复项（用户名）
			var repeat models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"username": authorizationRegisterForm.Username}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户名")
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"nickname": authorizationRegisterForm.Nickname}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(authorizationRegisterForm.Password), 14)

			// 保存新用户
			account := models.AccountModel{
				Username:                authorizationRegisterForm.Username,
				Password:                string(bytes),
				Nickname:                authorizationRegisterForm.Nickname,
				AccountStatusUniqueCode: "DEFAULT",
			}

			if ret := (&databases.MySql{}).GetMySqlConn().Omit(clause.Associations).Create(&account); ret.Error != nil {
				panic(ret.Error)
			}

			ctx.JSON(tools.CorrectIns("注册成功").Created(gin.H{"account": account}))
		})

		// 登录
		r.POST("/login", func(ctx *gin.Context) {
			// 表单验证
			var form AuthorizationLoginForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(err)
			}

			// 获取用户
			var account models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetPreloads(tools.Strings{clause.Associations}).
				SetWheres(tools.Map{"username": form.Username}).
				Prepare().
				First(&account)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
				panic(errors.ThrowUnAuthorization("账号或密码错误"))
			}

			// 生成Jwt
			token, err := tools.GenerateJwt(account.UUID)
			if err != nil {
				// 生成jwt错误
				panic(err)
			}
			ctx.JSON(tools.CorrectIns("登陆成功").OK(gin.H{"token": token}))
		})
	}
}
