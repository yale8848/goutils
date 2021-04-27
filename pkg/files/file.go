package files

import (
	"io"
	"os"
	"path/filepath"
)

func MkDirs(fp string) {
	fd := filepath.Base(fp)
	os.MkdirAll(fd, 0666)
}

func IsExist(fpath string) (bool,error) {
	_, err := os.Stat(fpath)
	if err==nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}
func IsFileExist(fpath string) (bool) {
	_, err := os.Stat(fpath)
	if err==nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Copy(srcFile,distFile string,isDeleteSrc bool) (err error) {
	fs,err:=os.Open(srcFile)
	if err!=nil {
		return
	}
	fd,err:=os.Create(distFile)
	if err!=nil {
		return
	}
	_,err=io.Copy(fd,fs)
	if err==nil&&isDeleteSrc {
		return os.Remove(srcFile)
	}
	return
}