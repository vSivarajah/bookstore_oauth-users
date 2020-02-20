package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	accessToken := GetNewAccessToken()
	assert.False(t, accessToken.isExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", accessToken.AccessToken, "new access token should not have defined access token id")
	assert.True(t, accessToken.UserId == 0, "New access token should not have an associated userid")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.isExpired(), "empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "accesstoken created three hours from now should NOT be expired")
}
