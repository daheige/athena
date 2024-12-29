#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

protoExec=$(which "protoc")
if [ -z $protoExec ]; then
    echo 'Please install protoc!'
    echo "Please look readme.md to install proto3"
    echo "if you use centos7,please look https://github.com/daheige/go-proj/blob/master/docs/centos7-protoc-install.md"
    exit 0
fi

# proto文件目录
proto_dir=$root_dir/proto
# 生成的pb文件目录
pb_dir=$root_dir/internal/pb

mkdir -p $pb_dir

#delete old pb code.
rm -rf $root_dir/clients/go/pb/*.pb.go
rm -rf $root_dir/clients/go/pb/*.gw.go

echo "\n\033[0;32mGenerating codes...\033[39;49;0m\n"

echo "generating golang stubs..."
cd $proto_dir

$protoExec -I $proto_dir \
    --go_out $pb_dir --go_opt paths=source_relative \
    --go-grpc_out $pb_dir --go-grpc_opt paths=source_relative \
    $proto_dir/*.proto

protoc -I $proto_dir --grpc-gateway_out $pb_dir \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    $proto_dir/*.proto

# proto文件中的@inject_tag标签注入，主要用于json标签和
sh $root_dir/bin/protoc-inject-tag.sh

# gen request validator code
sh $root_dir/bin/validator-generate.sh

# cp golang client code
mkdir -p $root_dir/clients/go/pb

cp -R $pb_dir/*.go $root_dir/clients/go/pb

echo "generating golang code success"

echo "\n\033[0;32mGenerate codes successfully!\033[39;49;0m\n"

exit 0
