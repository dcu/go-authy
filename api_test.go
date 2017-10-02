package authy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildURL(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	c.Equal("https://api.authy.com/path/to/endpoint", api.buildURL("/path/to/endpoint"))
	c.Equal("https://api.authy.com/path/to/endpoint", api.buildURL("path/to/endpoint"))
}
