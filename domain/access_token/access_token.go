package access_token

import (
	"github.com/privatesquare/bkst-go-utils/utils/dateutils"
	"time"
)

const (
	defaultExpirationTimeInSeconds = 90
)

type AccessToken struct {
	ClientId     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	AccessToken  string    `json:"access_token"`
	UserId       int64     `json:"user_id"`
	Expires      time.Time `json:"expires"`
}

func (at *AccessToken) New() error {
	at.Expires = dateutils.GetDateTimeNow().Add(defaultExpirationTimeInSeconds * time.Second)
	return nil
}

func (at *AccessToken) IsExpired() bool {
	return dateutils.GetDateTimeNow().After(at.Expires)
}
