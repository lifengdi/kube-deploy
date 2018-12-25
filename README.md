# b2c-deploy

# 参数说明
* -f app.ini配置文件
* -kubeConfs kubeconfig目录


# 调用k8s服务部署进行部署删除服务等操作
* 创建服务
* 删除服务
* 更新服务
* 重启服务
* 更改服务实例数量
* watch deployment


# api详解

## 创建服务
```
curl -X POST -H "Content-Type: application/json"  -d '{
    "serviceName":"nginx",
    "image":"nginx",
    "port":80,
    "targetPort":80,
    "requestCpu":"0.5",
    "requestMemory":"1Gi"
    "instanceNum":3,
    "kubeType":"test",
    "namespace":"default",
    "limitCpu":1,
    "limitMemory":"2Gi",
    "env":{
            "app":"appname"
    }
}' "http://192.168.177.224:9000/deploy/service"
```

## 删除服务
```


curl -X DELETE -H "Content-Type: application/json" -d
'{
     "serviceName":"nginx"
 }'  "http://localhost:8080/deploy/service"
```

## 更新服务
```
curl -X PUT -H "Content-Type: application/json" -d
'{
    "serviceName":"nginx",
    "image":"nginx",
    "port":80,
    "targetPort":80,
    "requestCpu":"0.5",
    "instanceNum":3
}'  "http://localhost:8080/deploy/service"
```

## 重启服务
```
curl -X PATCH -H "Content-Type: application/json" -d '{
    "serviceName":"kanche-platform-gateway"
}' "http://localhost:8080/deploy/service"
```


# post 请求参数说明
# 创建及修改 请求参数说明
|参数|描述|是否必填|默认值|参考值|
|--|--|--|--|--|--|
|serviceName|服务名称|y||nginx|
|image|服务镜像|y||nginx|
|port|服务端口|n|8080|8080|
|targetPort|容器端口|n|8080|8080|
|requestCpu|cpu|n|0.2|0.2|
|requestMemory|内存大小|n|1Gi|1Gi|
|limitCpu|限制cpu|n|1|1|
|limitMemory|限制内存大小|n|2Gi|1Gi|
|instanceNum|pod实例数量|n|1|3|
|kubeType|集群类型-用于查找kubeconfig|n|test|test|
|namespace|命名空间|n|default|default|
|env|容器环境变量|n|{}|{"app":"appname"}
|nodes|启动容器的节点,与node上的label对应|n|{}|{"attach":"default"}|

