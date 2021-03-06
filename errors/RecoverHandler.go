package errors

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoverHandler(c *gin.Context) {
	defer func() {
		if reco := recover(); reco != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", reco)

			// 判断错误类型
			switch fmt.Sprintf("%T", reco) {
			case "*validator.Validationerrors":
				// 表单验证错误
				c.JSON(Ins().Validate("", errorToString(reco)))
			case "*errors.ForbiddenError":
				// 禁止操作
				c.JSON(Ins().Forbidden(errorToString(reco)))
			case "*errors.EmptyError":
				// 空数据
				c.JSON(Ins().Empty(errorToString(reco)))
			case "*errors.UnAuthorizationError":
				// 未授权
				c.JSON(Ins().UnAuthorization(errorToString(reco)))
			case "*errors.UnLoginError":
				// 未登录
				c.JSON(Ins().ErrUnLogin())
			default:
				// 其他错误
				c.JSON(Ins().Accident(errorToString(reco), reco))
				debug.PrintStack() // 打印堆栈信息
			}

			c.Abort()
		}
	}()

	c.Next()
}

// recover错误，转string
func errorToString(recover interface{}) string {
	switch errorType := recover.(type) {
	case error:
		return errorType.Error()
	default:
		return recover.(string)
	}
}
