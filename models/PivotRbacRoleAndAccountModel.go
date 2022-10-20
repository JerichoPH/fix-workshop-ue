package models

type PivotRbacRoleAndAccountModel struct {
	RbacRoleUuid string        `json:"rbac_role_uuid"`
	RbacRole     RbacRoleModel `json:"rbac_role"`
	AccountUuid  string        `json:"account_uuid"`
	Account      AccountModel  `json:"account"`
}

func (PivotRbacRoleAndAccountModel) TableName() string {
	return "pivot_rbac_role_and_accounts"
}
