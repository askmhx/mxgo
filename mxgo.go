package mxgo

var (
	mxLog = NewMxLogger()
)

func Run(){
	mxGoApp := NewMxGoApp()
	mxGoApp.Run()
}

