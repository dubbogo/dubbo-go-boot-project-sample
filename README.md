[中文](./README_CN.md)

# github.com/dubbogo/dubbo-go-boot-project-sample

## Directory Structure

```
|- .run -------------------- Intellij run configuration
|- cmd --------------------- Application entry
|- controller -------------- Router handler layer
   |- controller_entry.go -- Web service router handler manager
|- dto --------------------- Data transfer object layer
|- resources --------------- Application config file and static files
   |- application.yml ------ Application config file
|- service ----------------- Service layer, includes common services and dubbo services
```

## Start

```shell
# Add Env
set CONF_APPLICATION_FILE_PATH=ProjectPath/resources/application.yml
go run cmd/main.go
```

## Dependence

| name                         | package                                                                                                                                          | description                 | version |
|:-----------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------|:-------:|
| dubbo-go-starter             | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter)                         | dubbo-go project boot suite |         |
| dubbo-go-middleware-database | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-database](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-database) | database middleware         |         |
| dubbo-go-middleware-redis    | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-redis](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-redis)       | memory cache middleware     |         |
| dubbo-go-middleware-web      | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-web](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-web)           | web service middleware      |         |
| dubbo-go-middleware-dubbo    | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-dubbo](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-dubbo)       | dubbo middleware            |         |

