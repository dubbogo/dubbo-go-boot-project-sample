[English](./README_CN.md)

# github.com/dubbogo/dubbo-go-boot-project-sample

## 项目目录结构
```
|- .run -------------------- Intellij运行配置
|- cmd --------------------- 程序入口
|- controller -------------- 控制器层
   |- controller_entry.go -- Web服务路由处理管理工具
|- dto --------------------- 数据传输对象层
|- resources --------------- 程序配置及静态资源
   |- application.yml ------ 程序配置文件
|- service ----------------- 服务层，包含了公共服务及Dubbo服务
```

## 启动项目

```shell
# 添加环境变量
set CONF_APPLICATION_FILE_PATH=项目地址/resources/application.yml
go run cmd/main.go
```

## 组件

| 名称                                | 包                                                                                                                                                | 描述            | 版本号 |
|:----------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------|:--------------|:---:|
| dubbo-go-boot-starter             | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter)                         | 项目启动器         |     |
| dubbo-go-boot-middleware-database | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-database](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-database) | 数据库中间件        |     |
| dubbo-go-boot-middleware-redis    | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-redis](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-redis)       | 内存缓存中间件       |     |
| dubbo-go-boot-middleware-web      | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-web](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-web)           | Web服务中间件      |     |
| dubbo-go-boot-middleware-dubbo    | [github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-dubbo](https://github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-middleware-dubbo)       | dubbo-go服务中间件 |     |
