# entity
存放表结构体信息,这里的go代码是自动生成的，尽量不要编辑

# tbox工具安装
```shell
go install github.com/daheige/tbox/cmd/tbox@latest
```
生成表结构体信息：
```shell
# 需要切换到 domain 目录，再运行如下命令
cd ../
tbox -dsn="root:root123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4" -t=user -p=entity -d=entity -tag=gorm -j=true
```
使用方式参考：https://github.com/daheige/tbox

# 数据表SQL
```
create database test;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名字',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
