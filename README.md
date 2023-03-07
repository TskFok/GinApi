# GinApi

``````
router:
app/router/router.go
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

获取配置内容:
app/utils/conf/conf.go
``````

``````
数据库:
database.Db.Where(condition).First(&u)

缓存:
cache.Set(key, string(res), 3600)
``````

``````
使用中间件:
//app/middleware/jwt.go
userApi.Use(middleware.Jwt())//userApi的路由组使用jwt的中间件
``````