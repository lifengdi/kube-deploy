package config


var config map[string]string = map[string]string{};

func Get(key string)string{
	return config[key]
}


func Set(key string,value string){
	config[key] = value
}