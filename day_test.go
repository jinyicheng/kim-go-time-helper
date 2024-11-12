package timeHelper

import (
	"log"
	"testing"
	"time"
)

func TestGetYesterday(t *testing.T) {
	log.Println("获取昨天")
	log.Println(GetYesterday(time.Now()))
}
func TestGetTomorrow(t *testing.T) {
	log.Println("获取明天")
	log.Println(GetTomorrow(time.Now()))

}
func TestGetLastMonday(t *testing.T) {
	log.Println("获取上个星期一")
	log.Println(GetLastMonday(time.Now()))

}
func TestGetLastTuesday(t *testing.T) {
	log.Println("获取上个星期二")
	log.Println(GetLastTuesday(time.Now()))

}
func TestGetLastWednesday(t *testing.T) {
	log.Println("获取上个星期三")
	log.Println(GetLastWednesday(time.Now()))

}
func TestGetLastThursday(t *testing.T) {
	log.Println("获取上个星期四")
	log.Println(GetLastThursday(time.Now()))
}
func TestGetLastFriday(t *testing.T) {
	log.Println("获取上个星期五")
	log.Println(GetLastFriday(time.Now()))
}
func TestGetLastSaturday(t *testing.T) {
	log.Println("获取上个星期六")
	log.Println(GetLastSaturday(time.Now()))
}
func TestGetLastSunday(t *testing.T) {
	log.Println("获取上个星期日")
	log.Println(GetLastSunday(time.Now()))
}

func TestGetThisWeekTuesday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期二")
	log.Println(GetThisWeekTuesday(time.Now()))
}

func TestGetThisWeekWednesday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期三")
	log.Println(GetThisWeekWednesday(time.Now()))
}

func TestGetThisWeekThursday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期四")
	log.Println(GetThisWeekThursday(time.Now()))
}

func TestGetThisWeekFriday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期五")
	log.Println(GetThisWeekFriday(time.Now()))
}

func TestGetThisWeekSaturday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期六")
	log.Println(GetThisWeekSaturday(time.Now()))
}

func TestGetThisWeekSunday(t *testing.T) {
	log.Println("获取传入时间对应的本周星期日")
	log.Println(GetThisWeekSunday(time.Now()))
}
