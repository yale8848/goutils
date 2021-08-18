package files

import (
	"github.com/yale8848/goutils/pkg/errs"
	"testing"
)

func TestCopy(t *testing.T) {
	err := Copy("./test_data/a.txt", "./test_data/b.txt", false)
	errs.Fatal(err)
}
func TestMkdirs(t *testing.T)  {
	pt:="test-data/aa/bb/cc.txt"
    //pt="aa"
	//d,f:=filepath.Split(pt)
	//fmt.Println("d "+d+" f "+f)
	//os.MkdirAll(pt, 0666)
	MustMkDirs(pt)
}
