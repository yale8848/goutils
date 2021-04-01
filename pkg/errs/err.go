package errs

import (
	"fmt"
	"log"
	"runtime"
)

func Fatal(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}
func Panic(e error) bool {
	if e != nil {
		panic(e)
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
type CallMessage func(message string)

func Recover( callMessage CallMessage){

	if err := recover(); err != nil {
		var stacktrace string
		for i := 1; ; i++ {
			_, f, l, got := runtime.Caller(i)
			if !got {
				break
			}
			stacktrace += fmt.Sprintf("%s:%d\n", f, l)
		}
		logMessage := fmt.Sprintf("Trace: %s\n", err)
		logMessage += fmt.Sprintf("\n%s", stacktrace)

		callMessage(logMessage)
	}
}