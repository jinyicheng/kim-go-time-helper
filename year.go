package timeHelper

import "time"

// GetLastDayOfLastYear 获取去年的最后一天
func GetLastDayOfLastYear() time.Time {
	// 获取当前时间
	now := time.Now()
	// 构造去年的第一天
	lastYearFirstDay := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, now.Location())
	// 返回去年最后一天的时间
	return lastYearFirstDay.AddDate(0, 11, 31)
}

// GetLastDayOfYear 获取指定日期对应的年最后一天
func GetLastDayOfYear(t time.Time) time.Time {
	// 获取下一年的第一天
	nextYear := time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
	// 返回下一年第一天的前一天，即为指定年的最后一天
	return nextYear.AddDate(0, 0, -1)
}

// GetYesterdayOfLastYear 获取去年的昨天
func GetYesterdayOfLastYear() time.Time {
	return time.Now().AddDate(-1, 0, -1)
}
