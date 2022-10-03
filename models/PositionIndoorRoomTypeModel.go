package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

type PositionIndoorRoomTypeModel struct {
	BaseModel
	UniqueCode          string                    `gorm:"type:CHAR(2);COMMENT:机房类型代码;" json:"unique_code"`
	Name                string                    `gorm:"type:VARCHAR(64);COMMENT:机房类型名称;" json:"name"`
	PositionIndoorRooms []PositionIndoorRoomModel `gorm:"foreignKey:PositionIndoorRoomTypeUuid;references:Uuid;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (PositionIndoorRoomTypeModel) TableName() string {
	return "position_indoor_room_types"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver ins
//  @param uuid
//  @return PositionIndoorRoomTypeModel
func (ins PositionIndoorRoomTypeModel) FindOneByUUID(uuid string) PositionIndoorRoomTypeModel {
	if ret := BootByModel(ins).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&ins); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "工区"))
	}

	return ins
}
