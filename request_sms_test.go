package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequestSms(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)

	verification, err := api.RequestSMS(user.ID, url.Values{"force": {"true"}})

	c.Nil(err)
	c.Equal("SMS is not enabled", verification.Message)
	c.False(verification.Valid())
}
