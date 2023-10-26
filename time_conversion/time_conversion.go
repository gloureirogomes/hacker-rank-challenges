package timeconversion

import (
	"time"
)

func TimeConversion(s string) string {
	parsedTime, err := time.Parse("15:04:05PM", s)
	if err != nil {
		return ""
	}

	return parsedTime.Format("15:04:05")
}
