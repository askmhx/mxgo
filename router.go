package mxgo

import (
	"github.com/menghx/mxgo/httplib"
	"strings"
)

type router struct {
	UriPattern string
	HttpMethod string
	CtrlName ControllerInterface
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

func (rm *RouterManager)Router(uriPattern string,ctrl ControllerInterface,funcName string){
	route := router{}
	p := strings.Split(uriPattern,":")//METHOD:URI
	route.HttpMethod = p[0]
	route.UriPattern = p[1]
	route.CtrlName = ctrl
	route.FuncName = funcName
	rm.routes = append(rm.routes,route)
}

func (rm *RouterManager)FindAction(request *httplib.Request,response *httplib.Response) *Action{
	inMethod := request.Method
	inUri := request.RequestURI
	//need cache here url->action
	for _,r := range rm.routes{
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
	action.CtrlName = &ErrorController{}
	action.FuncName = "Handle"
	return action;
}
