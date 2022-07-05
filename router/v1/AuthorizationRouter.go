package v1

import (
	"fix-workshop-ue/database"
	"fix-workshop-ue/error"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	Router *gin.Engine
}

func (cls *AuthorizationRouter) Load() {
	r := cls.Router.Group("/api/v1/authorization")
	{
		// 注册
		r.POST("/register", func(ctx *gin.Context) {
			// 表单验证
			var authorizationRegisterForm AuthorizationRegisterForm
			if err := ctx.ShouldBind(&authorizationRegisterForm); err != nil {
				panic(err)
			}

			if authorizationRegisterForm.Password != authorizationRegisterForm.PasswordConfirmation {
				panic(error.ThrowForbidden("两次密码输入不一致"))
			}

			// 检查重复项（用户名）
			accountRepeat := (&model.AccountModel{
				BaseModel: model.BaseModel{},
			}).FindOneByUsername(authorizationRegisterForm.Username)
			tool.ThrowErrorWhenIsRepeat(accountRepeat, model.AccountModel{}, "用户名")
			// 检查重复项（昵称）
			accountRepeat = (&model.AccountModel{
				BaseModel: model.BaseModel{},
			}).FindOneByUsername(authorizationRegisterForm.Nickname)
			tool.ThrowErrorWhenIsRepeat(accountRepeat, model.AccountModel{}, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(authorizationRegisterForm.Password), 14)

			// 保存新用户
			account := model.AccountModel{
				Username:                authorizationRegisterForm.Username,
				Password:                string(bytes),
				Nickname:                authorizationRegisterForm.Nickname,
				AccountStatusUniqueCode: "DEFAULT",
			}

			if ret := (&database.MySql{}).GetMySqlConn().Omit(clause.Associations).Create(&account); ret.Error != nil {
				panic(ret.Error)
			}

			ctx.JSON(tool.CorrectIns("注册成功").Created(gin.H{"account": account}))
		})

		// 登录
		r.POST("/login", func(ctx *gin.Context) {
			// 表单验证
			var authorizationLoginForm AuthorizationLoginForm
			if err := ctx.ShouldBind(&authorizationLoginForm); err != nil {
				panic(err)
			}

			// 获取用户
			account := (&model.AccountModel{
				BaseModel: model.BaseModel{
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUsername(authorizationLoginForm.Username)
			tool.ThrowErrorWhenIsEmpty(account, model.AccountModel{}, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(authorizationLoginForm.Password)); err != nil {
				panic(error.ThrowUnAuthorization("账号或密码错误"))
			}

			// 生成Jwt
			token, err := tool.GenerateJwt(account.UUID)
			if err != nil {
				// 生成jwt错误
				panic(err)
			}
			ctx.JSON(tool.CorrectIns("登陆成功").OK(gin.H{"token": token}))
		})
	}
}