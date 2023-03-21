package holidays

import (
	"testing"
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
