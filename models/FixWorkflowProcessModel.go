package models

type FixWorkflowProcessModel struct {
	BaseModel
	SerialNumber                  string                   `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:检修过程单流水号;" json:"serial_number"`
	OperatorId                    string                   `gorm:"type:BIGINT UNSIGNED;NOT NULL;COMMENT:操作人编号;" json:"operator_id"`
	Operator                      AccountModel             `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OperatorID;references:ID;COMMENT:操作人;" json:"operator"`
	FixWorkflowReportStage        string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:阶段;" json:"fix_workflow_report_stage"`
	BeAllow                       bool                     `gorm:"type:BOOLEAN;DEFAULT:0;NOT NULL;COMMENT:是否合格;" json:"be_allow"`
	FixWorkflowReportSerialNumber string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属检修单流水号;" json:"fix_workflow_report_serial_number"`
	FixWorkflowReport             FixWorkflowReportModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FixWorkflowReportSerialNumber;references:SerialNumber;COMMENT:所属检修单;" json:"fix_workflow_report"`
	FixWorkflowRecodes            []FixWorkflowRecodeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FixWorkflowProcessSerialNumber;references:SerialNumber;COMMENT:相关实测值;" json:"fix_workflow_recodes"`
}

// TableName 表名称
func (cls *FixWorkflowProcessModel) TableName() string {
	return "FixWorkflowProcesses"
}
