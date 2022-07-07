package middlewares

import (
	"fix-workshop-ue/configs"
	"fix-workshop-ue/errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

func CheckJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account map[string]interface{}

		// 获取令牌
		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		tokenType := split[0]
		token := split[1]

		cfg := (&configs.Config{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			if token == "" {
				panic(errors.ThrowUnAuthorization("令牌不存在"))
			} else {
				switch tokenType {
				case "JWT":
					claims, err := tools.ParseJwt(token)

					// 判断令牌是否有效
					if err != nil {
						panic(errors.ThrowUnAuthorization("令牌解析失败"))
					} else if time.Now().Unix() > claims.ExpiresAt {
						panic(errors.ThrowUnAuthorization("令牌过期"))
					}

					// 判断用户是否存在
					if reflect.DeepEqual(claims, tools.Claims{}) {
						panic(errors.ThrowUnAuthorization("令牌解析失败：用户不存在"))
					}

					// 获取用户信息
					var account = make(map[string]interface{})
					var ret *gorm.DB
					ret = (&models.BaseModel{}).
						SetWheres(tools.Map{"uuid": claims.UUID}).
						Prepare().
						First(&account)
					tools.ThrowErrorWhenIsEmptyByDB(ret, "用户")
				default:
					panic(errors.ThrowForbidden("权鉴认证方式不支持"))

				}
			}
		}

		ctx.Set("__currentAccount", account)
		ctx.Next()
	}
}
