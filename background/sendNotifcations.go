package background

import (
	"random_wikipedia/database"
	"random_wikipedia/general"
	"random_wikipedia/login"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"
	"time"
)

// SendNotifications checks if there are any notifications
// that need to be send in the current part
func SendNotifications(loginSession login.Session, wikipediaSession wikipedia.Session, notificationSession notifications.Session, dbSession database.Session) {
	interval := 15 * time.Minute
	waitDuraiton := general.CalculateTimeToInterval(time.Now().UTC(), interval)
	time.Sleep(waitDuraiton)

	for {
		startTime := time.Now()

		currentTime := startTime.UTC()
		users, err := dbSession.LoadUsersNotifyTime(general.TimeToInteger(currentTime))
		if err == nil {
			for _, user := range users {
				sendUserNotification(user, wikipediaSession, notificationSession)
			}
		}

		time.Sleep(interval - time.Now().Sub(startTime))
	}
}
