package test

import "github.com/whereabouts/sdk/database/redisc"

type EmailCache struct {
	*redisc.Cache
}

func GetEmailCache() *EmailCache {
	return &EmailCache{redisc.New("email")}
}

func (ec *EmailCache) AddEmailCode(id string, code string) redisc.Result {
	return ec.Set(id, code)
}

func (ec *EmailCache) GetEmailCode(id string) redisc.Result {
	return ec.Get(id)
}
