package wrongs

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Wrong struct{ errorMessage string }

// EmptyWrong 空数据异常
type EmptyWrong struct{ Wrong }

// Error 获取异常信息
//  @receiver cls
//  @return string
func (cls *Wrong) Error() string {
	return cls.errorMessage
}

// ValidateWrong 表单验证错误
type ValidateWrong struct{ Wrong }

// PanicValidate 421错误
//  @param text
func PanicValidate(text string) {
	panic(&ValidateWrong{Wrong{errorMessage: text}})
}

// PanicEmpty 404错误
//  @param text
//  @return error
func PanicEmpty(text string) {
	panic(&EmptyWrong{Wrong{errorMessage: text}})
}

// ForbiddenWrong
type ForbiddenWrong struct{ Wrong }

// PanicForbidden 403错误
//  @param text
//  @return error
func PanicForbidden(text string) {
	panic(&ForbiddenWrong{Wrong{errorMessage: text}})
}

// UnAuthWrong 未授权异常
type UnAuthWrong struct{ Wrong }

// PanicUnAuth 未授权错误
//  @param text
//  @return error
func PanicUnAuth(text string) {
	panic(&UnAuthWrong{Wrong{errorMessage: text}})
}

// UnLoginWrong 未登录异常
type UnLoginWrong struct{ Wrong }

// PanicUnLogin 未登录错误
//  @param text
func PanicUnLogin(text string) {
	panic(&UnLoginWrong{Wrong{errorMessage: text}})
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
