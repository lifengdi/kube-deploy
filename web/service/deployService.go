package service

type Demo struct {
	SourceId   string
	UserId     string
}

func (d *Demo) GetDemo() (string,error) {
	println("i am deploy service")
	return "hello",nil
}