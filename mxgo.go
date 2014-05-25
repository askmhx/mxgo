package mxgo

var (
	mxLog = NewMxLogger()
	app = NewMxGoApp()
	AppCfg = app.Cfg
)

func AddFilter(filter ...Filter){
	app.Fm.AddFilter(filter...)
}

func Router(uriPattern string,ctrl ControllerInterface,funcName string){
	app.Rm.Router(uriPattern,ctrl,funcName)
}

func EnableAdmin(enabled bool){
	app.EnableAdmin(enabled)
}

func Run(){
	app.Run()
}

