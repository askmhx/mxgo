package mxgo

const (
	MXGO_SERVER_NAME = "MxGo"
	MXGO_ROOT_PATH = "com.github/menghx/mxgo/"
	CONFIG_FILE = "/conf/app.conf"
	CONFIG_KEY_ADDR = "addr"
	CONFIG_KEY_PORT = "port"
	CONFIG_KEY_SSL_KEY_FILE = "ssl.key_file"
	CONFIG_KEY_SSL_CERT_FILE = "ssl.cert_file"
	ROUTER_FILE = "/conf/routes"
)

var (
	mxLog = NewMxLogger()
)

func Run(){
	mxGoApp := NewMxGoApp()
	mxGoApp.Run()
}

