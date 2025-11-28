package token

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

type CookieBased struct{}

// GenerateToken expects a length of the token and returns
// the token encoded on base64 or an error if any.
func (cb CookieBased) GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		// should be impossible, so the only option is to kill the program
		logrus.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
