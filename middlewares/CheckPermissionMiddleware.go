package middlewares

import (
	"errors"
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	"fix-workshop-ue/settings"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccountUUID, exists := ctx.Get("__ACCOUNT__")
		if !exists {
			panic(wrongs.PanicUnLogin("未登录"))
		}

		cfg := (&settings.Setting{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			var (
				ret               *gorm.DB
				currentRbacRoleID int64
			)

			// 获取权限
			var rbacPermission models.RbacPermissionModel
			ret = models.Init(models.RbacPermissionModel{}).
				SetWheres(tools.Map{
					"uri":    ctx.FullPath(),
					"method": ctx.Request.Method,
				}).
				Prepare().
				First(&rbacPermission)
			fmt.Println(ret.Error,ctx.Request.Method)
			if ret.Error != nil {
				if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
					wrongs.PanicUnAuth(fmt.Sprintf("权限不存在（%s %s）", ctx.Request.Method, ctx.FullPath()))
				}
			}

			(&databases.MySql{}).
				GetConn().
				Raw(`select prp.rbac_role_id
from pivot_rbac_role_and_rbac_permissions as prp
         join rbac_roles r on prp.rbac_role_id = r.id
         join rbac_permissions p on prp.rbac_permission_id = p.id
         join pivot_rbac_role_and_accounts pra on prp.rbac_role_id = pra.rbac_role_id
         join accounts a on pra.account_id = a.id
where p.uri = ? and p.method = ? and a.uuid = ?`, ctx.FullPath(), ctx.Request.Method, currentAccountUUID).
				Scan(&currentRbacRoleID)

			if currentRbacRoleID == 0 {
				wrongs.PanicUnAuth(fmt.Sprintf("当前用户没有权限进行此操作 %s %s", ctx.Request.Method, ctx.FullPath()))
			}
		}

		ctx.Next()
	}
}
