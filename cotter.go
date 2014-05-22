package mxgo

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"os/exec"
	"strings"
	"os"
)

type Cotter struct {
}

func NewCotter() *Cotter {
	cotter := &Cotter{}
	return cotter
}

func (cotter Cotter) startWatchTask() {
	WatchDoTask(func(fileName string, event *fsnotify.FileEvent) {
			if event.IsCreate() || event.IsModify() || event.IsRename() {
				if strings.HasSuffix(fileName,".go") {
					cotter.goExec("build", fileName)
				}
			}
		},
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/conf",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app/controller",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app/admin",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app/filter",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app/model",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog/app/view",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo/httplib",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo/controller")
}

func (cotter Cotter)isExecWithBin() bool{
	return false
}


func (cotter Cotter) goExec(subCommand string, params ...string) {
	p :=[]string{subCommand}
	for _,file := range params {
		p = append(p,file)
	}
	os.Setenv("GOPATH","/Users/MengHX/WorkSpace/GOWork/mxgo")
	goCmd := exec.Command("go",p ...)
	goCmd.Env = []string{"GOPATH=/Users/MengHX/WorkSpace/GOWork/mxgo"}
	mxLog.Info(p,os.Environ())
	cOut,err := goCmd.CombinedOutput()
	if err != nil {
		mxLog.Debug(string(cOut))
	}
}

func (this Cotter) genAction() {
	fmt.Println("genAction")
}
