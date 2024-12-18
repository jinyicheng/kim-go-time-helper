package timeHelper

import (
	"log"
	"testing"
	"time"
)

func TestGetFirstDayOfLastMonth(t *testing.T) {
	log.Println("获取上个月第一天")
	log.Println(GetFirstDayOfLastMonth(time.Now()))
}
func TestGetFirstDayOfCurrentMonth(t *testing.T) {
	log.Println("获取本月第一天")
	log.Println(GetFirstDayOfCurrentMonth())
}
func TestGetFirstDayOfNextMonth(t *testing.T) {
	log.Println("获取下个月第一天")
	log.Println(GetFirstDayOfNextMonth(time.Now()))
}
func TestGetFirstDayOfMonth(t *testing.T) {
	log.Println("获取本月第一天")
	log.Println(GetFirstDayOfMonth(time.Now()))
}
func TestGetLastDayOfCurrentMonth(t *testing.T) {
	log.Println("获取本月最后一天")
	log.Println(GetLastDayOfCurrentMonth(time.Now()))
}
func TestGetLastDayOfLastMonth(t *testing.T) {
	log.Println("获取上个月最后一天")
	log.Println(GetLastDayOfLastMonth(time.Now()))
}
func TestGetLastDayOfNextMonth(t *testing.T) {
	log.Println("获取下个月最后一天")
	log.Println(GetLastDayOfNextMonth(time.Now()))
}
func TestGetLastDayOfMonth(t *testing.T) {
	log.Println("获取本月最后一天")
	log.Println(GetLastDayOfMonth(time.Now()))
}

func TestGetNext3MonthsRange(t *testing.T) {
	log.Println("之后3个月起止")
	log.Println(GetNext3MonthsRange(time.Now()))
	log.Println(GetNext3MonthsRange(time.Date(2024, 12, 1, 0, 0, 0, 0, time.Local)))
	log.Println(GetNext3MonthsRange(time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)))
}
func TestGetNextMonthRange(t *testing.T) {
	log.Println("下个月起止")
	log.Println(GetNextMonthRange(time.Now()))
}
func TestGetNextFewMonthsRange(t *testing.T) {
	log.Println("之后6个月起止")
	log.Println(GetNextFewMonthsRange(time.Now(), 6))
}

func TestGetLastMonthRange(t *testing.T) {
	log.Println("获取指定日期对应的上个月时间范围")
	log.Println(GetLastMonthRange(time.Now()))
}

func TestGetPast1MonthRange(t *testing.T) {
	log.Println("获取指定日期对应的过去1个月时间范围")
	log.Println(GetPast1MonthRange(time.Now()))
}

func TestGetPast3MonthsRange(t *testing.T) {
	log.Println("获取指定日期对应的过去3个月时间范围")
	log.Println(GetPast3MonthsRange(time.Now()))
}
