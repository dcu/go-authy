package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SmsRequest struct {
	HttpResponse *http.Response
	Message      string `json:"message"`
}

func NewSmsRequest(response *http.Response) (*SmsRequest, error) {
	smsRequest := &SmsRequest{HttpResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return smsRequest, err
	}

	err = json.Unmarshal(body, &smsRequest)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return smsRequest, err
	}

	return smsRequest, nil
}

func (smsRequest *SmsRequest) Valid() bool {
	if smsRequest.HttpResponse.StatusCode == 200 {
		return true
	}

	return false
}
