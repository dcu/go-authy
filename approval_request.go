package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ApprovalRequest is the approval request response.
type ApprovalRequest struct {
	HTTPResponse *http.Response

	Status   string `json:"status"`
	UUID     string `json:"uuid"`
	Notified bool   `json:"notified"`
}

// NewApprovalRequest returns an instance of ApprovalRequest.
func NewApprovalRequest(response *http.Response) (*ApprovalRequest, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonResponse := struct {
		Success         bool             `json:"success"`
		ApprovalRequest *ApprovalRequest `json:"approval_request"`
	}{}

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, err
	}

	approvalRequest := jsonResponse.ApprovalRequest
	approvalRequest.HTTPResponse = response

	return approvalRequest, nil
}

// Valid returns true if the approval request was valid.
func (request *ApprovalRequest) Valid() bool {
	if request.HTTPResponse.StatusCode == 200 {
		return true
	}

	return false
}
