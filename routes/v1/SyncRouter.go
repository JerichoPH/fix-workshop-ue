package v1

import (
	"fix-workshop-ue/controllers"
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

type SyncRouter struct{}

func (SyncRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/v1/sync",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		// 同步仓库位置（段中心 → 检修车间）
		r.POST("positionDepotFromParagraphCenter", func(ctx *gin.Context) { new(controllers.SyncController).PostPositionDepotFromParagraphCenter(ctx) })
	}
}
