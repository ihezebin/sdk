package proto

import (
	"github.com/whereabouts/sdk/httpserver/result"
	"mime/multipart"
)

type StandardHandlerReq struct {
	Name string `json:"name,default=jsonName" form:"name,default=formName"`
}

type StandardHandlerResp struct {
	Code    bool   `json:"coder"`
	Message string `json:"message"`
}

type StandardFileHandlerReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Host string                `json:"host,default=whereabouts.icu" form:"host,default=whereabouts.icu"`
}

type StandardFileHandlerResp struct {
	Code bool   `json:"coder"`
	Url  string `json:"url"`
}

type StandardMultipleFilesHandlerReq struct {
}

type StandardMultipleFilesHandlerResp struct {
	result.DefaultResp
}
