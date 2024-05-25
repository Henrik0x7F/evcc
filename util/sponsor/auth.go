package sponsor

import (
	"time"
)

var (
	Subject, Token string
	ExpiresAt      time.Time
)

const unavailable = "sponsorship unavailable"

func IsAuthorized() bool {
	return true
}

func IsAuthorizedForApi() bool {
	return true
}

// check and set sponsorship token
func ConfigureSponsorship(token string) error {
	return nil
}
