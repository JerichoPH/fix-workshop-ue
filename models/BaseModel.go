package models

import (
	"database/sql"
	"fix-workshop-ue/databases"
	"gorm.io/gorm"
	"time"
)

// BaseModel 出厂数据、财务数据、检修数据、仓储数据、流转数据、运用数据
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
func (cls *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	cls.CreatedAt = time.Now()
	cls.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
func (cls *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
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
