package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/httpserver/result"
	"github.com/whereabouts/sdk/logger"
	"net/http"
	"reflect"
)

func init() {
}

type requestKey struct{}
type responseKey struct{}
type ginContextKey struct{}

var (
	contextType           = reflect.TypeOf((*context.Context)(nil)).Elem()
	returnErrorType       = reflect.TypeOf((*error)(nil)).Elem()
	returnResultErrorType = reflect.TypeOf((*result.HttpError)(nil))

	RequestKey    = requestKey{}
	ResponseKey   = responseKey{}
	GinContextKey = ginContextKey{}
)

type Context struct {
	ctx *gin.Context
}

func (c Context) GetContext() *gin.Context {
	return c.ctx
}

func ExtractRequestWithCtx(ctx context.Context) (*http.Request, error) {
	request, ok := ctx.Value(RequestKey).(*http.Request)
	if !ok {
		return nil, errors.New("context does not contain *http.Request")
	}
	return request, nil
}

func ExtractResponseWithCtx(ctx context.Context) (*http.Response, error) {
	response, ok := ctx.Value(ResponseKey).(*http.Response)
	if !ok {
		return nil, errors.New("context does not contain *http.Response")
	}
	return response, nil
}

func ExtractGinContextWithCtx(ctx context.Context) (*gin.Context, error) {
	ginCtx, ok := ctx.Value(GinContextKey).(*gin.Context)
	if !ok {
		return nil, errors.New("context does not contain *gin.Context")
	}
	return ginCtx, nil
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	return createHandlerFuncWithLogger(method, logger.StandardLogger())
}

func createHandlerFuncWithLogger(method interface{}, l *logger.Logger) gin.HandlerFunc {
	mV, reqT, respT, err := checkHandleMethod(method)
	if err != nil {
		l.Errorf("CreateHandlerFunc panic: %v", err)
		panic(err)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, RequestKey, c.Request)
		ctx = context.WithValue(ctx, ResponseKey, c.Writer)
		ctx = context.WithValue(ctx, GinContextKey, c)

		req := reflect.New(reqT)
		if err = c.ShouldBind(req.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, result.Err2HttpError(err, result.CodeBoolFail))
			l.Errorf("method(%T) failed to bind: %v", method, err)
			return
		}
		//inheritGinContext(req, c)

		resp := reflect.New(respT)
		//inheritGinContext(resp, c)

		results := mV.Call([]reflect.Value{reflect.ValueOf(ctx), req, resp})
		errValue := results[0].Interface()
		// response contains error
		if errValue != nil {
			switch v := errValue.(type) {
			case *result.HttpError:
				if v.HttpStatusCode != 0 {
					c.JSON(v.HttpStatusCode, v)
					return
				}
				c.JSON(http.StatusOK, v)
			case error:
				c.JSON(http.StatusOK, result.Err2HttpError(v, result.CodeBoolFail))
			}
			return
		}
		c.PureJSON(http.StatusOK, resp.Elem().Interface())
	}
}

func CreateRouteHandlerFunc(method interface{}) gin.HandlerFunc {
	return createRouteHandlerFuncWithLogger(method, logger.StandardLogger())
}

func createRouteHandlerFuncWithLogger(method interface{}, l *logger.Logger) gin.HandlerFunc {
	mV, reqT, err := checkRouteMethod(method)
	if err != nil {
		l.Errorf("CreateRouteHandlerFunc panic: %v", err)
		panic(err)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, RequestKey, c.Request)
		ctx = context.WithValue(ctx, ResponseKey, c.Writer)
		ctx = context.WithValue(ctx, GinContextKey, c)

		req := reflect.New(reqT)
		if err = c.ShouldBind(req.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, result.Err2HttpError(err, result.CodeBoolFail))
			l.Errorf("method(%T) failed to bind: %v", method, err)
		}

		results := mV.Call([]reflect.Value{reflect.ValueOf(ctx), req})
		resp, errValue := results[0].Interface(), results[1].Interface()
		// response contains http_error
		if errValue != nil {
			switch v := errValue.(type) {
			case *result.HttpError:
				if v.HttpStatusCode != 0 {
					c.JSON(v.HttpStatusCode, v)
					return
				}
				c.JSON(http.StatusOK, v)
			case error:
				c.JSON(http.StatusOK, result.Err2HttpError(v, result.CodeBoolFail))
			}
			return
		}
		c.PureJSON(http.StatusOK, resp)
	}
}

// Deprecated, use ExtractGinContextWithCtx
func inheritGinContext(v reflect.Value, c *gin.Context) {
	contextV := reflect.ValueOf(Context{c})
	vCtxChild := v.Elem().FieldByName(contextV.Type().Name())
	if ok := vCtxChild.CanSet(); ok {
		vCtxChild.Set(contextV)
	}
}

func checkHandleMethod(method interface{}) (mV reflect.Value, reqT, respT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = errors.Errorf("handle method(%T) not found", method)
		return
	}

	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = errors.Errorf("handle method(%T) type must be func", method)
		return
	}

	if mT.NumIn() != 3 {
		err = errors.Errorf("method(%T) must has 3 ins", method)
		return
	}

	ctxT := mT.In(0)
	if ctxT != contextType {
		err = errors.Errorf("the first in of method(%T) must be context.Context", method)
		return
	}

	reqT = mT.In(1)
	if reqT.Kind() != reflect.Ptr {
		err = errors.Errorf("the second in of method(%T) must be pointer", method)
		return
	}
	if reqT.Elem().Kind() != reflect.Struct {
		err = errors.Errorf("the second in of method(%T) must be struct", method)
		return
	}
	reqT = reqT.Elem()

	respT = mT.In(2)
	if respT.Kind() != reflect.Ptr {
		err = errors.Errorf("the third in of method(%T) must be pointer", method)
		return
	}
	if respT.Elem().Kind() != reflect.Struct {
		err = errors.Errorf("the third in of method(%T) must be struct", method)
		return
	}
	respT = respT.Elem()

	if mT.NumOut() != 1 {
		err = errors.Errorf("method(%T) must has 1 out", method)
		return
	}
	retT := mT.Out(0)
	if retT != returnErrorType && retT != returnResultErrorType {
		err = errors.Errorf("the out of method(%T) must be result.HttpError or error", method)
		return
	}
	return mV, reqT, respT, err
}

func checkRouteMethod(method interface{}) (mV reflect.Value, reqT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = errors.Errorf("handle method(%T) not found", method)
		return
	}

	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = errors.Errorf("handle method(%T) type must be func", method)
		return
	}

	if mT.NumIn() != 2 {
		err = errors.Errorf("method(%T) must has 2 ins", method)
		return
	}

	ctxT := mT.In(0)
	if ctxT != contextType {
		err = errors.Errorf("the first in of method(%T) must be context.Context", method)
		return
	}

	reqT = mT.In(1)
	if reqT.Kind() != reflect.Ptr {
		err = errors.Errorf("the second in of method(%T) must be pointer", method)
		return
	}

	if reqT.Elem().Kind() != reflect.Struct {
		err = errors.Errorf("the second in of method(%T) must be struct", method)
		return
	}
	reqT = reqT.Elem()

	if mT.NumOut() != 2 {
		err = errors.Errorf("method(%T) must has 2 out", method)
		return
	}

	respT := mT.Out(0)
	if respT.Kind() != reflect.Ptr {
		err = errors.Errorf("the first out of method(%T) must be pointer", method)
		return
	}

	if respT.Elem().Kind() != reflect.Struct {
		err = errors.Errorf("the first out of method(%T) must be struct", method)
		return
	}
	respT = respT.Elem()

	errT := mT.Out(1)
	if errT != returnErrorType && errT != returnResultErrorType {
		err = errors.Errorf("the second out of method(%T) must be result.HttpError or error", method)
		return
	}

	return mV, reqT, err
}
