package middlewares

import (
	"fix-workshop-ue/configs"
	"fix-workshop-ue/errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccount, exists := ctx.Get("__currentAccount")
		fmt.Println(currentAccount)
		if !exists {
			panic(errors.ThrowUnLogin("未登录"))
		}
		tools.ThrowErrorWhenIsEmpty(currentAccount, models.AccountModel{}, "用户")

		config := (&configs.Config{}).Init()

		if !config.App.Section("app").Key("production").MustBool(true){
			// 获取权限
			rbacPermissions := (&models.RbacPermissionModel{
				BaseModel: models.BaseModel{
					Preloads: []string{
						"RbacRoles",
						"RbacRoles.Accounts",
					},
				},
			}).
				FindOneByURIAndMethod(ctx.FullPath(), ctx.Request.Method)
			tools.ThrowErrorWhenIsEmpty(rbacPermissions, models.RbacPermissionModel{}, "权限")

			ok := false
			if len(rbacPermissions.RbacRoles) > 0 {
				for _, rbacRole := range rbacPermissions.RbacRoles {
					if len(rbacRole.Accounts) > 0 {
						for _, account := range rbacRole.Accounts {
							if account.UUID == currentAccount.(map[string]interface{})["uuid"] {
								ok = true
							}
						}
					}
				}
			}

			if !ok{
				panic(errors.ThrowUnAuthorization("未授权"))
			}
		}

		ctx.Next()
	}
}
