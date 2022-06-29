package middlewares

import (
	"fix-workshop-go/errors"
	"fix-workshop-go/models"
	"fix-workshop-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type GetJWTMiddleware struct {
}

func JwtCheck(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account models.Account

		token := tools.GetJwtFromHeader(ctx)

		ok := false

		if token == "" {
			panic(errors.ThrowUnAuthorization("令牌不存在"))
		} else {
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
			account = (&models.Account{BaseModel: models.BaseModel{DB: db}}).FindOneByUUID(claims.UUID)
			if reflect.DeepEqual(account, models.Account{}) {
				panic(errors.ThrowUnAuthorization("用户不存在"))
			}
		}

		ctx.Set("__currentAccount", account)

		ok = true
		if !ok {
			ctx.Abort()
		}

		ctx.Next()
	}
}
