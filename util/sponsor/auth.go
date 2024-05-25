package sponsor

import (
	"context"
	"fmt"
	"time"

	"github.com/evcc-io/evcc/api/proto/pb"
	"github.com/evcc-io/evcc/util/cloud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
