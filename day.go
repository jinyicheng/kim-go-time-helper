package timeHelper

import (
	"time"
)

// GetThisWeekMonday 获取传入时间对应的本周星期一
func GetThisWeekMonday(t time.Time) time.Time {
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(t.Weekday())
	// Calculate how many days to go back to this week's Monday
	daysToMonday := 0
	if currentWeekDay == 0 { // If today is Sunday, go back 6 days to Monday
		daysToMonday = 6
	} else { // Otherwise, go back to Monday
		daysToMonday = currentWeekDay - 1
	}
	// Calculate the date of this week's Monday
	thisWeekMonday := t.AddDate(0, 0, -daysToMonday)
	// Return the date at midnight of this week's Monday
	return startOfDate(thisWeekMonday)
}

// GetThisWeekTuesday 获取传入时间对应的本周星期二
func GetThisWeekTuesday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 1))
}

// GetThisWeekWednesday 获取传入时间对应的本周星期三
func GetThisWeekWednesday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 2))
}

// GetThisWeekThursday 获取传入时间对应的本周星期四
func GetThisWeekThursday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 3))
}

// GetThisWeekFriday 获取传入时间对应的本周星期五
func GetThisWeekFriday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 4))
}

// GetThisWeekSaturday 获取传入时间对应的本周星期六
func GetThisWeekSaturday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 5))
}

// GetThisWeekSunday 获取传入时间对应的本周星期日
func GetThisWeekSunday(t time.Time) time.Time {
	return startOfDate(GetThisWeekMonday(t).AddDate(0, 0, 6))
}

// GetYesterday 获取昨天
func GetYesterday(t time.Time) time.Time {
	return startOfDate(t.AddDate(0, 0, -1))
}

// GetTomorrow 获取明天
func GetTomorrow(t time.Time) time.Time {
	return startOfDate(t.AddDate(0, 0, 1))
}

// GetLastMonday 获取上个星期一
func GetLastMonday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -7)
}

// GetLastTuesday 获取上个星期二
func GetLastTuesday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -6)
}

// GetLastWednesday 获取上个星期三
func GetLastWednesday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -5)
}

// GetLastThursday 获取上个星期四
func GetLastThursday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -4)
}

// GetLastFriday 获取上个星期五
func GetLastFriday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -3)
}

// GetLastSaturday 获取上个星期六
func GetLastSaturday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -2)
}

// GetLastSunday 获取上个星期日
func GetLastSunday(t time.Time) time.Time {
	return GetThisWeekMonday(t).AddDate(0, 0, -1)
}

// GetNext31DaysRange 获取指定日期对应的之后31天时间范围
func GetNext31DaysRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return GetNextFewDaysRange(t, 31)
}

// GetNextFewDaysRange 获取指定日期对应的之后几天时间范围
func GetNextFewDaysRange(t time.Time, numberOfDays int) (startTime time.Time, endTime time.Time) {
	startDate := startOfDate(t).AddDate(0, 0, 1)
	endDate := startDate.AddDate(0, 0, numberOfDays)
	return startDate, endOfDate(endDate)
}
