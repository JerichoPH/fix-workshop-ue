package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time    `gorm:"type:DATETIME;auto_now_add;" json:"created_at"`
	UpdatedAt time.Time    `gorm:"type:DATETIME;" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
	Preloads  []string     `gorm:"-:all"`
	Selects   []string     `gorm:"-:all"`
	Omits     []string     `gorm:"-:all"`
	DB        *gorm.DB     `gorm:"-:all"`
}

// Boot 初始化
func (cls *BaseModel) Boot() *gorm.DB {
	if cls.Preloads != nil && len(cls.Preloads) > 0 {
		for _, v := range cls.Preloads {
			cls.DB = cls.DB.Preload(v)
		}
	}

	if cls.Selects != nil && len(cls.Selects) > 0 {
		cls.DB = cls.DB.Select(cls.Selects)
	}

	if cls.Omits != nil && len(cls.Omits) > 0 {
		cls.DB = cls.DB.Omit(cls.Omits...)
	}

	return cls.DB
}
