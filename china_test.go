package holidays

import (
	"testing"
	"time"
)

func TestGetChinaHolidays(t *testing.T) {
	holidays := GetChinaHolidays(2023)
	if _, ok := holidays["2023-10-01"]; !ok {
		t.Fail()
	}
	if _, ok := holidays["2023-10-15"]; ok {
		t.Fail()
	}
}

func TestIsChinaHoliday(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2023-10-01")
	holidays := GetChinaHolidays(date.Year())
	if _, ok := holidays[date.Format("2006-01-02")]; !ok {
		t.Fail()
	}
}
