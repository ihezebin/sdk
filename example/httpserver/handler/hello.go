package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ihezebin/sdk/example/httpserver/proto"
	"github.com/ihezebin/sdk/httpserver/handler"
	"github.com/ihezebin/sdk/httpserver/handler/result"
	"net/http"
)

type HelloHandler struct {
}

func (h *HelloHandler) Init(router gin.IRouter) {
	helloGroup := router.Group("hello")
	helloGroup.GET("/:name", handler.New(h.Hello))
	helloGroup.GET("/", handler.New(h.Hello))
	helloGroup.GET("/had_err", handler.New(h.HelloHadError))
	helloGroup.GET("/use_err", handler.New(h.HelloUseErr))
	helloGroup.GET("/with_gin_ctx", handler.NewWithOptions(h.HelloWithGinCtx, handler.WithContext()))
	helloGroup.GET("/with_gin_ctx_and_no_response", handler.NewWithOptions(h.HelloWithGinCtxAndNoReturn, handler.WithContext(), handler.DisableReturn()))
	helloGroup.GET("/with_result", handler.NewWithOptions(h.HelloWithResult, handler.WithResult()))
	helloGroup.POST("/file", handler.New(h.HelloFile))
	helloGroup.POST("/multiple_files", handler.NewWithOptions(h.HelloMultipleFiles, handler.WithContext()))
	helloGroup.GET("/with_message", handler.NewWithOptions(h.HelloWithMessage, handler.WithResult()))
}

func (h *HelloHandler) Hello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
	return &resp, nil
}

func (h *HelloHandler) HelloHadError(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	err := errors.New("one unknown err happened")
	if err != nil {
		return nil, err
	}
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
	return &resp, nil
}

func (h *HelloHandler) HelloUseErr(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, *result.Err) {
	err := errors.New("one unknown err happened")
	if err != nil {
		return nil, result.Error2Err(err)
	}
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
	return &resp, nil
}

func (h *HelloHandler) HelloWithGinCtx(c *gin.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s! the uri is %s", req.Name, c.Request.RequestURI)
	return &resp, nil
}

func (h *HelloHandler) HelloWithGinCtxAndNoReturn(c *gin.Context, req *proto.HelloReq) {
	var err error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s! the uri is %s", req.Name, c.Request.RequestURI)
	c.JSON(http.StatusOK, resp)
}

func (h *HelloHandler) HelloWithResult(ctx context.Context, req *proto.HelloReq) *result.Result {
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
	return result.Succeed(resp)
}

func (h *HelloHandler) HelloFile(ctx context.Context, req *proto.HelloFileReq) (*proto.HelloFileResp, error) {
	file := req.File
	if file == nil {
		return nil, result.Error("fail to find any file")
	}
	resp := proto.HelloFileResp{}
	resp.Url = fmt.Sprintf("http://file.%s/%s", req.Host, file.Filename)
	return &resp, nil
}

func (h *HelloHandler) HelloMultipleFiles(c *gin.Context, req *proto.HelloMultipleFilesReq) (*proto.HelloMultipleFilesResp, error) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		return nil, result.Error2Err(err)
	}
	// get files
	files := multipartForm.File["files"]
	var message string
	for _, file := range files {
		message = fmt.Sprintf("%s[%s]", message, file.Filename)
	}

	resp := proto.HelloMultipleFilesResp{}
	resp.Result = fmt.Sprintf("success to upload these files : %+v", message)
	return &resp, nil
}

func (h *HelloHandler) HelloWithMessage(ctx context.Context, req *proto.HelloReq) *result.Result {
	resp := proto.HelloResp{}
	resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
	return result.Succeed(resp).WithMessage("hello world")
}
