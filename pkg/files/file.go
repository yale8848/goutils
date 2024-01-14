package files

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func MkDirs(fp string) error {
	dirPath := filepath.Base(fp)
	return os.MkdirAll(dirPath, 0666)
}
func MkDirsWithoutFile(fp string) error {
	dirPath := ""
	dir, fl := filepath.Split(fp)
	if strings.Contains(fl, ".") {
		dirPath = dir
	} else {
		dirPath = fp
	}
	if len(dirPath) == 0 {
		return errors.New("dir empty")
	}
	return os.MkdirAll(dirPath, 0666)
}
func MustMkDirs(fp string) error {
	return os.MkdirAll(fp, 0666)
}

func IsExist(fpath string) (bool, error) {
	_, err := os.Stat(fpath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func IsFileExist(fpath string) bool {
	_, err := os.Stat(fpath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Copy(srcFile, distFile string, isDeleteSrc bool) (err error) {
	fs, err := os.Open(srcFile)
	if err != nil {
		return
	}

	fd, err := os.Create(distFile)
	if err != nil {
		return
	}
	_, err = io.Copy(fd, fs)

	fs.Close()
	fd.Close()

	if err == nil && isDeleteSrc {
		return os.Remove(srcFile)
	}
	return
}
func DeleteFile(f string) error {
	return os.RemoveAll(f)
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

// 获取当前目录下所有文件
func GetFileList(path string) []string {
	var fileList []string
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if !f.IsDir() {
			fileList = append(fileList, f.Name())
		}
	}
	return fileList
}

// 判断所给路径文件/文件夹是否存在
func filexists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
