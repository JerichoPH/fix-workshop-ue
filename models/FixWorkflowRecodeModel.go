package models

type FixWorkflowRecodeModel struct {
	BaseModel
	SerialNumber                   string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:检修单流水号;" json:"serial_number"`
	Note                           string                  `gorm:"type:LONGTEXT;COMMENT:备注说明;" json:"note"`
	StandardValue                  string                  `gorm:"type:VARCHAR(64);COMMENT:标准值;" json:"standard_value"`
	Unit                           string                  `gorm:"type:VARCHAR(64);COMMENT:单位;" json:"unit"`
	TestValue                      string                  `gorm:"type:VARCHAR(64);COMMENT:实测值;" json:"test_value"`
	BeAllow                        bool                    `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否通过;" json:"be_allow"`
	FixWorkflowProcessSerialNumber string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属检测过程流水号;" json:"fix_workflow_process_serial_number"`
	FixWorkflowProcess             FixWorkflowProcessModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FixWorkflowProcessSerialNumber;references:SerialNumber;COMMENT:所属检修过程;" json:"fix_workflow_process"`
}

// TableName 表名称
func (cls *FixWorkflowRecodeModel) TableName() string {
	return "fix_workflow_recodes"
}