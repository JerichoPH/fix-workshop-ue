package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

type MenuModel struct {
	BaseModel
	Name       string           `gorm:"type:VARCHAR(64);COMMENT:菜单名称;" json:"name"`
	URL        string           `gorm:"type:VARCHAR(128);COMMENT:url;" json:"url"`
	URIName    string           `gorm:"type:VARCHAR(64);COMMENT:路由标识;" json:"uri_name"`
	ParentUuid string           `gorm:"type:VARCHAR(36);COMMENT:父级编码;" json:"parent_uuid"`
	Icon       string           `gorm:"type:VARCHAR(64);COMMENT:图标;" json:"icon"`
	Parent     *MenuModel       `gorm:"foreignKey:ParentUuid;references:Uuid;COMMENT:所属父级;" json:"parent"`
	Subs       []*MenuModel     `gorm:"foreignKey:ParentUuid;references:Uuid;COMMENT:相关子集;" json:"subs"`
	RbacRoles  []*RbacRoleModel `gorm:"many2many:pivot_rbac_role_and_menus;foreignKey:uuid;joinForeignKey:menu_uuid;references:uuid;joinReferences:rbac_role_uuid;COMMENT:角色与菜单多对多;" json:"rbac_roles"`
}

// TableName 表名称
func (MenuModel) TableName() string {
	return "menus"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver ins
//  @param uuid
//  @return MenuModel
func (ins MenuModel) FindOneByUUID(uuid string) MenuModel {
	if ret := BootByModel(ins).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&ins); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "菜单"))
	}

	return ins
}
