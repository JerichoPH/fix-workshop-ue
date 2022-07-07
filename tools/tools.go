package tools

import (
	"fix-workshop-ue/errors"
	"gorm.io/gorm"
	"strconv"
)

// ThrowErrorWhenIsNotInt 文字转整型
func ThrowErrorWhenIsNotInt(v string, errMsg string) (intValue int) {
	intValue, err := strconv.Atoi(v)
	if err != nil && errMsg != "" {
		panic(errors.ThrowForbidden(errMsg))
	}
	return
}

// ThrowErrorWhenIsNotUint 文字转无符号整型
func ThrowErrorWhenIsNotUint(v string, errMsg string) (uintValue uint) {
	intValue := ThrowErrorWhenIsNotInt(v, errMsg)
	uintValue = uint(intValue)
	return
}

// ThrowErrorWhenIsEmptyByDB 当数据库返回空则报错
func ThrowErrorWhenIsEmptyByDB(db *gorm.DB, name string) bool {
	switch db.Error.Error() {
	case "record not found":
		if name != "" {
			panic(errors.ThrowEmpty(name + "不存在"))
			return false
		} else {
			return false
		}
	}
	return true
}

// ThrowErrorWhenIsRepeatByDB 当数据库返回不空则报错
func ThrowErrorWhenIsRepeatByDB(db *gorm.DB, name string) bool {
	if db.Error == nil {
		if name != "" {
			panic(errors.ThrowForbidden(name + "重复"))
			return false
		} else {
			return false
		}
	}
	return true
}
