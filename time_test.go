package timeHelper

import (
	"log"
	"testing"
	"time"
)

func TestGetDateTimeRange(t *testing.T) {
	log.Println("获取指定日期对应的时间范围")
	log.Println(GetDateTimeRange(time.Now()))
}
