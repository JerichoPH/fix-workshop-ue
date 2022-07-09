package tools

import (
	"errors"
	"fix-workshop-ue/exceptions"
	"gorm.io/gorm"
	"strconv"
)

type Map map[string]interface{}
type Strings []string

// ThrowExceptionWhenIsNotInt 文字转整型
func ThrowExceptionWhenIsNotInt(v string, errMsg string) (intValue int) {
	intValue, err := strconv.Atoi(v)
	if err != nil && errMsg != "" {
		panic(exceptions.ThrowForbidden(errMsg))
	}
	return
}

// ThrowExceptionWhenIsNotUint 文字转无符号整型
func ThrowExceptionWhenIsNotUint(v string, errMsg string) (uintValue uint) {
	intValue := ThrowExceptionWhenIsNotInt(v, errMsg)
	uintValue = uint(intValue)
	return
}

// ThrowExceptionWhenIsEmptyByDB 当数据库返回空则报错
func ThrowExceptionWhenIsEmptyByDB(db *gorm.DB, name string) bool {
	if db.Error != nil {
		if errors.Is(db.Error,gorm.ErrRecordNotFound){
			if name != "" {
				panic(exceptions.ThrowEmpty(name + "不存在"))
				return false
			} else {
				return false
			}
		}
	}
	return true
}

// ThrowExceptionWhenIsRepeatByDB 当数据库返回不空则报错
func ThrowExceptionWhenIsRepeatByDB(db *gorm.DB, name string) bool {
	if db.Error == nil {
		if name != "" {
			panic(exceptions.ThrowForbidden(name + "重复"))
			return false
		} else {
			return false
		}
	}
	return true
}
