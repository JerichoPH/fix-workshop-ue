package web

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) Load(engine *gin.Engine) {
	// 用户与权鉴
	(&AuthorizationRouter{}).Load(engine) // 权鉴

	// Excel测试
	(&ExcelReadTestRouter{}).Load(engine) // Excel测试
}
