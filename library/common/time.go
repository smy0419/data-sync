package common

import (
	"time"
)

const (
	SecondsOfDay int64 = 24 * 60 * 60
	ExpireDay    int64 = 30
)

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func NowSecond() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimestampToDateTime(t int64) string {
	tm := time.Unix(t/1000, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func SecondToDay(t int64) int64 {
	template := "2006-01-02"
	s := time.Unix(t, 0).Format(template)
	zero, _ := time.ParseInLocation(template, s, time.Local)
	return zero.Unix()
}
