package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	http_error2 "github.com/whereabouts/sdk-go/httpserver/old/http_error"
	"github.com/whereabouts/sdk-go/logger"
	"net/http"
	"reflect"
)

func init() {
	//structs.DefaultTagName = "json"
}

var (
	errTypeMustPtr             = errors.New("param type must be pointer")
	errTypeMustStruct          = errors.New("param type must be struct")
	errMethodMustHasThreeParam = errors.New("method must has three func param")
	errMethodMustHasTwoParam   = errors.New("method must has two func param")
	errTypeMustFunc            = errors.New("method type must be func")
	errMethodMustValid         = errors.New("method must be valid")
	errReturnMustError         = errors.New("method must return value which type is *chassis/http_error.HttpError")
	errReturnMustOneValue      = errors.New("method must return one value")

	returnErrorType = reflect.TypeOf((*http_error2.HttpError)(nil))
)

type Context struct {
	context *gin.Context
}

func (c Context) GetContext() *gin.Context {
	return c.context
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	mV, reqT, respT, err := checkMethod(method)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		req := reflect.New(reqT)
		if err := c.ShouldBind(req.Interface()); err != nil {
			logger.Errorf("req param err: %s", err.Error())
			c.JSON(http.StatusOK, http_error2.Error(http_error2.CodeBoolFail, fmt.Sprintf("req param err: %s", err.Error())))
			return
		}
		setContext(req, c)
		resp := reflect.New(respT)
		setContext(resp, c)
		results := mV.Call([]reflect.Value{req, resp})
		respErr, _ := results[0].Interface().(*http_error2.HttpError)
		// response contains http_error
		if respErr != nil {
			if respErr.HttpStatusCode != 0 {
				c.JSON(respErr.HttpStatusCode, respErr)
			} else {
				c.JSON(http.StatusOK, respErr)
			}
			return
		}
		c.JSON(http.StatusOK, resp.Interface())
	}
}

func setContext(v reflect.Value, c *gin.Context) {
	contextV := reflect.ValueOf(Context{c})
	contextChild := v.Elem().FieldByName(contextV.Type().Name())
	if ok := contextChild.IsValid(); ok {
		contextChild.Set(contextV)
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
	if mT.NumIn() != 2 {
		err = errMethodMustHasTwoParam
		return
	}
	reqT = mT.In(0)
	if reqT.Kind() != reflect.Ptr {
		err = errTypeMustPtr
		return
	}
	if reqT.Elem().Kind() != reflect.Struct {
		err = errTypeMustStruct
		return
	}
	reqT = reqT.Elem()
	respT = mT.In(1)
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
	if retT != returnErrorType && reqT != reflect.TypeOf(nil) {
		err = errReturnMustError
		return
	}
	return mV, reqT, respT, err
}
