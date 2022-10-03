package wrongs

import (
	"sync"
)

type inCorrect struct {
	msg       string
	content   interface{}
	status    int
	errorCode int
}

var responseIns *inCorrect
var once sync.Once

func InCorrectIns() *inCorrect {
	once.Do(func() { responseIns = &inCorrect{} })
	return responseIns
}

func (ins *inCorrect) Get() map[string]interface{} {
	ret := map[string]interface{}{
		"msg":        ins.msg,
		"content":    ins.content,
		"status":     ins.status,
		"error_code": ins.errorCode,
	}
	return ret
}

func (ins *inCorrect) Set(msg string, content interface{}, status int, errorCode int) *inCorrect {
	ins.msg = msg
	ins.content = content
	if status == 0 {
		ins.status = 200
	} else {
		ins.status = status
	}
	ins.errorCode = errorCode
	return ins
}

func (ins *inCorrect) UnAuthorization(msg string) (int, interface{}) {
	if msg == "" {
		msg = "未授权"
	}
	return 406, ins.Set(msg, map[string]interface{}{}, 406, 1).Get()
}

func (ins *inCorrect) ErrUnLogin() (int, map[string]interface{}) {
	return 401, ins.Set("未登录", map[string]interface{}{}, 401, 2).Get()
}

func (ins *inCorrect) Forbidden(msg string) (int, interface{}) {
	if msg == "" {
		msg = "禁止操作"
	}

	return 403, ins.Set(msg, map[string]interface{}{}, 403, 3).Get()
}

func (ins *inCorrect) Empty(msg string) (int, interface{}) {
	if msg == "" {
		msg = "数不存在"
	}

	return 404, ins.Set(msg, map[string]interface{}{}, 404, 4).Get()
}

func (ins *inCorrect) Validate(msg string, content interface{}) (int, map[string]interface{}) {
	if msg == "" {
		msg = "表单验证错误"
	}

	return 421, ins.Set(msg, content, 421, 5).Get()
}

func (ins *inCorrect) Accident(msg string, err interface{}) (int, map[string]interface{}) {
	if msg == "" {
		msg = "意外错误"
	}
	return 500, ins.Set(msg, err, 500, 6).Get()
}
