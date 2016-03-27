package authy

import (
	"net/url"
	"testing"
)

func Test_VerifyTokenWithUnregisteredUser(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	verification, err := api.VerifyToken("0", "000000", url.Values{})

	if err == nil {
		t.Log("No comm error found")
	}

	if verification.Valid() {
		t.Error("Verification should not be valid.")
	}

	if verification.Message != "User doesn't exist." {
		t.Error("Invalid test using a registered user.")
	}
}

func Test_VerifyTokenWithInvalidToken(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	userResponse, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	verification, err := api.VerifyToken(userResponse.ID, "000000", url.Values{})

	if err != nil {
		t.Log("No comm error found")
	}

	if verification.Valid() {
		t.Error("Verification should not be valid.")
	}

	if verification.Message == "User doesn't exist." {
		t.Error("Invalid test using an unregistered user.")
	}
}
