package controller

import (
	"github.com/menghx/mxgo"
	"strings"
)

type StaticController struct {
	*mxgo.Controller
}

func (stc *StaticController)Handle() Result{
	uriPath := strings.Split(stc.request.RequestURI,"/")
	if uriPath[0] == "favicon.ico" {

	}else if uriPath[0]==mxgo.AppCfg.String("static.uri","static"){
		stc.Data = ""
	}
	return stc.Static()
}
