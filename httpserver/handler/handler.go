package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/httpserver/handler/result"
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/sdk/utils/mapper"
	"net/http"
	"reflect"
)

func init() {
}

type requestKey struct{}
type responseKey struct{}
type ginContextKey struct{}

var (
	contextType    = reflect.TypeOf((*context.Context)(nil)).Elem()
	ginContextType = reflect.TypeOf((*gin.Context)(nil))
	errorType      = reflect.TypeOf((*error)(nil)).Elem()
	errType        = reflect.TypeOf((*result.Err)(nil))
	resultType     = reflect.TypeOf((*result.Result)(nil))

	RequestKey    = requestKey{}
	ResponseKey   = responseKey{}
	GinContextKey = ginContextKey{}
)

//New
// example:
// func Hello(ctx context.Context, req *proto.HelloHandlerReq) (*proto.HelloHandlerResp, error) {
// 	 resp := proto.HelloHandlerResp{}
// 	 resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
// 	 return &resp, nil
// }
func New(method interface{}) gin.HandlerFunc {
	return NewWithOptions(method)
}

func NewWithOptions(method interface{}, options ...Option) gin.HandlerFunc {
	conf := newConfig(options...)
	return newHandlerFuncWithLogger(method, conf, logger.StandardLogger())
}

func newHandlerFuncWithLogger(method interface{}, conf config, l *logger.Logger) gin.HandlerFunc {
	mV, reqT, checkErr := checkMethod(method, conf)
	if checkErr != nil {
		l.Errorf("checkMethod err: %v", checkErr)
		panic(checkErr)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, RequestKey, c.Request)
		ctx = context.WithValue(ctx, ResponseKey, c.Writer)
		ctx = context.WithValue(ctx, GinContextKey, c)

		var res *result.Result

		// bind request param
		req := reflect.New(reqT)
		if bindErr := c.ShouldBind(req.Interface()); bindErr != nil {
			res = result.Failed(bindErr).WithStatusCode(http.StatusBadRequest)
			c.JSON(res.StatusCode(), res)
			l.Errorf("method(%T) failed to bind: %v", method, bindErr)
			return
		}
		// bind path param
		reqM, convertErr := mapper.Struct2Map(req.Interface())
		if convertErr != nil {
			res = result.Failed(convertErr).WithStatusCode(http.StatusBadRequest)
			c.JSON(res.StatusCode(), res)
			l.Errorf("method(%T) failed to reconvert path param: %v", method, convertErr)
			return
		}
		for _, param := range c.Params {
			reqM[param.Key] = param.Value
		}
		if convertErr = mapper.Map2Struct(reqM, req.Interface()); convertErr != nil {
			res = result.Failed(convertErr).WithStatusCode(http.StatusBadRequest)
			c.JSON(res.StatusCode(), res)
			l.Errorf("method(%T) failed to reconvert path param: %v", method, convertErr)
			return
		}

		// handle request
		var resultV []reflect.Value
		if conf.withCtx {
			resultV = mV.Call([]reflect.Value{reflect.ValueOf(ctx), req, reflect.ValueOf(c)})
		} else {
			resultV = mV.Call([]reflect.Value{reflect.ValueOf(ctx), req})
		}

		// do response
		if conf.withResult {
			res = resultV[0].Interface().(*result.Result)
			if res == nil {
				res = result.New()
			}
			c.PureJSON(res.StatusCode(), res)
			return
		}

		var resp, err interface{}
		if conf.withoutResponse {
			err = resultV[0].Interface()
		} else {
			resp, err = resultV[0].Interface(), resultV[1].Interface()
		}
		// response contains err
		if err != nil {
			switch e := err.(type) {
			case *result.Err:
				res = result.Failed(e).WithCode(e.Code).WithStatusCode(e.StatusCode())
			case error:
				res = result.Failed(e)
			}
			c.JSON(res.StatusCode(), res)
			return
		}

		if conf.withoutResponse {
			return
		}

		res = result.Succeed(resp)
		c.PureJSON(res.StatusCode(), res)
	}
}

func checkMethod(method interface{}, conf config) (mV reflect.Value, reqT reflect.Type, err error) {
	// check method
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

	// check in params of method
	if conf.withCtx {
		if mT.NumIn() != 3 {
			err = errors.Errorf("method(%T) must has 3 ins", method)
			return
		}
		ginCtxT := mT.In(2)
		if ginCtxT != ginContextType {
			err = errors.Errorf("the third in of method(%T) must be *gin.Context", method)
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
		err = errors.Errorf("the second in of method(%T) must be a pointer", method)
		return
	}
	reqT = reqT.Elem()
	if reqT.Kind() != reflect.Struct {
		err = errors.Errorf("the second in of method(%T) must be a struct to bind param", method)
		return
	}

	// check out params of method return
	if conf.withResult {
		if mT.NumOut() != 1 {
			err = errors.Errorf("method(%T) must has 1 out", method)
			return
		}
		if mT.Out(0) != resultType {
			err = errors.Errorf("the first out of method(%T) must be *result.Result or error", method)
			return
		}
		return
	}

	errOutIndex := 0
	if conf.withoutResponse {
		if mT.NumOut() != 1 {
			err = errors.Errorf("method(%T) must has 1 out", method)
			return
		}
	} else {
		if mT.NumOut() != 2 {
			err = errors.Errorf("method(%T) must has 2 out", method)
			return
		}
		errOutIndex = 1
	}

	errT := mT.Out(errOutIndex)
	if errT != errorType && errT != errType {
		err = errors.Errorf("the first out of method(%T) must be *result.Err or error", method)
		return
	}

	return
}
