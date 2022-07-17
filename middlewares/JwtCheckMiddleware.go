package middlewares

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/models"
	"fix-workshop-ue/settings"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

func CheckJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取令牌
		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		if len(split) != 2 {
			abnormals.BombUnAuth("令牌格式错误")
		}
		tokenType := split[0]
		token := split[1]

		cfg := (&settings.Setting{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			var account models.AccountModel
			if token == "" {
				abnormals.BombUnAuth("令牌不存在")
			} else {
				switch tokenType {
				case "JWT":
					claims, err := tools.ParseJwt(token)

					// 判断令牌是否有效
					if err != nil {
						abnormals.BombUnAuth("令牌解析失败")
					} else if time.Now().Unix() > claims.ExpiresAt {
						abnormals.BombUnAuth("令牌过期")
					}

					// 判断用户是否存在
					if reflect.DeepEqual(claims, tools.Claims{}) {
						abnormals.BombUnAuth("令牌解析失败：用户不存在")
					}

					// 获取用户信息
					var ret *gorm.DB
					ret = (&models.BaseModel{}).
						SetModel(models.AccountModel{}).
						SetWheres(tools.Map{"uuid": claims.UUID}).
						Prepare().
						First(&account)

					abnormals.BombWhenIsEmpty(ret, "用户")
				default:
					abnormals.BombForbidden("权鉴认证方式不支持")
				}
			}
			ctx.Set("__ACCOUNT__", account.UUID)
		}

		ctx.Next()
	}
}
