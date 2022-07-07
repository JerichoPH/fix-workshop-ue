package v1

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/errors"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountStatusRouter struct {
	Router *gin.Engine
}

func (cls *AccountStatusRouter) Load() {
	r := cls.Router.Group(
		"/api/v1/accountStatus",
		middlewares.CheckJWT(),
		middlewares.CheckPermission(),
	)
	{
		// 新建用户状态
		r.POST(
			"",
			func(ctx *gin.Context) {
				// 表单验证
				var form models.AccountStatusStoreForm
				if err := ctx.ShouldBind(&form); err != nil {
					panic(err)
				}

				// 重复验证
				var repeat models.AccountStatusModel
				var ret *gorm.DB
				ret = (&models.BaseModel{
					Wheres: map[string]interface{}{"unique_code": form.UniqueCode},
				}).
					Prepare().
					First(&repeat)
				tools.ThrowErrorWhenIsRepeatByDB(ret, "用户代码")
				ret = (&models.BaseModel{
					Wheres: map[string]interface{}{"name": form.Name},
				}).
					Prepare().
					First(&repeat)
				tools.ThrowErrorWhenIsRepeatByDB(ret, "用户状态名称")

				if ret := (&databases.MySql{}).GetMySqlConn().Create(&models.AccountStatusModel{
					UniqueCode: form.UniqueCode,
					Name:       form.Name,
				});
					ret.Error != nil {
					panic(ret.Error)
				}

				ctx.JSON(tools.CorrectIns("").Created(nil))
			},
		)

		// 用户状态详情
		r.GET(
			":unique_code",
			func(ctx *gin.Context) {
				uniqueCode := ctx.Param("unique_code")

				var accountStatus models.AccountStatusModel
				if ret := (&models.BaseModel{
					Wheres: map[string]interface{}{"unique_code": uniqueCode},
				}).
					Prepare().
					First(&accountStatus);
					ret.Error != nil {
					panic(errors.ThrowEmpty(""))
				}

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account_status": accountStatus}))
			})

		// 用户状态列表
		r.GET(
			"",
			func(ctx *gin.Context) {
				var accountStatuses []models.AccountStatusModel

				(&models.BaseModel{
					Ctx: ctx,
					WhereFields: []string{
						"id",
						"created_at",
						"updated_at",
						"deleted_at",
						"unique_code",
						"name",
					},
				}).
					PrepareQuery().
					Find(&accountStatuses)

				ctx.JSON(tools.CorrectIns("").OK(gin.H{"account_statuses": accountStatuses}))
			},
		)
	}
}
