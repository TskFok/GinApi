# GinApi

``````
bootstrap:
引导

cmd:
main文件

command:
命令行

controller:
控制器

router:
路由

runtime:
运行时产生文件

server:
服务

utils:
第三方服务

``````

``````
正确返回:
ctx.JSON(err.SUCCESS, tool.GetSuccess(something))

错误返回:
ctx.JSON(err.RUNTIME_ERROR, tool.GetErrorInfo(err.ROUTE_CREATE_ERROR))
``````

``````
code定义:
app/err/code.go

codeMsg定义:
app/err/msg.go
``````

``````
配置文件:
app/utils/conf/app.yaml

初始化配置内容:
app/utils/conf/init.go
``````

``````
数据库:
global.MysqlClient.Where(condition).First(&u)

缓存:
global.RedisClient.Set(key, string(res), 3600)
``````

``````
使用中间件:
//app/middleware/jwt.go
userApi.Use(middleware.Jwt())//userApi的路由组使用jwt的中间件
``````