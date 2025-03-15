#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-rpc
version=v1.0
cd $root_dir

container_name=athena-rpc
container=$(docker ps -a | grep $container_name | awk '{print $1}')
if [ ${#container} -gt 0 ]; then
    docker rm -f $container_name
fi

docker run --name=$container_name -p 3337:3337 -p 8081:8081 -v $root_dir/config/app.yaml:/app/app.yaml -itd \
   $image_name:$version
