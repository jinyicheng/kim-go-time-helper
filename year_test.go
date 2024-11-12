package timeHelper

import (
	"log"
	"testing"
	"time"
)

func TestGetLastDayOfLastYear(t *testing.T) {
	log.Println("获取指定日期对应去年的最后一天")
	log.Println(GetLastDayOfLastYear(time.Now()))
}
func TestGetLastDayOfYear(t *testing.T) {
	log.Println("获取指定日期对应的年最后一天")
	log.Println(GetLastDayOfYear(time.Now()))
}
func TestGetYesterdayOfLastYear(t *testing.T) {
	log.Println("获取去年的昨天")
	log.Println(GetYesterdayOfLastYear(time.Now()))
}
func TestGetFirstDayOfYear(t *testing.T) {
	log.Println("获取指定日期对应的年第一天")
	log.Println(GetFirstDayOfYear(time.Now()))
}
func TestGetNextYearRange(t *testing.T) {
	log.Println("获取指定日期对应的明年")
	log.Println(GetNextYearRange(time.Now()))
}
