package authy

import (
	"net/url"
	"testing"
)

func Test_RequestSms(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")

	user, err := api.RegisterUser("foo@example.com", 1, "432-123-1111", url.Values{})
	verification, err := api.RequestSMS(user.ID, url.Values{"force": {"true"}})

	if err != nil {
		t.Error("External error found", err)
	}

	if !verification.Valid() {
		t.Error("Verification should be valid.")
	}
}
