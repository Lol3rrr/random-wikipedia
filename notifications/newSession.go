package notifications

import (
	"github.com/Lol3rrr/cvault"
)

// NewSession is used to obtain a new notification session
func NewSession(subEmail string, vault cvault.Session) (Session, error) {
	publicKey, privateKey, found := loadKeys(vault)
	if !found {
		pubKey, privKey, err := generateKeys(vault)
		if err != nil {
			return nil, err
		}
		publicKey = pubKey
		privateKey = privKey
	}

	return &session{
		PublicKey:       publicKey,
		PrivateKey:      privateKey,
		SubscriberEmail: subEmail,
	}, nil
}
