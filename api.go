package authy

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	// Logger is the default logger of this package. You can override it with your own.
	Logger = log.New(os.Stderr, "[authy] ", log.LstdFlags)
	client = &http.Client{}
)

// Authy contains credentials to connect to the Authy's API
type Authy struct {
	APIKey string
	APIURL string
}

// NewAuthyAPI returns an instance of Authy pointing to production.
func NewAuthyAPI(apiKey string) *Authy {
	apiURL := "https://api.authy.com"
	return &Authy{
		APIKey: apiKey,
		APIURL: apiURL,
	}
}

// NewSandboxAuthyAPI returns an instance of Authy pointing to sandbox. Use this to implement automated tests.
func NewSandboxAuthyAPI(apiKey string) *Authy {
	apiURL := "https://sandbox-api.authy.com"
	return &Authy{
		APIKey: apiKey,
		APIURL: apiURL,
	}
}

// RegisterUser register a new user given an email and phone number.
func (authy *Authy) RegisterUser(email string, countryCode int, phoneNumber string, params url.Values) (*User, error) {
	Logger.Println("Creating Authy user with", email, ",", phoneNumber, "and", countryCode)

	path := "/protected/json/users/new"

	params.Set("user[cellphone]", phoneNumber)
	params.Set("user[country_code]", strconv.Itoa(countryCode))
	params.Set("user[email]", email)

	response, err := authy.DoRequest("POST", path, params)

	if err != nil {
		return nil, err
	}

	userResponse, err := NewUser(response)
	return userResponse, err
}

// VerifyToken verifies the given token
func (authy *Authy) VerifyToken(userID string, token string, params url.Values) (*TokenVerification, error) {
	path := "/protected/json/verify/" + url.QueryEscape(token) + "/" + url.QueryEscape(userID)

	response, err := authy.DoRequest("GET", path, params)

	if err != nil {
		Logger.Println("Error while contacting the API:", err)
		return nil, err
	}

	defer response.Body.Close()

	tokenVerification, err := NewTokenVerification(response)
	return tokenVerification, err
}

// RequestSMS requests a SMS for the given userID
func (authy *Authy) RequestSMS(userID string, params url.Values) (*SMSRequest, error) {
	path := "/protected/json/sms/" + url.QueryEscape(userID)
	response, err := authy.DoRequest("GET", path, params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	smsVerification, err := NewSMSRequest(response)
	return smsVerification, err
}

// RequestPhoneCall requests a phone call for the given user
func (authy *Authy) RequestPhoneCall(userID string, params url.Values) (*PhoneCallRequest, error) {
	path := "/protected/json/call/" + url.QueryEscape(userID)

	response, err := authy.DoRequest("GET", path, params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	smsVerification, err := NewPhoneCallRequest(response)
	return smsVerification, err
}

// DoRequest performs a HTTP request to the Authy API
func (authy *Authy) DoRequest(method string, path string, params url.Values) (*http.Response, error) {
	apiURL := authy.buildURL(path)

	// Set api_key to all requests.
	params.Set("api_key", authy.APIKey)

	var bodyReader io.Reader
	switch method {
	case "POST":
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
		}
	case "GET":
		{
			apiURL += "?" + params.Encode()
		}
	}

	request, err := http.NewRequest(method, apiURL, bodyReader)
	if method == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if err != nil {
		Logger.Println("Error creating HTTP request:", err)
		return nil, err
	}
	response, err := client.Do(request)

	return response, err
}

func (authy *Authy) buildURL(path string) string {
	url := authy.APIURL + "/" + path

	return url
}
