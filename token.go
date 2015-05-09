package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TokenVerification struct {
	HttpResponse *http.Response
	Message      string      `json:"message"`
	Token        string      `json:"token"`
	Success      interface{} `json:"success"`
}

func NewTokenVerification(response *http.Response) (*TokenVerification, error) {
	tokenVerification := &TokenVerification{HttpResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return tokenVerification, err
	}

	err = json.Unmarshal(body, &tokenVerification)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return tokenVerification, err
	}

	return tokenVerification, nil
}

func (verification *TokenVerification) Valid() bool {
	if verification.HttpResponse.StatusCode == 200 && verification.Token == "is valid" {
		return true
	}

	return false
}
