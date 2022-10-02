package models

type PivotRbacRoleAndAccountModel struct {
	RbacRoleId uint64        `json:"rbac_role_id"`
	RbacRole   RbacRoleModel `json:"rbac_role"`
	AccountId  uint64        `json:"account_id"`
	Account    AccountModel  `json:"account"`
}

func (PivotRbacRoleAndAccountModel) TableName() string {
	return "pivot_rbac_role_and_accounts"
}
