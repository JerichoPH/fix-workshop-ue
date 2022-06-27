package tools

import (
	"fix-workshop-go/errors"
	"fmt"
	"reflect"
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

// ThrowErrorWhenIsEmpty 判断是否为空
func ThrowErrorWhenIsEmpty(ins interface{}, class interface{}, name string) (isEmpty bool) {
	isEmpty = reflect.DeepEqual(ins, class)

	if name != "" && isEmpty {
		panic(errors.ThrowEmpty(fmt.Sprintf("%v不存在", name)))
	}

	return
}

// ThrowErrorWhenIsRepeat 判断是否重复
func ThrowErrorWhenIsRepeat(ins interface{}, class interface{}, name string) (isRepeat bool) {
	isRepeat = !reflect.DeepEqual(ins, class)

	if name != "" && isRepeat {
		panic(errors.ThrowForbidden(fmt.Sprintf("%v重复", name)))
	}

	return
}
