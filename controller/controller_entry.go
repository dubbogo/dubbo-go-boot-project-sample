package controller

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	apiGroup := router.Group("/api")
	apiPingGroup := apiGroup.Group("/ping")
	apiPingGroup.GET("local", ApiPingLocal)
	apiPingGroup.GET("remote", ApiPingRemote)
}
