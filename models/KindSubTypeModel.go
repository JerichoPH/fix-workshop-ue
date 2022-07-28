package models

type KindSubTypeModel struct {
	BaseModel
	UniqueCode         string                `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:型号代码;" json:"unique_code"`
	Name               string                `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:型号名称;" json:"name"`
	Nickname           string                `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	BeEnable           bool                `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindEntireTypeUUID string                `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属类型代码;" json:"kind_entire_type_uuid"`
	KindEntireType     KindEntireTypeModel   `gorm:"foreignKey:KindEntireTypeUUID;references:UUID;" json:"kind_entire_type"`
	EntireInstances    []EntireInstanceModel `gorm:"foreignKey:KindSubTypeUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (KindSubTypeModel) TableName() string {
	return "kind_sub_types"
}
