package mxgo

import (
	"com.github/menghx/mxgo/httplib"
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

func (rm *RouterManager)AddRoute(uriPattern string,httpMethod string,action string){
	route := router{}
	route.UriPattern = uriPattern
	route.HttpMethod = httpMethod
	actionPattern := strings.Split(action,".")
	route.CtrlName = actionPattern[0]
	route.FuncName = actionPattern[1]
	append(rm.routes,route)
}

func (rm *RouterManager)DelRoute(uri string){
	for i := range rm.routes{
		r := rm.routes[i]
		if r.UriPattern == uri {
			//delete(r)
			//need delete cache here url->action
		}
	}
}

func (rm *RouterManager)FindAction(request *httplib.Request,response *httplib.Response) Action{
	inMethod := request.Method
	inUri := request.RequestURI
	//need cache here url->action
	for r := range rm.routes{
		if r.UriPattern == inUri {
			if r.HttpMethod == inMethod || r.HttpMethod== "*" {
				action := &Action{}
				action.CtrlName = r.CtrlName
				action.FuncName = r.FuncName
				return action
			}else{
				return ErrorAction(405,inUri+":use http:"+inMethod+" not allowed")
			}
		}
	}
	return ErrorAction(404,inUri+":action not found");
}

func ErrorAction(errorCode int,errMsg string) Action{
	action := &Action{}
	action.CtrlName = "ErrorController"
	action.FuncName = "handle"
	return action;
}
