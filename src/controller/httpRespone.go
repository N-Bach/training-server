package controller

import (
	"encoding/json"
	"net/http"
)

type IResponse interface {
	WithJson(data interface{}) IResponse
	WithString(message string) IResponse
	WithContent(content []byte) IResponse
	WithError(err error) IResponse
	WithStatusCode(status int) IResponse
	Excute(w http.ResponseWriter)
}

type Response struct {
	StatusCode int
	Content    []byte
	Error      error
}

var BaseResponse = func(content interface{}, status int) IResponse {
	response := &Response{}
	response.WithStatusCode(status)
	response.WithJson(content)
	return response
}

var ResponseOk = func(content interface{}) IResponse {
	response := &Response{}
	response.WithStatusCode(http.StatusOK)
	response.WithJson(content)
	return response
}

var ResponseInteralError = func(content interface{}, err error) IResponse {
	response := &Response{}
	response.WithStatusCode(http.StatusInternalServerError)
	response.WithError(err)
	response.WithJson(content)
	return response
}

var ResponseBadRequest = func(content interface{}, err error) IResponse {
	response := &Response{}
	response.WithStatusCode(http.StatusBadRequest)
	response.WithError(err)
	response.WithJson(content)
	return response
}

func (r *Response) WithJson(data interface{}) IResponse {
	js, err := json.Marshal(data)
	if err != nil {
		r.Content = nil
		return r
	}
	r.Content = js
	return r
}
func (r *Response) WithString(message string) IResponse {
	r.Content = []byte(message)
	return r
}
func (r *Response) WithContent(content []byte) IResponse {
	r.Content = content
	return r
}
func (r *Response) WithError(err error) IResponse {
	r.Error = err
	return r
}
func (r *Response) WithStatusCode(status int) IResponse {
	r.StatusCode = status
	return r
}

func (r *Response) Excute(w http.ResponseWriter) {
	if r.StatusCode == 0 {
		r.StatusCode = http.StatusInternalServerError
	}
	if r.Content == nil {
		r.Content = []byte("")
	}
	if r.Error != nil {
		//fmt.Println("Error in writing response")
	}

	w.WriteHeader(r.StatusCode)
	w.Write(r.Content)
}
