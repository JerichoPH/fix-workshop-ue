package models

type MeasurementModel struct {
	BaseModel
}

// TableName 表名称
func (MeasurementModel) TableName() string {
	return "measurements"
}
