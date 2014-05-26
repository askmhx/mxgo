package mxgo

import "reflect"


type Action struct {
	CtrlName ControllerInterface
	FuncName string
}

func (action *Action)Execute(mxGo *MxGoApp){
	mxGo.Fm.BeforeAction(nil)
	mxLog.Debug("ACTION:",action.CtrlName,action.FuncName)
	ctrlValue := reflect.ValueOf(action.CtrlName).Elem()
	if method := ctrlValue.MethodByName(action.FuncName);method.IsValid() {
		params := make([]reflect.Value,0)
		result := method.Call(params)[0]
		mxLog.Debug(result)
	}
	mxGo.Fm.AfterAction(nil)
}

func (action *Action)handleError(err error){

}

