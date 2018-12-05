#!/usr/bin/env bash

echo -e "执行go-script"

nohup /go/src/b2c-deploy/docker/main > monitor.log &

go run /go/src/b2c-deploy/web/main.go
