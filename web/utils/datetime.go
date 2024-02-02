package utils

import (
	"time"
)

var Loc, _ = time.LoadLocation("Asia/Shanghai")
var Format = "2006-01-02 15:04:05"

func TimeToString(t time.Time) string {
	return t.Local().Format(Format)
}

func StringToTime(s string) (time.Time, error) {
	// time.Parse()
	return time.ParseInLocation(Format, s, Loc)
}
