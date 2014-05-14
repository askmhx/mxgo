package mxgo

import (
	"com.github/menghx/mxgo/httplib"
	"reflect"
	"strings"
)

type Router struct {
	 pathPatterns map[string]map[string]string
}

func NewRouter() *Router{
	route := &Router{}
	return route
}

func (router *Router)InitContainer(){

}

func (route *Router)FindAction(request httplib.Request,response httplib.Response) Action{
	method := request.Method
	uri := request.RequestURI
	pathPattern := route.pathPatterns[uri]
	if pathPattern == nil {
		return ErrorAction(404,uri+":action not found")
	}
	pAction := pathPattern[method]
	if pAction == nil {
		return ErrorAction(405,uri+":use http:"+method+" not allowed")
	}
	cName := strings.Split(pAction,".")
	controller := reflect.New(cName[0])
	controller.MethodByName(cName[1])

	return nil
}

func ErrorAction(errorCode int,errMsg string) Action{
	return nil;
}


//if strings.Split(request.RequestURI, "/")[0] == mxGo.staticUrl || request.RequestURI == "/favicon.ico" {
//mxGo.execStatic(request, response)
//return;
//}
