package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TOTPTokenVerification encapsulates the response from Authy API when requesting verifying TOTP.
type TOTPTokenVerification struct {
	HTTPResponse      *http.Response
	Token             string `json:"token"`
	Message           string `json:"message"`
	DeviceInformation struct {
		City                 string `json:"city"`
		Country              string `json:"country"`
		IP                   string `json:"ip"`
		Region               string `json:"region"`
		RegistrationCity     string `json:"registration_city"`
		RegistrationCountry  string `json:"registration_country"`
		RegistrationDeviceID string `json:"registration_device_id"`
		RegistrationIP       string `json:"registration_ip"`
		RegistrationMethod   string `json:"registration_method"`
		RegistrationRegion   string `json:"registration_region"`
		OSType               string `json:"os_type"`
		ID                   string `json:"id"`
		RegistrationDate     string `json:"registration_date"`
	} `json:"device"`
	Success interface{} `json:"success"`
}

// NewTOTPTokenVerification creates an instance of a TOTPTokenVerification
func NewTOTPTokenVerification(response *http.Response) (*TOTPTokenVerification, error) {
	totpVerification := &TOTPTokenVerification{HTTPResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return totpVerification, err
	}

	err = json.Unmarshal(body, &totpVerification)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return totpVerification, err
	}

	return totpVerification, nil
}

// Valid returns true if the token verification was valid.
func (tokenVerification *TOTPTokenVerification) Valid() bool {
	if tokenVerification.HTTPResponse.StatusCode == 200 && tokenVerification.Token == "is valid" {
		return true
	}

	return false
}
