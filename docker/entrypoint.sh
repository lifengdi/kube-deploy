#!/usr/bin/env bash

echo -e "执行go-script"


go get -u github.com/kardianos/govendor

govendor init

govendor list +missing |awk '{print $2}'|xargs go get

# 获取环境变量，添加启动参数

echo $kubeconfs
echo $f
echo $imagePullSecrets
echo $log

go run /go/src/kube-deploy/web/main.go -kubeconfs $kubeconfs -f $f -imagePullSecrets $imagePullSecrets -log $log
