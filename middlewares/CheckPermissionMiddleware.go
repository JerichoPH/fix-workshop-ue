package middlewares

import (
	"errors"
	"fix-workshop-ue/models"
	"fix-workshop-ue/services"
	"fix-workshop-ue/settings"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckPermission 检查用户权限
func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上下文中的用户
		currentAccountUuid, exists := ctx.Get(tools.AccountUuid)
		if !exists {
			wrongs.PanicUnLogin("未登录")
		}

		new(services.EntireInstanceService).GetAccountOrganizationLevel(ctx)

		cfg := (&settings.Setting{}).Init()

		if cfg.App.Section("app").Key("production").MustBool(true) {
			var (
				ret                 *gorm.DB
				currentRbacPermission models.RbacPermissionModel
			)

			// 获取权限
			var rbacPermission models.RbacPermissionModel
			ret = models.BootByModel(models.RbacPermissionModel{}).
				SetWheres(tools.Map{
					"uri":    ctx.FullPath(),
					"method": ctx.Request.Method,
				}).
				Prepare("").
				First(&rbacPermission)
			if ret.Error != nil {
				if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
					wrongs.PanicUnAuth(fmt.Sprintf("权限不存在（%s %s）", ctx.Request.Method, ctx.FullPath()))
				}
			}

			models.BootByModel(models.RbacPermissionModel{}).
				PrepareByDefaultDbDriver().
				Joins("join pivot_rbac_role_and_rbac_permissions prrarp on rbac_permissions.uuid = prrarp.rbac_permission_uuid").
				Joins("join rbac_roles rr on prrarp.rbac_role_uuid = rr.uuid").
				Joins("join pivot_rbac_role_and_accounts prraa on rr.uuid = prraa.rbac_role_uuid").
				Joins("join accounts a on prraa.account_uuid = a.uuid").
				Where("rbac_permissions.uri = ?", ctx.FullPath()).
				Where("rbac_permissions.method = ?",ctx.Request.Method).
				Where("a.uuid = ?",currentAccountUuid).
				First(&currentRbacPermission)

			if currentRbacPermission.BaseModel.Id == 0 {
				wrongs.PanicUnAuth(fmt.Sprintf("当前用户没有权限进行此操作 %s %s", ctx.Request.Method, ctx.FullPath()))
			}
		}

		ctx.Next()
	}
}
