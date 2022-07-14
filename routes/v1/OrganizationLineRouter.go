package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationLineRouter struct{}

type OrganizationLineStoreForm struct {
	Sort       int64  `form:"sort" json:"sort"`
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
	BeEnable   bool   `form:"be_enable" json:"be_enable"`
}

type OrganizationLineUpdateForm struct {
	Sort       int64  `form:"sort" json:"sort"`
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
	BeEnable   bool   `form:"be_enable" json:"be_enable"`
}

func (cls *OrganizationLineRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("line", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			var form OrganizationLineStoreForm
			if err := ctx.ShouldBind(&form); err != nil {
				panic(exceptions.ThrowForbidden(err.Error()))
			}
			if form.UniqueCode == "" {
				panic(exceptions.ThrowForbidden("线别代码必填"))
			}
			if form.Name == "" {
				panic(exceptions.ThrowForbidden("线别名称必填"))
			}

			// 查重
			var repeat models.OrganizationLineModel
			ret = (&models.BaseModel{}).
				SetModel(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别代码")
			ret = (&models.BaseModel{}).
				SetModel(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "线别名称")

			// 新建
			if ret = (&models.BaseModel{}).
				SetModel(models.OrganizationLineModel{}).
				DB().
				Create(&models.OrganizationLineModel{
					BaseModel:  models.BaseModel{Sort: form.Sort},
					UniqueCode: form.UniqueCode,
					Name:       form.Name,
					BeEnable:   form.BeEnable,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})

		// 详情
		r.GET("line/:uuid", func(ctx *gin.Context) {
			var ret *gorm.DB
			uuid := ctx.Param("uuid")

			var organizationLine models.OrganizationLineModel
			ret = (&models.BaseModel{}).
				SetModel(models.OrganizationLineModel{}).
				SetWheres(tools.Map{"uuid": uuid}).
				Prepare().
				First(&organizationLine)
			tools.ThrowExceptionWhenIsEmptyByDB(ret, "线别")

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_line": organizationLine}))
		})

		// 列表
		r.GET("line", func(ctx *gin.Context) {
			var organizationLines []models.OrganizationLineModel
			(&models.BaseModel{}).
				SetModel(models.OrganizationLineModel{}).
				SetWhereFields(tools.Strings{"unique_code", "name", "be_enable", "sort"}).
				PrepareQuery(ctx).
				Find(&organizationLines)

			ctx.JSON(tools.CorrectIns("").OK(tools.Map{"organization_lines": organizationLines}))
		})
	}
}
