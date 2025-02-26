# etcd实现服务注册和发现
    etcd是一个开源的、高可用的分布式key-value存储系统，可以用于配置共享和服务的注册和发现。

# 本地运行etcd
```shell
docker run -d \
  --name etcd_test \
  --restart=always \
  -p 2379:2379 \
  -p 2380:2380 \
  quay.io/coreos/etcd:v3.5.1 \
  /usr/local/bin/etcd \
  --name etcd_test \
  --data-dir /etcd-data \
  --advertise-client-urls http://localhost:2379 \
  --listen-client-urls http://0.0.0.0:2379
```

# 服务注册和发现使用
参考 etcd_test.go 文件