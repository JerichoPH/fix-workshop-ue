package models

type KindSubTypeModel struct {
	BaseModel
	UniqueCode         string              `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:型号代码;" json:"unique_code"`
	Name               string              `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:型号名称;" json:"name"`
	Nickname           string              `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	BeEnable           bool                `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindEntireTypeUUID string              `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属类型代码;" json:"kind_entire_type_uuid"`
	KindEntireType     KindEntireTypeModel `gorm:"foreignKey:KindEntireTypeUUID;references:UUID;" json:"kind_entire_type"`
	CycleRepairYear    int16               `gorm:"type:INT2;NOT NULL;DEFAULT:0;COMMENT:周期修年;" json:"cycle_repair_year"`
	LifeYear           int16               `gorm:"type:INT2;NOT NULL;DEFAULT:15;COMMENT:寿命;" json:"life_year"`
}

// TableName 表名称
func (KindSubTypeModel) TableName() string {
	return "kind_sub_types"
}
