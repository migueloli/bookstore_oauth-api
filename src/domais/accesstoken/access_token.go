package accesstoken

import (
	"strings"
	"time"

	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

const expirationTime = 24

// AccessToken is the OAuth token struct.
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate is a method to validate the AccessToken.
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token.")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid user ID.")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client ID.")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expires.")
	}

	return nil
}

// GetNewAccessToken is a function to generate a new access token.
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired is a method to verify if the access token is expired or not.
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
