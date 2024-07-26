package timeHelper

import (
	"time"
)

// GetLastWeekDay 获取上周的某一天
func GetLastWeekDay(dayOfWeek time.Weekday) time.Time {
	now := time.Now()

	// 计算今天的星期几
	today := now.Weekday()

	// 如果今天已经是所求的那一天，我们需要回到上上周
	if today == dayOfWeek {
		return now.AddDate(0, 0, -7)
	}

	// 计算偏移量
	offset := int(dayOfWeek) - int(today)

	// 如果偏移量是正数，说明所求的那天在这周之后，因此需要减去7天
	if offset > 0 {
		offset -= 7
	}

	// 返回上周的某一天
	return now.AddDate(0, 0, offset)
}
