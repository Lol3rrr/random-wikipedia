package background

import (
	"random_wikipedia/wikipedia"
	"time"
)

// UpdateWikipediaLists is used to update the Wikipedia lists periodically
func UpdateWikipediaLists(wikipediaSession wikipedia.Session, sleepTime time.Duration) {
	for {
		wikipediaSession.UpdateLists()

		time.Sleep(sleepTime)
	}
}
