package files

import (
	"os"
	"path/filepath"
)

func MkDirs(fp string) {
	fd := filepath.Base(fp)
	os.MkdirAll(fd, 0666)
}
