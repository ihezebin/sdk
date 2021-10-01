package proto

import (
	"mime/multipart"
)

type HelloReq struct {
	Name string `json:"name,default=jsonName" form:"name,default=formName"`
}

type HelloResp struct {
	Welcome string `json:"welcome"`
}

type HelloFileReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Host string                `json:"host,default=whereabouts.icu" form:"host,default=whereabouts.icu"`
}

type HelloFileResp struct {
	Url string `json:"url"`
}

type HelloMultipleFilesReq struct {
}

type HelloMultipleFilesResp struct {
	Result string `json:"result"`
}
