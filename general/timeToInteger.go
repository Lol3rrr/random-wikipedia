package general

import "time"

// TimeToInteger is used to convert time.Time to the integer used for
// storing in the Database
func TimeToInteger(input time.Time) int {
	hour := input.Hour()
	minute := input.Minute()

	return hour*100 + minute
}
