package mxgo

import (
	"net/http"
	"fmt"
	"com.github/menghx/mxgo/httplib"
	"com.github/menghx/mxgo/config"
	"os"
	"time"
)

type MxGoApp struct{
	addr      string
	port      int
	enableSSL bool
	staticUrl string
	Rm *RouterManager
	Fm *FilterManager
	cfg *config.Config
	path      string
}

func NewMxGoApp() *MxGoApp {
	mxGo := &MxGoApp{}
	currentPath, _ := os.Getwd()
	mxGo.path = currentPath+"/src/blgo"
	mxGo.cfg = config.NewConfig(mxGo.path+MXGO_APP_CONFIG_FILE)
	mxGo.addr = mxGo.cfg.String(MXGO_APP_CONFIG_KEY_ADDR)
	mxGo.port = mxGo.cfg.Int(MXGO_APP_CONFIG_KEY_PORT)
	mxGo.Rm = NewRouterManager()
	mxGo.initRouter()
	mxGo.Fm = NewFilterManager()
	return mxGo
}

func (mxGO *MxGoApp)initRouter(){
	mxGO.Rm.AddRoute("/error/*","*","ErrorController.Handle")//erorr
	mxGO.Rm.AddRoute("/static/*","*","StaticController.Handle")//static
}

func (mxGo *MxGoApp)EnableAdmin(enable bool){
	if enable {
		mxGo.Rm.AddRoute("/admin/*","*","AdminController.Handle")
	}else{
		mxGo.Rm.DelRoute("/admin/*")
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
		certFile := mxGo.cfg.String(MXGO_APP_CONFIG_KEY_SSL_CERT_FILE) //parser from config
		keyFile := mxGo.cfg.String(MXGO_APP_CONFIG_KEY_SSL_KEY_FILE) //parser from config
		err = server.ListenAndServeTLS(certFile, keyFile)
	}else {
		err = server.ListenAndServe()
	}
	if err != nil {
		mxLog.Error("start server failed ", server.Addr)
	}
}


func (mxGo *MxGoApp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Server", MXGO_SERVER_NAME)
	var request = httplib.NewRequest(req)
	var response = httplib.NewResponse(rw)
	mxLog.Debug(request.URL, request.RequestURI)
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
