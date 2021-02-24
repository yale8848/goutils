package files

import (
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