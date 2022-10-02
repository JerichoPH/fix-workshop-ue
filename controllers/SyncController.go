package controllers

import (
	"fix-workshop-ue/services/sync"
	"github.com/gin-gonic/gin"
)

type SyncController struct{}

func (SyncController) PostPositionDepotFromParagraphCenter(ctx *gin.Context) {
	(&sync.PositionDepotService{}).FromParagraphCenter(ctx)
}
