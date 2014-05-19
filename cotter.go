package mxgo

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"os/exec"
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
				cotter.goExec("build", fileName)
			}
			mxLog.Info(fileName, event)
		},
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/conf",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app/controller",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app/admin",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app/filter",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app/model",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/blgo/app/view",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo/httplib",
		"/Users/MengHX/WorkSpace/GOWork/mxgo/src/github.com/menghx/mxgo/controller")
}


func (cotter Cotter) goExec(subCommand string, params ...string) {
	goCmd := exec.Command("go",append(append([]string{},subCommand),params))
	cOut,err := goCmd.CombinedOutput()
	if err != nil {
		mxLog.Debug(string(cOut))
	}
	mxLog.Error("ERROR",err)
}

func (this Cotter) genAction() {
	fmt.Println("")
}
