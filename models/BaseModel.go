package models

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// BaseModel 出厂数据、财务数据、检修数据、仓储数据、流转数据、运用数据
type BaseModel struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `gorm:"type:DATETIME;auto_now_add;" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:DATETIME;" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	preloads       []string
	selects        []string
	omits          []string
	whereFields    []string
	notWhereFields []string
	ignoreFields   []string
	wheres         map[string]interface{}
	notWheres      map[string]interface{}
	model          interface{}
}

// demoFindOne 获取单条数据演示
func (cls *BaseModel) demoFindOne() {
	var b BaseModel
	ret := cls.
		SetModel(BaseModel{}).
		SetWheres(tools.Map{}).
		SetNotWheres(tools.Map{}).
		Prepare().
		First(b)
	tools.ThrowErrorWhenIsEmptyByDB(ret, "XX")
}

// demoFind 获取多条数据演示
func (cls *BaseModel) demoFind() {
	var b BaseModel
	var ctx *gin.Context
	cls.
		SetModel(BaseModel{}).
		SetWhereFields(tools.Strings{"a", "b", "c"}).
		PrepareQuery(ctx).
		Find(&b)
}

// SetModel 设置使用的模型
func (cls *BaseModel) SetModel(model interface{}) *BaseModel {
	cls.model = model
	return cls
}

// SetPreloads 设置Preloads
func (cls *BaseModel) SetPreloads(preloads []string) *BaseModel {
	cls.preloads = preloads
	return cls
}

// SetPreloadsDefault 设置Preloads为默认
func (cls *BaseModel) SetPreloadsDefault() *BaseModel {
	cls.preloads = tools.Strings{clause.Associations}
	return cls
}

// SetSelects 设置Selects
func (cls *BaseModel) SetSelects(selects []string) *BaseModel {
	cls.selects = selects
	return cls
}

// SetOmits 设置Omits
func (cls *BaseModel) SetOmits(omits []string) *BaseModel {
	cls.omits = omits
	return cls
}

// SetWhereFields 设置WhereFields
func (cls *BaseModel) SetWhereFields(whereFields []string) *BaseModel {
	cls.whereFields = whereFields
	return cls
}

// SetNotWhereFields 设置NotWhereFields
func (cls *BaseModel) SetNotWhereFields(notWhereFields []string) *BaseModel {
	cls.notWhereFields = notWhereFields
	return cls
}

// SetIgnoreFields 设置IgnoreFields
func (cls *BaseModel) SetIgnoreFields(ignoreFields []string) *BaseModel {
	cls.ignoreFields = ignoreFields
	return cls
}

// SetWheres 通过Map设置Wheres
func (cls *BaseModel) SetWheres(wheres map[string]interface{}) *BaseModel {
	cls.wheres = wheres
	return cls
}

// SetNotWheres 设置NotWheres
func (cls *BaseModel) SetNotWheres(notWheres map[string]interface{}) *BaseModel {
	cls.notWheres = notWheres
	return cls
}

// BeforeCreate 插入数据前
func (cls *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	cls.CreatedAt = time.Now()
	cls.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
func (cls *BaseModel) BeforeSave(db *gorm.DB) (err error) {
	cls.UpdatedAt = time.Now()
	return
}

// DB 获取对象
func (cls *BaseModel) DB() (dbSession *gorm.DB) {
	dbSession = (&databases.MySql{}).GetMySqlConn()
	return
}

// Prepare 初始化
func (cls *BaseModel) Prepare() (dbSession *gorm.DB) {
	dbSession = (&databases.MySql{}).GetMySqlConn().Where(cls.wheres).Not(cls.notWheres)

	if cls.model != nil {
		dbSession = dbSession.Model(&cls.model)
	}

	// 拼接preloads关系
	if cls.preloads != nil && len(cls.preloads) > 0 {
		for _, v := range cls.preloads {
			dbSession = dbSession.Preload(v)
		}
	}

	// 拼接selects字段
	if cls.selects != nil && len(cls.selects) > 0 {
		dbSession = dbSession.Select(cls.selects)
	}

	// 拼接omits字段
	if cls.omits != nil && len(cls.omits) > 0 {
		dbSession = dbSession.Omit(cls.omits...)
	}

	return dbSession
}

// PrepareQuery 根据Query参数初始化
func (cls *BaseModel) PrepareQuery(ctx *gin.Context) *gorm.DB {
	dbSession := cls.Prepare()

	wheres := make(map[string]interface{})
	notWheres := make(map[string]interface{})

	// 拼接需要跳过的字段
	ignoreFields := make(map[string]int8)
	if len(cls.ignoreFields) > 0 {
		for _, v := range cls.ignoreFields {
			ignoreFields[v] = 1
		}
	}

	// 拼接Where条件
	for _, v := range cls.whereFields {
		if _, ok := ignoreFields[v]; !ok {
			if val, ok := ctx.GetQuery(v); ok {
				wheres[v] = val
			}
		}
	}

	// 拼接NotWhere条件
	for _, v := range cls.notWhereFields {
		if _, ok := ignoreFields[v]; !ok {
			if val, ok := ctx.GetQuery(v); ok == true {
				notWheres[v] = val
			}
		}
	}
	dbSession = dbSession.Where(wheres).Not(notWheres)

	// 排序
	if order, ok := ctx.GetQuery("__order__"); ok {
		dbSession.Order(order)
	}

	// offset
	if offset, ok := ctx.GetQuery("__offset__"); ok {
		offset := tools.ThrowErrorWhenIsNotInt(offset, "offset参数只能填写整数")
		dbSession.Offset(offset)
	}

	// limit
	if limit, ok := ctx.GetQuery("__limit__"); ok {
		limit := tools.ThrowErrorWhenIsNotInt(limit, "limit参数只能填写整数")
		dbSession.Limit(limit)
	}

	return dbSession
}
