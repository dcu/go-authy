package authy

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_SendApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1211", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	c.Nil(err)
	c.NotNil(approvalRequest)
	c.True(approvalRequest.Valid())
}

func Test_FindApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1131", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	c.Nil(err)
	c.True(approvalRequest.Valid())

	uuid := approvalRequest.UUID
	approvalRequest, err = api.FindApprovalRequest(uuid, url.Values{})

	c.Nil(err)
	c.Equal(OneTouchStatusPending, approvalRequest.Status)
	c.Equal(uuid, approvalRequest.UUID)
}

func Test_WaitForApprovalRequest(t *testing.T) {
	c := require.New(t)

	api := newAPI()

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1114", url.Values{})
	c.Nil(err)
	c.NotNil(user)

	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})
	c.Nil(err)
	c.NotNil(approvalRequest)
	c.True(approvalRequest.Valid())

	now := time.Now()
	maxDuration := 2 * time.Second
	status, err := api.WaitForApprovalRequest(approvalRequest.UUID, maxDuration, url.Values{"user_ip": {"234.78.25.2"}})
	c.Nil(err)

	elapsedTime := time.Since(now)
	c.True(elapsedTime < maxDuration+(1*time.Second))
	c.Equal(status, OneTouchStatusExpired)
}
