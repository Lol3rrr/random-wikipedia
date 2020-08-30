package wikipedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomArticle(t *testing.T) {
	tables := []struct {
		Name        string
		InputList   list
		ResultError bool
	}{
		{
			Name: "Valid",
			InputList: list{
				Articles: []Article{
					{
						ID:    123,
						Title: "testTitle1",
						URL:   "testURL1",
					},
					{
						ID:    234,
						Title: "testTitle2",
						URL:   "testURL2",
					},
				},
			},
			ResultError: false,
		},
		{
			Name: "Empty Articles List",
			InputList: list{
				Articles: []Article{},
			},
			ResultError: true,
		},
	}

	for _, table := range tables {
		inList := table.InputList
		resError := table.ResultError

		t.Run(table.Name, func(t *testing.T) {
			outArticle, outError := inList.GetRandomArticle()

			if resError {
				assert.NotNil(t, outError)
			} else {
				assert.Contains(t, inList.Articles, outArticle)
				assert.Nil(t, outError)
			}
		})
	}
}
