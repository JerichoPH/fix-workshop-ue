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
//  @receiver ins
func (ins *BaseModel) demoFindOne() {
	var b BaseModel
	ret := ins.
		SetModel(BaseModel{}).
		SetWheres(tools.Map{}).
		SetNotWheres(tools.Map{}).
		Prepare("").
		First(b)
	wrongs.PanicWhenIsEmpty(ret, "XX")
}

// demoFind 获取多条数据演示
//  @receiver ins
func (ins *BaseModel) demoFind() {
	var b BaseModel
	var ctx *gin.Context
	ins.
		SetModel(BaseModel{}).
		SetWhereFields("a", "b", "c").
		PrepareUseQuery(ctx, "").
		Find(&b)
}

// ScopeBeEnableTrue 启用（查询域）
//  @receiver ins
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
func (ins *BaseModel) SetCtx(ctx *gin.Context) *BaseModel {
	ins.ctx = ctx

	return ins
}

// SetModel 设置使用的模型
//  @receiver ins
//  @param model
//  @return *BaseModel
func (ins *BaseModel) SetModel(model interface{}) *BaseModel {
	ins.model = model
	return ins
}

// SetDistinct 设置不重复字段
func (ins *BaseModel) SetDistinct(distinctFieldNames ...string) *BaseModel {
	ins.distinctFieldNames = distinctFieldNames

	return ins
}

// SetPreloads 设置Preloads
//  @receiver ins
//  @param preloads
//  @return *BaseModel
func (ins *BaseModel) SetPreloads(preloads ...string) *BaseModel {
	ins.preloads = preloads
	return ins
}

// SetPreloadsByDefault 设置Preloads为默认
//  @receiver ins
//  @return *BaseModel
func (ins *BaseModel) SetPreloadsByDefault() *BaseModel {
	ins.preloads = tools.Strings{clause.Associations}
	return ins
}

// SetSelects 设置Selects
//  @receiver ins
//  @param selects
//  @return *BaseModel
func (ins *BaseModel) SetSelects(selects ...string) *BaseModel {
	ins.selects = selects
	return ins
}

// SetOmits 设置Omits
//  @receiver ins
//  @param omits
//  @return *BaseModel
func (ins *BaseModel) SetOmits(omits ...string) *BaseModel {
	ins.omits = omits
	return ins
}

// SetWhereFields 设置WhereFields
//  @receiver ins
//  @param whereFields
//  @return *BaseModel
func (ins *BaseModel) SetWhereFields(whereFields ...string) *BaseModel {
	ins.whereFields = whereFields
	return ins
}

// SetNotWhereFields 设置NotWhereFields
//  @receiver ins
//  @param notWhereFields
//  @return *BaseModel
func (ins *BaseModel) SetNotWhereFields(notWhereFields ...string) *BaseModel {
	ins.notWhereFields = notWhereFields
	return ins
}

// SetIgnoreFields 设置IgnoreFields
//  @receiver ins
//  @param ignoreFields
//  @return *BaseModel
func (ins *BaseModel) SetIgnoreFields(ignoreFields ...string) *BaseModel {
	ins.ignoreFields = ignoreFields
	return ins
}

// SetWheres 通过Map设置Wheres
//  @receiver ins
//  @param wheres
//  @return *BaseModel
func (ins *BaseModel) SetWheres(wheres map[string]interface{}) *BaseModel {
	ins.wheres = wheres
	return ins
}

// SetNotWheres 设置NotWheres
//  @receiver ins
//  @param notWheres
//  @return *BaseModel
func (ins *BaseModel) SetNotWheres(notWheres map[string]interface{}) *BaseModel {
	ins.notWheres = notWheres
	return ins
}

// SetScopes 设置Scopes
//  @receiver ins
//  @param scopes
//  @return *BaseModel
func (ins *BaseModel) SetScopes(scopes ...func(*gorm.DB) *gorm.DB) *BaseModel {
	ins.scopes = scopes
	return ins
}

// SetExtraWheres 设置额外搜索条件字段
func (ins *BaseModel) SetExtraWheres(extraWheres map[string]func(string, *gorm.DB) *gorm.DB) *BaseModel {
	ins.extraWheres = extraWheres

	return ins
}

// BeforeCreate 插入数据前
//  @receiver ins
//  @param db
//  @return err
func (ins *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	ins.CreatedAt = time.Now()
	ins.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
//  @receiver ins
//  @param db
//  @return err
func (ins *BaseModel) BeforeSave(db *gorm.DB) (err error) {
	ins.UpdatedAt = time.Now()
	return
}

// Prepare 初始化
//  @receiver ins
//  @param dbDriver
//  @return query
func (ins *BaseModel) Prepare(dbDriver string) (query *gorm.DB) {
	query = (&databases.Launcher{DbDriver: dbDriver}).GetDatabaseConn()

	query = query.Where(ins.wheres).Not(ins.notWheres)

	if ins.model != nil {
		query = query.Model(&ins.model)
	}

	// 设置scopes
	if len(ins.scopes) > 0 {
		query = query.Scopes(ins.scopes...)
	}

	// 拼接preloads关系
	if len(ins.preloads) > 0 {
		for _, v := range ins.preloads {
			query = query.Preload(v)
		}
	}

	// 拼接distinct
	if len(ins.distinctFieldNames) > 0 {
		query = query.Distinct(ins.distinctFieldNames)
	}

	// 拼接selects字段
	if len(ins.selects) > 0 {
		query = query.Select(ins.selects)
	}

	// 拼接omits字段
	if len(ins.omits) > 0 {
		query = query.Omit(ins.omits...)
	}

	return query
}

// PrepareUseQuery 根据Query参数初始化
//  @receiver ins
//  @param ctx
//  @param dbDriver
//  @return *gorm.DB
func (ins *BaseModel) PrepareUseQuery(ctx *gin.Context, dbDriver string) *gorm.DB {
	dbSession := ins.Prepare(dbDriver)

	wheres := make(map[string]interface{})
	notWheres := make(map[string]interface{})

	// 拼接需要跳过的字段
	ignoreFields := make(map[string]int32)
	if len(ins.ignoreFields) > 0 {
		for _, v := range ins.ignoreFields {
			ignoreFields[v] = 1
		}
	}

	// 拼接Where条件
	for _, v := range ins.whereFields {
		if _, ok := ignoreFields[v]; !ok {
			if val, ok := ctx.GetQuery(v); ok {
				wheres[v] = val
			}
		}
	}

	// 拼接NotWhere条件
	for _, v := range ins.notWhereFields {
		if _, ok := ignoreFields[v]; !ok {
			if val, ok := ctx.GetQuery(v); ok == true {
				notWheres[v] = val
			}
		}
	}
	dbSession = dbSession.Where(wheres).Not(notWheres)

	// 拼接额外搜索条件
	for fieldName, v := range ins.extraWheres {
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
//  @receiver ins
//  @return dbSession
func (ins *BaseModel) PrepareByDefaultDbDriver() (query *gorm.DB) {
	return ins.Prepare("")
}

// PrepareUseQueryByDefaultDbDriver 通过默认数据库初始化
//  @receiver ins
//  @param ctx
//  @return dbSession
func (ins *BaseModel) PrepareUseQueryByDefaultDbDriver(ctx *gin.Context) (query *gorm.DB) {
	return ins.PrepareUseQuery(ctx, "")
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
