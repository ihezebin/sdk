package main

type config struct {
	Port int    `json:"port"`
	Env  string `json:"env"`
}

var gConfig config
