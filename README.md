# athena
go web/job/rpc framework
- web 使用gin http框架，参考文档：https://github.com/gin-gonic/gin
- job 使用cobra框架，参考文档：https://cobra.dev/
- rpc 使用grpc框架(通过gmicro框架: https://github.com/daheige/gmicro 定制化开发，支持validator、prometheus接入，同时支持go、rust、php、nodejs等语言的客户端代码生成)
- gateway 用于grpc微服务gateway http proxy请求转发（这个需要提前启动好grpc微服务，然后再运行该gateway http proxy服务）

# 为什么我要开源这个项目
从2017年开始写go，到现在已经有7个年头了，我深刻体会到要把go语言（这里暂时不讨论go runtime底层和go plan9的方方面面等）在实际项目中快速用起来还是有一定的难度的。因此，这7年来我一直在探索、验证、总结经验。
我开源这个项目的目的是：让go开发者或者想转go语言的开发者能够在短时间内快速上手web/rpc/job等实战开发（少走一些弯路，让开发更加顺畅），仅此而已。

# 支持的特性
目前该项目支持config配置读取、validator参数校验（web和grpc都支持）、logger日志记录、metrics/prometheus接入（服务监控）、grpc微服务和gateway http服务，以及MySQL和redis初始化和数据读写等操作。
至于trace功能，后续再持续加入。

# grpc相关工具
    见bin目录中的shell脚本

# linux centos环境安装protoc工具
https://github.com/daheige/rs-rpc?tab=readme-ov-file#centos7-install-protoc

# 开始运行
1. 配置好app.yaml文件，可以根据实际情况更改。
2. 执行如下命令分别启动web,job,rpc服务。

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

# 本地运行etcd
```shell
docker run -d \
  --name etcd_test \
  --restart=always \
  -p 12379:2379 \
  -p 12380:2380 \
  quay.io/coreos/etcd:v3.5.1 \
  /usr/local/bin/etcd \
  --name etcd_test \
  --data-dir /etcd-data \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-client-urls http://0.0.0.0:2379
```
当etcd运行后，就可以修改app.yaml配置文件的服务发现和注册如下：
```yaml
  # 服务注册和发现配置
  enable_discovery: true # 是否开启服务发现和注册，本地开发时可以设置为false
  discovery:
    target_type: etcd
    endpoints:
      - "127.0.0.1:12379"
```
随后就可以运行rpc应用程序

# 查看etcd注册的服务列表
服务启动后，执行如下命令进入etcd容器中
```shell
docker exec -it etcd_test /bin/bash
```
接着执行如下命令获取服务列表
```shell
etcdctl get athena/registry-etcd/athena_grpc --prefix
```
运行效果如下：
```
etcdctl get athena/registry-etcd/athena_grpc --prefix
athena/registry-etcd/athena_grpc/02bde795-ef8e-4594-bff8-a21dd07c97a7
{
"name":"athena_grpc","address":"localhost:8081",
"instance_id":"02bde795-ef8e-4594-bff8-a21dd07c97a7",
"version":"v1","created":"2025-03-05 21:54:18","metadata":null
}
```

# rust语言的grpc微服务解决方案
https://github.com/daheige/rs-rpc
