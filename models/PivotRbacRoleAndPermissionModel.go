package models

type PivotRbacRoleAndPermissionModel struct {
	Id                 uint64              `json:"id"`
	RbacRoleUuid       string              `json:"rbac_role_uuid"`
	RbacRole           RbacRoleModel       `json:"rbac_role"`
	RbacPermissionUuid string              `json:"rbac_permission_uuid"`
	RbacPermission     RbacPermissionModel `json:"rbac_permission"`
}

func (PivotRbacRoleAndPermissionModel) TableName() string {
	return "pivot_rbac_role_and_rbac_permissions"
}
