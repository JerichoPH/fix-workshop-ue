package abnormals

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Abnormal struct{ errorMessage string }

// EmptyAbnormal 空数据异常
type EmptyAbnormal struct{ Abnormal }

// Error 获取异常信息
//  @receiver cls
//  @return string
func (cls *Abnormal) Error() string {
	return cls.errorMessage
}

// BombEmpty 404错误
//  @param text
//  @return error
func BombEmpty(text string) error {
	return &EmptyAbnormal{Abnormal{errorMessage: text}}
}

// ForbiddenAbnormal
type ForbiddenAbnormal struct{ Abnormal }

// BombForbidden 403错误
//  @param text
//  @return error
func BombForbidden(text string) error {
	return &ForbiddenAbnormal{Abnormal{errorMessage: text}}
}

// UnAuthAbnormal 未授权异常
type UnAuthAbnormal struct{ Abnormal }

// BombUnAuth 未授权错误
//  @param text
//  @return error
func BombUnAuth(text string) error {
	return &UnAuthAbnormal{Abnormal{errorMessage: text}}
}

// UnLoginAbnormal 未登录异常
type UnLoginAbnormal struct{ Abnormal }

// BombUnLogin 未登录错误
//  @param text
//  @return error
func BombUnLogin(text string) error {
	return &UnLoginAbnormal{Abnormal{errorMessage: text}}
}

// BombWhenIsNotInt 文字转整型
func BombWhenIsNotInt(v string, errMsg string) (intValue int) {
	intValue, err := strconv.Atoi(v)
	if err != nil && errMsg != "" {
		panic(BombForbidden(errMsg))
	}
	return
}

// BombWhenIsNotUint 文字转无符号整型
func BombWhenIsNotUint(v string, errMsg string) (uintValue uint) {
	intValue := BombWhenIsNotInt(v, errMsg)
	uintValue = uint(intValue)
	return
}

// BombWhenIsEmptyByDB 当数据库返回空则报错
func BombWhenIsEmptyByDB(db *gorm.DB, name string) bool {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			if name != "" {
				panic(BombEmpty(name + "不存在"))
				return false
			} else {
				return false
			}
		}
	}
	return true
}

// BombWhenIsRepeatByDB 当数据库返回不空则报错
func BombWhenIsRepeatByDB(db *gorm.DB, name string) bool {
	if db.Error == nil {
		if name != "" {
			panic(BombForbidden(name + "重复"))
			return false
		} else {
			return false
		}
	}
	return true
}
