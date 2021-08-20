package proto

import (
	"github.com/whereabouts/sdk/httpserver/result"
	"mime/multipart"
)

type HelloStandardHandlerReq struct {
	Name string `json:"name,default=jsonName" form:"name,default=formName"`
}

type HelloStandardHandlerResp struct {
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
