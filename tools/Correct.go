package tools

import (
	"strconv"
	"sync"
)

type correct struct {
	msg        string
	content    interface{}
	pagination interface{}
	status     uint
	errorCode  uint
}

var responseIns *correct
var correctOnce sync.Once

// CorrectBoot 正确返回值
//  @param msg
//  @return *correct
func CorrectBoot(msg string) *correct {
	correctOnce.Do(func() { responseIns = &correct{msg: ""} })
	responseIns.msg = msg
	return responseIns
}

// CorrectBootByDefault
//  @return *correct
func CorrectBootByDefault() *correct {
	correctOnce.Do(func() { responseIns = &correct{msg: ""} })
	responseIns.msg = ""
	return responseIns
}

func (cls *correct) get() map[string]interface{} {
	ret := map[string]interface{}{
		"msg":        cls.msg,
		"content":    cls.content,
		"pagination": cls.pagination,
		"status":     cls.status,
		"error_code": cls.errorCode,
	}
	return ret
}

func (cls *correct) set(content, pagination interface{}, status uint, errorCode uint) *correct {
	cls.content = content
	cls.pagination = pagination
	if status == 0 {
		cls.status = 200
	} else {
		cls.status = status
	}
	cls.errorCode = errorCode
	return cls
}

// Ok 读取成功
func (cls *correct) Ok(content interface{}) (int, map[string]interface{}) {
	if cls.msg == "" {
		cls.msg = "OK"
	}
	return 200, cls.set(content, nil, 200, 0).get()
}

// OkForPagination 返回分页数据
func (cls *correct) OkForPagination(content interface{}, pageStr string, count int64) (int, map[string]interface{}) {
	if cls.msg == "" {
		cls.msg = "OK"
	}

	page, _ := strconv.Atoi(pageStr)

	return 200, cls.set(content, map[string]interface{}{"page": page, "previous": page - 1, "next": page + 1, "count": count}, 200, 0).get()
}

// Created 新建成功
func (cls *correct) Created(content interface{}) (int, map[string]interface{}) {
	if cls.msg == "" {
		cls.msg = "新建成功"
	}
	return 201, cls.set(content, nil, 201, 0).get()
}

// Updated 更新成功
func (cls *correct) Updated(content interface{}) (int, map[string]interface{}) {
	if cls.msg == "" {
		cls.msg = "编辑成功"
	}

	return 202, cls.set(content, nil, 202, 0).get()
}

// Deleted 删除成功
func (cls *correct) Deleted() (int, interface{}) {
	if cls.msg == "" {
		cls.msg = "删除成功"
	}
	return 204, cls.set(nil, nil, 204, 0).get()
}
