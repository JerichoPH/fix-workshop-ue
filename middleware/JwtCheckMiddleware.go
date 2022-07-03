package middleware

import (
	"fix-workshop-ue/error"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"time"
)

func CheckJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account map[string]interface{}

		split := strings.Split(tool.GetJwtFromHeader(ctx), " ")
		tokenType := split[0]
		token := split[1]

		if token == "" {
			panic(error.ThrowUnAuthorization("令牌不存在"))
		} else {
			switch tokenType {
			case "JWT":
				claims, err := tool.ParseJwt(token)

				// 判断令牌是否有效
				if err != nil {
					panic(error.ThrowUnAuthorization("令牌解析失败"))
				} else if time.Now().Unix() > claims.ExpiresAt {
					panic(error.ThrowUnAuthorization("令牌过期"))
				}

				// 判断用户是否存在
				if reflect.DeepEqual(claims, tool.Claims{}) {
					panic(error.ThrowUnAuthorization("令牌解析失败：用户不存在"))
				}

				// 获取用户信息
				account = (&model.AccountModel{}).FindOneByUUIDAsMap(claims.UUID)
				if reflect.DeepEqual(account, model.AccountModel{}) {
					panic(error.ThrowUnAuthorization("用户不存在"))
				}
			default:
				panic(error.ThrowForbidden("权鉴认证方式不支持"))

			}
		}

		ctx.Set("__currentAccount", account)
		ctx.Next()
	}
}
