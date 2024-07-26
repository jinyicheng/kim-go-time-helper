package timeHelper

import (
	"time"
)

// GetYesterday 获取昨天
func GetYesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// GetTomorrow 获取明天
func GetTomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

func GetLastMonday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Monday
	daysToLastMonday := 0
	if currentWeekDay == 0 { // If today is Sunday, go back 6 days to Monday
		daysToLastMonday = 6
	} else { // Otherwise, go back to the last Monday
		daysToLastMonday = currentWeekDay - 1
	}
	// Calculate the date of the last Monday
	lastMonday := now.AddDate(0, 0, -daysToLastMonday-7)
	// Return the date at midnight of last Monday
	return time.Date(lastMonday.Year(), lastMonday.Month(), lastMonday.Day(), 0, 0, 0, 0, now.Location())
}

// GetLastTuesday returns the date of last Tuesday.
func GetLastTuesday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Tuesday
	daysToLastTuesday := 0
	if currentWeekDay < 2 { // If today is Sunday or Monday, go back to last week's Tuesday
		daysToLastTuesday = currentWeekDay + 5
	} else { // Otherwise, go back to Tuesday of this week and then an additional 7 days
		daysToLastTuesday = currentWeekDay - 2
	}
	// Calculate the date of the last Tuesday
	lastTuesday := now.AddDate(0, 0, -daysToLastTuesday-7)
	// Return the date at midnight of last Tuesday
	return time.Date(lastTuesday.Year(), lastTuesday.Month(), lastTuesday.Day(), 0, 0, 0, 0, now.Location())
}

func GetLastWednesday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Wednesday
	daysToLastWednesday := 0
	if currentWeekDay < 3 { // If today is Sunday, Monday, or Tuesday, go back to last week's Wednesday
		daysToLastWednesday = currentWeekDay + 4
	} else { // Otherwise, go back to Wednesday of this week and then an additional 7 days
		daysToLastWednesday = currentWeekDay - 3
	}
	// Calculate the date of the last Wednesday
	lastWednesday := now.AddDate(0, 0, -daysToLastWednesday-7)
	// Return the date at midnight of last Wednesday
	return time.Date(lastWednesday.Year(), lastWednesday.Month(), lastWednesday.Day(), 0, 0, 0, 0, now.Location())
}

func GetLastThursday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Thursday
	daysToLastThursday := 0
	if currentWeekDay < 4 { // If today is before Thursday, go back to last week's Thursday
		daysToLastThursday = currentWeekDay + 3
	} else { // Otherwise, go back to Thursday of this week and then an additional 7 days
		daysToLastThursday = currentWeekDay - 4
	}
	// Calculate the date of the last Thursday
	lastThursday := now.AddDate(0, 0, -daysToLastThursday-7)
	// Return the date at midnight of last Thursday
	return time.Date(lastThursday.Year(), lastThursday.Month(), lastThursday.Day(), 0, 0, 0, 0, now.Location())
}

// GetLastFriday returns the date of last Friday.
func GetLastFriday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Friday
	daysToLastFriday := 0
	if currentWeekDay < 5 { // If today is before Friday, go back to last week's Friday
		daysToLastFriday = currentWeekDay + 2
	} else { // Otherwise, go back to Friday of this week and then an additional 7 days
		daysToLastFriday = currentWeekDay - 5
	}
	// Calculate the date of the last Friday
	lastFriday := now.AddDate(0, 0, -daysToLastFriday-7)
	// Return the date at midnight of last Friday
	return time.Date(lastFriday.Year(), lastFriday.Month(), lastFriday.Day(), 0, 0, 0, 0, now.Location())
}

// GetLastSaturday 返回上一个星期的星期六
func GetLastSaturday() time.Time {
	now := time.Now()
	currentWeekDay := int(now.Weekday()) // 0 = Sunday, 1 = Monday, ..., 6 = Saturday

	// Calculate how many days ago was the last Saturday
	daysSinceLastSaturday := (currentWeekDay - 6 + 7) % 7

	// Calculate the date of the last Saturday
	lastSaturday := now.AddDate(0, 0, -daysSinceLastSaturday)

	// Return the date at midnight of last Saturday
	return time.Date(lastSaturday.Year(), lastSaturday.Month(), lastSaturday.Day(), 0, 0, 0, 0, now.Location())
}

func GetLastSunday() time.Time {
	now := time.Now()
	// Calculate the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	currentWeekDay := int(now.Weekday())
	// Calculate how many days to go back to the last Sunday
	daysToLastSunday := 0
	if currentWeekDay == 0 { // If today is Sunday, go back 7 days to last Sunday
		daysToLastSunday = 7
	} else { // Otherwise, go back to the previous Sunday
		daysToLastSunday = currentWeekDay
	}
	// Calculate the date of the last Sunday
	lastSunday := now.AddDate(0, 0, -daysToLastSunday)
	// Return the date at midnight of last Sunday
	return time.Date(lastSunday.Year(), lastSunday.Month(), lastSunday.Day(), 0, 0, 0, 0, now.Location())
}
