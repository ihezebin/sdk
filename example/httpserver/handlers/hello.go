package handlers

import (
	"context"
	"fmt"
	"github.com/whereabouts/sdk/example/httpserver/proto"
	result2 "github.com/whereabouts/sdk/httpserver/handler/result"
	"github.com/whereabouts/sdk/httpserver/middleware"
	"github.com/whereabouts/sdk/httpserver/result"
	//"mime/multipart"
)

func HelloHandler(ctx context.Context, req *proto.HelloHandlerReq, resp *proto.HelloHandlerResp) error {
	// test panic happened
	//a := 0
	//_ = 1 / a
	// the way to get *gin.Context
	c, err := middleware.ExtractGinContextWithCtx(ctx)
	if err != nil {
		return err
	}
	// set response value
	resp.Code = result2.CodeBoolOk
	resp.Message = fmt.Sprintf("hello, %s! the request uri is %s", req.Name, c.Request.RequestURI)
	fmt.Println("hello standard")
	return nil
}

func HelloFileHandler(ctx context.Context, req *proto.HelloFileHandlerReq, resp *proto.HelloFileHandlerResp) error {
	file := req.File
	if file == nil {
		return result2.Error(result2.CodeBoolFail, "fail to find any file")
	}
	resp.Code = result2.CodeBoolOk
	resp.Url = fmt.Sprintf("http://file.%s/%s", req.Host, file.Filename)
	fmt.Println("hello file")
	return nil
}

func HelloMultipleFilesHandler(ctx context.Context, req *proto.HelloMultipleFilesHandlerReq, resp *proto.HelloMultipleFilesHandlerResp) *result.HttpError {
	c, err := middleware.ExtractGinContextWithCtx(ctx)
	if err != nil {
		return result.Err2HttpError(err, false)
	}
	multipartForm, err := c.MultipartForm()
	if err != nil {
		return result.Err2HttpError(err, result2.CodeBoolFail)
	}
	// get files
	files := multipartForm.File["files"]
	var message string
	for _, file := range files {
		message = fmt.Sprintf("%s[%s]", message, file.Filename)
	}
	resp.Code = result2.CodeBoolOk
	resp.Message = fmt.Sprintf("success to upload these files : %+v", message)
	fmt.Println("hello multiple file")
	return nil
}
