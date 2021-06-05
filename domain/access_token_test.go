package domain

import (
	"github.com/privatesquare/bkst-go-utils/utils/dateutils"
	"testing"
	"time"
)

func TestAccessToken_New(t *testing.T) {
	at := new(AccessToken)
	at.SetExpiration()
	if at.IsExpired() {
		t.Error("New access token should not be expired")
	}
	if at.UserId != 0 {
		t.Error("New access token should not have an associated user id")
	}
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := new(AccessToken)
	at.Expires = dateutils.GetDateTimeNow().Add(90 * time.Minute)
	if at.IsExpired() {
		t.Error("expected access token to be valid")
	}
}
