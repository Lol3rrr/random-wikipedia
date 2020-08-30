package general

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToInteger(t *testing.T) {
	tables := []struct {
		Name      string
		InputTime time.Time
		Result    int
	}{
		{
			Name:      "Valid",
			InputTime: time.Unix(0, 0),
			Result:    0,
		},
		{
			Name:      "Valid",
			InputTime: time.Unix(30*60, 0),
			Result:    30,
		},
		{
			Name:      "Valid",
			InputTime: time.Unix(1.5*60*60, 0),
			Result:    130,
		},
	}

	for _, table := range tables {
		inTime := table.InputTime
		res := table.Result

		t.Run(table.Name, func(t *testing.T) {
			out := TimeToInteger(inTime)

			assert.Equal(t, res, out)
		})
	}
}
