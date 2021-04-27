package str

func IsEmpty(s string) bool {
	return len(s) == 0
}

func SubStr(src string , start,endExclude int)string  {

	runes:=[]rune(src)
	rLen:=len(runes)
	if start<0||endExclude>rLen||start>endExclude {
		return ""
	}
	if start == 0 && endExclude == rLen {
		return src
	}
	return string(runes[start:endExclude])

}
