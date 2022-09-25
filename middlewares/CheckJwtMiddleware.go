package middlewares

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/settings"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

// CheckJwt 检查Jwt是否合法
func CheckJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取令牌
		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		if len(split) != 2 {
			wrongs.PanicUnAuth("令牌格式错误")
		}
		tokenType := split[0]
		token := split[1]

		cfg := (&settings.Setting{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			var account models.AccountModel
			if token == "" {
				wrongs.PanicUnAuth("令牌不存在")
			} else {
				switch tokenType {
				case "JWT":
					claims, err := tools.ParseJwt(token)

					// 判断令牌是否有效
					if err != nil {
						wrongs.PanicUnAuth("令牌解析失败")
					} else if time.Now().Unix() > claims.ExpiresAt {
						wrongs.PanicUnAuth("令牌过期")
					}

					// 判断用户是否存在
					if reflect.DeepEqual(claims, tools.Claims{}) {
						wrongs.PanicUnAuth("令牌解析失败：用户不存在")
					}

					// 获取用户信息
					var ret *gorm.DB
					ret = (&models.BaseModel{}).
						SetModel(models.AccountModel{}).
						SetWheres(tools.Map{"uuid": claims.Uuid}).
						SetPreloadsByDefault().
						SetPreloads("OrganizationWorkshop.OrganizationWorkshopType",
							"OrganizationWorkArea.OrganizationWorkAreaType",
							"OrganizationWorkArea.OrganizationWorkAreaProfession").
						PrepareByDefaultDbDriver().
						First(&account)

					wrongs.PanicWhenIsEmpty(ret, "令牌指向用户不存在")
				default:
					wrongs.PanicForbidden("权鉴认证方式不支持")
				}
			}
			ctx.Set(tools.AccountUuid, account.Uuid)                                           // 设置用户Uuid
			ctx.Set(tools.AccountOrganizationRailwayUuid, account.OrganizationRailwayUuid)     // 设置用户所属路局
			ctx.Set(tools.AccountOrganizationParagraphUuid, account.OrganizationParagraphUuid) // 设置用户所属站段
			ctx.Set(tools.AccountOrganizationWorkshopUuid, account.OrganizationWorkshopUuid)   // 设置用户所属车间
			ctx.Set(tools.AccountOrganizationWorkAreaUuid, account.OrganizationWorkAreaUuid)   // 设置用户所属工区
			ctx.Set(tools.AccountOrganizationLevel, "")                                        // 设置用户归属级别
			ctx.Set(tools.AccountOrganizationWorkshopTypeUniqueCode, "")                       // 设置所属车间类型
			ctx.Set(tools.AccountOrganizationWorkAreaTypeUniqueCode, "")                       // 设置工区类型
			ctx.Set(tools.AccountOrganizationWorkAreaProfessionUniqueCode, "")                 // 设置用户归属工区专业

			if (!reflect.DeepEqual(account.OrganizationRailway, &models.OrganizationRailwayModel{})) {
				ctx.Set(tools.AccountOrganizationLevel, tools.OrganizationLevelRailway)
			}
			if (!reflect.DeepEqual(account.OrganizationParagraph, models.OrganizationParagraphModel{})) {
				ctx.Set(tools.AccountOrganizationLevel, tools.OrganizationLevelParagraph)
			}
			if (!reflect.DeepEqual(account.OrganizationWorkshop, models.OrganizationWorkshopModel{})) {
				ctx.Set(tools.AccountOrganizationLevel, tools.OrganizationLevelWorkshop)
				ctx.Set(tools.AccountOrganizationWorkshopTypeUniqueCode, account.OrganizationWorkshop.OrganizationWorkshopType.UniqueCode)
			}
			if (!reflect.DeepEqual(account.OrganizationWorkArea, models.OrganizationWorkAreaModel{})) {
				ctx.Set(tools.AccountOrganizationLevel, tools.OrganizationLevelWorkArea)
				ctx.Set(tools.AccountOrganizationWorkAreaTypeUniqueCode, account.OrganizationWorkArea.OrganizationWorkAreaType.UniqueCode)
				ctx.Set(tools.AccountOrganizationWorkAreaProfessionUniqueCode, account.OrganizationWorkArea.OrganizationWorkAreaProfession.UniqueCode)
			}
			if account.Username == "admin" {
				ctx.Set(tools.AccountOrganizationLevel, tools.OrganizationLevelAll)
				ctx.Set(tools.AccountOrganizationWorkshopTypeUniqueCode, tools.OrganizationLevelAll)
				ctx.Set(tools.AccountOrganizationWorkAreaTypeUniqueCode, tools.OrganizationLevelAll)
				ctx.Set(tools.AccountOrganizationWorkAreaProfessionUniqueCode, tools.OrganizationLevelAll)
			}
		}

		ctx.Next()
	}
}