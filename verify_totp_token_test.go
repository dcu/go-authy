package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CheckTOTPVerification(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)

	genericAppAuthenticator, err := api.CheckTOTPVerification(user.ID, "3333", url.Values{})

	c.Nil(err)
	c.Equal("Token is invalid", genericAppAuthenticator.Message)
	c.False(genericAppAuthenticator.Valid())
}
