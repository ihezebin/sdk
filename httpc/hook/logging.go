package hook

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/utils/mapper"
	"io/ioutil"
	"net/http"
	"strings"
)

const maxBodyLen = 1024

func LoggingRequest() RequestHook {
	return LoggingRequestWithLogger(logger.StandardLogger(), false)
}

func LoggingResponse() ResponseHook {
	return LoggingResponseWithLogger(logger.StandardLogger(), false)
}

func LoggingSimplyRequest() RequestHook {
	return LoggingRequestWithLogger(logger.StandardLogger(), true)
}

func LoggingSimplyResponse() ResponseHook {
	return LoggingResponseWithLogger(logger.StandardLogger(), true)
}

func LoggingRequestWithLogger(l *logger.Logger, simply bool) RequestHook {
	return func(c *resty.Client, r *resty.Request) error {
		fields := field.Fields{
			"method": r.RawRequest.Method,
			"url":    r.RawRequest.URL.String(),
			"body":   requestBody(r),
		}
		if !simply {
			fields["headers"] = convertHeaders2JSON(r.RawRequest.Header)
		}
		l.WithFields(fields).Info("outgoing http request")
		return nil
	}
}

func LoggingResponseWithLogger(l *logger.Logger, simply bool) ResponseHook {
	return func(c *resty.Client, r *resty.Response) error {
		fields := field.Fields{
			"status": r.Status(),
			"body":   responseBody(r),
		}
		if !simply {
			fields["headers"] = convertHeaders2JSON(r.Header())
		}
		l.WithFields(fields).Info("incoming http response")
		return nil
	}
}

func requestBody(r *resty.Request) string {
	if r.RawRequest.Body == nil || r.RawRequest.Body == http.NoBody {
		return ""
	}
	body, _ := ioutil.ReadAll(r.RawRequest.Body)
	_ = r.RawRequest.Body.Close()
	r.RawRequest.Body = ioutil.NopCloser(bytes.NewReader(body))
	bodySize := len(body)
	if bodySize > maxBodyLen {
		bodySize = maxBodyLen
	}
	return string(body[:bodySize])
}

func responseBody(r *resty.Response) string {
	cnt := len(r.String())
	if cnt > maxBodyLen {
		cnt = maxBodyLen
	}
	return r.String()[:cnt]
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
