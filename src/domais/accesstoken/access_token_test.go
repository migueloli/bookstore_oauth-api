package accesstoken

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("Expiration time should be 24 hours.")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("Brand new AccessToken should not be expired.")
	}

	if at.AccessToken != "" {
		t.Error("Brand new AccessToken should not have defined AccessToken ID.")
	}

	if at.UserID != 0 {
		t.Error("Brand new AccessToken should not have associated a UserID.")
	}
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("Empty AccessToken should be expired by default.")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("AccessToken expiring 3 hours from now should not be expired by default.")
	}
}
