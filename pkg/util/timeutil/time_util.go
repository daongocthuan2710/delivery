package timeutil

import "time"

const DateISOFormat = "2006-01-02T15:04:05.000Z"

func FormatInLocation(layout string, t time.Time, loc *time.Location) string {
	return t.In(loc).Format(layout)
}

// ParseISODate ...
func ParseISODate(value string) time.Time {
	t, _ := time.Parse(DateISOFormat, value)
	return t
}
