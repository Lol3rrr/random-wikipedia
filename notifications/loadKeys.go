package notifications

import "github.com/Lol3rrr/cvault"

func loadKeys(vault cvault.Session) (string, string, bool) {
	keyData, err := vault.ReadMap("/kv/data/wikipedia/vapidKeys")
	if err != nil {
		return "", "", false
	}

	pubKey, worked := keyData["PublicKey"].(string)
	if !worked || len(pubKey) <= 0 {
		return "", "", false
	}
	privKey, worked := keyData["PrivateKey"].(string)
	if !worked || len(privKey) <= 0 {
		return "", "", false
	}

	return pubKey, privKey, true
}
