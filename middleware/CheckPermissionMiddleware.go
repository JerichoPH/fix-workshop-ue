package middleware

import (
	"fix-workshop-ue/config"
	"fix-workshop-ue/error"
	"fix-workshop-ue/model"
	"fix-workshop-ue/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccount, exists := ctx.Get("__currentAccount")
		fmt.Println(currentAccount)
		if !exists {
			panic(error.ThrowUnLogin("未登录"))
		}
		tool.ThrowErrorWhenIsEmpty(currentAccount, model.AccountModel{}, "用户")

		config := (&config.Config{}).Init()

		if !config.App.Section("app").Key("production").MustBool(true){
			// 获取权限
			rbacPermissions := (&model.RbacPermissionModel{
				BaseModel: model.BaseModel{
					Preloads: []string{
						"RbacRoles",
						"RbacRoles.Accounts",
					},
				},
			}).
				FindOneByURIAndMethod(ctx.FullPath(), ctx.Request.Method)
			tool.ThrowErrorWhenIsEmpty(rbacPermissions, model.RbacPermissionModel{}, "权限")

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
				panic(error.ThrowUnAuthorization("未授权"))
			}
		}

		ctx.Next()
	}
}
