package notifications

import (
	"github.com/Lol3rrr/cvault"
	"github.com/SherClockHolmes/webpush-go"
)

func generateKeys(vault cvault.Session) (string, string, error) {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		return "", "", err
	}

	keyMap := map[string]interface{}{
		"PublicKey":  publicKey,
		"PrivateKey": privateKey,
	}
	err = vault.WriteMapData("/kv/data/wikipedia/vapidKeys", keyMap)
	if err != nil {
		return "", "", err
	}

	return publicKey, privateKey, nil
}
