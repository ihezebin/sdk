package redisc

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
	value, _ := result.Value()
	return value
}

func (result Result) Error() error {
	_, err := result.Value()
	return err
}

func (result Result) Value() (interface{}, error) {
	if result.reply == nil {
		return nil, ErrNil
	}
	return result.reply, nil
}

func (result Result) Int() (int, error) {
	return redis.Int(result.reply, result.err)
}

func (result Result) Int64() (int64, error) {
	return redis.Int64(result.reply, result.err)
}

func (result Result) Uint64() (uint64, error) {
	return redis.Uint64(result.reply, result.err)
}

func (result Result) Float64() (float64, error) {
	return redis.Float64(result.reply, result.err)
}

func (result Result) String() (string, error) {
	return redis.String(result.reply, result.err)
}

func (result Result) Bytes() ([]byte, error) {
	return redis.Bytes(result.reply, result.err)
}

func (result Result) Bool() (bool, error) {
	return redis.Bool(result.reply, result.err)
}

func (result Result) Values() ([]interface{}, error) {
	return redis.Values(result.reply, result.err)
}

func (result Result) Float64s() ([]float64, error) {
	return redis.Float64s(result.reply, result.err)
}

func (result Result) Strings() ([]string, error) {
	return redis.Strings(result.reply, result.err)
}

func (result Result) ByteSlices() ([][]byte, error) {
	return redis.ByteSlices(result.reply, result.err)
}

func (result Result) Int64s() ([]int64, error) {
	return redis.Int64s(result.reply, result.err)
}

func (result Result) Ints() ([]int, error) {
	return redis.Ints(result.reply, result.err)
}

func (result Result) Uint64s() ([]uint64, error) {
	return redis.Uint64s(result.reply, result.err)
}

func (result Result) StringMap() (map[string]string, error) {
	return redis.StringMap(result.reply, result.err)
}

func (result Result) IntMap() (map[string]int, error) {
	return redis.IntMap(result.reply, result.err)
}

func (result Result) Int64Map() (map[string]int64, error) {
	return redis.Int64Map(result.reply, result.err)
}

func (result Result) Positions() ([]*[2]float64, error) {
	return redis.Positions(result.reply, result.err)
}

func (result Result) Uint64Map() (map[string]uint64, error) {
	return redis.Uint64Map(result.reply, result.err)
}

func (result Result) SlowLogs() ([]redis.SlowLog, error) {
	return redis.SlowLogs(result.reply, result.err)
}
