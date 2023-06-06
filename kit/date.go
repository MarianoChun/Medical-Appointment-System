package kit

import (
	"fmt"
	"strconv"
	"time"
)

func TimeToDateString(time time.Time) string {
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(time.Year()), strconv.Itoa(int(time.Month())), strconv.Itoa(time.Day()))
}

func TimeToDateTimeString(time time.Time) string {
	return fmt.Sprintf("%s-%s-%s %s:%s:%s", strconv.Itoa(time.Year()), strconv.Itoa(int(time.Month())), strconv.Itoa(time.Day()), strconv.Itoa(time.Hour()), strconv.Itoa(time.Minute()), strconv.Itoa(time.Second()))
}
