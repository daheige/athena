#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-job
version=v1.0
cd $root_dir

docker rm -f athena-job
docker run --name=athena-job -v $root_dir/config/app.yaml:/app/app.yaml -itd \
   athena-project-job:v1.0
