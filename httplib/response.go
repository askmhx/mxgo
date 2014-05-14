package httplib

import (
	"net/http"
)

type Response struct{
	*http.Response
	responseWriter http.ResponseWriter
}

func NewResponse(rw http.ResponseWriter) *Response{
	response := &Response{}
	response.responseWriter = rw
	return response
}

func (resp *Response) WriteFile(filePath string){
	resp.Write(resp.responseWriter)
}


func (resp *Response) WriteText(text string){
	resp.Write(resp.responseWriter)
}
