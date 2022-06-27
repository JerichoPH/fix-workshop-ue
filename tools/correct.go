package tools

import "sync"

type correct struct {
	m  string
	c  interface{}
	s  uint
	ec uint
}

var responseIns *correct
var correctOnce sync.Once

func CorrectIns(msg string) *correct {
	correctOnce.Do(func() { responseIns = &correct{m: ""} })
	responseIns.m = msg
	return responseIns
}

func (cls *correct) get() map[string]interface{} {
	ret := map[string]interface{}{
		"message":    cls.m,
		"content":    cls.c,
		"status":     cls.s,
		"error_code": cls.ec,
	}
	return ret
}

func (cls *correct) set(content interface{}, status uint, errorCode uint) *correct {
	cls.c = content
	if status == 0 {
		cls.s = 200
	} else {
		cls.s = status
	}
	cls.ec = errorCode
	return cls
}

func (cls *correct) OK(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "OK"
	}
	return 200, cls.set(content, 200, 0).get()
}

func (cls *correct) Created(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "新建成功"
	}
	return 201, cls.set(content, 201, 0).get()
}

func (cls *correct) Updated(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "编辑成功"
	}

	return 202, cls.set(content, 202, 0).get()
}

func (cls *correct) Deleted() (int, interface{}) {
	if cls.m == "" {
		cls.m = "删除成功"
	}
	return 204, cls.set(nil, 204, 0).get()
}
