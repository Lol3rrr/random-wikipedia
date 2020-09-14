package general

import "time"

// CalculateTimeToInterval is used to calculate the duration until you are back on the interval
func CalculateTimeToInterval(current time.Time, interval time.Duration) time.Duration {
	nextTime := current.Round(interval)
	if nextTime.After(current) {
		return nextTime.Sub(current)
	}
	if nextTime.Before(current) {
		return nextTime.Add(interval).Sub(current)
	}

	return interval
}
