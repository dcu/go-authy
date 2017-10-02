package authy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RegisterUserWithInvalidData(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	userResponse, err := api.RegisterUser("foo", 1, "123", url.Values{})
	c.Nil(err)
	c.False(userResponse.Valid())

	t.Log("Errors:", userResponse.Errors)
	c.Equal("is invalid", userResponse.Errors["email"])
}

func Test_RegisterUserWithValidData(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	userResponse, err := api.RegisterUser("foo@example.com", 1, "432-123-1115", url.Values{})
	c.Nil(err)
	c.True(userResponse.Valid())
	t.Log("Errors:", userResponse.Errors)
	c.NotEmpty(userResponse.ID)
}

func Test_UserStatus(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	userResponse, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)
	c.True(userResponse.Valid())

	t.Log("Errors:", userResponse.Errors)

	c.NotEmpty(userResponse.ID)

	userStatus, err := api.UserStatus(userResponse.ID, url.Values{})
	c.Nil(err)
	c.Equal(200, userStatus.HTTPResponse.StatusCode)
	c.Equal(1, userStatus.StatusData.Country)
	c.True(userStatus.Success)
}
