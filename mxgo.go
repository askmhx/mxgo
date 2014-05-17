package mxgo


var (
	mxLog = NewMxLogger()
	app = NewMxGoApp()
	AppCfg = app.Cfg
)

func AddFilter(filter ...Filter){
	app.Fm.AddFilter(filter...)
}

func Router(uri string,method string,action string){
	app.Rm.Router(uri,method,action)
}

func EnableAdmin(enabled bool){
	app.EnableAdmin(enabled)
}

func Run(){
	app.Run()
}

