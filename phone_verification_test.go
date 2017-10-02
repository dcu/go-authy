package authy

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func Test_StartPhoneVerification(t *testing.T) {
	api := newAPI()

	verification, err := api.StartPhoneVerification(1, "555-555-5555", SMS, url.Values{})

	require.Nil(t, err)
	require.NotNil(t, verification)
}

func Test_CheckPhoneVerification(t *testing.T) {
	api := newAPI()

	verification, err := api.CheckPhoneVerification(1, "555-555-5555", "000000", url.Values{})

	require.Nil(t, err)
	require.NotNil(t, verification)
}
