package v1

import (
	"fix-workshop-go/errors"
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	gcasbin "github.com/maxwellhertz/gin-casbin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
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
	Router     *gin.Engine
	MySqlConn  *gorm.DB
	MsSqlConn  *gorm.DB
	AppConfig  *ini.File
	DBConfig   *ini.File
	AuthCasbin *gcasbin.CasbinMiddleware
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
				panic(errors.ThrowForbidden("两次密码输入不一致"))
			}

			// 检查重复项（用户名）
			accountRepeat := (&models.Account{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
				},
			}).FindOneByUsername(authorizationRegisterForm.Username)
			tools.ThrowErrorWhenIsRepeat(accountRepeat, models.Account{}, "用户名")
			// 检查重复项（昵称）
			accountRepeat = (&models.Account{
				BaseModel: models.BaseModel{
					DB: cls.MySqlConn,
				},
			}).FindOneByUsername(authorizationRegisterForm.Nickname)
			tools.ThrowErrorWhenIsRepeat(accountRepeat, models.Account{}, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(authorizationRegisterForm.Password), 14)

			// 保存新用户
			account := models.Account{
				Username:                authorizationRegisterForm.Username,
				Password:                string(bytes),
				Nickname:                authorizationRegisterForm.Nickname,
				AccountStatusUniqueCode: "DEFAULT",
			}

			if ret := cls.MySqlConn.Omit(clause.Associations).Create(&account); ret.Error != nil {
				panic(ret.Error)
			}

			ctx.JSON(tools.CorrectIns("注册成功").Created(gin.H{"account": account}))
		})

		// 登录
		r.POST("/login", func(ctx *gin.Context) {
			// 表单验证
			var authorizationLoginForm AuthorizationLoginForm
			if err := ctx.ShouldBind(&authorizationLoginForm); err != nil {
				panic(err)
			}

			// 获取用户
			account := (&models.Account{
				BaseModel: models.BaseModel{
					DB:       cls.MySqlConn,
					Preloads: []string{clause.Associations},
				},
			}).FindOneByUsername(authorizationLoginForm.Username)
			tools.ThrowErrorWhenIsEmpty(account, models.Account{}, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(authorizationLoginForm.Password)); err != nil {
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
