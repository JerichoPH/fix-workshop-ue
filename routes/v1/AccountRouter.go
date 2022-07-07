package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRouter struct{}

// AccountUpdateForm 用户编辑表单
type AccountUpdateForm struct {
	Username                string `form:"username" json:"string" uri:"username"`
	Nickname                string `form:"nickname" json:"nickname" uri:"nickname"`
	AccountStatusUniqueCode string `form:"account_status_unique_code" json:"account_status_unique_code" uri:"account_status_unique_code"`
}

// Load 加载路由
func (cls *AccountRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/account",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 编辑用户
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form AccountUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(err)
			}

			// 查重
			var repeat models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheresMap(tools.Map{"username": form.Username}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户账号")
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheresMap(tools.Map{"nickname": form.Nickname}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户昵称")

			// 查询
			var account models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheresMap(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")

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

		// 用户详情
		r.GET(":uuid", func(ctx *gin.Context) {
			uuid := ctx.Param("uuid")

			var account models.AccountModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetPreloads(tools.Strings{"AccountStatus"}).
				SetWheresMap(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account": account}))
		})

		// 用户列表
		r.GET("", func(ctx *gin.Context) {
			var accounts models.AccountModel
			(&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetPreloads(tools.Strings{"AccountStatus"}).
				PrepareQuery(ctx).
				Find(&accounts)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"accounts": accounts}))
		})
	}
}
