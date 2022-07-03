package middlewares

import (
	"fix-workshop-ue/errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"time"
)

func CheckJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account map[string]interface{}

		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		tokenType := split[0]
		token := split[1]

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
				account = (&models.AccountModel{}).FindOneByUUIDAsMap(claims.UUID)
				if reflect.DeepEqual(account, models.AccountModel{}) {
					panic(errors.ThrowUnAuthorization("用户不存在"))
				}
			default:
				panic(errors.ThrowForbidden("权鉴认证方式不支持"))

			}
		}

		ctx.Set("__currentAccount", account)
		ctx.Next()
	}
}
