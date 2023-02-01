package timeutil

import "time"

func FormatInLocation(layout string, t time.Time, loc *time.Location) string {
	return t.In(loc).Format(layout)
}
