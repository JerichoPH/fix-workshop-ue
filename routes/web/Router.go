package web

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) Load(engine *gin.Engine) {
	(&AuthorizationRouter{}).Load(engine) // 用户与权鉴
	(&CommandRouter{}).Load(engine)       // Command控制台
}
