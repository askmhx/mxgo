package httplib

import (
	"net/http"
)

type Response struct{
	*http.Response
	ResponseWriter http.ResponseWriter
}

func NewResponse(rw http.ResponseWriter) *Response{
	response := &Response{}
	response.ResponseWriter = rw
	return response
}

func (resp *Response) WriteFile(filePath string){
	resp.Write(resp.ResponseWriter)
}


func (resp *Response) WriteText(text string){
	resp.Write(resp.ResponseWriter)
}
