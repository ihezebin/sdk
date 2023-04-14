package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/utils/mapper"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	maxBodyLen = 1024
)

func LoggingRequest() Middleware {
	return LoggingRequestWithLogger(logger.StandardLogger(), false)
}

func LoggingResponse() Middleware {
	return LoggingResponseWithLogger(logger.StandardLogger(), false)
}

func LoggingSimplyRequest() Middleware {
	return LoggingRequestWithLogger(logger.StandardLogger(), true)
}

func LoggingSimplyResponse() Middleware {
	return LoggingResponseWithLogger(logger.StandardLogger(), true)
}

func LoggingRequestWithLogger(l *logger.Logger, simply bool) Middleware {
	return func(c *gin.Context) {
		fields := field.Fields{
			"method": c.Request.Method,
			"uri":    c.Request.URL.RequestURI(),
			"remote": c.Request.RemoteAddr,
			"body":   requestBody(c),
		}
		if !simply {
			fields["headers"] = convertHeaders2JSON(c.Request.Header)
		}
		l.WithFields(fields).Info("incoming http request")
	}
}

func LoggingResponseWithLogger(l *logger.Logger, simply bool) Middleware {
	return func(c *gin.Context) {
		rw := &responseWriter{Body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = rw

		c.Next()

		fields := field.Fields{
			"status": fmt.Sprintf("%v %s", c.Writer.Status(), http.StatusText(c.Writer.Status())),
			"body":   responseBody(rw),
		}
		if !simply {
			fields["headers"] = convertHeaders2JSON(c.Writer.Header())
		}
		l.WithFields(fields).Info("outgoing http response")
	}
}

func requestBody(c *gin.Context) string {
	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return ""
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return fmt.Sprintf("read request body err: %s", err.Error())
	}
	_ = c.Request.Body.Close()
	// create a new buffer and replace the original request.Body, put the read byte stream back into the request.Body
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	bodySize := len(body)
	if bodySize > maxBodyLen {
		bodySize = maxBodyLen
	}
	return string(body[:bodySize])
}

func responseBody(rw *responseWriter) string {
	body := rw.Body.Bytes()
	bodyLen := len(body)
	if bodyLen > maxBodyLen {
		bodyLen = maxBodyLen
	}
	return string(body[:bodyLen])
}

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// Write rewrite gin.ResponseWriter to store body before write
func (w responseWriter) Write(body []byte) (int, error) {
	// store body
	w.Body.Write(body)
	// write
	return w.ResponseWriter.Write(body)
}

func convertHeaders2JSON(headers http.Header) string {
	headerM := make(map[string]string, len(headers))
	for key, values := range headers {
		headerM[key] = strings.Join(values, "; ")
	}
	headerJson, err := mapper.Map2Json(headerM)
	if err != nil {
		return fmt.Sprintf(`{"convert headers err": "%s"}`, err.Error())
	}
	return headerJson
}
