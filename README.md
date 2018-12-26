# b2c-deploy

# 参数说明
* -f app.ini配置文件
* -kubeConfs kubeconfig目录
* -log 日志目录
* -imagePullSecrets 下载镜像密钥


# 调用k8s服务部署进行部署删除服务等操作
* 创建服务
* 删除服务
* 更新服务
* 重启服务
* 获取服务
* 更改实例数量(待完善)


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

## 获取服务
### 请求
```
curl -X OPTIONS -H "Content-Type: application/json"  -d '{
    "serviceName":"consul-test",
    "kubeType":"taoche-test"
}' "http://192.168.177.224:9000/deploy/service"
```
### 响应
```
{
  "code": 200,
  "data": {
    "ServiceName": "consul-client",
    "Image": "consul:1.3.0",
    "InstanceNum": 3,
    "Namespace": "default",
    "Running": true
  },
  "msg": "SUCCESS"
}
```
* ServiceName:服务名
* Image:服务镜像
* InstanceNum:应启动pod数量
* Namespace:命名空间
* Running:运行状态



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
|volume|挂载目录|n|[]|[{"name":"logs","hostPath":"/data1/logs/consul-client/"}]|
|volumeMount|挂载目录至容器|n|[]|[{"name":"logs","mountPath":"/mnt"}]|
|ports|端口映射，当该值有值时，port&targetPort失效|n|[]|[{"port":8080,"targetPort":8080}]|
|args|启动参数|n|[]|["agent","-ui","-client=0.0.0.0","-join=192.168.177.224"]|

# 公共响应结构
```
{
  "code": 500,
  "data": null,
  "msg": "镜像不能为空"
}
```
* code: 非200表示失败
* data: 响应内容
* msg: code 描述