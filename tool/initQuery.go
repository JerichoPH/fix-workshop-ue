package tool

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryBuilder struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Init 初始化查询器
func (cls *QueryBuilder) Init(w, n map[string]interface{}) *gorm.DB {
	tx := cls.DB.Where(w).Not(n)

	// 排序
	if order := cls.CTX.Query("order"); order != "" {
		tx.Order(order)
	}

	// offset
	if offset := cls.CTX.Query("offset"); offset != "" {
		offset := ThrowErrorWhenIsNotInt(offset, "offset参数只能填写整数")
		tx.Offset(offset)
	}

	// limit
	if limit := cls.CTX.Query("limit"); limit != "" {
		limit := ThrowErrorWhenIsNotInt(limit, "limit参数只能填写整数")
		tx.Limit(limit)
	}

	return tx.Preload("~~~as~~~")
}
