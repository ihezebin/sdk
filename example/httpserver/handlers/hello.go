package handlers

import (
	"context"
	"fmt"
	"github.com/whereabouts/sdk/example/httpserver/proto"
	"github.com/whereabouts/sdk/httpserver/result"
	//"mime/multipart"
)

// HelloStandardHandler
// request parameters are automatically mapped and bound to req,
// If the return value is nil, the resp structure is mapped to the response body in JSON format,
// Otherwise, respond with JSON *result.HttpError or error content
func HelloStandardHandler(ctx context.Context, req *proto.HelloStandardHandlerReq, resp *proto.HelloStandardHandlerResp) error {
	// the first way to get *gin.Context
	//c := ctx.Value(middleware.GinContextKey).(*gin.Context)
	// the second way to get *gin.Context is to make req inherit middleware.Context
	c := req.GetContext()
	// set response value
	resp.Code = result.CodeBoolOk
	resp.Message = fmt.Sprintf("hello, %s! the request uri is %s", req.Name, c.Request.RequestURI)
	fmt.Println("hello standard")
	return nil
}

// HelloFileHandler
// Upload a single file:
// it can be directly encapsulated into the req structure for parsing *multipart.FileHeader
func HelloFileHandler(ctx context.Context, req *proto.HelloFileHandlerReq, resp *proto.HelloFileHandlerResp) error {
	file := req.File
	if file == nil {
		return result.Error(result.CodeBoolFail, "fail to find any file")
	}
	resp.Code = result.CodeBoolOk
	resp.Url = fmt.Sprintf("http://file.%s/%s", req.Host, file.Filename)
	fmt.Println("hello file")
	return nil
}

// HelloMultipleFilesHandler
// Multiple file upload:
// cannot be resolved by mapping, get multiple files through gin native operation
func HelloMultipleFilesHandler(ctx context.Context, req *proto.HelloMultipleFilesHandlerReq, resp *proto.HelloMultipleFilesHandlerResp) *result.HttpError {
	multipartForm, err := req.GetContext().MultipartForm()
	if err != nil {
		return result.Err2HttpError(err, result.CodeBoolFail)
	}
	// get files
	files := multipartForm.File["files"]
	var message string
	for _, file := range files {
		message = fmt.Sprintf("%s[%s]", message, file.Filename)
	}
	resp.Code = result.CodeBoolOk
	resp.Message = fmt.Sprintf("success to upload these files : %+v", message)
	fmt.Println("hello multiple file")
	return nil
}
