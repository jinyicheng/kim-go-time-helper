package timeHelper

import (
	"time"
)

// startOfDate 将时间截断到凌晨时间
func startOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// startOfDate 将时间截断到凌晨时间
func endOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}
