package models

type SiteDepotStorehouseModel struct{
	BaseModel
	UniqueCode string `gorm:"type:CHAR(4);NOT NULL;COMMENT:仓库库房代码;" json:""`
	Name string `gorm:"type:VARCHAR(36);NOT NULL;COMMENT:仓库库房名称;" json:""`
}