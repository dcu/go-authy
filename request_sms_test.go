package authy

import(
    "testing"
)

func Test_RequestSms(t *testing.T) {
    api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    user, err := api.RegisterUser(UserOpts{
		Email: "foo@example.com",
		PhoneNumber: "432-123-1111",
		CountryCode: 1,
	})
    verification, err := api.RequestSms(user.Id, true)

    if err != nil {
        t.Error("External error found", err)
    }

    if !verification.Valid() {
        t.Error("Verification should be valid.")
    }
}
