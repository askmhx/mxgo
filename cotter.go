package mxgo

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"os/exec"
	"strings"
	"os"
	"path/filepath"
)

type Cotter struct {
	appHome string
	binPath string
	ctrlsMap map[string]Controller
}

func NewCotter() *Cotter {
	cotter := &Cotter{}
	cotter.appHome = cotter.getAppHome()
	cotter.binPath = cotter.getBinPath()
	return cotter
}

func (cotter Cotter) watchApp() {
	cotter.genAction()
	watchPaths := []string{}
	watchPaths = append(watchPaths,cotter.appHome)
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"conf"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app","admin"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app","controllers"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app","models"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app","filters"))
	watchPaths = append(watchPaths,filepath.Join(cotter.appHome,"app","views"))

	watchAction := func(fileName string, event *fsnotify.FileEvent) {
		if event.IsCreate() || event.IsModify() || event.IsRename() {
			if strings.HasSuffix(fileName,".go") {
				cotter.goExec("build", fileName)
			}
		}
	}
	WatchDoTask(watchAction,watchPaths...)
}

func (cotter Cotter) goExec(subCommand string, params ...string) {
	p :=[]string{subCommand}
	for _,file := range params {
		p = append(p,file)
	}
	dir,_ := filepath.Abs(filepath.Join(cotter.appHome,"..",".."))
	goCmd := exec.Command("go",p ...)
	goCmd.Env = []string{"GOPATH="+dir}
	mxLog.Info(goCmd.Env)
	cOut,err := goCmd.CombinedOutput()
	if err != nil {
		mxLog.Debug(string(cOut))
	}
}

func (cotter Cotter) genAction() {
	ctrlPath := filepath.Join(cotter.appHome,"app","controllers")
	tmp,_ := os.Create(filepath.Join(cotter.appHome,"tmps","ctrl.go"))
	walkFunc := func(path string, info os.FileInfo, err error) error {
		tmp.WriteString(path+"\n")
		return nil
	}
	filepath.Walk(ctrlPath, (filepath.WalkFunc)(walkFunc))
	fmt.Println("genAction")
}

func (cotter Cotter)getAppHome() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	mxLog.Debug(dir,err,os.Args)
	return "/Users/MengHX/WorkSpace/GOWork/mxgo/src/blog"
}

func (cotter Cotter)getBinPath() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	mxLog.Debug(dir,err,os.Args)
	return "/Users/MengHX/WorkSpace/GOWork/mxgo/bin/blog"
}

func (cotter Cotter)startApp(){
	mxLog.Debug("cotter start app")
}

func (cotter Cotter)stopApp(){
	mxLog.Debug("cotter stop app")
}

var action = "var CtrlsMap = map[string]Controller{\"%s\":%s}"

