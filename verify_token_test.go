package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_VerifyTokenWithUnregisteredUser(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	verification, err := api.VerifyToken("0", "000000", url.Values{})
	c.Nil(err)
	c.Equal("User doesn't exist", verification.Message)
	c.False(verification.Valid())
}

func Test_VerifyTokenWithInvalidToken(t *testing.T) {
	c := require.New(t)
	api := newAPI()

	userResponse, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)
	c.True(userResponse.Valid())

	verification, err := api.VerifyToken(userResponse.ID, "000000", url.Values{})
	c.Nil(err)
	c.False(verification.Valid())
	c.Equal("Token is invalid", verification.Message)
}
