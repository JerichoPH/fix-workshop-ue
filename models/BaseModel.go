package models

import (
	"database/sql"
	"fix-workshop-ue/databases"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// BaseModel 出厂数据、财务数据、检修数据、仓储数据、流转数据、运用数据
type BaseModel struct {
	ID             uint                   `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time              `gorm:"type:DATETIME;auto_now_add;" json:"created_at"`
	UpdatedAt      time.Time              `gorm:"type:DATETIME;" json:"updated_at"`
	DeletedAt      sql.NullTime           `gorm:"index" json:"deleted_at"`
	Preloads       []string               `gorm:"-:all"`
	Selects        []string               `gorm:"-:all"`
	Omits          []string               `gorm:"-:all"`
	Ctx            *gin.Context           `gorm:"-:all"`
	WhereFields    []string               `gorm:"-:all"`
	NotWhereFields []string               `gorm:"-:all"`
	IgnoreFields   []string               `gorm:"-:all"`
	Wheres         map[string]interface{} `gorm:"-:all"`
	NotWheres      map[string]interface{} `gorm:"-:all"`
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

// Prepare 初始化
func (cls *BaseModel) Prepare() (tx *gorm.DB) {
	tx = (&databases.MySql{}).GetMySqlConn().Where(cls.Wheres).Not(cls.NotWheres)

	// 拼接preloads关系
	if cls.Preloads != nil && len(cls.Preloads) > 0 {
		for _, v := range cls.Preloads {
			tx = tx.Preload(v)
		}
	}

	// 拼接selects字段
	if cls.Selects != nil && len(cls.Selects) > 0 {
		tx = tx.Select(cls.Selects)
	}

	// 拼接omits字段
	if cls.Omits != nil && len(cls.Omits) > 0 {
		tx = tx.Omit(cls.Omits...)
	}

	return tx
}

// PrepareQuery 根据Query参数初始化
func (cls *BaseModel) PrepareQuery() *gorm.DB {
	tx := cls.Prepare()

	//wheres := make(map[string]interface{})
	//notWheres := make(map[string]interface{})
	//
	//// 拼接需要跳过的字段
	//ignoreFields := make(map[string]int8)
	//if len(cls.IgnoreFields) > 0 {
	//	for _, v := range cls.IgnoreFields {
	//		ignoreFields[v] = 1
	//	}
	//}
	//
	//// 拼接Where条件
	//for _, v := range cls.WhereFields {
	//	if _, ok := ignoreFields[v]; !ok {
	//		if val, ok := cls.Ctx.GetQuery(v); ok {
	//			wheres[v] = val
	//		}
	//	}
	//}
	//
	//// 拼接NotWhere条件
	//for _, v := range cls.NotWhereFields {
	//	if _, ok := ignoreFields[v]; !ok {
	//		if val, ok := cls.Ctx.GetQuery(v); ok == true {
	//			notWheres[v] = val
	//		}
	//	}
	//}
	//tx = tx.Where(wheres).Not(notWheres)

	// 排序
	if order, ok := cls.Ctx.GetQuery("order"); ok {
		tx.Order(order)
	}

	// offset
	if offset, ok := cls.Ctx.GetQuery("offset"); ok {
		offset := tools.ThrowErrorWhenIsNotInt(offset, "offset参数只能填写整数")
		tx.Offset(offset)
	}

	// limit
	if limit, ok := cls.Ctx.GetQuery("limit"); ok {
		limit := tools.ThrowErrorWhenIsNotInt(limit, "limit参数只能填写整数")
		tx.Limit(limit)
	}

	return tx
}
