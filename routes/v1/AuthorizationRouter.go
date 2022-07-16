package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AuthorizationRegisterForm
func (cls AuthorizationRegisterForm) ShouldBind(ctx *gin.Context) AuthorizationRegisterForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Username == "" {
		panic(exceptions.ThrowForbidden("账号必填"))
	}
	if cls.Password == "" {
		panic(exceptions.ThrowForbidden("密码必填"))
	}
	if len(cls.Password) < 6 || len(cls.Password) > 18 {
		panic(exceptions.ThrowForbidden("密码不可小于6位或大于18位"))
	}
	if cls.Password != cls.PasswordConfirmation {
		panic(exceptions.ThrowForbidden("两次密码输入不一致"))
	}

	return cls
}

// AuthorizationLoginForm 登录表单
type AuthorizationLoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (cls AuthorizationLoginForm) ShouldBind(ctx *gin.Context) AuthorizationLoginForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.Username == "" {
		panic(exceptions.ThrowForbidden("账号必填"))
	}
	if cls.Password == "" {
		panic(exceptions.ThrowForbidden("密码必填"))
	}
	if len(cls.Password) < 6 || len(cls.Password) > 18 {
		panic(exceptions.ThrowForbidden("密码不可小于6位或大于18位"))
	}

	return cls
}

type AuthorizationRouter struct{}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *AuthorizationRouter) Load(router *gin.Engine) {
	r := router.Group("/api/v1/authorization")
	{
		// 注册
		r.POST("register", func(ctx *gin.Context) {
			// 表单验证
			form := (&AuthorizationRegisterForm{}).ShouldBind(ctx)

			// 检查重复项（用户名）
			var repeat models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				Prepare().
				First(&repeat)
			exceptions.ThrowWhenIsRepeatByDB(ret, "用户名")
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				Prepare().
				First(&repeat)
			exceptions.ThrowWhenIsRepeatByDB(ret, "昵称")

			// 密码加密
			bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

			// 保存新用户
			if ret = models.Init(models.AccountModel{}).
				SetOmits(clause.Associations).
				DB().
				Create(&models.AccountModel{
					BaseModel: models.BaseModel{UUID: uuid.NewV4().String()},
					Username:  form.Username,
					Password:  string(bytes),
					Nickname:  form.Nickname,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden("创建失败：" + ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("注册成功").Created(tools.Map{}))
		})

		// 登录
		r.POST("login", func(ctx *gin.Context) {
			// 表单验证
			form := (&AuthorizationLoginForm{}).ShouldBind(ctx)

			// 获取用户
			var account models.AccountModel
			var ret *gorm.DB
			ret = models.Init(models.AccountModel{}).
				SetPreloads(tools.Strings{clause.Associations}).
				SetWheres(tools.Map{"username": form.Username}).
				Prepare().
				First(&account)
			exceptions.ThrowWhenIsEmptyByDB(ret, "用户")

			// 验证密码
			if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
				panic(exceptions.ThrowUnAuthorization("账号或密码错误"))
			}

			// 生成Jwt
			if token, err := tools.GenerateJwt(account.UUID, account.Password); err != nil {
				// 生成jwt错误
				panic(exceptions.ThrowForbidden(err.Error()))
			} else {
				ctx.JSON(tools.CorrectIns("登陆成功").OK(tools.Map{
					"token":    token,
					"username": account.Username,
					"nickname": account.Nickname,
					"uuid":     account.UUID,
				}))
			}
		})

		// 获取当前账号相关菜单
		r.GET(
			"menus",
			middlewares.CheckJwt(),
			func(ctx *gin.Context) {
				var ret *gorm.DB
				if accountUUID, exists := ctx.Get("__ACCOUNT__"); !exists {
					panic(exceptions.ThrowUnLogin("用户未登录"))
				} else {
					// 获取当前用户信息
					var account models.AccountModel
					ret = models.Init(models.AccountModel{}).
						SetWheres(tools.Map{"uuid": accountUUID}).
						SetPreloads(tools.Strings{"RbacRoles", "RbacRoles.Menus"}).
						Prepare().
						First(&account)
					exceptions.ThrowWhenIsEmptyByDB(ret, "当前令牌指向用户")

					menuUUIDs := make([]string, 50)
					if len(account.RbacRoles) > 0 {
						for _, rbacRole := range account.RbacRoles {
							if len(rbacRole.Menus) > 0 {
								for _, menu := range rbacRole.Menus {
									menuUUIDs = append(menuUUIDs, menu.UUID)
								}
							}
						}
					}

					var menus []models.MenuModel
					models.Init(models.MenuModel{}).
						DB().
						Where("uuid in ?", menuUUIDs).
						Where("parent_uuid is null").
						Preload("Subs").
						Find(&menus)

					ctx.JSON(tools.CorrectIns("").OK(tools.Map{"menus": menus}))
				}
			},
		)
	}
}
