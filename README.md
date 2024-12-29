# athena
go web/job/rpc framework
- web 使用gin http框架
- job 使用corba框架
- rpc 使用grpc框架(通过gmicro框架定制化开发，支持validator、prometheus接入)

# 为什么我要开源这个项目
从2017年开始写go，到现在已经有7个年头了，我深刻体会到要把go语言（这里暂时不讨论go runtime底层的方方面面等），在实际项目中快速用起来还是有一定的难度的，因此这7年来我一直在探索、验证、总结经验。
我开源这个项目实战的目的是：让go开发者能够在短时间内快速上手web/rpc/job等实战开发，仅此而已。

# 支持的特性
目录该项目支持config配置读取、validator参数校验（web和grpc都支持）、logger日志记录、metrics/prometheus接入（服务监控）、grpc微服务和gateway http服务，以及MySQL和redis初始化和数据读写等操作。
至于trace功能，后续再持续加入。

# grpc相关工具
    见bin目录中的shell脚本

# 开始运行
1. 配置好app.yaml文件
2. 执行如下命名启动web,job,rpc

运行web
```shell
cd cmd/web
go run main.go
```

运行grpc
```shell
cd cmd/rpc
go run main.go

# 运行客户端
go clients/go/client.go
```

运行job
```shell
cd cmd/job
go build -o athena-job main.go
cp ../../app.yaml ./
./athena-job version
```

# monitor服务监控
web服务监控
- http://localhost:2337/metrics
- http://localhost:2337/debug/pprof

rpc服务监控
- http://localhost:3337/metrics
- http://localhost:3337/debug/pprof

以上监控的服务端口号，可自行更改配置文件

# dockerfile
根据不同的业务场景进行构建，参考bin目录中的docker开头的shell脚本，或者执行Makefile中的命令。
