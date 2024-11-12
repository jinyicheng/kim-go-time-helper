package timeHelper

import "time"

// GetLastDayOfLastWeek 获取指定日期对应上周的最后一天
func GetLastDayOfLastWeek(t time.Time) time.Time {
	return GetLastSunday(t)
}

// GetYesterdayOfLastWeek 获取上周的昨天
func GetYesterdayOfLastWeek(t time.Time) time.Time {
	return startOfDate(t.AddDate(0, 0, -8))
}

// GetFirstDayOfWeek 获取指定日期对应的周第一天
func GetFirstDayOfWeek(t time.Time) time.Time {
	return GetThisWeekMonday(t)
}

// GetNextWeekRange 获取指定日期对应的下周范围
func GetNextWeekRange(t time.Time) (startTime time.Time, endTime time.Time) {
	// 获取本周的第一天
	thisWeekFirstDay := GetFirstDayOfWeek(t)
	// 下周的第一天是本周第一天的后七天
	startTime = thisWeekFirstDay.AddDate(0, 0, 7)
	// 下周的最后一天是下周的第一天之后的第六天
	endTime = startTime.AddDate(0, 0, 6)
	// 返回下周第一天和最后一天的凌晨时间
	return startOfDate(startTime), endOfDate(endTime)
}
