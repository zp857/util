package timeutil

import "time"

func IsBetween(t time.Time, start, end time.Time) bool {
	return t.Equal(start) || t.After(start) && t.Before(end) || t.Equal(end)
}
