package models

import (
	"time"
)

// EntireInstanceLockModel 器材锁模型
type EntireInstanceLockModel struct {
	BaseModel
	EntireInstanceIdentityCode string              `gorm:"type:VARCHAR(20);COMMENT:唯一编号;" json:"entire_instance_identity_code"`
	EntireInstance             EntireInstanceModel `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	ExpireAt                   time.Time           `gorm:"COMMENT:过期时间;" json:"expire_at"`
	LockName                   string              `gorm:"type:VARCHAR(64);COMMENT:所名称;" json:"lock_name"`
	LockDescription            string              `gorm:"type:TEXT;COMMENT:锁描述;" json:"lock_description"`
	BusinessOrderTableName string `gorm:"type:VARCHAR(256);COMMENT:业务表名称;" json:"business_order_table_name"`
	BusinessOrderUuid      string `gorm:"type:VARCHAR(36);COMMENT:业务表UUID;" json:"business_order_uuid"`
	BusinessItemTableName string `gorm:"type:VARCHAR(256);COMMENT:业务子项表名称;" json:"business_item_table_name"`
	BusinessItemUuid      string `gorm:"type:VARCHAR(36);COMMENT:业务子项表UUID;" json:"business_item_uuid"`
}

// TableName 表名称
//  @receiver EntireInstanceLockModel
//  @return string
func (EntireInstanceLockModel) TableName() string {
	return "entire_instance_locks"
}
