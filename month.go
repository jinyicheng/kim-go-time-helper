package timeHelper

import "time"

// GetFirstDayOfLastMonth 获取上个月的第一天
func GetFirstDayOfLastMonth() time.Time {
	return GetFirstDayOfCurrentMonth().AddDate(0, -1, 0)
}

// GetFirstDayOfCurrentMonth 获取本月的第一天
func GetFirstDayOfCurrentMonth() time.Time {
	return GetFirstDayOfMonth(time.Now())
}

// GetFirstDayOfMonth 获取指定日期对应的月第一天
func GetFirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetLastDayOfCurrentMonth 获取本月的最后一天
func GetLastDayOfCurrentMonth() time.Time {
	return GetFirstDayOfCurrentMonth().AddDate(0, 1, -1)
}

// GetLastDayOfLastMonth 获取上个月的最后一天
func GetLastDayOfLastMonth() time.Time {
	return GetFirstDayOfLastMonth().AddDate(0, 0, -1)
}

// GetLastDayOfMonth 获取指定日期对应的月最后一天
func GetLastDayOfMonth(t time.Time) time.Time {
	return GetFirstDayOfMonth(t).AddDate(0, 1, -1)
}
