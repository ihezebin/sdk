package configure

import (
	"flag"
)

type CmdParam struct {
	Name  string
	Value interface{}
}

func (cp CmdParam) String() string {
	value, ok := cp.Value.(*string)
	if !ok {
		return ""
	}
	return *value
}

func (cp CmdParam) Bool() bool {
	value, ok := cp.Value.(*bool)
	if !ok {
		return false
	}
	return *value
}

func (cp CmdParam) Int() int {
	value, ok := cp.Value.(*int)
	if !ok {
		return 0
	}
	return *value
}

func GetCmdParam(name string, defaultValue interface{}) CmdParam {
	cmdParam := CmdParam{Name: name}
	switch defaultValue.(type) {
	case int:
		cmdParam.Value = flag.Int(name, defaultValue.(int), "go run main.go -[name] [value]")
	case string:
		cmdParam.Value = flag.String(name, defaultValue.(string), "go run main.go -[name] [value]")
	case bool:
		cmdParam.Value = flag.Bool(name, defaultValue.(bool), "go run main.go -[name] [value]")
	default:
		cmdParam.Value = flag.String(name, defaultValue.(string), "go run main.go -[name] [value]")
	}
	flag.Parse()
	return cmdParam
}
