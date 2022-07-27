package models

// LocationDepotRowModel 仓储仓库排
type LocationDepotRowModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:仓储仓库排代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓储仓库排名称;" json:"name"`
	LocationDepotSectionUUID string                      `gorm:"type:CHAR(36);NOT NULL;COMMENT:仓储仓库区域UUID;" json:"location_depot_section_uuid"`
	LocationDepotSection     LocationDepotSectionModel   `gorm:"foreignKey:LocationDepotSectionUUID;references:UUID;COMMENT:所属仓储仓库区域;" json:"location_depot_section"`
	LocationDepotCabinets    []LocationDepotCabinetModel `gorm:"foreignKey:LocationDepotRowUUID;references:UUID;COMMENT:相关仓储柜架;" json:"location_depot_cabinet"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (LocationDepotRowModel) TableName() string {
	return "location_depot_rows"
}
