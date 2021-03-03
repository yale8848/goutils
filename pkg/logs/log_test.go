package logs

import (
	"sync"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {



	wg:=sync.WaitGroup{}
	go func() {
		wg.Add(1)
		lg:=NewLogger("logsdir",1,10)

		lg.Info("aaaa")
		wg.Done()
	}()
	time.Sleep(3*time.Second)
	wg.Wait()

}
