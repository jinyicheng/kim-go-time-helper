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

// GetLastWeekRange 获取指定日期对应的上周时间范围
func GetLastWeekRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetLastMonday(t), endOfDate(GetLastSunday(t))
}

// GetPast1WeekRange 获取指定日期对应的过去1周时间范围
func GetPast1WeekRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetPastFewWeeksRange(t, 1)
}

// GetPast3WeeksRange 获取指定日期对应的过去3周时间范围
func GetPast3WeeksRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetPastFewWeeksRange(t, 3)
}

// GetPastFewWeeksRange 获取指定日期对应的过去几周时间范围
func GetPastFewWeeksRange(t time.Time, numberOfWeeks int) (startTime time.Time, endTime time.Time) {
	endDate := endOfDate(t.AddDate(0, 0, -1))
	startDate := endDate.AddDate(0, 0, -numberOfWeeks*7)
	return startOfDate(startDate), endDate
}
