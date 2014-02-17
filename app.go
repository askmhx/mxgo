package mxgo

import "net/http"

type MxGoApp struct{

}

func NewMxGoApp() *MxGoApp{
	return &MxGoApp{}
}

func (mxgo *MxGoApp)Run(){
	http.Serve(mxgo)
}


