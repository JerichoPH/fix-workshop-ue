package models

type PivotRbacRoleAndMenuModel struct {
	RbacRoleUuid string        `json:"rbac_role_uuid"`
	RbacRole     RbacRoleModel `json:"rbac_role"`
	MenuUuid     string        `json:"menu_uuid"`
	Menu         MenuModel     `json:"menu"`
}

// TableName 获取表名称
func (PivotRbacRoleAndMenuModel) TableName() string {
	return "pivot_rbac_role_and_menus"
}
