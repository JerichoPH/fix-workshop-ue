package v1

import (
	"fix-workshop-ue/middlewares"
	"github.com/gin-gonic/gin"
)

type OrganizationWorkshopRouter struct{}

func (OrganizationWorkshopRouter) Load(router *gin.Engine) {
	r := router.Group(
		"api/v1/organization",
		middlewares.CheckJwt(),
		middlewares.CheckPermission(),
	)
	{
		r.GET("workshop", func(ctx *gin.Context) {})
	}
}
