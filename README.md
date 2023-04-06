# Holidays

The lib provide few methods to get china holidays.

## Usage

```shell
go get -u github.com/blackstorm/holidays
```

```golang
import "github.com/blackstorm/holidays"

holidays.IsChinaHoliday(time.Now())

// With location
loc, _ := time.LoadLocation("Asia/Shanghai")
now := time.Now().In(loc)
holidays.IsChinaHoliday(now)
```

## Holiday data
The China holiday data from https://github.com/NateScarlet/holiday-cn

## China Holiday API

https://jiejiariapi.com
