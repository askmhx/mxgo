package httplib

import (
	"net/http"
)

type Request struct{
	*http.Request
	UploadFiles map[string]http.File
	Session map[string]interface {}
	Cookie map[string]string
}

func NewRequest(req *http.Request) *Request{
	request := &Request{}
	request.Request = req
	return request
}
