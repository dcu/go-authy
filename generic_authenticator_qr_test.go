package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateAuthenticatorQR(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)

	genericAppAuthenticator, err := api.GenerateGenericAuthenticatorQR(user.ID, "foo app", 300, url.Values{})

	c.Nil(err)
	c.False(genericAppAuthenticator.Valid())
}
