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
	Logger = log.New(os.Stderr, "[authy] ", log.LstdFlags)
	client = &http.Client{}
)

type Authy struct {
	ApiKey string
	ApiUrl string
}

func NewAuthyApi(apiKey string) *Authy {
	apiUrl := "https://api.authy.com"
	return &Authy{
		ApiKey: apiKey,
		ApiUrl: apiUrl,
	}
}

func NewSandboxAuthyApi(apiKey string) *Authy {
	apiUrl := "https://sandbox-api.authy.com"
	return &Authy{
		ApiKey: apiKey,
		ApiUrl: apiUrl,
	}
}

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

func (authy *Authy) VerifyToken(userId int, token string, params url.Values) (*TokenVerification, error) {
	path := "/protected/json/verify/" + url.QueryEscape(token) + "/" + url.QueryEscape(strconv.Itoa(userId))

	response, err := authy.DoRequest("GET", path, params)

	if err != nil {
		Logger.Println("Error while contacting the API:", err)
		return nil, err
	}

	defer response.Body.Close()

	tokenVerification, err := NewTokenVerification(response)
	return tokenVerification, err
}

func (authy *Authy) RequestSms(userId int, params url.Values) (*SmsRequest, error) {
	path := "/protected/json/sms/" + url.QueryEscape(strconv.Itoa(userId))
	response, err := authy.DoRequest("GET", path, params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	smsVerification, err := NewSmsRequest(response)
	return smsVerification, err
}

func (authy *Authy) RequestPhoneCall(userId int, params url.Values) (*PhoneCallRequest, error) {
	path := "/protected/json/call/" + url.QueryEscape(strconv.Itoa(userId))

	response, err := authy.DoRequest("GET", path, params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	smsVerification, err := NewPhoneCallRequest(response)
	return smsVerification, err
}

func (authy *Authy) DoRequest(method string, path string, params url.Values) (*http.Response, error) {
	apiUrl := authy.buildUrl(path)

	// Add api_key to all requests.
	params.Add("api_key", authy.ApiKey)

	var bodyReader io.Reader
	switch method {
	case "POST":
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
		}
	case "GET":
		{
			apiUrl += "?" + params.Encode()
		}
	}

	request, err := http.NewRequest(method, apiUrl, bodyReader)
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

func (authy *Authy) buildUrl(path string) string {
	url := authy.ApiUrl + "/" + path

	return url
}
