package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

type LocationIndoorRoomTypeModel struct {
	BaseModel
	UniqueCode           string                    `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:机房类型代码;" json:"unique_code"`
	Name                 string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:机房类型名称;" json:"name"`
	LocationIndoorRooms []LocationIndoorRoomModel `gorm:"foreignKey:LocationIndoorRoomTypeUUID;references:UUID;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (LocationIndoorRoomTypeModel) TableName() string {
	return "location_indoor_room_types"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return LocationIndoorRoomTypeModel
func (cls LocationIndoorRoomTypeModel) FindOneByUUID(uuid string) LocationIndoorRoomTypeModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "工区"))
	}

	return cls
}