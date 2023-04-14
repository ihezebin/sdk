package v1

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

var ErrNil = errors.New("redisc returned nil")

type Result struct {
	reply interface{}
	err   error
}

func (result Result) Reply() interface{} {
	return result.reply
}

func (result Result) Value() interface{} {
	return result.Reply()
}

func (result Result) Error() error {
	return result.err
}

func (result Result) Int() int {
	if reply, err := redis.Int(result.reply, result.err); err == nil {
		return reply
	}
	return 0
}

func (result Result) Int64() int64 {
	if reply, err := redis.Int64(result.reply, result.err); err == nil {
		return reply
	}
	return 0
}

func (result Result) Uint64() uint64 {
	if reply, err := redis.Uint64(result.reply, result.err); err == nil {
		return reply
	}
	return 0
}

func (result Result) Float64() float64 {
	if reply, err := redis.Float64(result.reply, result.err); err == nil {
		return reply
	}
	return 0
}

func (result Result) String() string {
	if reply, err := redis.String(result.reply, result.err); err == nil {
		return reply
	}
	return ""
}

func (result Result) Bytes() []byte {
	if reply, err := redis.Bytes(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Bool() bool {
	if reply, err := redis.Bool(result.reply, result.err); err == nil {
		return reply
	}
	return false
}

func (result Result) Values() []interface{} {
	if reply, err := redis.Values(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Float64s() []float64 {
	if reply, err := redis.Float64s(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Strings() []string {
	if reply, err := redis.Strings(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) ByteSlices() [][]byte {
	if reply, err := redis.ByteSlices(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Int64s() []int64 {
	if reply, err := redis.Int64s(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Ints() []int {
	if reply, err := redis.Ints(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Uint64s() []uint64 {
	if reply, err := redis.Uint64s(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) StringMap() map[string]string {
	if reply, err := redis.StringMap(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) IntMap() map[string]int {
	if reply, err := redis.IntMap(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Int64Map() map[string]int64 {
	if reply, err := redis.Int64Map(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Positions() []*[2]float64 {
	if reply, err := redis.Positions(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) Uint64Map() map[string]uint64 {
	if reply, err := redis.Uint64Map(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}

func (result Result) SlowLogs() []redis.SlowLog {
	if reply, err := redis.SlowLogs(result.reply, result.err); err == nil {
		return reply
	}
	return nil
}
