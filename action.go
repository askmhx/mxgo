package mxgo

import "reflect"

type Action struct {
	controller *Controller
	method string
}

func (action *Action)Execute(){
	result := reflect.ValueOf(&action.controller).MethodByName(action.method)
	err := result.render()
	if err {
		action.handleError(err)
	}
}

func (action *Action)handleError(err error){

}

