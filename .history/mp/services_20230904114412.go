package mp

import "time"

func ToIso8601(targetDate time.Time) string {
	return targetDate.Format("2006-01-02T15:04:05")
}
