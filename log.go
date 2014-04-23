package mxgo

import (
	"fmt"
//	"log"
)

type MxLogger struct {

}

func NewMxLogger() *MxLogger{
	return &MxLogger{}
}

func (mxLog MxLogger)Error(strs ...interface{}){
	fmt.Println("ERR ",strs)
}

func (mxLog MxLogger)Trace(strs ...interface{}){
	fmt.Println("TRA ",strs)
}

func (mxLog MxLogger)Info(strs ...interface{}){
	fmt.Println("INF ",strs)
}

func (mxLog MxLogger)Debug(strs ...interface{}){
	fmt.Println("DEB ",strs)
}
