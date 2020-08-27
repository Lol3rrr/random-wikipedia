package wikipedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetID(t *testing.T) {
	testList := &list{
		ID: 123,
	}

	outputID := testList.GetID()
	assert.Equal(t, testList.ID, outputID)
}

func TestGetTitle(t *testing.T) {
	testList := &list{
		Title: "testTitle",
	}

	outputTitle := testList.GetTitle()
	assert.Equal(t, testList.Title, outputTitle)
}

func TestGetArticles(t *testing.T) {
	testList := &list{
		Articles: []Article{
			{
				ID:    123,
				Title: "testTitle",
				URL:   "testURL",
			},
		},
	}

	outputArticles := testList.GetArticles()
	assert.Equal(t, testList.Articles, outputArticles)
}

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
