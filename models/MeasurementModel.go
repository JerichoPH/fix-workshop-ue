package models

type MeasurementModel struct {
	BaseModel
}

// TableName 表名称
func (cls *MeasurementModel) TableName() string {
	return "measurements"
}
