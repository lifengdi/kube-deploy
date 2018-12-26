package responses

type DeploymentResp struct {
	ServiceName string
	Image string
	InstanceNum int32
	Namespace string
	Running bool
}
