package middlewares

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/models"
	"fix-workshop-ue/settings"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccountUUID, exists := ctx.Get("__ACCOUNT__")
		if !exists {
			panic(abnormals.PanicUnLogin("未登录"))
		}

		cfg := (&settings.Setting{}).Init()

		if !cfg.App.Section("app").Key("production").MustBool(true) {
			var ret *gorm.DB

			// 获取权限
			var rbacPermission models.RbacPermissionModel
			ret = (&models.BaseModel{}).
				SetPreloads(tools.Strings{"RbacRoles"}).
				SetWheres(tools.Map{
					"uri":    ctx.FullPath(),
					"method": ctx.Request.Method,
				}).
				Prepare().
				First(&rbacPermission)
			abnormals.PanicWhenIsEmpty(ret, "权限")

			ok := false
			if len(rbacPermission.RbacRoles) > 0 {
				for _, rbacRole := range rbacPermission.RbacRoles {
					if len(rbacRole.Accounts) > 0 {
						for _, account := range rbacRole.Accounts {
							if account.UUID == currentAccountUUID {
								ok = true
							}
						}
					}
				}
			}

			if !ok {
				abnormals.PanicUnAuth("未授权")
			}
		}

		ctx.Next()
	}
}
