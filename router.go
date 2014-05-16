package mxgo

import (
	"github.com/menghx/mxgo/httplib"
	"strings"
)

type router struct {
	UriPattern string
	HttpMethod string
	CtrlName string
	FuncName string
}

type RouterManager struct {
	 routes []router
}

func NewRouterManager() *RouterManager{
	rm := &RouterManager{}
	rm.routes = make([]router,0)
	return rm
}

func (rm *RouterManager)Router(uriPattern string,httpMethod string,action string){
	route := router{}
	route.UriPattern = uriPattern
	route.HttpMethod = httpMethod
	actionPattern := strings.Split(action,".")
	route.CtrlName = actionPattern[0]
	route.FuncName = actionPattern[1]
	rm.routes = append(rm.routes,route)
}

func (rm *RouterManager)FindAction(request *httplib.Request,response *httplib.Response) *Action{
	inMethod := request.Method
	inUri := request.RequestURI
	//need cache here url->action
	for i := range rm.routes{
		r := rm.routes[i]
		if rm.matchPattern(r.UriPattern,inUri) {
			if r.HttpMethod == inMethod || r.HttpMethod== "*" {
				action := &Action{}
				action.CtrlName = r.CtrlName
				action.FuncName = r.FuncName
				return action
			}else{
				return rm.errorAction(405,inUri+":use http:"+inMethod+" not allowed")
			}
		}
	}
	return rm.errorAction(404,inUri+":action not found");
}

func (rm *RouterManager)matchPattern(pattern,uri string) bool{
	if pattern==uri || strings.HasPrefix(uri,pattern) {
		return true
	}
	return false
}

func (rm *RouterManager)errorAction(errorCode int,errMsg string) *Action{
	action := &Action{}
	action.CtrlName = "ErrorController"
	action.FuncName = "Handle"
	return action;
}
