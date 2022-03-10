package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	databaseMiddleware "github.com/dubbogo/dubbo-go-boot-middleware-database/middleware"
	dubboMiddleware "github.com/dubbogo/dubbo-go-boot-middleware-dubbo/middleware"
	redisMiddleware "github.com/dubbogo/dubbo-go-boot-middleware-redis/middleware"
	webMiddleware "github.com/dubbogo/dubbo-go-boot-middleware-web/middleware"
	"github.com/dubbogo/dubbo-go-boot-project-sample/controller"
	starter "github.com/dubbogo/dubbo-go-boot-starter"
	"github.com/gin-gonic/gin"
)

func onWebMiddlewareSetup(router *gin.Engine) {
	controller.SetRouter(router)
	logger.Info("web middleware setup")
}

func onDatabaseMiddlewareSetup() {
	logger.Info("database middleware setup")
}

func onRedisMiddlewareSetup() {
	logger.Info("redis middleware setup")
}

func onDubboMiddlewareSetup() {
	logger.Info("dubbo-go middleware setup")
}

// export CONF_APPLICATION_FILE_PATH=ProjectPath/resources/application.yml
func main() {
	err := starter.NewStarter().AddMiddlewareSetupHooks(
		webMiddleware.NewWebSetupHook(onWebMiddlewareSetup),
		databaseMiddleware.NewDatabaseSetupHook(onDatabaseMiddlewareSetup),
		redisMiddleware.NewRedisSetupHook(onRedisMiddlewareSetup),
		dubboMiddleware.NewDubboSetupHook(onDubboMiddlewareSetup),
	).Start()
	if err != nil {
		logger.Error(err)
		return
	}
}
