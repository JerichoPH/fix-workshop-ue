package models

// LocationDepotRowModel 库房排模型
type LocationDepotRowModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(12);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationDepotSectionUUID string                      `gorm:"type:CHAR(36);COMMENT:所属区代码;" json:"location_warehouse_section_uuid"`
	LocationDepotSection     LocationDepotSectionModel   `gorm:"foreignKey:LocationDepotSectionUUID;references:UUID;NOT NULL;COMMENT:所属区;" json:"location_warehouse_section"`
	LocationDepotRowTypeUUID string                      `gorm:"type:VARCHAR(36);NOT NULL;COMMENT:所属排类型;"`
	LocationDepotRowType     LocationDepotRowTypeModel   `gorm:"foreignKey:LocationDepotRowTypeUUID;references:UUID;NOT NULL;COMMENT:所属排类型;" json:"location_warehouse_platoon_type"`
	LocationDepotCabinets    []LocationDepotCabinetModel `gorm:"foreignKey:LocationDepotRowUUID;references:UUID;NOT NULL;COMMENT:相关柜架;" json:"location_warehouse_shelves"`
}

// TableName 表名称
func (cls *LocationDepotRowModel) TableName() string {
	return "location_depot_rows"
}
