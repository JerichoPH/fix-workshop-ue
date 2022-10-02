package models

type PivotRbacRoleAndPermissionModel struct {
	RbacRoleId       uint64              `json:"rbac_role_id"`
	RbacRole         RbacRoleModel       `json:"rbac_role"`
	RbacPermissionId uint64              `json:"rbac_permission_id"`
	RbacPermission   RbacPermissionModel `json:"rbac_permission"`
}

func (PivotRbacRoleAndPermissionModel) TableName() string {
	return "pivot_rbac_role_and_rbac_permissions"
}
