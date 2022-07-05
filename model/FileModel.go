package model

type FileModel struct {
	BaseModel
	UUID               string       `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:UUID;" json:"uuid"`
	Filename           string       `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:文件存储名;" json:"filename"`
	Type               string       `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:文件类型;" json:"type"`
	OriginalFilename   string       `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:原始文件名;" json:"original_filename"`
	OriginalExtension  string       `gorm:"type:VARCHAR(32);NOT NULL;COMMENT:原始文件扩展名;" json:"original_extension"`
	Size               string       `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:文件大小;" json:"size"`
	UploadOperatorUUID string       `gorm:"type:CHAR(36);NOT NULL;COMMENT:上传人uuid;" json:"upload_operator_uuid"`
	UploadOperator     AccountModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UploadOperatorUUID;references:UUID;COMMENT:所属操作人;" json:"upload_operator"`
}

// TableName 表名称
func (cls *FileModel) TableName() string {
	return "files"
}