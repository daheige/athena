DEV_IMAGE_NAME := athena-project-dev
VERSION := v1.0

RPC_IMAGE_NAME := athena-project-rpc
RPC_SERVICE := athena-rpc

WEB_IMAGE_NAME := athena-project-web
WEB_SERVICE := athena-web
go-gen:
	# 生成go grpc相关代码
	sh bin/go-generate.sh
nodejs-gen:
	# 生成nodejs client grpc相关代码
	sh bin/nodejs-generate.sh
php-gen:
	# 生成php client grpc相关代码
	sh bin/nodejs-generate.sh
dev-build:
	docker build . -f Dockerfile-dev -t ${DEV_IMAGE_NAME}:${VERSION}
rpc-build:
	docker build . -f Dockerfile-rpc -t ${RPC_IMAGE_NAME}:${VERSION}
rpc-start: rpc-stop
	docker run --name=${RPC_SERVICE} -p 3337:3337 -p 8081:8081 -v ./config/app.yaml:/app/app.yaml -itd ${RPC_IMAGE_NAME}:v1.0
rpc-stop:
	docker rm -f ${RPC_SERVICE}
rpc-restart: rpc-stop rpc-build rpc-run

web-build:
	docker build . -f Dockerfile-web -t ${WEB_IMAGE_NAME}:${VERSION}
web-start: web-stop
	docker run --name=athena-web -p 1337:1337 -p 2337:2337 -v ./config/app.yaml:/app/app.yaml -itd ${WEB_IMAGE_NAME}:v1.0
web-stop:
	docker rm -f ${WEB_SERVICE}
rpc-restart: web-stop web-build web-start
