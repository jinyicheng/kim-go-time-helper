package timeHelper

import "time"

// GetFirstDayOfLastMonth 获取上个月的第一天
func GetFirstDayOfLastMonth() time.Time {
	// 直接使用GetFirstDayOfMonth并向前推一个月
	return GetFirstDayOfMonth(time.Now().AddDate(0, -1, 0))
}

// GetFirstDayOfCurrentMonth 获取本月的第一天
func GetFirstDayOfCurrentMonth() time.Time {
	// 调用GetFirstDayOfMonth并传入当前时间
	return GetFirstDayOfMonth(time.Now())
}

// GetFirstDayOfNextMonth 获取下个月的第一天
func GetFirstDayOfNextMonth() time.Time {
	// 同样直接调用GetFirstDayOfMonth并加上一个月
	return GetFirstDayOfMonth(time.Now().AddDate(0, 1, 0))
}

// GetFirstDayOfMonth 获取指定日期对应的月第一天
func GetFirstDayOfMonth(t time.Time) time.Time {
	// 返回指定月份的第一天
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetLastDayOfCurrentMonth 获取本月的最后一天
func GetLastDayOfCurrentMonth() time.Time {
	// 获取下个月的第一天然后减去一天
	return GetFirstDayOfNextMonth().AddDate(0, 0, -1)
}

// GetLastDayOfLastMonth 获取上个月的最后一天
func GetLastDayOfLastMonth() time.Time {
	// 获取本月的第一天然后减去一天
	return GetFirstDayOfCurrentMonth().AddDate(0, 0, -1)
}

// GetLastDayOfNextMonth 获取下个月的最后一天
func GetLastDayOfNextMonth() time.Time {
	// 获取下下个月的第一天然后减去一天
	return GetFirstDayOfNextMonth().AddDate(0, 1, -1)
}

// GetLastDayOfMonth 获取指定日期对应的月最后一天
func GetLastDayOfMonth(t time.Time) time.Time {
	// 获取指定月份的下个月的第一天然后减去一天
	return GetFirstDayOfMonth(t.AddDate(0, 1, 0)).AddDate(0, 0, -1)
}
