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
//  @receiver ins
//  @param ctx
//  @return AccountStoreForm
func (ins AccountStoreForm) ShouldBind(ctx *gin.Context) AccountStoreForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if len(ins.Username) > 64 {
		wrongs.PanicValidate("账号不能超过64位")
	}
	if ins.Nickname == "" {
		wrongs.PanicValidate("昵称必填")
	}
	if len(ins.Nickname) > 64 {
		wrongs.PanicValidate("昵称不能超过64位")
	}
	if len(ins.Password) > 32 || len(ins.Password) < 6 {
		wrongs.PanicValidate("密码不能小于6位或大于32位")
	}
	if ins.Password != ins.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
	}
	if ins.OrganizationRailwayUUID != "" {
		ret = models.BootByModel(models.OrganizationRailwayModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationRailwayUUID}).PrepareByDefaultDbDriver().First(&ins.OrganizationRailway)
		wrongs.PanicWhenIsEmpty(ret, "路局")
	}
	if ins.OrganizationParagraphUUID != "" {
		ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationParagraphUUID}).PrepareByDefaultDbDriver().First(&ins.OrganizationParagraph)
		wrongs.PanicWhenIsEmpty(ret, "站段")
	}
	if ins.OrganizationWorkshopUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationWorkshopUUID}).PrepareByDefaultDbDriver().First(&ins.OrganizationWorkshop)
		wrongs.PanicWhenIsEmpty(ret, "车间")
	}
	if ins.OrganizationWorkAreaUUID != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationWorkAreaUUID}).PrepareByDefaultDbDriver().First(&ins.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return ins
}

// AccountUpdateForm 编辑用户表单
type AccountUpdateForm struct {
	Username                  string `form:"username" json:"username" uri:"username"`
	Nickname                  string `form:"nickname" json:"nickname" uri:"nickname"`
	OrganizationRailwayUuid   string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway       models.OrganizationRailwayModel
	OrganizationParagraphUuid string `form:"organization_paragraph_uuid" json:"organization_paragraph_uuid"`
	OrganizationParagraph     models.OrganizationParagraphModel
	OrganizationWorkshopUuid  string `form:"organization_workshop_uuid" json:"organization_workshop_uuid"`
	OrganizationWorkshop      models.OrganizationWorkshopModel
	OrganizationWorkAreaUuid  string `form:"organization_work_area_uuid" json:"organization_work_area_uuid"`
	OrganizationWorkArea      models.OrganizationWorkAreaModel
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return AccountUpdateForm
func (ins AccountUpdateForm) ShouldBind(ctx *gin.Context) AccountUpdateForm {
	var ret *gorm.DB

	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if ins.Nickname == "" {
		wrongs.PanicValidate("昵称必填")
	}
	if ins.OrganizationRailwayUuid != "" {
		ret = models.BootByModel(models.OrganizationRailwayModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationRailwayUuid}).PrepareByDefaultDbDriver().First(&ins.OrganizationRailway)
		wrongs.PanicWhenIsEmpty(ret, "路局")
	}
	if ins.OrganizationParagraphUuid != "" {
		ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationParagraphUuid}).PrepareByDefaultDbDriver().First(&ins.OrganizationParagraph)
		wrongs.PanicWhenIsEmpty(ret, "站段")
	}
	if ins.OrganizationWorkshopUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationWorkshopUuid}).PrepareByDefaultDbDriver().First(&ins.OrganizationWorkshop)
		wrongs.PanicWhenIsEmpty(ret, "车间")
	}
	if ins.OrganizationWorkAreaUuid != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).SetWheres(tools.Map{"uuid": ins.OrganizationWorkAreaUuid}).PrepareByDefaultDbDriver().First(&ins.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "工区")
	}

	return ins
}

// AccountUpdatePasswordForm 修改密码表单
type AccountUpdatePasswordForm struct {
	OldPassword          string `form:"old_password" binding:"required"`
	NewPassword          string `form:"new_password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}

// ShouldBind 绑定表单
//  @receiver ins
//  @param ctx
//  @return AccountUpdatePasswordForm
func (ins AccountUpdatePasswordForm) ShouldBind(ctx *gin.Context) AccountUpdatePasswordForm {
	if err := ctx.ShouldBind(&ins); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if ins.OldPassword == "" {
		wrongs.PanicValidate("原始密码必填")
	}
	if ins.NewPassword == "" {
		wrongs.PanicValidate("新密码必填")
	}
	if ins.PasswordConfirmation == "" {
		wrongs.PanicValidate("确认密码必填")
	}
	if ins.NewPassword != ins.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
	}

	return ins
}

// N 新建用户
func (ins AccountController) N(ctx *gin.Context) {
	// 表单
	form := (&AccountStoreForm{}).ShouldBind(ctx)

	// 查重
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

	account := &models.AccountModel{
		BaseModel:             models.BaseModel{Uuid: uuid.NewV4().String()},
		Username:              form.Username,
		Nickname:              form.Nickname,
		Password:              string(bytes),
		OrganizationRailway:   form.OrganizationRailway,
		OrganizationParagraph: form.OrganizationParagraph,
		OrganizationWorkshop:  form.OrganizationWorkshop,
		OrganizationWorkArea:  form.OrganizationWorkArea,
	}

	if ret = models.BootByModel(models.AccountModel{}).
		PrepareByDefaultDbDriver().
		Create(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("新建成功").Created(tools.Map{}))
}

// R 删除用户
func (AccountController) R(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)

	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		Delete(&account)

	ctx.JSON(tools.CorrectBootByDefault().Deleted())
}

// E 编辑用户
func (AccountController) E(ctx *gin.Context) {
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
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户账号")
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"nickname": form.Nickname}).
		SetNotWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		PrepareByDefaultDbDriver().
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户昵称")

	// 查询
	ret = models.BootByModel(models.AccountModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 编辑
	account.Username = form.Username
	account.Nickname = form.Nickname
	account.OrganizationRailwayUuid = form.OrganizationRailway.Uuid
	account.OrganizationParagraphUuid = form.OrganizationParagraph.Uuid
	account.OrganizationWorkshopUuid = form.OrganizationWorkshop.Uuid
	account.OrganizationWorkAreaUuid = form.OrganizationWorkArea.Uuid
	if ret = models.BootByModel(models.AccountModel{}).SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).PrepareByDefaultDbDriver().Debug().Save(&account); ret.Error != nil {
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
		PrepareByDefaultDbDriver().
		Updates(map[string]interface{}{
			"password": string(bytes),
		}); ret.Error != nil {
		wrongs.PanicForbidden("编辑失败：" + ret.Error.Error())
	}

	ctx.JSON(tools.CorrectBoot("密码修改成功").Updated(tools.Map{}))
}

// D 详情
func (AccountController) D(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)
	ret = models.BootByModel(models.AccountModel{}).
		SetWheres(tools.Map{"uuid": ctx.Param("uuid")}).
		SetPreloads("RbacRoles", "RbacRoles.RbacPermissions").
		PrepareByDefaultDbDriver().
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"account": account}))
}

// L 列表
func (AccountController) L(ctx *gin.Context) {
	var (
		accounts []models.AccountModel
		count    int64
		db       *gorm.DB
	)
	db = models.BootByModel(models.AccountModel{}).
		PrepareUseQueryByDefaultDbDriver(ctx)

	if ctx.Query("__page__") == "" {
		db.Find(&accounts)
		ctx.JSON(tools.CorrectBootByDefault().Ok(tools.Map{"accounts": accounts}))
	} else {
		db.Count(&count)
		models.Pagination(db, ctx).Find(&accounts)
		ctx.JSON(tools.CorrectBootByDefault().OkForPagination(tools.Map{"accounts": accounts}, ctx.Query("__page__"), count))
	}
}
