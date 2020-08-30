package notifications

import "random_wikipedia/general"

// Session is used to abstract away the actual implementation
type Session interface {
	GetPublicKey() string
	SendNotification(text []byte, user general.User) error
}

type session struct {
	PublicKey       string
	PrivateKey      string
	SubscriberEmail string
}
