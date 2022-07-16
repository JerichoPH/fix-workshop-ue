package exceptions

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

// ThrowWhenIsNotInt 文字转整型
func ThrowWhenIsNotInt(v string, errMsg string) (intValue int) {
	intValue, err := strconv.Atoi(v)
	if err != nil && errMsg != "" {
		panic(ThrowForbidden(errMsg))
	}
	return
}

// ThrowExceptionWhenIsNotUint 文字转无符号整型
func ThrowExceptionWhenIsNotUint(v string, errMsg string) (uintValue uint) {
	intValue := ThrowWhenIsNotInt(v, errMsg)
	uintValue = uint(intValue)
	return
}

// ThrowWhenIsEmptyByDB 当数据库返回空则报错
func ThrowWhenIsEmptyByDB(db *gorm.DB, name string) bool {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			if name != "" {
				panic(ThrowEmpty(name + "不存在"))
				return false
			} else {
				return false
			}
		}
	}
	return true
}

// ThrowWhenIsRepeatByDB 当数据库返回不空则报错
func ThrowWhenIsRepeatByDB(db *gorm.DB, name string) bool {
	if db.Error == nil {
		if name != "" {
			panic(ThrowForbidden(name + "重复"))
			return false
		} else {
			return false
		}
	}
	return true
}
