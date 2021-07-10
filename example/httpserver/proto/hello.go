package proto

import (
	"github.com/whereabouts/sdk-go/httpserver/middleware"
	"github.com/whereabouts/sdk-go/httpserver/result"
	"mime/multipart"
)

type HelloStandardHandlerReq struct {
	middleware.Context
	Name string `json:"name,default=whereabouts.icu" form:"name,default=whereabouts.icu"`
}

type HelloStandardHandlerResp struct {
	Code    bool   `json:"code"`
	Message string `json:"message"`
}

type HelloFileHandlerReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Host string                `json:"host,default=whereabouts.icu" form:"host,default=whereabouts.icu"`
}

type HelloFileHandlerResp struct {
	Code bool   `json:"code"`
	Url  string `json:"url"`
}

type HelloMultipleFilesHandlerReq struct {
	middleware.Context
}

type HelloMultipleFilesHandlerResp struct {
	result.DefaultResp
}
