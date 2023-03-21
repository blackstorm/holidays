package holidays

import (
	_ "embed"
	"encoding/json"
	"time"
)

//go:embed china2023.json
var holidays2023 []byte

type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	IsOffDay bool   `json:"isOffDay"`
}

type Holidays struct {
	Schema string    `json:"$schema"`
	Id     string    `json:"$id`
	Year   int       `json:"year"`
	Papers []string  `json:"papers"`
	Days   []Holiday `json:"days`
}

var (
	chinaHolidaysMap map[int]map[string]Holiday
)

func init() {
	chinaHolidaysMap = make(map[int]map[string]Holiday)

	var holidays Holidays
	err := json.Unmarshal(holidays2023, &holidays)
	if err != nil {
		panic(err)
	} else {
		holidays2023 = nil
	}

	m := make(map[string]Holiday)
	for _, day := range holidays.Days {
		m[day.Date] = day
	}

	chinaHolidaysMap[2023] = m
}

func GetChinaHolidays(year int) map[string]Holiday {
	return chinaHolidaysMap[year]
}

func GetNowYearChinaHolidays(loc *time.Location) map[string]Holiday {
	year := time.Now().In(loc).Year()
	return GetChinaHolidays(year)

}

func TodayIsChinaHoliday(loc *time.Location) bool {
	now := time.Now().In(loc)
	return IsChinaHoliday(now)
}

func IsChinaHoliday(time time.Time) bool {
	date := time.Format("2006-01-02")
	holidays := GetNowYearChinaHolidays(time.Location())
	_, ok := holidays[date]
	return ok
}
