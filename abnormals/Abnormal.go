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
func BombEmpty(text string) {
	panic(&EmptyAbnormal{Abnormal{errorMessage: text}})
}

// ForbiddenAbnormal
type ForbiddenAbnormal struct{ Abnormal }

// BombForbidden 403错误
//  @param text
//  @return error
func BombForbidden(text string) {
	panic(&ForbiddenAbnormal{Abnormal{errorMessage: text}})
}

// UnAuthAbnormal 未授权异常
type UnAuthAbnormal struct{ Abnormal }

// BombUnAuth 未授权错误
//  @param text
//  @return error
func BombUnAuth(text string) {
	panic(&UnAuthAbnormal{Abnormal{errorMessage: text}})
}

// UnLoginAbnormal 未登录异常
type UnLoginAbnormal struct{ Abnormal }

// BombUnLogin 未登录错误
//  @param text
//  @return error
func BombUnLogin(text string) error {
	panic(&UnLoginAbnormal{Abnormal{errorMessage: text}})
}

// BombWhenIsNotInt 文字转整型
//  @param v
//  @param errMsg
//  @return intValue
func BombWhenIsNotInt(strValue string, errorMessage string) (intValue int) {
	intValue, err := strconv.Atoi(strValue)
	if err != nil && errorMessage != "" {
		BombForbidden(errorMessage)
	}
	return
}

// BombWhenIsNotUint 文字转无符号整型
//  @param v
//  @param errMsg
//  @return uintValue
func BombWhenIsNotUint(strValue string, errorMessage string) (uintValue uint) {
	intValue := BombWhenIsNotInt(strValue, errorMessage)
	uintValue = uint(intValue)
	return
}

// BombWhenIsEmptyByDB 当数据库返回空则报错
//  @param db
//  @param name
//  @return bool
func BombWhenIsEmptyByDB(db *gorm.DB, errorField string) bool {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			if errorField != "" {
				BombEmpty(errorField + "不存在")
				return false
			} else {
				return false
			}
		}
	}
	return true
}

// BombWhenIsRepeatByDB 当数据库返回不空则报错
//  @param db
//  @param name
//  @return bool
func BombWhenIsRepeatByDB(db *gorm.DB, errorField string) bool {
	if db.Error == nil {
		if errorField != "" {
			BombForbidden(errorField + "重复")
			return false
		} else {
			return false
		}
	}
	return true
}
