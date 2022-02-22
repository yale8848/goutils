package md5s

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(string string) string {
	h := md5.New()
	h.Write([]byte(string))
	return hex.EncodeToString(h.Sum(nil))
}
