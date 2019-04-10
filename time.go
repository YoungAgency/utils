package utils

import (
	"time"
)

// GetMidnight returns the midnight related to t
func GetMidnight(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// GetYesterday returns a time rapresenting t yesterday at midnight
func GetYesterday(t time.Time) (r time.Time) {
	day, _ := time.ParseDuration("24h")
	r = GetMidnight(t.Add(day * -1))
	return
}

func GetTomorrow(t time.Time) (r time.Time) {
	day, _ := time.ParseDuration("24h")
	r = GetMidnight(t.Add(day))
	return
}

// ToMilliseconds returns t as milliseconds
func ToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func NsToMillis(ns int64) int64 {
	return ns / int64(1e6)
}

func TimeWithTimestamp(milliseconds int64) time.Time {
	return time.Unix(milliseconds/1e3, (milliseconds%1e3)*1e6).UTC()
}
