package config

var gConfigurator = New()

func Load(file interface{}, config interface{}) error {
	return gConfigurator.Load(file, config)
}

func LoadWithDefault(config interface{}) error {
	return gConfigurator.LoadWithDefault(config)
}

func LoadWithCmd(key, value string, config interface{}) error {
	return gConfigurator.LoadWithCmd(key, value, config)
}

func LoadWithEnv(key string, config interface{}) error {
	return gConfigurator.LoadWithEnv(key, config)
}

//func LoadWithURI(uri string, config interface{}) error {
//	return gConfigurator.LoadWithURI(uri, config)
//}

func SetDefaultConfigPath(path string) {
	defaultConfigPath = path
}
