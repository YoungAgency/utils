package utils

import (
	"time"
)

const (
	dayMillis int64 = 1000 * 60 * 60 * 24
)

// Midnight returns t UTC midnight
func Midnight(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// Yesterday subtract 24h from t
func Yesterday(t time.Time) (r time.Time) {
	day, _ := time.ParseDuration("24h")
	r = t.Add(day * -1)
	return
}

// Tomorrow subtract 24h from t
func Tomorrow(t time.Time) (r time.Time) {
	day, _ := time.ParseDuration("24h")
	r = t.Add(day)
	return
}

// TimeToTimestamp returns t as milliseconds
func TimeToTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// TimestampToTime retuns time.Time rapresentation of milliseconds
func TimestampToTime(milliseconds int64) time.Time {
	return time.Unix(milliseconds/1e3, (milliseconds%1e3)*1e6).UTC()
}

// NsToMs returns ns as ms
func NsToMs(ns int64) int64 {
	return ns / int64(time.Millisecond)
}

// GetMidnight returns timestamp date at midnight
func GetMidnight(timestamp int64) int64 {
	offset := timestamp % dayMillis
	return timestamp - offset
}
