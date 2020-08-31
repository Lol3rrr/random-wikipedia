package background

import (
	"random_wikipedia/database"
	"random_wikipedia/general"
	"random_wikipedia/login"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"
	"time"

	"github.com/sirupsen/logrus"
)

// SendNotifications checks if there are any notifications
// that need to be send in the current part
func SendNotifications(loginSession login.Session, wikipediaSession wikipedia.Session, notificationSession notifications.Session, dbSession database.Session) {
	interval := 15 * time.Minute
	waitDuraiton := general.CalculateTimeToInterval(time.Now(), interval)

	logrus.Infof("Waiting %d to be back on the interval", waitDuraiton)

	time.Sleep(waitDuraiton)

	for {
		startTime := time.Now()

		currentTime := startTime.UTC()
		currentTimeStamp := general.TimeToInteger(currentTime)

		logrus.Infof("Sending Notifications at %d", currentTimeStamp)

		users, err := dbSession.LoadUsersNotifyTime(currentTimeStamp)
		if err == nil {
			logrus.Infof("Sending %d Notifications", len(users))
			for _, user := range users {
				sendUserNotification(user, wikipediaSession, notificationSession)
			}
		}

		time.Sleep(interval - time.Now().Sub(startTime))
	}
}
