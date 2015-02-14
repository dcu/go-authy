package authy

import(
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

type UserOpts struct {
    Email string
    PhoneNumber string
    CountryCode int
}

type UserId struct {
	Id int `json:"id"`
}

type User struct {
	HttpResponse *http.Response
	Id int
	UserId UserId `json:"user"`
    Errors map[string]string `json:"errors"`
    Message string `json:"message"`
	success bool `json:"success"`
}

func NewUser(httpResponse *http.Response) (*User, error) {
	userResponse := &User{HttpResponse: httpResponse}

    defer httpResponse.Body.Close()
    body, err := ioutil.ReadAll(httpResponse.Body)

    if err != nil {
        log.Fatal("Error reading from API:", err)
        return userResponse, err
    }

    err = json.Unmarshal(body, userResponse)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return userResponse, err
    }

	userResponse.Id = userResponse.UserId.Id
	return userResponse, nil
}

func (response *User) Valid() bool {
	if response.HttpResponse.StatusCode != 200 {
		return false
	}

	return true
}

