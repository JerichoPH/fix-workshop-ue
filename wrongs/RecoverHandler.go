package wrongs

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
				c.JSON(InCorrectIns().Validate("", errorToString(reco)))
			case "*wrongs.ValidateWrong":
				c.JSON(InCorrectIns().Validate(errorToString(reco), map[string]string{}))
			case "*wrongs.ForbiddenWrong":
				// 禁止操作
				c.JSON(InCorrectIns().Forbidden(errorToString(reco)))
			case "*wrongs.EmptyWrong":
				// 空数据
				c.JSON(InCorrectIns().Empty(errorToString(reco)))
			case "*wrongs.UnAuthWrong":
				// 未授权
				c.JSON(InCorrectIns().UnAuthorization(errorToString(reco)))
			case "*wrongs.UnLoginWrong":
				// 未登录
				c.JSON(InCorrectIns().ErrUnLogin())
			default:
				// 其他错误
				c.JSON(InCorrectIns().Accident(errorToString(reco), reco))
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
