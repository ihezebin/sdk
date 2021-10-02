package result

import "encoding/json"

type Json map[string]interface{}

func (j Json) Set(key string, val interface{}) Json {
	j[key] = val
	return j
}

func (j Json) Get(key string) interface{} {
	return j[key]
}

func (j Json) String() string {
	data, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(data)
}
