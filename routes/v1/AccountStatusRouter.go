package v1

import (
	"fix-workshop-ue/errors"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountStatusRouter struct{}

// AccountStatusStoreForm 用户状态创建表单
type AccountStatusStoreForm struct {
	UniqueCode string `form:"unique_code" json:"unique_code" uri:"unique_code"`
	Name       string `form:"name" json:"name" uri:"name"`
}

// AccountStatusUpdateForm 用户状态编辑表单
type AccountStatusUpdateForm struct {
	Name string `form:"name" json:"name" uri:"name"`
}

func (cls *AccountStatusRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/accountStatus",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建用户状态
		r.POST("", func(ctx *gin.Context) {
			// 表单验证
			var form AccountStatusStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(errors.ThrowForbidden(err.Error()))
			}

			// 重复验证
			var repeat models.AccountStatusModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户代码")
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户状态名称")

			ret = (&models.BaseModel{}).DB().Create(&models.AccountStatusModel{
				UniqueCode: form.UniqueCode,
				Name:       form.Name,
			})
			if ret.Error != nil {
				panic(ret.Error)
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 删除用户状态
		r.DELETE(":unique_code", func(ctx *gin.Context) {
			// 查询
			uniqueCode := ctx.Param("unique_code")
			var accountStatus models.AccountStatusModel
			var ret *gorm.DB
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"unique_code": uniqueCode}).
				Prepare().
				First(&accountStatus)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户状态")

			// 删除
			(&models.BaseModel{}).DB().Delete(&accountStatus)

			ctx.JSON(tools.CorrectIns("").Deleted())
		})

		// 修改用户状态
		r.PUT(":unique_code", func(ctx *gin.Context) {
			var ret *gorm.DB
			uniqueCode := ctx.Param("unique_code")

			// 表单
			var form AccountStatusUpdateForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(errors.ThrowForbidden(err.Error()))
			}

			// 查重
			var repeat models.AccountStatusModel
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				SetNotWheres(tools.Map{"unique_code": uniqueCode}).
				Prepare().
				First(&repeat)
			tools.ThrowErrorWhenIsRepeatByDB(ret, "用户状态名称")

			// 查询
			var accountStatus models.AccountStatusModel
			ret = (&models.BaseModel{}).
				SetWheres(tools.Map{"unique_code": uniqueCode}).
				Prepare().
				First(&accountStatus)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "用户状态")

			// 修改
			accountStatus.Name = form.Name
			ret = (&models.BaseModel{}).DB().Save(&accountStatus)
			if ret.Error != nil {
				panic(errors.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Updated(nil))
		})

		// 用户状态列表
		r.GET("", func(ctx *gin.Context) {
			var accountStatuses []models.AccountStatusModel

			(&models.BaseModel{}).
				SetWhereFields(tools.Strings{"id", "created_at", "updated_at", "deleted_at", "unique_code", "name"}).
				PrepareQuery(ctx).
				Find(&accountStatuses)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account_statuses": accountStatuses}))
		})

		// 用户状态详情
		r.GET(":unique_code", func(ctx *gin.Context) {
			uniqueCode := ctx.Param("unique_code")

			var accountStatus models.AccountStatusModel
			if ret := (&models.BaseModel{}).
				SetWheres(tools.Map{"unique_code": uniqueCode}).
				Prepare().
				First(&accountStatus);
				ret.Error != nil {
				panic(errors.ThrowEmpty(""))
			}

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"account_status": accountStatus}))
		})
	}
}
