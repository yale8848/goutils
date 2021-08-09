package files

import (
	"github.com/yale8848/goutils/pkg/errs"
	"testing"
)

func TestCopy(t *testing.T) {
	err := Copy("./test_data/a.txt", "./test_data/b.txt", false)
	errs.Fatal(err)
}
