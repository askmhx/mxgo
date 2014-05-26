package mxgo

import (
	"strings"
)

type StaticController struct {
	Controller
}

func (ctrl StaticController)Handle() Result{
	uriPath := strings.Split(ctrl.Request.RequestURI,"/")
	if uriPath[0] == "favicon.ico" {

	}else if uriPath[0]==AppCfg.String("static.uri"){
		ctrl.Data = ""
	}
	return ctrl.Static()
}
