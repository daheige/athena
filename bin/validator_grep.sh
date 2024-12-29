#!/usr/bin/env bash

root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

echo $root_dir;
proto_dir=$root_dir/proto

grep "\w*@validator=*" $proto_dir/*.proto > $root_dir/validator.log
