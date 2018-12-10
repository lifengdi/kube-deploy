package reqBody


type CreateRequest struct {
	ServiceName string
	InstanceNum int
	Image string
	Port int32
	TargetPort int
}