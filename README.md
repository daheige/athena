# athena
go web/job/rpc framework
- web 使用gin http框架
- job 使用corba框架
- rpc 使用grpc框架

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
```

运行job
```shell
cd cmd/job
go build -o athena-job main.go
cp ../../app.yaml ./
./athena-job version
```

# monitor服务监控

