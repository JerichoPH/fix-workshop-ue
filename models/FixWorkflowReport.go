package models

type FixWorkflowReport struct {
	BaseModel
	SerialNumber               string               `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:检修单流水号;" json:"serial_number"`
	//OperatorID                 string               `gorm:"type:BIGINT UNSIGNED;NOT NULL;COMMENT:操作人编号;" json:"operator_id"`
	//Operator                   Account              `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OperatorID;references:ID;COMMENT:操作人;" json:"operator"`
	FixWorkflowReportStage     string               `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:阶段;" json:"fix_workflow_report_stage"`
	BeAllow                    bool                 `gorm:"type:BOOLEAN;DEFAULT:0;NOT NULL;COMMENT:是否合格;" json:"be_allow"`
	//FixWorkflowProcesses       []FixWorkflowProcess `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FixWorkflowReportSerialNumber;references:SerialNumber;COMMENT:相关检修过程;" json:"fix_workflow_processes"`
	//EntireInstanceIdentityCode string               `gorm:"type:VARCHAR(20);NOT NULL;COMMENT:所属器材唯一编号;" json:"entire_instance_identity_code"`
	//EntireInstance             EntireInstance       `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
}
