#!/usr/bin/env bash

echo -e "执行go-script"


go get -u github.com/kardianos/govendor

govendor init

govendor list +missing |awk '{print $2}'|xargs go get

go run /go/src/kube-deploy/web/main.go
