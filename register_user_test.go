package authygo

import(
    "testing"
)

func Test_RegisterUserWithInvalidData(t *testing.T) {
    api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    userResponse, err := api.RegisterUser(UserOpts{
		Email: "foo",
		PhoneNumber: "123",
		CountryCode: 1,
	})

    if err == nil {
        t.Log("No comm error found")
    }

    if userResponse.Valid == true {
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

    userResponse, err := api.RegisterUser(UserOpts{
		Email: "foo@example.com",
		PhoneNumber: "432-123-1111",
		CountryCode: 1,
	})

    if err != nil {
        t.Log("Comm error found:",err)
    }

    if userResponse.Valid != true {
        t.Error("User should be valid.")
    }

    t.Log("Errors:", userResponse.Errors)

    if userResponse.User.Id == 0 {
        t.Error("User id should be set.")
    }
}


