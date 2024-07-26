package timeHelper

import (
	"time"
)

// GetYesterday 获取昨天
func GetYesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// GetTomorrow 获取明天
func GetTomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}
