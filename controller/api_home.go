package controller

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/dubbogo/dubbo-go-boot-project-sample/dto"
	"github.com/dubbogo/dubbo-go-boot-project-sample/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiPingLocal(ctx *gin.Context) {
	result := service.ApiPing()
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func ApiPingRemote(ctx *gin.Context) {
	result, err := service.ApiPingService.Ping(context.Background(), &dto.ApiPingRequestDTO{})
	if err != nil {
		logger.Errorf(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": "ERROR",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "SUCCESS",
		"result": result,
	})
}
