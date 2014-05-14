package mxgo

import "reflect"

type Action struct {
	CtrlName string
	FuncName string
}

func (action *Action)Execute(){
	ctl := reflect.ValueOf(&action.CtrlName)
	result := ctl.Call()
	err := result.render()
	if err {
		action.handleError(err)
	}
}

func (action *Action)handleError(err error){

}

