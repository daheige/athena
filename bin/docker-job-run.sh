#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-job
version=v1.0
cd $root_dir

container_name=athena-job
container=$(docker ps -a | grep $container_name | awk '{print $1}')
if [ ${#container} -gt 0 ]; then
    docker rm -f $container_name
fi

docker run --name=$container_name -v $root_dir/config/app.yaml:/app/app.yaml -itd \
   athena-project-job:v1.0
