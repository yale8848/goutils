package errs

import "log"

func Panic(e error) bool {
	if e != nil {
		log.Fatal(e.Error())
		return true
	}
	return false
}

func Check(e error) bool {
	if e != nil {
		log.Println(e.Error())
		return true
	}
	return false
}
