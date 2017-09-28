package authy

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_SendApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	c.Nil(err)
	c.NotNil(approvalRequest)
	c.True(approvalRequest.Valid())
}

func Test_FindApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	c.Nil(err)
	c.True(approvalRequest.Valid())

	uuid := approvalRequest.UUID
	approvalRequest, err = api.FindApprovalRequest(uuid, url.Values{})

	c.Nil(err)
	c.Equal("pending", approvalRequest.Status)
	c.Equal(uuid, approvalRequest.UUID)
}

func Test_WaitForApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	c.Nil(err)
	c.NotNil(user)

	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})
	c.Nil(err)
	c.NotNil(approvalRequest)
	c.True(approvalRequest.Valid())

	now := time.Now()
	status, err := api.WaitForApprovalRequest(approvalRequest.UUID, 1*time.Second, url.Values{"user_ip": {"234.78.25.2"}})
	c.Nil(err)

	elapsedTime := time.Since(now)

	c.True(elapsedTime < 1)
	c.Equal(status, OneTouchStatusExpired)
}
