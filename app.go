package mxgo

import (
	"net/http"
	"fmt"
	"github.com/menghx/mxgo/httplib"
	"github.com/menghx/mxgo/config"
	"os"
	"time"
	"path"
)

const (
	MxGoVersion = "0.0.1"
	MxGoServerName = "MxGo"
	MxGoImportPath = "github.com/menghx/mxgo"
)


type MxGoApp struct{
	AppName string
	AppPath   string
	addr      string
	port      int
	enableSSL bool
	StaticUri string
	Cfg *config.Config
	Rm *RouterManager
	Fm *FilterManager
}

func NewMxGoApp() *MxGoApp {
	mxGo := &MxGoApp{}
	currentPath, _ := os.Getwd()
	mxGo.AppPath = path.Join(currentPath,"src","blgo")
	mxGo.Cfg = config.NewConfig(path.Join(mxGo.AppPath,"conf","app.conf"))
	mxGo.AppName = mxGo.Cfg.String("name")
	mxGo.StaticUri = mxGo.Cfg.String("static.uri")
	mxGo.addr = mxGo.Cfg.String("addr")
	mxGo.port = mxGo.Cfg.Int("port")
	mxGo.Rm = NewRouterManager()
	mxGo.initRouter()
	mxGo.Fm = NewFilterManager()
	return mxGo
}

func (mxGO *MxGoApp)initRouter(){
	mxGO.Rm.Router("/error/*","*","ErrorController.Handle")//erorr
	mxGO.Rm.Router(mxGO.StaticUri,"*","StaticController.Handle")//static
	mxGO.Rm.Router("/favicon.ico","*","StaticController.Handle")//static
}

func (mxGo *MxGoApp)EnableAdmin(enable bool){
	if enable {
		mxGo.Rm.Router("/admin/*","*","AdminController.Handle")
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
		certFile := mxGo.Cfg.String("ssl.cert_file") //parser from config
		keyFile := mxGo.Cfg.String("ssl.key_file") //parser from config
		err = server.ListenAndServeTLS(certFile, keyFile)
	}else {
		err = server.ListenAndServe()
	}
	if err != nil {
		mxLog.Error("start server failed ", server.Addr)
	}
}


func (mxGo *MxGoApp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Server", MxGoServerName)
	var request = httplib.NewRequest(req)
	var response = httplib.NewResponse(rw)
	mxLog.Debug("ServeHTTP:",request.URL)
	if SecurityVerify(request, response){
		mxGo.execAction(request, response)
	}
}


func (mxGo *MxGoApp) execAction(request *httplib.Request, response *httplib.Response) {
	action := mxGo.Rm.FindAction(request,response)
	mxGo.Fm.BeforeAction(action)
	action.Execute()
	mxGo.Fm.AfterAction(action)
}
