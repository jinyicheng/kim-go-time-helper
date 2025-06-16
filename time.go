package timeHelper

import (
	"time"
)

// GetDateTimeRange 获取指定日期对应的时间范围
func GetDateTimeRange(t time.Time) (startTime time.Time, endTime time.Time) {
	return startOfDate(t), endOfDate(t)
}
