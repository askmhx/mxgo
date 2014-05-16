package mxgo

import "github.com/menghx/mxgo/httplib"

type Controller struct {
	Request *httplib.Request
	Response *httplib.Response
	Data interface {}
}

func (ctrl *Controller)Json() Result{
	result := JSONResult{}
	result.Data = ctrl.Data
	return result
}

func (ctrl *Controller)XML() Result{
	result := XMLResult{}
	result.Data = ctrl.Data
	return result
}

func (ctrl *Controller)Template() Result{
	result := TemplateResult{}
	result.Data = ctrl.Data
	return result
}

func (ctrl *Controller)Redirect() Result{
	result := RedirectResult{}
	result.Data = ctrl.Data
	return result
}

func (ctrl *Controller)Forward() Result{
	result := ForwardResult{}
	result.Data = ctrl.Data
	return result
}


func (ctrl *Controller)Plain() Result{
	result := PlainResult{}
	result.Data = ctrl.Data
	return result
}

func (ctrl *Controller)Static() Result{
	result := StaticResult{}
	result.Data = ctrl.Data
	return result
}
