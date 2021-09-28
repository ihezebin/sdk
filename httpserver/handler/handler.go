package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	result2 "github.com/whereabouts/sdk/httpserver/handler/result"
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
	returnResultErrorType = reflect.TypeOf((*result2.Err)(nil))

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

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	return CreateHandlerFuncWithOptions(method)
}

func CreateHandlerFuncWithOptions(method interface{}, options ...Option) gin.HandlerFunc {
	conf := newConfig(options...)
	return createHandlerFuncWithLogger(method, conf, logger.StandardLogger())
}

func createHandlerFuncWithLogger(method interface{}, conf config, l *logger.Logger) gin.HandlerFunc {
	mV, reqT, err := checkMethod(method, conf)
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
			c.JSON(http.StatusBadRequest, result.Err2HttpError(err, result2.CodeBoolFail))
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
				c.JSON(http.StatusOK, result.Err2HttpError(v, result2.CodeBoolFail))
			}
			return
		}
		c.PureJSON(http.StatusOK, resp)
	}
}

func checkMethod(method interface{}, conf config) (mV reflect.Value, reqT reflect.Type, err error) {
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

	if conf.withGinCtx {
		if mT.NumIn() != 3 {
			err = errors.Errorf("method(%T) must has 3 ins", method)
			return
		}
	} else {
		if mT.NumIn() != 2 {
			err = errors.Errorf("method(%T) must has 2 ins", method)
			return
		}
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
