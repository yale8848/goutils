package catch

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"time"
)

func Dmp() {
	errstr := ""
	if err := recover(); err != nil {
		errstr += (fmt.Sprintf("%v\r\n", err))
		errstr += ("--------------------------------------------\r\n")
	}
	errstr += (string(debug.Stack()))

}

func OnWriteErrToFile(errstring string) {
	path := GetModelPath() + "/err"
	if !PathExists(path) {
		os.MkdirAll(path, os.ModePerm)
	}

	now := time.Now()
	pid := os.Getpid()
	time_str := now.Format("2006-01-02")
	fname := fmt.Sprintf("%s/panic_%s-%x.log", path, time_str, pid)
	fmt.Println("panic to file ", fname)

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString("=========================" + now.Format("2006-01-02 15:04:05 ========================= \r\n"))
	f.WriteString(errstring)
	f.WriteString("=========================end=========================")
}

func GetModelPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(file)
	path, _ = filepath.Abs(path)

	return path
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
