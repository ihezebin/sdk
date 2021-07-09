package handlers

import (
	"context"
	"fmt"
	"github.com/whereabouts/sdk-go/example/httpserver/proto"
	"github.com/whereabouts/sdk-go/httpserver/middleware"
	"github.com/whereabouts/sdk-go/httpserver/result"
	"github.com/whereabouts/web-template/engine/http_error"
	"net/http"
	//"mime/multipart"
)

// normal request：
// request parameters are automatically mapped and bound to req,
// If the return value is nil, the resp structure is mapped to the response body in JSON format,
// Otherwise, respond with JSON *http_error.HttpError content
func SayHello(ctx context.Context, req *proto.SayHelloReq, resp *proto.SayHelloResp) error {
	value := ctx.Value(middleware.RequestKey).(*http.Request)
	fmt.Println(value.URL.String())
	fmt.Println("say hello")
	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("hello, %s! your age is %d", req.Name, req.Age)
	resp.GetContext().Header("a", "b")
	return result.Error(false, "123").WithHttpStatusCode(202)
}

// Upload a single file:
// it can be directly encapsulated into the req structure for parsing *multipart.FileHeader
func FileHello(req *proto.FileHelloReq, resp *proto.FileHelloResp) *http_error.HttpError {
	if file := req.File; file == nil {
		return http_error.Error(http.StatusBadRequest, "fail to find any file")
	}
	resp.Code = http.StatusOK
	fmt.Println(req.Name)
	resp.Message = fmt.Sprintf("success to upload file : %s", req.File.Filename)
	return nil
}

// Multiple file upload:
// cannot be resolved by mapping, req needs to be inherited handler.Context Structure,
// so that it has access *gin.context The ability to obtain multiple files through the native operation of gin
func FilesHello(req *proto.FilesHelloReq, resp *proto.FilesHelloResp) *http_error.HttpError {
	fmt.Println(req.GetContext().Request.URL)
	form, err := req.GetContext().MultipartForm()
	if err != nil {
		return http_error.Err2HttpError(err, http.StatusBadRequest)
	}
	files := form.File["files"]
	var message string
	for _, file := range files {
		message = fmt.Sprintf("%s[%s]", message, file.Filename)
	}
	resp.Code = http.StatusOK
	fmt.Println(req.Name)
	resp.Message = fmt.Sprintf("success to upload these files : %+v", message)
	return nil
}
