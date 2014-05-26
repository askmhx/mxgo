package mxgo

import "reflect"


type Action struct {
	CtrlName ControllerInterface
	FuncName string
}

func NewAction(ctrlName ControllerInterface,funcName string) Action{
	action := &Action{}
	action.CtrlName = ctrlName
	action.FuncName = funcName
	return action;
}

func ErrorAction(errorCode int,errorMsg string) Action{
	action := &Action{}
	action.CtrlName = &ErrorController{}
	action.FuncName = "Handle"
	return action;
}


func (action *Action)Execute(mxGo *MxGoApp){
	mxGo.Fm.BeforeAction(nil)
	ctrlValue := reflect.ValueOf(action.CtrlName).Elem()
	mxLog.Debug("ACTION:",ctrlValue,action.FuncName)
	if ctrlFunc := ctrlValue.MethodByName(action.FuncName);ctrlFunc.IsValid() {
		reflect.TypeOf(ctrlFunc)
		request := ctrlValue.FieldByName("Request")
		if parseFormFunc := request.MethodByName("ParseForm");parseFormFunc.IsValid() {
			parseFormFunc.Call(make([]reflect.Value,0))//Call ParseForm Func
		}
		formValues := request.FieldByName("Form").Elem()
		params := make([]reflect.Value,0)
		result := ctrlFunc.Call(params)[0]
		mxLog.Debug(result)
	}
	mxGo.Fm.AfterAction(nil)
}
