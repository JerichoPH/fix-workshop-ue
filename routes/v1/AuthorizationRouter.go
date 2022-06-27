package v1

import (
	"fix-workshop-go/errors"
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type AuthorizationRouter struct {
	Router    *gin.Engine
	MySqlConn *gorm.DB
	MsSqlConn *gorm.DB
	AppConfig *ini.File
	DBConfig  *ini.File
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
			accountRepeat := (&models.AccountService{CTX: ctx}).FindOneByUsername(authorizationRegisterForm.Username)
			tools.ThrowErrorWhenIsRepeat(accountRepeat, models.Account{}, "用户名")
			// 检查重复项（昵称）
			accountRepeat = (&models.AccountService{CTX: ctx, MySqlConn: cls.MySqlConn}).FindOneByUsername(authorizationRegisterForm.Nickname)
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
			account := (&models.AccountService{
				CTX:       ctx,
				MySqlConn: cls.MySqlConn,
				Preloads:  []string{clause.Associations},
			}).FindOneByUsername(authorizationLoginForm.Username)
			tools.ThrowErrorWhenIsEmpty(account, models.Account{}, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(authorizationLoginForm.Password)); err != nil {
				panic(errors.ThrowUnAuthorization("账号或密码错误"))
			}

			// 生成Jwt
			token, err := cls.GenerateJwt(account.Username, account.Password)
			if err != nil {
				// 生成jwt错误
				panic(err)
			}
			ctx.JSON(tools.CorrectIns("登陆成功").OK(gin.H{"token": token}))
		})
	}
}

// GenerateJwt 生成Jwt
func (cls *AuthorizationRouter) GenerateJwt(username, password string) (string, error) {
	// 设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(168 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-learn",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseJwt 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func (cls *AuthorizationRouter) ParseJwt(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(account models.Account) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "jkdev.cn",
			Subject:   "account token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
