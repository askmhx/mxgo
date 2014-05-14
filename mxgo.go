package mxgo

const (
	MXGO_VERSION = "0.0.1"
	MXGO_SERVER_NAME = "MxGo"
	MXGO_ROOT_PATH = "com.github/menghx/mxgo/"
	MXGO_APP_CONFIG_FILE = "/conf/app.conf"
	MXGO_APP_CONFIG_KEY_ADDR = "addr"
	MXGO_APP_CONFIG_KEY_PORT = "port"
	MXGO_APP_CONFIG_KEY_SSL_KEY_FILE = "ssl.key_file"
	MXGO_APP_CONFIG_KEY_SSL_CERT_FILE = "ssl.cert_file"
)

var (
	mxLog = NewMxLogger()
	mxGoApp = NewMxGoApp()
)


func AddFilter(filters ...Filter){
	mxGoApp.Fm.AddFilter(filters)
}

func AddRoute(uri string,method string,action string){
	mxGoApp.Rm.AddRoute(uri,method,action)
}

func EnableAdmin(enabled bool){
	mxGoApp.EnableAdmin(enabled)
}

func Run(){
	mxGoApp.Run()
}

