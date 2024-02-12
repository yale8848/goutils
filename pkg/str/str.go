package str

import (
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

func ToPtr(s string) *string {
	return &s
}
func IsNotEmpty(s string) bool {

	return len(strings.TrimSpace(s)) != 0
}
func IsEmpty(s string) bool {

	return len(strings.TrimSpace(s)) == 0
}
func IsHaveEmpty(s ...string) bool {
	for _, v := range s {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}
func IsAllEmpty(s ...string) bool {
	for _, v := range s {
		if !IsEmpty(v) {
			return false
		}
	}
	return true
}
func SubStr(src string, start, endExclude int) string {
	runes := []rune(src)
	rLen := len(runes)
	if start < 0 || endExclude > rLen || start > endExclude {
		return ""
	}
	if start == 0 && endExclude == rLen {
		return src
	}
	return string(runes[start:endExclude])

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

func IsChinese(str string) bool {
	var hzRegexp = regexp.MustCompile("[\u4e00-\u9fa5]+")
	return hzRegexp.MatchString(str)
}

func IsSmallChars(str string) bool {
	var hzRegexp = regexp.MustCompile("[a-z]+")
	return hzRegexp.MatchString(str)
}
func StrToInt64(str string, def int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}
