package models

type PivotRbacRoleAndPermissionModel struct {
	RbacRoleUuid       string              `json:"rbac_role_uuid"`
	RbacRole           RbacRoleModel       `json:"rbac_role"`
	RbacPermissionUuid string              `json:"rbac_permission_uuid"`
	RbacPermission     RbacPermissionModel `json:"rbac_permission"`
}

// TableName 获取表名称
func (PivotRbacRoleAndPermissionModel) TableName() string {
	return "pivot_rbac_role_and_rbac_permissions"
}
