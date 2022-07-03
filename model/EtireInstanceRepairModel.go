package model

import (
	"time"
)

type EntireInstanceRepairModel struct {
	BaseModel
	EntireInstanceIdentityCode string              `gorm:"type:VARCHAR(19);UNIQUE;NOT NULL;COMMENT:所属器材唯一编号;" json:"entire_instance_identity_code"`
	EntireInstance             EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	FixerName                  string              `gorm:"type:VARCHAR(64);COMMENT:检修人;" json:"fixer_name"`
	FixedAt                    time.Time           `gorm:"type:DATETIME;COMMENT:检修时间;" json:"fixed_at"`
	PrevFixerName              string              `gorm:"type:VARCHAR(64);COMMENT:上一次检修人;" json:"prev_fixer_name"`
	PrevFixedAt                time.Time           `gorm:"type:DATETIME;COMMENT:上一次检修时间;" json:"prev_fixed_at"`
	CheckerName                string              `gorm:"type:VARCHAR(64);COMMENT:验收人;" json:"checker_name"`
	CheckedAt                  time.Time           `gorm:"type:DATETIME;COMMENT:验收时间;" json:"checked_at"`
	PrevCheckerName            string              `gorm:"type:VARCHAR(64);COMMENT:上一次验收人;" json:"prev_checker_name"`
	PrevCheckedAt              time.Time           `gorm:"type:DATETIME;COMMENT:上一次验收时间;" json:"prev_checked_at"`
	SpotCheckerName            string              `gorm:"type:VARCHAR(64);COMMENT:抽验人;" json:"spot_checker_name"`
	SpotCheckedAt              time.Time           `gorm:"type:DATETIME;COMMENT:抽验时间;" json:"spot_checked_at"`
	PrevSpotCheckerName           string            `gorm:"type:VARCHAR(64);COMMENT:上一次抽验人;" json:"prev_spot_checker_name"`
	PrevSpotCheckedAt             time.Time         `gorm:"type:DATETIME;COMMENT:上一次抽验时间;" json:"prev_spot_checked_at"`
	//FixWorkflowReportSerialNumber string            `gorm:"type:VARCHAR(64);COMMENT:所属检修单流水号;" json:"fix_workflow_report_serial_number"`
	//FixWorkflowReportModel             FixWorkflowReportModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FixWorkflowReportSerialNumber;references:SerialNumber;COMMENT:所属检修单;" json:"fix_workflow_report"`
	//PrevFixWorkflowSerialNumber   string            `gorm:"type:VARCHAR(64);COMMENT:上一次所属检修单号;" json:"prev_fix_workflow_serial_number"`
	//PrevFixWorkflowReport         FixWorkflowReportModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevFixWorkflowSerialNumber;references:SerialNumber;COMMENT:上一次所属检修单;" json:"prev_fix_workflow_report"`
}

// TableName 表名称
func (cls *EntireInstanceRepairModel) TableName() string {
	return "entire_instance_repairs"
}
