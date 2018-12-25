package reqBody


type ServiceRequest struct {
	ServiceName string
	InstanceNum int32
	Image string
	Port int32
	TargetPort int
	KubeType string
	Namespace string
	LimitCpu string
	LimitMemory string
	RequestCpu string
	RequestMemory string
	Env map[string]string
	Nodes map[string]string
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
	return request
}
