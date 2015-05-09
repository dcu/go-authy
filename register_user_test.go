package authy

import (
	"net/url"
	"testing"
)

func Test_RegisterUserWithInvalidData(t *testing.T) {
	api := NewSandboxAuthyApi("bf12974d70818a08199d17d5e2bae630")

	userResponse, err := api.RegisterUser("foo", 1, "123", url.Values{})

	if err == nil {
		t.Log("No comm error found")
	}

	if userResponse.Valid() {
		t.Error("User should not be valid.")
	}

	t.Log("Errors:", userResponse.Errors)
	if userResponse.Errors["email"] != "is invalid" {
		t.Error("Invalid error returned by server.")
	}
}

func Test_RegisterUserWithValidData(t *testing.T) {
	api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
	api.ApiUrl = "http://sandbox-api.authy.com"

	userResponse, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})

	if err != nil {
		t.Log("Comm error found:", err)
	}

	if !userResponse.Valid() {
		t.Error("User should be valid.")
	}

	t.Log("Errors:", userResponse.Errors)

	if userResponse.Id == 0 {
		t.Error("User id should be set.")
	}
}
