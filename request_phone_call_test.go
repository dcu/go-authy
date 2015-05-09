package authy

import (
	"net/url"
	"testing"
)

func Test_RequestPhoneCall(t *testing.T) {
	api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
	api.ApiUrl = "http://sandbox-api.authy.com"

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	verification, err := api.RequestPhoneCall(user.Id, url.Values{"force": {"true"}})

	if err != nil {
		t.Error("External error found", err)
	}

	if !verification.Valid() {
		t.Error("Verification should be valid.")
	}
}
