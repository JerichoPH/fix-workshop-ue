package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizationRouter struct{}

func (AuthorizationRouter) Load(engine *gin.Engine) {
	r := engine.Group("")
	{
		r.GET("", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/index.html")
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
}
