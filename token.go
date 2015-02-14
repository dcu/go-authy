package authy

import(
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

type TokenVerification struct {
	HttpResponse *http.Response
    Message string `json:"message"`
	token string `json:"token"`
}

func NewTokenVerification(response *http.Response) (*TokenVerification, error) {
	tokenVerification := &TokenVerification{HttpResponse: response}
    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal("Error reading from API:", err)
        return tokenVerification, err
    }

    err = json.Unmarshal(body, &tokenVerification)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return tokenVerification, err
    }

	return tokenVerification, nil
}

func (verification *TokenVerification) Valid() bool {
	if verification.HttpResponse.StatusCode == 200 && verification.token == "is valid" {
		return true
	}

	return false
}

