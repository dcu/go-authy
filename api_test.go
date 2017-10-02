package authy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildURL(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	c.Equal(api.buildURL("/path/to/endpoint"), "https://api.authy.com/path/to/endpoint")
	c.Equal(api.buildURL("path/to/endpoint"), "https://api.authy.com/path/to/endpoint")
}
