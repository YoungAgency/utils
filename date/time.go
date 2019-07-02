package date

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

// Tomorrow Add 24h from t
func Tomorrow(t time.Time) (r time.Time) {
	day, _ := time.ParseDuration("24h")
	r = t.Add(day)
	return
}

// WeekStart get the week start
func WeekStart(t time.Time) (r time.Time) {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	r = Midnight(t)
	r = r.Add(-1 * (weekday - 1) * 24 * time.Hour)
	return
}

// WeekEnd get the week end
func WeekEnd(t time.Time) (r time.Time) {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	off := 7 - weekday
	r = Midnight(t)
	r = r.Add((off + 1) * 24 * time.Hour)
	return
}

// Timestamp returns t timestamp
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// NowTimestamp returns time.Now timestamp
func NowTimestamp() int64 {
	return Timestamp(time.Now())
}

// Time returns time.Time rapresentation of milliseconds
func Time(milliseconds int64) time.Time {
	return time.Unix(milliseconds/1e3, (milliseconds%1e3)*1e6).UTC()
}

// Time returns time.Time rapresentation of ns
func NsTime(ns int64) time.Time {
	return time.Unix(ns/1e9, ns%1e9).UTC()
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

// GetYesterday subtracts 24h to timestamp
func GetYesterday(timestamp int64) int64 {
	return timestamp - dayMillis
}

// GetTomorrow adds 24h to timestamp
func GetTomorrow(timestamp int64) int64 {
	return timestamp + dayMillis
}

// GetWeekStart get the week start
func GetWeekStart(timestamp int64) (r int64) {
	return Timestamp(WeekStart(Time(timestamp)))
}

// GetWeekEnd get the week end
func GetWeekEnd(timestamp int64) (r int64) {
	return Timestamp(WeekEnd(Time(timestamp)))
}
