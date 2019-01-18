package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	// HTTPResult is the result of an http request
	HTTPResult struct {
		Res *http.Response
		Err error
	}

	// HTTPError is an error for http requests
	HTTPError struct {
		Request    *http.Request
		StatusCode int
		Body       string
	}

	// ResponseWriter extends the standard http.ResponseWriter
	ResponseWriter interface {
		http.ResponseWriter
		StatusCode() int
		StatusClass() string
	}

	responseWriter struct {
		http.ResponseWriter
		statusCode  int
		statusClass string
	}
)

// NewHTTPError creates a new instance of HTTPError
func NewHTTPError(res *http.Response) *HTTPError {
	err := &HTTPError{
		Request:    res.Request,
		StatusCode: res.StatusCode,
	}

	if res.Body != nil {
		if data, e := ioutil.ReadAll(res.Body); e == nil {
			err.Body = string(data)
		}
	}

	return err
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s %s %d: %s", e.Request.Method, e.Request.URL.Path, e.StatusCode, e.Body)
}

// NewResponseWriter creates a new response writer
func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	return &responseWriter{
		ResponseWriter: rw,
	}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)

	if rw.statusCode == 0 {
		rw.statusCode = statusCode
		rw.statusClass = fmt.Sprintf("%dxx", statusCode/100)
	}
}

func (rw *responseWriter) StatusCode() int {
	return rw.statusCode
}

func (rw *responseWriter) StatusClass() string {
	return rw.statusClass
}
