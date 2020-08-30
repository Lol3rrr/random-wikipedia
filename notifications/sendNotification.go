package notifications

import (
	"encoding/json"
	"random_wikipedia/general"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/sirupsen/logrus"
)

func (s *session) SendNotification(text []byte, user general.User) error {
	sub := &webpush.Subscription{}
	err := json.Unmarshal([]byte(user.Subscription), &sub)
	if err != nil {
		return err
	}

	options := &webpush.Options{
		Subscriber:      s.SubscriberEmail,
		VAPIDPublicKey:  s.PublicKey,
		VAPIDPrivateKey: s.PrivateKey,
	}

	resp, err := webpush.SendNotification(text, sub, options)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logrus.Errorf("Closing Notifcation body: %v", err)
		}
	}()

	return nil
}
