package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

/*
key                = webhook_secret
message            = webhook_body // raw webhook request body
received_signature = webhook_signature

expected_signature = hmac('sha256', message, key)

if expected_signature != received_signature
	throw SecurityErrorend

*/

// default go implementation avaliable here: https://pkg.go.dev/crypto/hmac#pkg-overview
func ValidMAC(message, receivedSignature, key []byte) bool {

	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectingSignature := mac.Sum(nil)
	sha := hex.EncodeToString(expectingSignature)

	return hmac.Equal(receivedSignature, []byte(sha))
}
