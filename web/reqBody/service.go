package reqBody


type ServiceRequest struct {
	ServiceName string// 服务名称
	InstanceNum int32 // 实例数量
	Image string  // 镜像
	Port int32  // pod端口
	TargetPort int  // 容器端口
	KubeType string  // 集群类型
	Namespace string  // 命名空间
	LimitCpu string // cpu
	LimitMemory string  // 内存
	RequestCpu string
	RequestMemory string
	Env map[string]string  // 环境变量
	Nodes map[string]string  // 部署节点
	Volume []Volume   // 挂载目录
	VolumeMount []VolumeMount
	Ports []PortMap  // 多端口映射
	Args []string  // 启动参数

}

type PortMap struct{
	Port int32
	TargetPort int
	Type string "TCP"
}

type Volume struct {
	Name string
	HostPath string
}

type VolumeMount struct{
	Name string
	MountPath string
}

func InitServiceRequest()ServiceRequest{
	var request ServiceRequest
	request.Port = 8080
	request.TargetPort=8080
	request.Namespace="default"
	request.InstanceNum= 1
	request.KubeType="test"
	request.RequestCpu= "0.2"
	request.LimitCpu="1"
	request.RequestMemory="1Gi"
	request.LimitMemory="2Gi"
	request.Env = map[string] string{}
	request.Nodes = map[string] string{}
	request.Volume = []Volume{}
	request.VolumeMount = []VolumeMount{}
	request.Ports = []PortMap{}
	request.Args = []string{}
	return request
}