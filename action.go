package mxgo


type Action struct {
	CtrlName string
	FuncName string
}

func (action *Action)Execute(mxGo *MxGoApp){
	mxGo.Fm.BeforeAction(nil)
	mxLog.Debug("ACTION:",action.CtrlName+"."+action.FuncName)
//	result := ctl.Call()
//	err := result.render()
//	if err {
//		action.handleError(err)
//	}
	mxGo.Fm.AfterAction(nil)
}

func (action *Action)handleError(err error){

}

