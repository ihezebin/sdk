package test

import "github.com/whereabouts/sdk-go/model/redis"

type EmailCache struct {
	*redis.Cache
}

func GetEmailCache() *EmailCache {
	return &EmailCache{redis.New("email")}
}

func (ec *EmailCache) AddEmailCode(id string, code string) redis.Result {
	return ec.Set(id, code)
}

func (ec *EmailCache) GetEmailCode(id string) redis.Result {
	return ec.Get(id)
}
