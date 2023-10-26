package timeconversion

import (
	"time"
)

func TimeConversion() string {
	input := "07:05:45PM"

	parsedTime, err := time.Parse("15:04:05PM", input)
	if err != nil {
		return ""
	}

	return parsedTime.Format("15:04:05")
}
