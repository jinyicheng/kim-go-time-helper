package timeHelper

import "time"

// GetFirstDayOfLastMonth 获取上个月的第一天
func GetFirstDayOfLastMonth(t time.Time) time.Time {
	// 直接使用GetFirstDayOfMonth并向前推一个月
	return GetFirstDayOfMonth(t.AddDate(0, -1, 0))
}

// GetFirstDayOfCurrentMonth 获取本月的第一天
func GetFirstDayOfCurrentMonth() time.Time {
	// 调用GetFirstDayOfMonth并传入当前时间
	return GetFirstDayOfMonth(time.Now())
}

// GetFirstDayOfNextMonth 获取下个月的第一天
func GetFirstDayOfNextMonth(t time.Time) time.Time {
	// 同样直接调用GetFirstDayOfMonth并加上一个月
	return GetFirstDayOfMonth(t.AddDate(0, 1, 0))
}

// GetFirstDayOfMonth 获取指定日期对应的月第一天
func GetFirstDayOfMonth(t time.Time) time.Time {
	// 返回指定月份的第一天
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetLastDayOfCurrentMonth 获取本月的最后一天
func GetLastDayOfCurrentMonth(t time.Time) time.Time {
	// 获取下个月的第一天然后减去一天
	return GetFirstDayOfNextMonth(t).AddDate(0, 0, -1)
}

// GetLastDayOfLastMonth 获取上个月的最后一天
func GetLastDayOfLastMonth(t time.Time) time.Time {
	// 获取本月的第一天然后减去一天
	return GetFirstDayOfMonth(t).AddDate(0, 0, -1)
}

// GetLastDayOfNextMonth 获取下个月的最后一天
func GetLastDayOfNextMonth(t time.Time) time.Time {
	// 获取下下个月的第一天然后减去一天
	return GetFirstDayOfNextMonth(t).AddDate(0, 1, -1)
}

// GetLastDayOfMonth 获取指定日期对应的月最后一天
func GetLastDayOfMonth(t time.Time) time.Time {
	// 获取指定月份的下个月的第一天然后减去一天
	return GetFirstDayOfMonth(t.AddDate(0, 1, 0)).AddDate(0, 0, -1)
}

// GetNext3MonthsRange 获取指定日期对应的后3个月时间范围
func GetNext3MonthsRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetNextFewMonthsRange(t, 3)
}

// GetNextMonthRange 获取指定日期对应的下个月时间范围
func GetNextMonthRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetNextFewMonthsRange(t, 1)
}

// GetNextFewMonthsRange 获取指定日期对应的后几个月时间范围
func GetNextFewMonthsRange(t time.Time, numberOfMonths int) (startTime time.Time, endTime time.Time) {
	startDate := GetFirstDayOfMonth(t).AddDate(0, 1, 0)
	endDate := startDate.AddDate(0, numberOfMonths, -1)
	return startDate, endOfDate(endDate)
}

// GetLastMonthRange 获取指定日期对应的上个月时间范围
func GetLastMonthRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetFirstDayOfLastMonth(t), endOfDate(GetLastDayOfLastMonth(t))
}

// GetPast1MonthRange 获取指定日期对应的过去1个月时间范围
func GetPast1MonthRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetPastFewMonthsRange(t, 1)
}

// GetPast3MonthsRange 获取指定日期对应的过去3个月时间范围
func GetPast3MonthsRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetPastFewMonthsRange(t, 3)
}

// GetPastFewMonthsRange 获取指定日期对应的过去几个月时间范围
func GetPastFewMonthsRange(t time.Time, numberOfMonths int) (startTime time.Time, endTime time.Time) {
	endDate := startOfDate(t.AddDate(0, 0, -1))
	startDate := endDate.AddDate(0, -numberOfMonths, 0)
	return startOfDate(startDate), endOfDate(endDate)
}
