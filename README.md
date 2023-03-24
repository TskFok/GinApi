# GinApi

``````
目录

bootstrap:
引导

bin:
入口文件

cmd:
cli文件

command:
命令行

controller:
控制器

router:
路由

runtime:
运行时产生的文件，当前目录下没有runtime目录时，自动在用户的系统目录下生成

server:
服务

utils:
第三方服务

``````

``````
正确返回:
response.Success()

错误返回:
response.Error()

code定义:
app/err/code.go

codeMsg定义:
app/err/msg.go

配置文件:
utils/conf/conf.yaml

初始化配置内容:
utils/conf/init.go

数据库:
global.MysqlClient.Where(condition).First(&u)

缓存:
global.RedisClient.Set(key, string(res), 3600)

使用中间件:
//app/middleware/jwt.go
userApi.Use(middleware.Jwt())//userApi的路由组使用jwt的中间件
``````

``````
cobra

新增命令:
cobra-cli add test

新增test命令的子命令:
cobra-cli add child
修改child.go里的init中的rootCmd->testCmd

额外字段:
Args: cobra.ExactArgs(2)//新增两个额外字段
go run bin/cli/main.go create rule 1 2

设置flags:
ruleCmd.Flags().StringVarP(&name, "name", "n", "", "rule name")//设置name的flag

以下等效:
go run bin/cli/main.go create rule 1 2 --name=a
go run bin/cli/main.go create rule 1 2 --name a
go run bin/cli/main.go create rule 1 2 -n=a
go run bin/cli/main.go create rule 1 2 -n a

打包后使用:
./xxx create rule 1 2 -n a
``````

``````
守护进程:
process.InitProcess()

关闭守护进程:
kill -2 pid

进程日志:
api.log
``````