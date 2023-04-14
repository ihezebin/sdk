package main

import (
	"bytes"
	"github.com/ihezebin/sdk/config"
	"github.com/ihezebin/sdk/logger"
	"os"
	"testing"
)

var (
	json = []byte(`
	{
	 "port": 8001,
	 "name": "Korbin",
	 "real_age": 18
	}
	`)
	toml = []byte(`
	port = 8002
	name = "Korbin"
	real_age = 18
	`)
	yaml = []byte(`
	port: 8003
	name: "Korbin"
	real_age: 18
	`)
)

func TestLoadWithFile(t *testing.T) {
	file, err := os.Open("./config.json")
	if err != nil {
		logger.Fatalf("open file err: %s", err.Error())
	}
	err = config.LoadWithReader(file, GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadWithReader(t *testing.T) {
	err := config.LoadWithReader(bytes.NewBuffer(json), GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadDifferentConfigType(t *testing.T) {
	config.SetConfigType("toml")
	err := config.LoadWithReader(bytes.NewReader(toml), GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadWithFilePath(t *testing.T) {
	configurator := config.New()
	err := configurator.LoadWithFilePath("./config.json", GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}
