package authy

import (
	"net/url"
	"testing"
)

func Test_SendApprovalRequest(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if !approvalRequest.Valid() {
		t.Error("Apprval request should be valid.")
	}
}

func Test_FindApprovalRequest(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if !approvalRequest.Valid() {
		t.Error("Apprval request should be valid.")
	}

	uuid := approvalRequest.UUID
	approvalRequest, err = api.FindApprovalRequest(uuid, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if approvalRequest.Status != "pending" {
		t.Error("Approval request status is wrong")
	}

	if uuid != approvalRequest.UUID {
		t.Error("Approval request doesn't match.")
	}
}
