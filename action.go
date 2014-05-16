package mxgo


type Action struct {
	CtrlName string
	FuncName string
}

func (action *Action)Execute(){
	mxLog.Debug("ACTION:",action.CtrlName+"."+action.FuncName)
//	result := ctl.Call()
//	err := result.render()
//	if err {
//		action.handleError(err)
//	}
}

func (action *Action)handleError(err error){

}

