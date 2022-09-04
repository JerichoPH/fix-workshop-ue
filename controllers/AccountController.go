package controllers

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountController struct{}

// AccountStoreForm 新建用户表单
type AccountStoreForm struct {
	Username                  string `form:"username" json:"username" binding:"required"`
	Password                  string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation      string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
	Nickname                  string `form:"nickname" json:"nickname" binding:"required"`
	OrganizationRailwayUUID   string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway       models.OrganizationRailwayModel
	OrganizationParagraphUUID string `form:"organization_paragraph_uuid" json:"organization_paragraph_uuid"`
	OrganizationParagraph     models.OrganizationParagraphModel
	OrganizationWorkshopUUID  string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop      models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID  string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea      models.OrganizationWorkAreaModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AccountStoreForm
func (cls AccountStoreForm) ShouldBind(ctx *gin.Context) AccountStoreForm {
	var ret *gorm.DB

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
	if cls.OrganizationRailwayUUID != "" {
		ret = models.BootByModel(models.OrganizationRailwayModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationRailwayUUID}).PrepareByDefault().First(&cls.OrganizationRailway)
		wrongs.PanicWhenIsEmpty(ret, "路局")
	}
	if cls.OrganizationParagraphUUID != "" {
		ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationParagraphUUID}).PrepareByDefault().First(&cls.OrganizationParagraph)
		wrongs.PanicWhenIsEmpty(ret, "站段")
	}
	if cls.OrganizationWorkshopUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).PrepareByDefault().First(&cls.OrganizationWorkshop)
		wrongs.PanicWhenIsEmpty(ret, "车间")
	}
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).PrepareByDefault().First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return cls
}

// AccountUpdateForm 编辑用户表单
type AccountUpdateForm struct {
	Username                  string `form:"username" json:"username" uri:"username"`
	Nickname                  string `form:"nickname" json:"nickname" uri:"nickname"`
	OrganizationRailwayUUID   string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway       models.OrganizationRailwayModel
	OrganizationParagraphUUID string `form:"organization_paragraph_uuid" json:"organization_paragraph_uuid"`
	OrganizationParagraph     models.OrganizationParagraphModel
	OrganizationWorkshopUUID  string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop      models.OrganizationWorkshopModel
	OrganizationWorkAreaUUID  string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea      models.OrganizationWorkAreaModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return AccountUpdateForm
func (cls AccountUpdateForm) ShouldBind(ctx *gin.Context) AccountUpdateForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&cls); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if cls.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if cls.Nickname == "" {
		wrongs.PanicValidate("昵称必填")
	}
	if cls.OrganizationRailwayUUID != "" {
		ret = models.BootByModel(models.OrganizationRailwayModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationRailwayUUID}).PrepareByDefault().First(&cls.OrganizationRailway)
		wrongs.PanicWhenIsEmpty(ret, "路局")
	}
	if cls.OrganizationParagraphUUID != "" {
		ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationParagraphUUID}).PrepareByDefault().First(&cls.OrganizationParagraph)
		wrongs.PanicWhenIsEmpty(ret, "站段")
	}
	if cls.OrganizationWorkshopUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkshopUUID}).PrepareByDefault().First(&cls.OrganizationWorkshop)
		wrongs.PanicWhenIsEmpty(ret, "车间")
	}
	if cls.OrganizationWorkAreaUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).SetWheres(tools.Map{"uuid": cls.OrganizationWorkAreaUUID}).PrepareByDefault().First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
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

// Post 新建用户
func (cls AccountController) Post(ctx *gin.Context) {
	// 表单
	form := (&AccountStoreForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.AccountModel
	var ret *gorm.DB
	ret = (&models.BaseModel{}).
		SetWheres(tools.Map{"username": form.Username}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户名")
	ret = (&models.BaseModel{}).
		SetWheres(tools.Map{"nickname": form.Nickname}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "昵称")

	// 密码加密
	bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	account := &models.AccountModel{
		BaseModel:             models.BaseModel{UUID: uuid.NewV4().String()},
		Username:              form.Username,
		Nickname:              form.Nickname,
		Password:              string(bytes),
		OrganizationRailway:   form.OrganizationRailway,
		OrganizationParagraph: form.OrganizationParagraph,
		OrganizationWorkshop:  form.OrganizationWorkshop,
		OrganizationWorkArea:  form.OrganizationWorkArea,
	}

	if ret = models.BootByModel(models.AccountModel{}).
		PrepareByDefault().
		Create(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("新建成功").Created(tools.Map{}))
}

// Put 编辑用户
func (AccountController) Put(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)

	// 表单
	form := (&AccountUpdateForm{}).ShouldBind(ctx)

	// 查重
	var repeat models.AccountModel
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"username": form.Username}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户账号")
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"nickname": form.Nickname}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户昵称")

	// 查询
	ret = models.BootByModel(models.AccountModel{}).SetWheres(map[string]interface{}{"uuid": ctx.Param("uuid")}).PrepareByDefault().First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 编辑
	account.Username = form.Username
	account.Nickname = form.Nickname
	account.OrganizationRailwayUUID = form.OrganizationRailway.UUID
	account.OrganizationParagraphUUID = form.OrganizationParagraph.UUID
	account.OrganizationWorkshopUUID = form.OrganizationWorkshop.UUID
	account.OrganizationWorkAreaUUID = form.OrganizationWorkArea.UUID
	if ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		//SetPreloads("OrganizationRailway", "OrganizationParagraph", "OrganizationWorkshop", "OrganizationWorkArea").
		PrepareByDefault().
		Debug().
		Save(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBootByDefault().Updated(tools.Map{}))
}

// PutPassword 编辑密码
func (AccountController) PutPassword(ctx *gin.Context) {
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

	if ret = models.BootByModel(models.AccountModel{}).
		PrepareByDefault().
		Updates(map[string]interface{}{
			"password": string(bytes),
		}); ret.Error != nil {
		wrongs.PanicForbidden("编辑失败：" + ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("密码修改成功").Updated(tools.Map{}))
}

// Destroy 删除用户
func (AccountController) Destroy(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)

	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefault().
		Delete(&account)

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// Show 详情
func (AccountController) Show(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetPreloads("RbacRoles", "RbacRoles.RbacPermissions").
		PrepareByDefault().
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"account": account}))
}

// Index 列表
func(AccountController) Index(ctx *gin.Context){
	var accounts []models.AccountModel
	models.BootByModel(models.AccountModel{}).
		PrepareUseQueryByDefault(ctx).
		Find(&accounts)

	ctx.JSON(tools.CorrectBootByDefault().OK(tools.Map{"accounts": accounts}))
}
