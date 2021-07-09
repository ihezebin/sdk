package proto

import (
	"github.com/whereabouts/sdk-go/httpserver/middleware"
	"mime/multipart"
)

type SayHelloReq struct {
	Name string `json:"name,default=korbin" form:"name,default=korbin"`
	Age  int    `json:"age" form:"age"`
}

type SayHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
	middleware.Context
}

type FileHelloReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Name string                `json:"name,default=hezebin" form:"name,default=hezebin"`
}

type FileHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}

type FilesHelloReq struct {
	middleware.Context
	Name string `json:"name" form:"name"`
}

type FilesHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}
