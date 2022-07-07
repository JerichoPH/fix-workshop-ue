package v1

import (
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRouter struct {
	Router *gin.Engine
}

// Load 加载路由
func (cls *AccountRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/account",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// 修改用户
		r.PUT(":uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			// 表单
			var form models.AccountUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(err)
			}

			// 查重
			var repeat models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"username": form.Username}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户账号")
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"nickname": form.Nickname}).
				SetNotWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户账号")

			// 查询
			var account models.AccountModel
			ret = (&models.BaseModel{}).
				SetModel(models.AccountModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&account)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")

			// 修改
			fmt.Println(form)
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
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account": account}))
		})
	}
}
