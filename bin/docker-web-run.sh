#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-web
version=v1.0
cd $root_dir

docker rm -f athena-web
docker run --name=athena-web -p 1337:1337 -p 2337:2337 -v $root_dir/config/app.yaml:/app/app.yaml -itd \
   athena-project-web:v1.0
