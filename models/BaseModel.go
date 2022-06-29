package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint         `gorm:"type:BIGINT UNSIGNED;primaryKey;" json:"id"`
	UUID      string       `gorm:"type:VARCHAR(128);UNIQUE;NOT NULL;" json:"uuid"`
	CreatedAt time.Time    `gorm:"type:DATETIME;auto_now_add;" json:"created_at"`
	UpdatedAt time.Time    `gorm:"type:DATETIME;" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
}

// Boot 预处理
func (cls *BaseModel) Boot(db *gorm.DB, preloads, selects, omits []string) *gorm.DB {
	if preloads != nil && len(preloads) > 0 {
		for _, v := range preloads {
			db = db.Preload(v)
		}
	}

	if selects != nil && len(selects) > 0 {
		db = db.Select(selects)
	}

	if omits != nil && len(omits) > 0 {
		db = db.Omit(omits...)
	}

	return db
}
