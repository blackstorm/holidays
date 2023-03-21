# Hlidays

The lib provide few methods to get holidays for a given country and year.

Only support China holidays now.

## Usage

```
go get -u github.com/blackstorm/holidays
```

```
import "github.com/blackstorm/holidays"

holidays.IsChinaHoliday(time.Now())

// With location
loc, _ := time.LoadLocation("Asia/Shanghai")
now := time.Now().In(loc)
holidays.IsChinaHoliday(now)
```

## Holiday data
The China holiday data from https://github.com/NateScarlet/holiday-cn
