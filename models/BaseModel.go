package models

import (
	"fix-workshop-ue/databases"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

// BaseModel 出厂数据、财务数据、检修数据、仓储数据、流转数据、运用数据
type BaseModel struct {
	Id                 uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time      `gorm:"timestamptz" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"timestamptz" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Uuid               string         `gorm:"type:VARCHAR(36);COMMENT:uuid;" json:"uuid"`
	Sort               int64          `gorm:"type:BIGINT;DEFAULT:0;COMMENT:排序;" json:"sort"`
	ctx                *gin.Context
	preloads           []string
	selects            []string
	omits              []string
	whereFields        []string
	notWhereFields     []string
	ignoreFields       []string
	distinctFieldNames []string
	wheres             map[string]interface{}
	notWheres          map[string]interface{}
	extraWheres        map[string]func(string, *gorm.DB) *gorm.DB
	scopes             []func(*gorm.DB) *gorm.DB
	model              interface{}
}

// BootByModel 通过model启动
//  @param model
//  @return *BaseModel
func BootByModel(model interface{}) *BaseModel {
	return new(BaseModel).SetModel(model)
}

// Boot 启动
func Boot() *BaseModel {
	return new(BaseModel)
}

// Pagination 分页器
func Pagination(dbSession *gorm.DB, ctx *gin.Context) *gorm.DB {
	if pageStr, ok := ctx.GetQuery("__page__"); ok {
		limitStr := ctx.DefaultQuery("__limit__", "50")
		limit, _ := strconv.Atoi(limitStr)
		page, _ := strconv.Atoi(pageStr)
		dbSession.Offset((page - 1) * limit).Limit(limit)
	}

	return dbSession
}

// demoFindOne 获取单条数据演示
//  @receiver cls
func (cls *BaseModel) demoFindOne() {
	var b BaseModel
	ret := cls.
		SetModel(BaseModel{}).
		SetWheres(tools.Map{}).
		SetNotWheres(tools.Map{}).
		Prepare("").
		First(b)
	wrongs.PanicWhenIsEmpty(ret, "XX")
}

// demoFind 获取多条数据演示
//  @receiver cls
func (cls *BaseModel) demoFind() {
	var b BaseModel
	var ctx *gin.Context
	cls.
		SetModel(BaseModel{}).
		SetWhereFields("a", "b", "c").
		PrepareUseQuery(ctx, "").
		Find(&b)
}

// ScopeBeEnableTrue 启用（查询域）
//  @receiver cls
//  @param db
//  @return *gorm.DB
func (BaseModel) ScopeBeEnableTrue(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is true")
}

// ScopeBeEnableFalse 不启用（查询域）
//  @receiver BaseModel
//  @param db
//  @return *gorm.DB
func (BaseModel) ScopeBeEnableFalse(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is false")
}

// SetCtx 设置Context
func (cls *BaseModel) SetCtx(ctx *gin.Context) *BaseModel {
	cls.ctx = ctx

	return cls
}

// SetModel 设置使用的模型
//  @receiver cls
//  @param model
//  @return *BaseModel
func (cls *BaseModel) SetModel(model interface{}) *BaseModel {
	cls.model = model
	return cls
}

// SetDistinct 设置不重复字段
func (cls *BaseModel) SetDistinct(distinctFieldNames ...string) *BaseModel {
	cls.distinctFieldNames = distinctFieldNames

	return cls
}

// SetPreloads 设置Preloads
//  @receiver cls
//  @param preloads
//  @return *BaseModel
func (cls *BaseModel) SetPreloads(preloads ...string) *BaseModel {
	cls.preloads = preloads
	return cls
}

// SetPreloadsByDefault 设置Preloads为默认
//  @receiver cls
//  @return *BaseModel
func (cls *BaseModel) SetPreloadsByDefault() *BaseModel {
	cls.preloads = tools.Strings{clause.Associations}
	return cls
}

// SetSelects 设置Selects
//  @receiver cls
//  @param selects
//  @return *BaseModel
func (cls *BaseModel) SetSelects(selects ...string) *BaseModel {
	cls.selects = selects
	return cls
}

// SetOmits 设置Omits
//  @receiver cls
//  @param omits
//  @return *BaseModel
func (cls *BaseModel) SetOmits(omits ...string) *BaseModel {
	cls.omits = omits
	return cls
}

// SetWhereFields 设置WhereFields
//  @receiver cls
//  @param whereFields
//  @return *BaseModel
func (cls *BaseModel) SetWhereFields(whereFields ...string) *BaseModel {
	cls.whereFields = whereFields
	return cls
}

// SetNotWhereFields 设置NotWhereFields
//  @receiver cls
//  @param notWhereFields
//  @return *BaseModel
func (cls *BaseModel) SetNotWhereFields(notWhereFields ...string) *BaseModel {
	cls.notWhereFields = notWhereFields
	return cls
}

// SetIgnoreFields 设置IgnoreFields
//  @receiver cls
//  @param ignoreFields
//  @return *BaseModel
func (cls *BaseModel) SetIgnoreFields(ignoreFields ...string) *BaseModel {
	cls.ignoreFields = ignoreFields
	return cls
}

// SetWheres 通过Map设置Wheres
//  @receiver cls
//  @param wheres
//  @return *BaseModel
func (cls *BaseModel) SetWheres(wheres map[string]interface{}) *BaseModel {
	cls.wheres = wheres
	return cls
}

// SetNotWheres 设置NotWheres
//  @receiver cls
//  @param notWheres
//  @return *BaseModel
func (cls *BaseModel) SetNotWheres(notWheres map[string]interface{}) *BaseModel {
	cls.notWheres = notWheres
	return cls
}

// SetScopes 设置Scopes
//  @receiver cls
//  @param scopes
//  @return *BaseModel
func (cls *BaseModel) SetScopes(scopes ...func(*gorm.DB) *gorm.DB) *BaseModel {
	cls.scopes = scopes
	return cls
}

// SetExtraWheres 设置额外搜索条件字段
func (cls *BaseModel) SetExtraWheres(extraWheres map[string]func(string, *gorm.DB) *gorm.DB) *BaseModel {
	cls.extraWheres = extraWheres

	return cls
}

// BeforeCreate 插入数据前
//  @receiver cls
//  @param db
//  @return err
func (cls *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	cls.CreatedAt = time.Now()
	cls.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
//  @receiver cls
//  @param db
//  @return err
func (cls *BaseModel) BeforeSave(db *gorm.DB) (err error) {
	cls.UpdatedAt = time.Now()
	return
}

// Prepare 初始化
//  @receiver cls
//  @param dbDriver
//  @return query
func (cls *BaseModel) Prepare(dbDriver string) (query *gorm.DB) {
	query = (&databases.Launcher{DbDriver: dbDriver}).GetDatabaseConn()

	query = query.Where(cls.wheres).Not(cls.notWheres)

	if cls.model != nil {
		query = query.Model(&cls.model)
	}

	// 设置scopes
	if len(cls.scopes) > 0 {
		query = query.Scopes(cls.scopes...)
	}

	// 拼接preloads关系
	if len(cls.preloads) > 0 {
		for _, v := range cls.preloads {
			query = query.Preload(v)
		}
	}

	// 拼接distinct
	if len(cls.distinctFieldNames) > 0 {
		query = query.Distinct(cls.distinctFieldNames)
	}

	// 拼接selects字段
	if len(cls.selects) > 0 {
		query = query.Select(cls.selects)
	}

	// 拼接omits字段
	if len(cls.omits) > 0 {
		query = query.Omit(cls.omits...)
	}

	return query
}

// PrepareUseQuery 根据Query参数初始化
//  @receiver cls
//  @param ctx
//  @param dbDriver
//  @return *gorm.DB
func (cls *BaseModel) PrepareUseQuery(ctx *gin.Context, dbDriver string) *gorm.DB {
	dbSession := cls.Prepare(dbDriver)

	wheres := make(map[string]interface{})
	notWheres := make(map[string]interface{})

	// 拼接需要跳过的字段
	ignoreFields := make(map[string]int32)
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

	// 拼接额外搜索条件
	for fieldName, v := range cls.extraWheres {
		if _, ok := ignoreFields[fieldName]; !ok {
			dbSession = v(fieldName, dbSession)
		}
	}

	// 排序
	if order, ok := ctx.GetQuery("__order__"); ok {
		dbSession.Order(order)
	} else {
		dbSession.Order("id asc, sort asc")
	}

	return dbSession
}

// PrepareByDefaultDbDriver 通过默认数据库初始化
//  @receiver cls
//  @return dbSession
func (cls *BaseModel) PrepareByDefaultDbDriver() (query *gorm.DB) {
	return cls.Prepare("")
}

// PrepareUseQueryByDefaultDbDriver 通过默认数据库初始化
//  @receiver cls
//  @param ctx
//  @return dbSession
func (cls *BaseModel) PrepareUseQueryByDefaultDbDriver(ctx *gin.Context) (query *gorm.DB) {
	return cls.PrepareUseQuery(ctx, "")
}

// BaseOption 基础查询条件
type BaseOption struct {
	Preloads       []string
	Selects        []string
	Omits          []string
	WhereFields    []string
	NotWhereFields []string
	IgnoreFields   []string
	Wheres         map[string]interface{}
	NotWheres      map[string]interface{}
	Scopes         []func(*gorm.DB) *gorm.DB
	Model          interface{}
}
