package models

type RbacPermissionModel struct {
	BaseModel
	Name      string           `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:权限名称;" json:"name"`
	URI       string           `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:指向路由;" json:"uri"`
	Method    string           `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:请求方法;" json:"method"`
	RbacRoles []*RbacRoleModel `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;COMMENT:角色与权限多对多;" json:"rbac_roles"`
}

// TableName 表名称
func (cls RbacPermissionModel) TableName() string {
	return "rbac_permissions"
}

// FindOneByURI 根据uri和method获取一条数据
func (cls *RbacPermissionModel) FindOneByURIAndMethod(uri string, method string) (rbacPermissionModel RbacPermissionModel) {
	cls.BaseModel.Boot().Where(map[string]interface{}{"uri": uri, "method": method}).First(&rbacPermissionModel)
	return
}
