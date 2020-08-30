package general

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTimeToInterval(t *testing.T) {
	tables := []struct {
		Name          string
		InputTime     time.Time
		InputInterval time.Duration
		Result        time.Duration
	}{
		{
			Name:          "Valid",
			InputTime:     time.Unix(0, 0),
			InputInterval: 5 * time.Minute,
			Result:        5 * time.Minute,
		},
		{
			Name:          "Valid",
			InputTime:     time.Unix(3*60, 0),
			InputInterval: 5 * time.Minute,
			Result:        2 * time.Minute,
		},
	}

	for _, table := range tables {
		inTime := table.InputTime
		inInterval := table.InputInterval
		res := table.Result

		t.Run(table.Name, func(t *testing.T) {
			out := CalculateTimeToInterval(inTime, inInterval)

			assert.Equal(t, res, out)
		})
	}
}
