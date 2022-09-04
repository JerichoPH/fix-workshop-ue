package webSocket

import (
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Router struct{}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ParseMessage struct {
	URI     string `json:"uri"`
	Context string `json:"context"`
}

// Load 加载路由
func (Router) Load(engine *gin.Engine) {
	// 长连接
	engine.GET("ws", func(ctx *gin.Context) {
		ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			wrongs.PanicForbidden("长连接错误：" + err.Error())
			return
		}
		defer ws.Close()
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				wrongs.PanicForbidden("接收消息错误：" + err.Error())
			}

			var parseMessage ParseMessage
			if err := ctx.ShouldBindJSON(&parseMessage); err != nil {
				wrongs.PanicForbidden("解析消息错误：" + err.Error())
			} else {
				if parseMessage.URI == "ping" {
					message = []byte("pong")
				}
			}

			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				wrongs.PanicForbidden("相应消息错误：" + err.Error())
			}
		}
	})
}
