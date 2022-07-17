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

// ValidateAbnormal 表单验证错误
type ValidateAbnormal struct{ Abnormal }

// PanicValidate 421错误
//  @param text
func PanicValidate(text string) {
	panic(&ValidateAbnormal{Abnormal{errorMessage: text}})
}

// PanicEmpty 404错误
//  @param text
//  @return error
func PanicEmpty(text string) {
	panic(&EmptyAbnormal{Abnormal{errorMessage: text}})
}

// ForbiddenAbnormal
type ForbiddenAbnormal struct{ Abnormal }

// PanicForbidden 403错误
//  @param text
//  @return error
func PanicForbidden(text string) {
	panic(&ForbiddenAbnormal{Abnormal{errorMessage: text}})
}

// UnAuthAbnormal 未授权异常
type UnAuthAbnormal struct{ Abnormal }

// PanicUnAuth 未授权错误
//  @param text
//  @return error
func PanicUnAuth(text string) {
	panic(&UnAuthAbnormal{Abnormal{errorMessage: text}})
}

// UnLoginAbnormal 未登录异常
type UnLoginAbnormal struct{ Abnormal }

// PanicUnLogin 未登录错误
//  @param text
//  @return error
func PanicUnLogin(text string) error {
	panic(&UnLoginAbnormal{Abnormal{errorMessage: text}})
}

// PanicWhenIsNotInt 文字转整型
//  @param v
//  @param errMsg
//  @return intValue
func PanicWhenIsNotInt(strValue string, errorMessage string) (intValue int) {
	intValue, err := strconv.Atoi(strValue)
	if err != nil && errorMessage != "" {
		PanicForbidden(errorMessage)
	}
	return
}

// PanicWhenIsNotUint 文字转无符号整型
//  @param v
//  @param errMsg
//  @return uintValue
func PanicWhenIsNotUint(strValue string, errorMessage string) (uintValue uint) {
	intValue := PanicWhenIsNotInt(strValue, errorMessage)
	uintValue = uint(intValue)
	return
}

// PanicWhenIsEmpty 当数据库返回空则报错
//  @param db
//  @param name
//  @return bool
func PanicWhenIsEmpty(db *gorm.DB, errorField string) bool {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			if errorField != "" {
				PanicEmpty(errorField + "不存在")
				return false
			} else {
				return false
			}
		}
	}
	return true
}

// PanicWhenIsRepeat 当数据库返回不空则报错
//  @param db
//  @param name
//  @return bool
func PanicWhenIsRepeat(db *gorm.DB, errorField string) bool {
	if db.Error == nil {
		if errorField != "" {
			PanicForbidden(errorField + "重复")
			return false
		} else {
			return false
		}
	}
	return true
}
