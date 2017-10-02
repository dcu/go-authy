package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequestPhoneCall(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1511", url.Values{})
	c.Nil(err)
	c.True(user.Valid())

	verification, err := api.RequestPhoneCall(user.ID, url.Values{"force": {"true"}})
	c.Nil(err)
	c.Equal("Phone calls are not enabled on this account", verification.Message)
	c.False(verification.Valid())
}
