package wikipedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterArticle(t *testing.T) {
	tables := []struct {
		Name       string
		InputTitle string
		Result     bool
	}{
		{
			Name:       "Valid Input",
			InputTitle: "Normal Article",
			Result:     true,
		},
		{
			Name:       "Template Article",
			InputTitle: "Template: Some Topic",
			Result:     false,
		},
		{
			Name:       "Template Talk Article",
			InputTitle: "Template talk: Some Topic",
			Result:     false,
		},
		{
			Name:       "Portal",
			InputTitle: "Portal: Some Topic",
			Result:     false,
		},
		{
			Name:       "Category",
			InputTitle: "Category: Some Category",
			Result:     false,
		},
		{
			Name:       "Wikipedia Article",
			InputTitle: "Wikipedia: Something",
			Result:     false,
		},
		{
			Name:       "Talk Article",
			InputTitle: "Talk: Some Topic",
			Result:     false,
		},
	}

	for _, table := range tables {
		inTitle := table.InputTitle
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := filterArticle(inTitle)

			assert.Equal(t, result, output)
		})
	}
}
