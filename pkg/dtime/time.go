package dtime

import (
	"sync"
	"time"
)

var (
	randSeek = int64(1)
	l        sync.Mutex
	zone     = "CST"
)

func TimeIntToDate(time_int int) string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Unix(int64(time_int), 0).In(cstZone).Format("2006-01-02 15:04:05")
}