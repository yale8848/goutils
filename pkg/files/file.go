package files

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)


func MkDirs(fp string) error {
	dirPath:=filepath.Base(fp)
	return os.MkdirAll(dirPath, 0666)
}
func MkDirsWithoutFile(fp string) error {
	dirPath:=""
	dir,fl:=filepath.Split(fp)
	if strings.Contains(fl,".") {
		dirPath=dir
	}else{
		dirPath=fp
	}
	if len(dirPath)==0 {
		return errors.New("dir empty")
	}
	return os.MkdirAll(dirPath, 0666)
}
func MustMkDirs(fp string) error {
	return os.MkdirAll(fp, 0666)
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

	fs.Close()
	fd.Close()

	if err==nil&&isDeleteSrc {
		return os.Remove(srcFile)
	}
	return
}