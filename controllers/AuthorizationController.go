package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorizationController struct{}

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
		wrongs.PanicValidate(err.Error())
	}
	if cls.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if cls.Password == "" {
		wrongs.PanicValidate("密码必填")
	}
	if len(cls.Password) < 6 || len(cls.Password) > 18 {
		wrongs.PanicValidate("密码不可小于6位或大于18位")
	}
	if cls.Password != cls.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
	}

	return cls
}

// AuthorizationLoginForm 登录表单
type AuthorizationLoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ShouldBind 绑定表单
func (cls AuthorizationLoginForm) ShouldBind(ctx *gin.Context) AuthorizationLoginForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if cls.Password == "" {
		wrongs.PanicValidate("密码必填")
	}
	if len(cls.Password) < 6 || len(cls.Password) > 18 {
		wrongs.PanicValidate("密码不可小于6位或大于18位")
	}

	return cls
}

// PostRegister 注册
func (AuthorizationController) PostRegister(ctx *gin.Context) {
	// 表单验证
	form := (&AuthorizationRegisterForm{}).ShouldBind(ctx)

	// 检查重复项（用户名）
	var repeat models.AccountModel
	var ret *gorm.DB
	ret = (&models.BaseModel{}).
		SetWheres(tools.Map{"username": form.Username}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户名")
	ret = (&models.BaseModel{}).
		SetWheres(tools.Map{"nickname": form.Nickname}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "昵称")

	// 密码加密
	bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	// 保存新用户
	account := &models.AccountModel{
		BaseModel: models.BaseModel{Uuid: uuid.NewV4().String()},
		Username:  form.Username,
		Password:  string(bytes),
		Nickname:  form.Nickname,
	}
	if ret = models.BootByModel(models.AccountModel{}).
		SetOmits(clause.Associations).
		PrepareByDefaultDbDriver().
		Create(&account); ret.Error != nil {
		wrongs.PanicForbidden("创建失败：" + ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("注册成功").Created(tools.Map{"account": account}))
}

// PostLogin 登录
func (AuthorizationController) PostLogin(ctx *gin.Context) {
	// 表单验证
	form := (&AuthorizationLoginForm{}).ShouldBind(ctx)

	// 获取用户
	var account models.AccountModel
	var ret *gorm.DB
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"username": form.Username}).
		PrepareByDefaultDbDriver().
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
		wrongs.PanicUnAuth("账号或密码错误")
	}

	// 生成Jwt
	if token, err := tools.GenerateJwt(account.Uuid, account.Password); err != nil {
		// 生成jwt错误
		wrongs.PanicForbidden(err.Error())
	} else {
		ctx.JSON(tools.CorrectBoot("登陆成功").Ok(tools.Map{
			"token":    token,
			"username": account.Username,
			"nickname": account.Nickname,
			"uuid":     account.Uuid,
		}))
	}
}

// GetMenus 获取当前用户菜单
func (AuthorizationController) GetMenus(ctx *gin.Context) {
	var ret *gorm.DB
	if accountUUID, exists := ctx.Get("__ACCOUNT__"); !exists {
		wrongs.PanicUnLogin("用户未登录")
	} else {
		// 获取当前用户信息
		var account models.AccountModel
		ret = models.BootByModel(models.AccountModel{}).
			SetWheres(tools.Map{"uuid": accountUUID}).
			SetPreloads("RbacRoles", "RbacRoles.Menus").
			PrepareByDefaultDbDriver().
			First(&account)
		if !wrongs.PanicWhenIsEmpty(ret, "") {
			wrongs.PanicUnLogin("当前令牌指向用户不存在")
		}

		var menus []models.MenuModel
		models.BootByModel(models.MenuModel{}).
			PrepareByDefaultDbDriver().
			Joins("join pivot_rbac_role_and_menus prram on menus.id = prram.menu_id").
			Joins("join rbac_roles r on prram.rbac_role_id = r.id").
			Joins("join pivot_rbac_role_and_accounts prraa on r.id = prraa.rbac_role_id").
			Joins("join accounts a on prraa.account_id = a.id").
			Where("a.uuid = ?", account.BaseModel.Uuid).
			Where("menus.deleted_at is null").
			Where("menus.parent_uuid = ''").
			Order("menus.sort asc").
			Order("menus.id asc").
			Preload("Subs").
			Find(&menus)

		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"menus": menus}))
	}
}
