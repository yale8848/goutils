package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Byte(s []byte) string {
	h := sha1.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256Byte(s []byte) string {
	h := sha256.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha512Byte(s []byte) string {
	h := sha512.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}
