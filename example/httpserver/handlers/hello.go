package handlers

import (
	"fmt"
	"github.com/whereabouts/web-template/engine/http_error"
	"github.com/whereabouts/web-template/proto"
	"net/http"
	//"mime/multipart"
)

// normal request：
// request parameters are automatically mapped and bound to req,
// If the return value is nil, the resp structure is mapped to the response body in JSON format,
// Otherwise, respond with JSON *http_error.HttpError content
func SayHello(req *proto.SayHelloReq, resp *proto.SayHelloResp) *http_error.HttpError {
	fmt.Println("say hello")
	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("hello, %s! your age is %d", req.Name, req.Age)
	resp.GetContext().Header("a", "b")
	return http_error.Error(1, "2")
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
