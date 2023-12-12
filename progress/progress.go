package progress

import (
	"fmt"
	"strconv"
	"time"
)

func Calculate(total, finished int, start time.Time) (doneString string, doneFloat, remainingFloat float64) {
	done := float64(finished) / float64(total)
	doneString = fmt.Sprintf("%.2f", done)
	doneFloat, _ = strconv.ParseFloat(doneString, 64)
	remaining := time.Since(start).Seconds() / done
	remainingString := fmt.Sprintf("%.2f", remaining)
	remainingFloat, _ = strconv.ParseFloat(remainingString, 64)
	if done == 1 {
		remainingFloat = 0
		doneString = "1.0"
	}
	return
}
