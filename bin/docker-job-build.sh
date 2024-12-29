#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=athena-project-job
version=v1.0
cd $root_dir
docker build . -f Dockerfile-job -t $image_name:$version