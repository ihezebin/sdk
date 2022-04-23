package redisc

import (
	"github.com/go-redis/redis/v8"
	"strings"
)

type Model interface {
	Name() string
	Key() string
}

type warpCmd = redis.Cmdable

const defaultModelKeySep = "."

type model struct {
	warpCmd
	name   string
	keySep string
}

func NewModel(c *Client, name string) *model {
	return &model{warpCmd: c, name: name, keySep: defaultModelKeySep}
}

func (m *model) Name() string {
	return m.name
}

func (m *model) Key(ext ...string) string {
	k := []string{m.Name()}
	k = append(k, ext...)
	return strings.Join(k, m.keySep)
}

func (m *model) Kernel() redis.Cmdable {
	return m.warpCmd
}

func (m *model) SetKeySep(sep string) {
	m.keySep = sep
}
