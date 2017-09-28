package authy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildURL(t *testing.T) {
	c := require.New(t)

	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	c.Equal(api.buildURL("/path/to/endpoint"), "https://sandbox-api.authy.com/path/to/endpoint")
	c.Equal(api.buildURL("path/to/endpoint"), "https://sandbox-api.authy.com/path/to/endpoint")
}
