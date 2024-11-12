package timeHelper

import "time"

// GetLastDayOfLastYear 获取指定日期对应去年的最后一天
func GetLastDayOfLastYear(t time.Time) time.Time { // 构造去年的第一天
	// 返回去年最后一天的时间
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
}

// GetLastDayOfYear 获取指定日期对应的年最后一天
func GetLastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
}

// GetYesterdayOfLastYear 获取去年的昨天
func GetYesterdayOfLastYear(t time.Time) time.Time {
	return startOfDate(t.AddDate(-1, 0, -1))
}

// GetFirstDayOfYear 获取指定日期对应的年第一天
func GetFirstDayOfYear(t time.Time) time.Time {
	// 返回指定月份的第一天
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetNextYearRange 获取指定日期对应的明年
func GetNextYearRange(t time.Time) (startTime time.Time, endTime time.Time) {
	startDate := GetFirstDayOfYear(t).AddDate(1, 0, 0)
	endDate := endOfDate(GetLastDayOfYear(startDate))
	return startDate, endDate
}
