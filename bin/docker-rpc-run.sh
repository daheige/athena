#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-rpc
version=v1.0
cd $root_dir

docker rm -f athena-rpc
docker run --name=athena-rpc -p 3337:3337 -p 8081:8081 -v $root_dir/config/app.yaml:/app/app.yaml -itd \
   athena-project-rpc:v1.0
