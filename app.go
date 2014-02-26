package mxgo

import (
	"net/http"
	"fmt"
	"strings"
	"os"
	"com.github/menghx/mxgo/mxhttp"
	"com.github/menghx/mxgo/router"
	"com.github/menghx/mxgo/config"
)

type MxGoApp struct{
	addr  string
	port  int
	enableSSL bool
	staticUrl string
	router *router.Router
	config map[string]string
}

func NewMxGoApp() *MxGoApp {
	mxGo := &MxGoApp{}
	mxGo.config = config.ParserConfig("app.conf")
	mxGo.router = router.NewRouter()
	mxGo.router.ParserConfig("routes")
	return &MxGoApp{}
}

func (mxGo *MxGoApp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	response := mxhttp.NewResponse(rw)
	request := *mxhttp.NewRequest(req)
	if Split(request.RequestURI,"/")[0] == mxGo.staticUrl || request.RequestURI == "favicon.ico"{
		mxGo.executeStatic(response,request)
		return;
	}

	mxGo.executeAction(response,request)
}


func (mxGo *MxGoApp) Run() {
	server := &Server{
		Addr:    fmt.Sprintf("%s:%d", mxGo.addr, mxGo.port),
		Handler:mxGo,
	}
	var err
	if mxGo.enableSSL {
		err = server.ListenAndServeTLS("", "")
	}else {
		err = server.ListenAndServe()
	}
	if err {
		mxLog.Error("start server failed ", mxGo.addr, ":", mxGo.port)
	}
}


func (mxGo *MxGoApp)executeStatic(resp mxhttp.Response, req *mxhttp.Request){
	//find static
	//write static to response
	//close
}


func (mxGo *MxGoApp)executeAction(resp mxhttp.Response, req *mxhttp.Request){
	action := mxGo.router.FindAction(req.RequestURI)
	action()
}
