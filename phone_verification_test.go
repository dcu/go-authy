package authy

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_StartPhoneVerification(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	verification, err := api.StartPhoneVerification(1, "555-555-5555", SMS, url.Values{})

	assert.Nil(t, err)
	assert.NotNil(t, verification)
}

func Test_CheckPhoneVerification(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	verification, err := api.CheckPhoneVerification(1, "555-555-5555", "000000", url.Values{})

	assert.Nil(t, err)
	assert.NotNil(t, verification)
}
