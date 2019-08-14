package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GenericAuthenticatorApp encapsulates the response from Authy API when requesting Generic TOTP for 2FA apps.
type GenericAuthenticatorApp struct {
	HTTPResponse *http.Response
	Label        string      `json:"label"`
	Issuer       string      `json:"issuer"`
	QRCode       string      `json:"qr_code"`
	Message      string      `json:"message"`
	Success      interface{} `json:"success"`
}

// NewGenericAuthenticatorQR creates an instance of a GenericAuthenticatorApp
func NewGenericAuthenticatorQR(response *http.Response) (*GenericAuthenticatorApp, error) {
	genericAuthenticatorQR := &GenericAuthenticatorApp{HTTPResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return genericAuthenticatorQR, err
	}

	err = json.Unmarshal(body, &genericAuthenticatorQR)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return genericAuthenticatorQR, err
	}

	return genericAuthenticatorQR, nil
}

// Valid returns true if the Generic Authenticator QR was generated
func (request *GenericAuthenticatorApp) Valid() bool {
	return request.HTTPResponse.StatusCode == 200
}
