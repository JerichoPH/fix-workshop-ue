package models

import (
	"database/sql"
	"fix-workshop-go/databases"
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
}

// BeforeCreate 插入数据前
func (cls *BaseModel) BeforeCreate() (err error) {
	cls.CreatedAt = time.Now()
	cls.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
func (cls *BaseModel) BeforeSave() (err error) {
	cls.UpdatedAt = time.Now()
	return
}

// Boot 初始化
func (cls *BaseModel) Boot() *gorm.DB {
	db := (&databases.MySql{}).GetMySqlConn()
	if cls.Preloads != nil && len(cls.Preloads) > 0 {
		for _, v := range cls.Preloads {
			db = db.Preload(v)
		}
	}

	if cls.Selects != nil && len(cls.Selects) > 0 {
		db = db.Select(cls.Selects)
	}

	if cls.Omits != nil && len(cls.Omits) > 0 {
		db = db.Omit(cls.Omits...)
	}

	return db
}
