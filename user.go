package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// User is an Authy User
type User struct {
	HTTPResponse *http.Response
	ID           string
	UserData     struct {
		ID int `json:"id"`
	} `json:"user"`
	Errors  map[string]string `json:"errors"`
	Message string            `json:"message"`
	success bool              `json:"success"`
}

// NewUser returns an instance of User
func NewUser(httpResponse *http.Response) (*User, error) {
	userResponse := &User{HTTPResponse: httpResponse}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return userResponse, err
	}

	err = json.Unmarshal(body, userResponse)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return userResponse, err
	}

	userResponse.ID = strconv.Itoa(userResponse.UserData.ID)
	return userResponse, nil
}

// Valid returns true if the user was created successfully
func (response *User) Valid() bool {
	if response.HTTPResponse.StatusCode != 200 {
		return false
	}

	return true
}
