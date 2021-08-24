package proto

import (
	"github.com/whereabouts/sdk/httpserver/result"
	"mime/multipart"
)

type HelloHandlerReq struct {
	Name string `json:"name,default=jsonName" form:"name,default=formName"`
}

type HelloHandlerResp struct {
	Code    bool   `json:"coder"`
	Message string `json:"message"`
}

type HelloFileHandlerReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Host string                `json:"host,default=whereabouts.icu" form:"host,default=whereabouts.icu"`
}

type HelloFileHandlerResp struct {
	Code bool   `json:"coder"`
	Url  string `json:"url"`
}

type HelloMultipleFilesHandlerReq struct {
}

type HelloMultipleFilesHandlerResp struct {
	result.DefaultResp
}
