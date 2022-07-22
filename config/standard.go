package config

import "io"

var gConfigurator = New()

func LoadWithFilePath(path string, config interface{}) error {
	return gConfigurator.LoadWithFilePath(path, config)
}

func LoadWithReader(reader io.Reader, config interface{}) error {
	return gConfigurator.LoadWithReader(reader, config)
}

func SetTagName(tagName string) *Configurator {
	return gConfigurator.SetTagName(tagName)
}

func SetConfigType(configType string) *Configurator {
	return gConfigurator.SetConfigType(configType)
}
