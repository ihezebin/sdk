package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/httpserver/result"
	"log"
	"net/http"
	"reflect"
)

func init() {
}

type requestKey struct{}
type responseKey struct{}
type ginContextKey struct{}

var (
	errTypeMustContext         = errors.New("param type must be context.Context")
	errTypeMustPtr             = errors.New("param type must be pointer")
	errTypeMustStruct          = errors.New("param type must be struct")
	errMethodMustHasThreeParam = errors.New("method must has three func param")
	errMethodMustHasTwoParam   = errors.New("method must has two func param")
	errTypeMustFunc            = errors.New("method type must be func")
	errMethodMustValid         = errors.New("method must be valid")
	errReturnMustError         = errors.New("method must return value which type is *chassis/http_error.HttpError")
	errReturnMustOneValue      = errors.New("method must return one value")

	returnErrorType       = reflect.TypeOf((error)(nil))
	returnResultErrorType = reflect.TypeOf((*result.HttpError)(nil))

	ctxType = reflect.TypeOf(context.Context(nil))

	RequestKey    = requestKey{}
	ResponseKey   = responseKey{}
	GinContextKey = ginContextKey{}
)

type Context struct {
	*gin.Context
}

func (c Context) GetContext() *gin.Context {
	return c.Context
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	mV, reqT, respT, err := checkMethod(method)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, RequestKey, c.Request)
		ctx = context.WithValue(ctx, ResponseKey, c.Writer)
		ctx = context.WithValue(ctx, GinContextKey, c)

		req := reflect.New(reqT)
		if err := c.ShouldBind(req.Interface()); err != nil {
			log.Printf("bind param failed: %s\n", err.Error())
			c.JSON(http.StatusBadRequest, result.Error(result.CodeFail, fmt.Sprintf("bind param failed: %s", err.Error())))
			return
		}
		setContext(req, c)

		resp := reflect.New(respT)
		setContext(resp, c)

		results := mV.Call([]reflect.Value{reflect.ValueOf(ctx), req, resp})
		errValue := results[0].Interface()
		// response contains http_error
		if errValue != nil {
			switch v := errValue.(type) {
			case *result.HttpError:
				if v.HttpStatusCode != 0 {
					c.JSON(v.HttpStatusCode, v)
				}
				return
			case error:
				c.JSON(http.StatusOK, result.Err2HttpError(v, result.CodeFail))
				return
			}
		}
		c.PureJSON(http.StatusOK, resp.Interface())
	}
}

func setContext(v reflect.Value, c *gin.Context) {
	contextV := reflect.ValueOf(Context{c})
	vCtxChild := v.Elem().FieldByName(contextV.Type().Name())
	if ok := vCtxChild.IsValid(); ok {
		vCtxChild.Set(contextV)
	}
}

func checkMethod(method interface{}) (mV reflect.Value, reqT, respT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = errMethodMustValid
		return
	}

	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = errTypeMustFunc
		return
	}

	if mT.NumIn() != 3 {
		err = errMethodMustHasThreeParam
		return
	}

	ctxT := mT.In(0)
	if !ctxT.Implements(ctxType) && ctxT != ctxType {
		err = errTypeMustContext
		return
	}

	reqT = mT.In(1)
	if reqT.Kind() != reflect.Ptr {
		err = errTypeMustPtr
		return
	}
	if reqT.Elem().Kind() != reflect.Struct {
		err = errTypeMustStruct
		return
	}
	reqT = reqT.Elem()

	respT = mT.In(2)
	if respT.Kind() != reflect.Ptr {
		err = errTypeMustPtr
		return
	}
	if respT.Elem().Kind() != reflect.Struct {
		err = errTypeMustStruct
		return
	}
	respT = respT.Elem()

	if mT.NumOut() != 1 {
		err = errReturnMustOneValue
		return
	}
	retT := mT.Out(0)
	if retT != returnErrorType && retT != returnResultErrorType {
		err = errReturnMustError
		return
	}
	return mV, reqT, respT, err
}
