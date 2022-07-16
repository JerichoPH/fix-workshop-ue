package v1

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/middlewares"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrganizationParagraphRouter 站段路由
type OrganizationParagraphRouter struct{}

// OrganizationParagraphStoreForm
type OrganizationParagraphStoreForm struct {
	Sort                      int64  `form:"sort" json:"sort"`
	UniqueCode                string `form:"unique_code" json:"unique_code"`
	Name                      string `form:"name" json:"name"`
	ShortName                 string `form:"short_name" json:"short_name"`
	BeEnable                  bool   `form:"be_enable" json:"be_enable"`
	OrganizationRailwayUUID   string `form:"organization_railway_uuid" json:"organization_railway_uuid"`
	OrganizationRailway       models.OrganizationRailwayModel
	OrganizationWorkshopUUIDs []string `form:"organization_workshop_uuids" json:"organization_workshop_uuids"`
	OrganizationWorkshops     []models.OrganizationWorkshopModel
	OrganizationLineUUIDs     []string `form:"organization_line_uuids" json:"organization_line_uuids"`
	OrganizationLines         []*models.OrganizationLineModel
}

// ShouldBind 绑定表单
//  @receiver cls
//  @param ctx
//  @return OrganizationParagraphStoreForm
func (cls OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	if err := ctx.ShouldBind(&cls); err != nil {
		panic(exceptions.ThrowForbidden(err.Error()))
	}
	if cls.UniqueCode == "" {
		panic(exceptions.ThrowForbidden("站段代码必填"))
	}
	if cls.Name == "" {
		panic(exceptions.ThrowForbidden("站段名称必填"))
	}
	if cls.OrganizationRailwayUUID == "" {
		panic(exceptions.ThrowForbidden("所属路局必选"))
	}
	var ret *gorm.DB
	(&models.BaseModel{}).
		SetModel(models.OrganizationParagraphModel{}).
		SetScopes((&models.OrganizationRailwayModel{}).ScopeBeEnable).
		SetWheres(tools.Map{"uuid": cls.OrganizationRailwayUUID}).
		Prepare().
		First(&cls.OrganizationRailway)
	tools.ThrowExceptionWhenIsEmptyByDB(ret, "路局不存在")
	if len(cls.OrganizationWorkshopUUIDs) > 0 {
		(&models.BaseModel{}).
			SetModel(models.OrganizationWorkshopModel{}).
			SetScopes((&models.OrganizationWorkshopModel{}).ScopeBeEnable).
			DB().
			Where("uuid in ?", cls.OrganizationWorkshopUUIDs).
			Find(cls.OrganizationWorkshops)
	}
	if len(cls.OrganizationLineUUIDs) > 0 {
		(&models.BaseModel{}).
			SetModel(models.OrganizationLineModel{}).
			SetScopes((&models.OrganizationLineModel{}).ScopeBeEnable).
			DB().
			Where("uuid in ?", cls.OrganizationLineUUIDs).
			Find(&cls.OrganizationLines)
	}

	return cls
}

// Load 加载路由
//  @receiver cls
//  @param router
func (cls *OrganizationParagraphRouter) Load(router *gin.Engine) {
	r := router.Group(
		"/api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("paragraph", func(ctx *gin.Context) {
			var ret *gorm.DB

			// 表单
			form := (&OrganizationParagraphStoreForm{}).ShouldBind(ctx)

			// 查重
			var repeat models.OrganizationParagraphModel
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"unique_code": form.UniqueCode}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "站段代码")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"name": form.Name}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "站段名称")
			ret = models.Init(models.OrganizationParagraphModel{}).
				SetWheres(tools.Map{"short_name": form.ShortName}).
				Prepare().
				First(&repeat)
			tools.ThrowExceptionWhenIsRepeatByDB(ret, "站段简称")

			// 保存
			if ret = models.Init(models.OrganizationParagraphModel{}).
				DB().
				Create(&models.OrganizationParagraphModel{
					BaseModel:                     models.BaseModel{Sort: form.Sort},
					UniqueCode:                    form.UniqueCode,
					Name:                          form.Name,
					ShortName:                     form.ShortName,
					BeEnable:                      form.BeEnable,
					OrganizationRailwayUniqueCode: form.OrganizationRailway.UniqueCode,
					OrganizationWorkshops:         form.OrganizationWorkshops,
					OrganizationLines:             form.OrganizationLines,
				}); ret.Error != nil {
				panic(exceptions.ThrowForbidden(ret.Error.Error()))
			}

			ctx.JSON(tools.CorrectIns("").Created(tools.Map{}))
		})
	}
}
