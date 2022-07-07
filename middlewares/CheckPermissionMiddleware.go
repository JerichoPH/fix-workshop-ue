package middlewares

import (
	"fix-workshop-ue/configs"
	"fix-workshop-ue/errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccount, exists := ctx.Get("__currentAccount")
		if !exists {
			panic(errors.ThrowUnLogin("未登录"))
		}
		fmt.Println(reflect.TypeOf(currentAccount))

		cfg := (&configs.Config{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			// 获取权限
			var rbacPermission models.RbacPermissionModel
			var ret *gorm.DB
			ret = (&models.BaseModel{
				Preloads: []string{"RbacRoles"},
				Wheres: map[string]interface{}{
					"uri":    ctx.FullPath(),
					"method": ctx.Request.Method,
				},
			}).
				Prepare().
				First(&rbacPermission)
			tools.ThrowErrorWhenIsEmptyByDB(ret, "权限")

			ok := false
			if len(rbacPermission.RbacRoles) > 0 {
				for _, rbacRole := range rbacPermission.RbacRoles {
					if len(rbacRole.Accounts) > 0 {
						for _, account := range rbacRole.Accounts {
							if account.UUID == currentAccount.(map[string]interface{})["uuid"] {
								ok = true
							}
						}
					}
				}
			}

			if !ok {
				panic(errors.ThrowUnAuthorization("未授权"))
			}
		}

		ctx.Next()
	}
}
