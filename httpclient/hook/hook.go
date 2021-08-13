package hook

import "github.com/go-resty/resty/v2"

type RequestHook func(*resty.Client, *resty.Request) error
type ResponseHook func(*resty.Client, *resty.Response) error
