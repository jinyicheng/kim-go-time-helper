package timeHelper

import (
	"log"
	"testing"
	"time"
)

func TestGetLastDayOfLastWeek(t *testing.T) {
	log.Println("获取指定日期对应上周的最后一天")
	log.Println(GetLastDayOfLastWeek(time.Now()))
}

func TestGetYesterdayOfLastWeek(t *testing.T) {
	log.Println("获取上周的昨天")
	log.Println(GetYesterdayOfLastWeek(time.Now()))
}

func TestGetFirstDayOfWeek(t *testing.T) {
	log.Println("获取指定日期对应的周第一天")
	log.Println(GetFirstDayOfWeek(time.Now()))
}

func TestGetNextWeekRange(t *testing.T) {
	log.Println("获取指定日期对应的下周范围")
	log.Println(GetNextWeekRange(time.Now()))
}
