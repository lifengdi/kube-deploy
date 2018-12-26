package responses

type DeploymentResp struct {
	ServiceName string `json:"serviceName"`
	Image string `json:"image"`
	InstanceNum int32 `json:"instancheNum"`
	Namespace string `json:"namespace"`
	Running bool `json:"running"`
}
