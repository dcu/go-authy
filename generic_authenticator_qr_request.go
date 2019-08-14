package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GenericAuthenticatorQRRequest encapsulates the response from Authy API when requesting Generic TOTP for 2FA apps.
type GenericAuthenticatorQRRequest struct {
	HTTPResponse *http.Response

	Label   string      `json:"label"`
	Issuer  string      `json:"issuer"`
	QRCode  string      `json:"qr_code"`
	Message string      `json:"message"`
	Success interface{} `json:"success"`
}

// NewGenericAuthenticatorQR creates an instance of a GenericAuthenticatorQRRequest
func NewGenericAuthenticatorQR(response *http.Response) (*GenericAuthenticatorQRRequest, error) {
	genericAuthenticatorQR := &GenericAuthenticatorQRRequest{HTTPResponse: response}
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
func (genericAuthenticatorApp *GenericAuthenticatorQRRequest) Valid() bool {
	return genericAuthenticatorApp.HTTPResponse.StatusCode == 200
}
