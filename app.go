package mxgo

import (
	"net/http"
	"fmt"
	"com.github/menghx/mxgo/httplib"
	"com.github/menghx/mxgo/router"
	"com.github/menghx/mxgo/config"
	"strings"
	"os"
	"time"
	"path"
)

type MxGoApp struct{
	addr      string
	port      int
	enableSSL bool
	staticUrl string
	router *Router
	config *config.Config
	path      string
}

func NewMxGoApp() *MxGoApp {
	mxGo := &MxGoApp{}
	currentPath, _ := os.Getwd()
	mxGo.path = currentPath+"/src/blgo"
	mxGo.config = config.NewConfig(mxGo.path+CONFIG_FILE)
	mxGo.addr = mxGo.config.String(CONFIG_KEY_ADDR)
	mxGo.port = mxGo.config.Int(CONFIG_KEY_PORT)
	mxGo.router = NewRouter()
	mxGo.router.InitContainer()
	return mxGo
}

func (mxGo *MxGoApp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Server", MXGO_SERVER_NAME)
	var request = httplib.NewRequest(req)
	var response = httplib.NewResponse(rw)
	mxLog.Debug(request.URL, request.RequestURI)
	if Verify(request, response){
		mxGo.execAction(request, response)
	}
}

func (mxGo *MxGoApp) Run() {
	server := &http.Server{
		Addr:fmt.Sprintf("%s:%d", mxGo.addr, mxGo.port),
		Handler:mxGo,
	}
	var err error

	go func() {
		time.Sleep(100 * time.Millisecond)
		mxLog.Info("Start server success listening ", server.Addr)
	}()

	if mxGo.enableSSL {
		certFile := mxGo.config.String(CONFIG_KEY_SSL_CERT_FILE) //parser from config
		keyFile := mxGo.config.String(CONFIG_KEY_SSL_KEY_FILE) //parser from config
		err = server.ListenAndServeTLS(certFile, keyFile)
	}else {
		err = server.ListenAndServe()
	}
	if err != nil {
		mxLog.Error("start server failed ", server.Addr)
	}
}


func (mxGo *MxGoApp) execAction(request *httplib.Request, response *httplib.Response) {
	action := mxGo.router.FindAction(request.RequestURI)
	action()
}
