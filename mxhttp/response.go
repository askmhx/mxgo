package mxhttp

import (
	"net/http"
)

type Response struct{
	*http.Response
	responseWriter http.ResponseWriter
}

func NewResponse(rw http.ResponseWriter) Response{
	response := &Response{}
	response.responseWriter = rw
	return response
}

func (resp *Response) writeFile(fileName string){
	resp.Write(resp.responseWriter)
}
