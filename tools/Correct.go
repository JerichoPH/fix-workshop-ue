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
	responseIns.msg = "OK"
	return responseIns
}

func (ins *correct) get() map[string]interface{} {
	ret := map[string]interface{}{
		"msg":        ins.msg,
		"content":    ins.content,
		"pagination": ins.pagination,
		"status":     ins.status,
		"error_code": ins.errorCode,
	}
	return ret
}

func (ins *correct) set(content, pagination interface{}, status uint, errorCode uint) *correct {
	ins.content = content
	ins.pagination = pagination
	if status == 0 {
		ins.status = 200
	} else {
		ins.status = status
	}
	ins.errorCode = errorCode
	return ins
}

// Ok 读取成功
func (ins *correct) Ok(content interface{}) (int, map[string]interface{}) {
	if ins.msg == "" {
		ins.msg = "OK"
	}
	return 200, ins.set(content, nil, 200, 0).get()
}

// OkForPagination 返回分页数据
func (ins *correct) OkForPagination(content interface{}, pageStr string, count int64) (int, map[string]interface{}) {
	if ins.msg == "" {
		ins.msg = "OK"
	}

	page, _ := strconv.Atoi(pageStr)

	return 200, ins.set(content, map[string]interface{}{"page": page, "previous": page - 1, "next": page + 1, "count": count}, 200, 0).get()
}

// Created 新建成功
func (ins *correct) Created(content interface{}) (int, map[string]interface{}) {
	if ins.msg == "" {
		ins.msg = "新建成功"
	}
	return 201, ins.set(content, nil, 201, 0).get()
}

// Updated 更新成功
func (ins *correct) Updated(content interface{}) (int, map[string]interface{}) {
	if ins.msg == "" {
		ins.msg = "编辑成功"
	}

	return 202, ins.set(content, nil, 202, 0).get()
}

// Deleted 删除成功
func (ins *correct) Deleted() (int, interface{}) {
	if ins.msg == "" {
		ins.msg = "删除成功"
	}
	return 204, ins.set(nil, nil, 204, 0).get()
}
